// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: bookService.proto

package protoFiles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookMgmtServiceClient is the client API for BookMgmtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookMgmtServiceClient interface {
	//Create Book
	CreateBook(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	//Get All Books
	GetAllBooks(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (BookMgmtService_GetAllBooksClient, error)
	//Search Book
	SearchBook(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (BookMgmtService_SearchBookClient, error)
	//Update Book
	UpdateBook(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	//Delete Book
	DeleteBook(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type bookMgmtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookMgmtServiceClient(cc grpc.ClientConnInterface) BookMgmtServiceClient {
	return &bookMgmtServiceClient{cc}
}

func (c *bookMgmtServiceClient) CreateBook(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/protoFiles.BookMgmtService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookMgmtServiceClient) GetAllBooks(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (BookMgmtService_GetAllBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookMgmtService_ServiceDesc.Streams[0], "/protoFiles.BookMgmtService/GetAllBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookMgmtServiceGetAllBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookMgmtService_GetAllBooksClient interface {
	Recv() (*GetAllResponse, error)
	grpc.ClientStream
}

type bookMgmtServiceGetAllBooksClient struct {
	grpc.ClientStream
}

func (x *bookMgmtServiceGetAllBooksClient) Recv() (*GetAllResponse, error) {
	m := new(GetAllResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookMgmtServiceClient) SearchBook(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (BookMgmtService_SearchBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookMgmtService_ServiceDesc.Streams[1], "/protoFiles.BookMgmtService/SearchBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookMgmtServiceSearchBookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookMgmtService_SearchBookClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type bookMgmtServiceSearchBookClient struct {
	grpc.ClientStream
}

func (x *bookMgmtServiceSearchBookClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookMgmtServiceClient) UpdateBook(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/protoFiles.BookMgmtService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookMgmtServiceClient) DeleteBook(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/protoFiles.BookMgmtService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookMgmtServiceServer is the server API for BookMgmtService service.
// All implementations must embed UnimplementedBookMgmtServiceServer
// for forward compatibility
type BookMgmtServiceServer interface {
	//Create Book
	CreateBook(context.Context, *CreateRequest) (*CreateResponse, error)
	//Get All Books
	GetAllBooks(*GetAllRequest, BookMgmtService_GetAllBooksServer) error
	//Search Book
	SearchBook(*SearchRequest, BookMgmtService_SearchBookServer) error
	//Update Book
	UpdateBook(context.Context, *UpdateRequest) (*UpdateResponse, error)
	//Delete Book
	DeleteBook(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedBookMgmtServiceServer()
}

// UnimplementedBookMgmtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookMgmtServiceServer struct {
}

func (UnimplementedBookMgmtServiceServer) CreateBook(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookMgmtServiceServer) GetAllBooks(*GetAllRequest, BookMgmtService_GetAllBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllBooks not implemented")
}
func (UnimplementedBookMgmtServiceServer) SearchBook(*SearchRequest, BookMgmtService_SearchBookServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchBook not implemented")
}
func (UnimplementedBookMgmtServiceServer) UpdateBook(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookMgmtServiceServer) DeleteBook(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookMgmtServiceServer) mustEmbedUnimplementedBookMgmtServiceServer() {}

// UnsafeBookMgmtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookMgmtServiceServer will
// result in compilation errors.
type UnsafeBookMgmtServiceServer interface {
	mustEmbedUnimplementedBookMgmtServiceServer()
}

func RegisterBookMgmtServiceServer(s grpc.ServiceRegistrar, srv BookMgmtServiceServer) {
	s.RegisterService(&BookMgmtService_ServiceDesc, srv)
}

func _BookMgmtService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookMgmtServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoFiles.BookMgmtService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookMgmtServiceServer).CreateBook(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookMgmtService_GetAllBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookMgmtServiceServer).GetAllBooks(m, &bookMgmtServiceGetAllBooksServer{stream})
}

type BookMgmtService_GetAllBooksServer interface {
	Send(*GetAllResponse) error
	grpc.ServerStream
}

type bookMgmtServiceGetAllBooksServer struct {
	grpc.ServerStream
}

func (x *bookMgmtServiceGetAllBooksServer) Send(m *GetAllResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BookMgmtService_SearchBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookMgmtServiceServer).SearchBook(m, &bookMgmtServiceSearchBookServer{stream})
}

type BookMgmtService_SearchBookServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type bookMgmtServiceSearchBookServer struct {
	grpc.ServerStream
}

func (x *bookMgmtServiceSearchBookServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BookMgmtService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookMgmtServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoFiles.BookMgmtService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookMgmtServiceServer).UpdateBook(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookMgmtService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookMgmtServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoFiles.BookMgmtService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookMgmtServiceServer).DeleteBook(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookMgmtService_ServiceDesc is the grpc.ServiceDesc for BookMgmtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookMgmtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoFiles.BookMgmtService",
	HandlerType: (*BookMgmtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookMgmtService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookMgmtService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookMgmtService_DeleteBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllBooks",
			Handler:       _BookMgmtService_GetAllBooks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchBook",
			Handler:       _BookMgmtService_SearchBook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bookService.proto",
}
