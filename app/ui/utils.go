package ui

import (
	"fmt"
	"net"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func Values[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func headerContainer(headerText string, backCallback func()) *fyne.Container {
	return container.New(
		layout.NewHBoxLayout(),
		widget.NewButtonWithIcon("Go back", theme.NavigateBackIcon(), backCallback),
		widget.NewLabelWithStyle(headerText, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
	)
}

func ipValidator(s string) error {
	ip := net.ParseIP(s)
	if ip == nil {
		return fmt.Errorf("invalid IP address")
	}
	return nil
}

func portValidator(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if i < 1000 || i > 65535 {
		return fmt.Errorf("invalid port number")
	}
	return nil
}
