// Package packetlogger is adapted based on code in the slog guide (https://github.com/golang/example/blob/master/slog-handler-guide/guide.md)
package packetlogger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// !+IndentHandler
type IndentHandler struct {
	opts           Options
	preformatted   []byte   // data from WithGroup and WithAttrs
	unopenedGroups []string // groups from WithGroup that haven't been opened
	indentLevel    int      // same as number of opened groups so far
	mu             *sync.Mutex
	out            io.Writer
}

// !-IndentHandler

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler

	AddTime   bool
	AddLevel  bool
	AddSource bool
}

func New(out io.Writer, opts *Options) *IndentHandler {
	h := &IndentHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}

func (h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

// !+WithGroup
func (h *IndentHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	h2 := *h
	// Add an unopened group to h2 without modifying h.
	h2.unopenedGroups = make([]string, len(h.unopenedGroups)+1)
	copy(h2.unopenedGroups, h.unopenedGroups)
	h2.unopenedGroups[len(h2.unopenedGroups)-1] = name
	return &h2
}

// !-WithGroup

// !+WithAttrs
func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	h2 := *h
	// Force an append to copy the underlying array.
	// pre := slices.Clip(h.preformatted)
	pre := []byte{}
	// Add all groups from WithGroup that haven't already been added.
	h2.preformatted = h2.appendUnopenedGroups(pre, h2.indentLevel)
	// Each of those groups increased the indent level by 1.
	h2.indentLevel += len(h2.unopenedGroups)
	// Now all groups have been opened.
	h2.unopenedGroups = nil
	// Pre-format the attributes.
	for _, a := range attrs {
		h2.preformatted = h2.appendAttr(h2.preformatted, a, h2.indentLevel)
	}
	return &h2
}

func (h *IndentHandler) appendUnopenedGroups(buf []byte, indentLevel int) []byte {
	for _, g := range h.unopenedGroups {
		buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", g)
		indentLevel++
	}
	return buf
}

// !-WithAttrs

// !+Handle
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	bufp := allocBuf()
	buf := *bufp
	defer func() {
		*bufp = buf
		freeBuf(bufp)
	}()
	if h.opts.AddTime {
		if !r.Time.IsZero() {
			buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
		}
	}
	if h.opts.AddLevel {
		buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	}
	if h.opts.AddSource {
		if r.PC != 0 {
			fs := runtime.CallersFrames([]uintptr{r.PC})
			f, _ := fs.Next()
			// Optimize to minimize allocation.
			srcbufp := allocBuf()
			defer freeBuf(srcbufp)
			*srcbufp = append(*srcbufp, f.File...)
			*srcbufp = append(*srcbufp, ':')
			*srcbufp = strconv.AppendInt(*srcbufp, int64(f.Line), 10)
			buf = h.appendAttr(buf, slog.String(slog.SourceKey, string(*srcbufp)), 0)
		}
	}

	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	// Insert preformatted attributes just after built-in ones.
	buf = append(buf, h.preformatted...)
	if r.NumAttrs() > 0 {
		buf = h.appendUnopenedGroups(buf, h.indentLevel)
		r.Attrs(func(a slog.Attr) bool {
			buf = h.appendAttr(buf, a, h.indentLevel+len(h.unopenedGroups))
			return true
		})
	}
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

// !-Handle

func (h *IndentHandler) appendAttr(buf []byte, a slog.Attr, indentLevel int) []byte {
	// Resolve the Attr's value before doing anything else.
	a.Value = a.Value.Resolve()
	// Ignore empty Attrs.
	if a.Equal(slog.Attr{}) {
		return buf
	}
	// Indent 4 spaces per level.
	buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
	switch a.Value.Kind() {
	case slog.KindString:
		// Quote string values, to make them easy to parse.
		buf = append(buf, a.Key...)
		buf = append(buf, ": "...)
		buf = strconv.AppendQuote(buf, a.Value.String())
		buf = append(buf, '\n')
	case slog.KindTime:
		// Write times in a standard way, without the monotonic time.
		buf = append(buf, a.Key...)
		buf = append(buf, ": "...)
		buf = a.Value.Time().AppendFormat(buf, time.RFC3339Nano)
		buf = append(buf, '\n')
	case slog.KindGroup:
		attrs := a.Value.Group()
		// Ignore empty groups.
		if len(attrs) == 0 {
			return buf
		}
		// If the key is non-empty, write it out and indent the rest of the attrs.
		// Otherwise, inline the attrs.
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s:\n", a.Key)
			indentLevel++
		}
		for _, ga := range attrs {
			buf = h.appendAttr(buf, ga, indentLevel)
		}

	default:
		buf = append(buf, a.Key...)
		buf = append(buf, ": "...)
		buf = append(buf, a.Value.String()...)
		buf = append(buf, '\n')
	}
	return buf
}

// !+pool
var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 1024)
		return &b
	},
}

func allocBuf() *[]byte {
	return bufPool.Get().(*[]byte)
}

func freeBuf(b *[]byte) {
	// To reduce peak allocation, return only smaller buffers to the pool.
	const maxBufferSize = 16 << 10
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufPool.Put(b)
	}
}
