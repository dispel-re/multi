// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: multi/v1/character.proto

package multiv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/dispel-re/dispel-multi/gen/multi/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// CharacterServiceName is the fully-qualified name of the CharacterService service.
	CharacterServiceName = "multi.v1.CharacterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CharacterServiceGetCharacterProcedure is the fully-qualified name of the CharacterService's
	// GetCharacter RPC.
	CharacterServiceGetCharacterProcedure = "/multi.v1.CharacterService/GetCharacter"
	// CharacterServiceListCharactersProcedure is the fully-qualified name of the CharacterService's
	// ListCharacters RPC.
	CharacterServiceListCharactersProcedure = "/multi.v1.CharacterService/ListCharacters"
	// CharacterServiceCreateCharacterProcedure is the fully-qualified name of the CharacterService's
	// CreateCharacter RPC.
	CharacterServiceCreateCharacterProcedure = "/multi.v1.CharacterService/CreateCharacter"
)

// CharacterServiceClient is a client for the multi.v1.CharacterService service.
type CharacterServiceClient interface {
	GetCharacter(context.Context, *connect.Request[v1.GetCharacterRequest]) (*connect.Response[v1.GetCharacterResponse], error)
	ListCharacters(context.Context, *connect.Request[v1.ListCharactersRequest]) (*connect.Response[v1.ListCharactersResponse], error)
	CreateCharacter(context.Context, *connect.Request[v1.CreateCharacterRequest]) (*connect.Response[v1.CreateCharacterResponse], error)
}

// NewCharacterServiceClient constructs a client for the multi.v1.CharacterService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCharacterServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CharacterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &characterServiceClient{
		getCharacter: connect.NewClient[v1.GetCharacterRequest, v1.GetCharacterResponse](
			httpClient,
			baseURL+CharacterServiceGetCharacterProcedure,
			opts...,
		),
		listCharacters: connect.NewClient[v1.ListCharactersRequest, v1.ListCharactersResponse](
			httpClient,
			baseURL+CharacterServiceListCharactersProcedure,
			opts...,
		),
		createCharacter: connect.NewClient[v1.CreateCharacterRequest, v1.CreateCharacterResponse](
			httpClient,
			baseURL+CharacterServiceCreateCharacterProcedure,
			opts...,
		),
	}
}

// characterServiceClient implements CharacterServiceClient.
type characterServiceClient struct {
	getCharacter    *connect.Client[v1.GetCharacterRequest, v1.GetCharacterResponse]
	listCharacters  *connect.Client[v1.ListCharactersRequest, v1.ListCharactersResponse]
	createCharacter *connect.Client[v1.CreateCharacterRequest, v1.CreateCharacterResponse]
}

// GetCharacter calls multi.v1.CharacterService.GetCharacter.
func (c *characterServiceClient) GetCharacter(ctx context.Context, req *connect.Request[v1.GetCharacterRequest]) (*connect.Response[v1.GetCharacterResponse], error) {
	return c.getCharacter.CallUnary(ctx, req)
}

// ListCharacters calls multi.v1.CharacterService.ListCharacters.
func (c *characterServiceClient) ListCharacters(ctx context.Context, req *connect.Request[v1.ListCharactersRequest]) (*connect.Response[v1.ListCharactersResponse], error) {
	return c.listCharacters.CallUnary(ctx, req)
}

// CreateCharacter calls multi.v1.CharacterService.CreateCharacter.
func (c *characterServiceClient) CreateCharacter(ctx context.Context, req *connect.Request[v1.CreateCharacterRequest]) (*connect.Response[v1.CreateCharacterResponse], error) {
	return c.createCharacter.CallUnary(ctx, req)
}

// CharacterServiceHandler is an implementation of the multi.v1.CharacterService service.
type CharacterServiceHandler interface {
	GetCharacter(context.Context, *connect.Request[v1.GetCharacterRequest]) (*connect.Response[v1.GetCharacterResponse], error)
	ListCharacters(context.Context, *connect.Request[v1.ListCharactersRequest]) (*connect.Response[v1.ListCharactersResponse], error)
	CreateCharacter(context.Context, *connect.Request[v1.CreateCharacterRequest]) (*connect.Response[v1.CreateCharacterResponse], error)
}

// NewCharacterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCharacterServiceHandler(svc CharacterServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	characterServiceGetCharacterHandler := connect.NewUnaryHandler(
		CharacterServiceGetCharacterProcedure,
		svc.GetCharacter,
		opts...,
	)
	characterServiceListCharactersHandler := connect.NewUnaryHandler(
		CharacterServiceListCharactersProcedure,
		svc.ListCharacters,
		opts...,
	)
	characterServiceCreateCharacterHandler := connect.NewUnaryHandler(
		CharacterServiceCreateCharacterProcedure,
		svc.CreateCharacter,
		opts...,
	)
	return "/multi.v1.CharacterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CharacterServiceGetCharacterProcedure:
			characterServiceGetCharacterHandler.ServeHTTP(w, r)
		case CharacterServiceListCharactersProcedure:
			characterServiceListCharactersHandler.ServeHTTP(w, r)
		case CharacterServiceCreateCharacterProcedure:
			characterServiceCreateCharacterHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCharacterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCharacterServiceHandler struct{}

func (UnimplementedCharacterServiceHandler) GetCharacter(context.Context, *connect.Request[v1.GetCharacterRequest]) (*connect.Response[v1.GetCharacterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("multi.v1.CharacterService.GetCharacter is not implemented"))
}

func (UnimplementedCharacterServiceHandler) ListCharacters(context.Context, *connect.Request[v1.ListCharactersRequest]) (*connect.Response[v1.ListCharactersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("multi.v1.CharacterService.ListCharacters is not implemented"))
}

func (UnimplementedCharacterServiceHandler) CreateCharacter(context.Context, *connect.Request[v1.CreateCharacterRequest]) (*connect.Response[v1.CreateCharacterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("multi.v1.CharacterService.CreateCharacter is not implemented"))
}
