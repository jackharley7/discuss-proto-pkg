// Code generated by protoc-gen-go. DO NOT EDIT.
// source: education.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Education struct {
	Id                     int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId                 int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	School                 string   `protobuf:"bytes,3,opt,name=school,proto3" json:"school,omitempty"`
	Concentration          string   `protobuf:"bytes,4,opt,name=concentration,proto3" json:"concentration,omitempty"`
	SecondaryConcentration string   `protobuf:"bytes,5,opt,name=secondaryConcentration,proto3" json:"secondaryConcentration,omitempty"`
	DegreeType             string   `protobuf:"bytes,6,opt,name=degreeType,proto3" json:"degreeType,omitempty"`
	GraduationYear         int32    `protobuf:"varint,7,opt,name=graduationYear,proto3" json:"graduationYear,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Education) Reset()         { *m = Education{} }
func (m *Education) String() string { return proto.CompactTextString(m) }
func (*Education) ProtoMessage()    {}
func (*Education) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b007cc3e896a8b, []int{0}
}

func (m *Education) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Education.Unmarshal(m, b)
}
func (m *Education) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Education.Marshal(b, m, deterministic)
}
func (m *Education) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Education.Merge(m, src)
}
func (m *Education) XXX_Size() int {
	return xxx_messageInfo_Education.Size(m)
}
func (m *Education) XXX_DiscardUnknown() {
	xxx_messageInfo_Education.DiscardUnknown(m)
}

var xxx_messageInfo_Education proto.InternalMessageInfo

func (m *Education) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Education) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Education) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *Education) GetConcentration() string {
	if m != nil {
		return m.Concentration
	}
	return ""
}

func (m *Education) GetSecondaryConcentration() string {
	if m != nil {
		return m.SecondaryConcentration
	}
	return ""
}

func (m *Education) GetDegreeType() string {
	if m != nil {
		return m.DegreeType
	}
	return ""
}

func (m *Education) GetGraduationYear() int32 {
	if m != nil {
		return m.GraduationYear
	}
	return 0
}

type CreateEducationRequest struct {
	Education            *Education `protobuf:"bytes,1,opt,name=education,proto3" json:"education,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateEducationRequest) Reset()         { *m = CreateEducationRequest{} }
func (m *CreateEducationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEducationRequest) ProtoMessage()    {}
func (*CreateEducationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b007cc3e896a8b, []int{1}
}

func (m *CreateEducationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEducationRequest.Unmarshal(m, b)
}
func (m *CreateEducationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEducationRequest.Marshal(b, m, deterministic)
}
func (m *CreateEducationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEducationRequest.Merge(m, src)
}
func (m *CreateEducationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEducationRequest.Size(m)
}
func (m *CreateEducationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEducationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEducationRequest proto.InternalMessageInfo

func (m *CreateEducationRequest) GetEducation() *Education {
	if m != nil {
		return m.Education
	}
	return nil
}

type CreateEducationResponse struct {
	Education            *Education `protobuf:"bytes,1,opt,name=education,proto3" json:"education,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateEducationResponse) Reset()         { *m = CreateEducationResponse{} }
func (m *CreateEducationResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEducationResponse) ProtoMessage()    {}
func (*CreateEducationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b007cc3e896a8b, []int{2}
}

func (m *CreateEducationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEducationResponse.Unmarshal(m, b)
}
func (m *CreateEducationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEducationResponse.Marshal(b, m, deterministic)
}
func (m *CreateEducationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEducationResponse.Merge(m, src)
}
func (m *CreateEducationResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEducationResponse.Size(m)
}
func (m *CreateEducationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEducationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEducationResponse proto.InternalMessageInfo

func (m *CreateEducationResponse) GetEducation() *Education {
	if m != nil {
		return m.Education
	}
	return nil
}

type UpdateEducationRequest struct {
	Id                   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Education            *Education `protobuf:"bytes,2,opt,name=education,proto3" json:"education,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateEducationRequest) Reset()         { *m = UpdateEducationRequest{} }
func (m *UpdateEducationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEducationRequest) ProtoMessage()    {}
func (*UpdateEducationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b007cc3e896a8b, []int{3}
}

func (m *UpdateEducationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEducationRequest.Unmarshal(m, b)
}
func (m *UpdateEducationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEducationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEducationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEducationRequest.Merge(m, src)
}
func (m *UpdateEducationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEducationRequest.Size(m)
}
func (m *UpdateEducationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEducationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEducationRequest proto.InternalMessageInfo

func (m *UpdateEducationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateEducationRequest) GetEducation() *Education {
	if m != nil {
		return m.Education
	}
	return nil
}

type UpdateEducationResponse struct {
	Education            *Education `protobuf:"bytes,1,opt,name=education,proto3" json:"education,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateEducationResponse) Reset()         { *m = UpdateEducationResponse{} }
func (m *UpdateEducationResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateEducationResponse) ProtoMessage()    {}
func (*UpdateEducationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b007cc3e896a8b, []int{4}
}

func (m *UpdateEducationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEducationResponse.Unmarshal(m, b)
}
func (m *UpdateEducationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEducationResponse.Marshal(b, m, deterministic)
}
func (m *UpdateEducationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEducationResponse.Merge(m, src)
}
func (m *UpdateEducationResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateEducationResponse.Size(m)
}
func (m *UpdateEducationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEducationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEducationResponse proto.InternalMessageInfo

func (m *UpdateEducationResponse) GetEducation() *Education {
	if m != nil {
		return m.Education
	}
	return nil
}

type DeleteEducationRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEducationRequest) Reset()         { *m = DeleteEducationRequest{} }
func (m *DeleteEducationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEducationRequest) ProtoMessage()    {}
func (*DeleteEducationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b007cc3e896a8b, []int{5}
}

func (m *DeleteEducationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEducationRequest.Unmarshal(m, b)
}
func (m *DeleteEducationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEducationRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEducationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEducationRequest.Merge(m, src)
}
func (m *DeleteEducationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEducationRequest.Size(m)
}
func (m *DeleteEducationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEducationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEducationRequest proto.InternalMessageInfo

func (m *DeleteEducationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteEducationResponse struct {
	Done                 bool     `protobuf:"varint,1,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEducationResponse) Reset()         { *m = DeleteEducationResponse{} }
func (m *DeleteEducationResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEducationResponse) ProtoMessage()    {}
func (*DeleteEducationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64b007cc3e896a8b, []int{6}
}

func (m *DeleteEducationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEducationResponse.Unmarshal(m, b)
}
func (m *DeleteEducationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEducationResponse.Marshal(b, m, deterministic)
}
func (m *DeleteEducationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEducationResponse.Merge(m, src)
}
func (m *DeleteEducationResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEducationResponse.Size(m)
}
func (m *DeleteEducationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEducationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEducationResponse proto.InternalMessageInfo

func (m *DeleteEducationResponse) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func init() {
	proto.RegisterType((*Education)(nil), "proto.Education")
	proto.RegisterType((*CreateEducationRequest)(nil), "proto.CreateEducationRequest")
	proto.RegisterType((*CreateEducationResponse)(nil), "proto.CreateEducationResponse")
	proto.RegisterType((*UpdateEducationRequest)(nil), "proto.UpdateEducationRequest")
	proto.RegisterType((*UpdateEducationResponse)(nil), "proto.UpdateEducationResponse")
	proto.RegisterType((*DeleteEducationRequest)(nil), "proto.DeleteEducationRequest")
	proto.RegisterType((*DeleteEducationResponse)(nil), "proto.DeleteEducationResponse")
}

func init() { proto.RegisterFile("education.proto", fileDescriptor_64b007cc3e896a8b) }

var fileDescriptor_64b007cc3e896a8b = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb1, 0xdb, 0xb8, 0x64, 0x10, 0x4d, 0xb4, 0x07, 0xd7, 0x4a, 0x69, 0x09, 0x46, 0xaa,
	0xa2, 0x8a, 0xd8, 0x52, 0x90, 0x7a, 0xe0, 0xd8, 0x82, 0x44, 0xaf, 0x06, 0xa4, 0xc2, 0x6d, 0xeb,
	0x1d, 0xb9, 0x2b, 0x12, 0xaf, 0xd9, 0x5d, 0xb7, 0x54, 0x88, 0x0b, 0x6f, 0x80, 0x78, 0x34, 0x1e,
	0x00, 0x09, 0xf1, 0x10, 0x1c, 0x51, 0xd6, 0x61, 0x83, 0x1d, 0x87, 0x7f, 0xa7, 0x5d, 0xef, 0x7c,
	0x33, 0xbf, 0xf1, 0x37, 0xf6, 0x42, 0x0f, 0x59, 0x99, 0x52, 0xcd, 0x45, 0x1e, 0x15, 0x52, 0x68,
	0x41, 0x3a, 0x66, 0x19, 0xdc, 0xc9, 0x84, 0xc8, 0xa6, 0x18, 0xd3, 0x82, 0xc7, 0x34, 0xcf, 0x85,
	0x36, 0x1a, 0x55, 0x89, 0x06, 0x47, 0x19, 0xd7, 0x17, 0xe5, 0x79, 0x94, 0x8a, 0x59, 0x3c, 0xbb,
	0xe2, 0xfa, 0xb5, 0xb8, 0x8a, 0x33, 0x31, 0x36, 0xc1, 0xf1, 0x25, 0x9d, 0x72, 0x46, 0xb5, 0x90,
	0x2a, 0xb6, 0xdb, 0x2a, 0x2f, 0xfc, 0xe8, 0x42, 0xf7, 0xc9, 0x4f, 0x20, 0xd9, 0x06, 0x97, 0xb3,
	0xc0, 0x19, 0x3a, 0xa3, 0x8d, 0xc4, 0xe5, 0x8c, 0xf8, 0xe0, 0x95, 0x0a, 0xe5, 0x29, 0x0b, 0x5c,
	0x73, 0xb6, 0x78, 0x22, 0xfb, 0xe0, 0xa9, 0xf4, 0x42, 0x88, 0x69, 0xb0, 0x31, 0x74, 0x46, 0xdd,
	0x63, 0xef, 0xeb, 0x97, 0xbb, 0xee, 0x99, 0x93, 0x2c, 0x4e, 0xc9, 0x03, 0xb8, 0x9d, 0x8a, 0x3c,
	0xc5, 0x5c, 0x4b, 0x53, 0x38, 0xd8, 0xac, 0xc9, 0xea, 0x41, 0x72, 0x04, 0xbe, 0xc2, 0x54, 0xe4,
	0x8c, 0xca, 0xeb, 0x93, 0x5a, 0x5a, 0x67, 0x9e, 0x96, 0xac, 0x89, 0x92, 0x03, 0x00, 0x86, 0x99,
	0x44, 0x7c, 0x7e, 0x5d, 0x60, 0xe0, 0xd5, 0x10, 0xbf, 0x44, 0x48, 0x04, 0xdb, 0x99, 0xa4, 0xac,
	0x34, 0x59, 0x2f, 0x91, 0xca, 0x60, 0x6b, 0xe8, 0x8c, 0x3a, 0x95, 0xb6, 0x7f, 0x23, 0x69, 0x44,
	0xc3, 0xa7, 0xe0, 0x9f, 0x48, 0xa4, 0x1a, 0xad, 0x31, 0x09, 0xbe, 0x29, 0x51, 0x69, 0x12, 0x41,
	0xd7, 0x4e, 0xc7, 0xd8, 0x74, 0x6b, 0xd2, 0xaf, 0x8c, 0x8c, 0x96, 0xda, 0xa5, 0x24, 0x3c, 0x85,
	0x9d, 0x95, 0x4a, 0xaa, 0x10, 0xb9, 0xc2, 0x7f, 0x2e, 0x75, 0x06, 0xfe, 0x8b, 0x82, 0xb5, 0x35,
	0xd5, 0x1c, 0x5a, 0xad, 0xb2, 0xfb, 0x57, 0x4d, 0xae, 0x54, 0xfe, 0xcf, 0x26, 0x47, 0xe0, 0x3f,
	0xc6, 0x29, 0xfe, 0xb9, 0xc9, 0x70, 0x0c, 0x3b, 0x2b, 0xca, 0x05, 0x94, 0xc0, 0x26, 0x13, 0x39,
	0x1a, 0xf1, 0xcd, 0xc4, 0xec, 0x27, 0xdf, 0x5d, 0xe8, 0x5b, 0xe5, 0x33, 0x94, 0x97, 0x3c, 0x45,
	0xa2, 0xa0, 0xd7, 0x70, 0x97, 0xec, 0x2d, 0xba, 0x6b, 0x9f, 0xdf, 0x60, 0x7f, 0x5d, 0xb8, 0x42,
	0x87, 0xf7, 0x3e, 0x7c, 0xfe, 0xf6, 0xc9, 0xdd, 0x0d, 0x7d, 0xf3, 0x97, 0xcd, 0x3f, 0xf6, 0x78,
	0x86, 0xb1, 0x7d, 0xbf, 0x47, 0xce, 0x21, 0x79, 0x0b, 0xbd, 0x86, 0x5b, 0x16, 0xda, 0x3e, 0x1f,
	0x0b, 0x5d, 0x63, 0x72, 0x78, 0x60, 0xa0, 0xc3, 0xc9, 0x6e, 0x3b, 0x34, 0x7e, 0xc7, 0xd9, 0xfb,
	0x39, 0xb9, 0x84, 0x5e, 0xc3, 0x32, 0x4b, 0x6e, 0x37, 0xdd, 0x92, 0xd7, 0x38, 0x1d, 0xde, 0x37,
	0xe4, 0xbd, 0xc3, 0xdf, 0x91, 0x8f, 0xb7, 0x5e, 0x55, 0x17, 0xd0, 0xb9, 0x67, 0x96, 0x87, 0x3f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xa9, 0xee, 0x03, 0xa1, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EducationServiceClient is the client API for EducationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EducationServiceClient interface {
	CreateEducation(ctx context.Context, in *CreateEducationRequest, opts ...grpc.CallOption) (*CreateEducationResponse, error)
	UpdateEducation(ctx context.Context, in *UpdateEducationRequest, opts ...grpc.CallOption) (*UpdateEducationResponse, error)
	DeleteEducation(ctx context.Context, in *DeleteEducationRequest, opts ...grpc.CallOption) (*DeleteEducationResponse, error)
}

type educationServiceClient struct {
	cc *grpc.ClientConn
}

func NewEducationServiceClient(cc *grpc.ClientConn) EducationServiceClient {
	return &educationServiceClient{cc}
}

func (c *educationServiceClient) CreateEducation(ctx context.Context, in *CreateEducationRequest, opts ...grpc.CallOption) (*CreateEducationResponse, error) {
	out := new(CreateEducationResponse)
	err := c.cc.Invoke(ctx, "/proto.EducationService/CreateEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *educationServiceClient) UpdateEducation(ctx context.Context, in *UpdateEducationRequest, opts ...grpc.CallOption) (*UpdateEducationResponse, error) {
	out := new(UpdateEducationResponse)
	err := c.cc.Invoke(ctx, "/proto.EducationService/UpdateEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *educationServiceClient) DeleteEducation(ctx context.Context, in *DeleteEducationRequest, opts ...grpc.CallOption) (*DeleteEducationResponse, error) {
	out := new(DeleteEducationResponse)
	err := c.cc.Invoke(ctx, "/proto.EducationService/DeleteEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EducationServiceServer is the server API for EducationService service.
type EducationServiceServer interface {
	CreateEducation(context.Context, *CreateEducationRequest) (*CreateEducationResponse, error)
	UpdateEducation(context.Context, *UpdateEducationRequest) (*UpdateEducationResponse, error)
	DeleteEducation(context.Context, *DeleteEducationRequest) (*DeleteEducationResponse, error)
}

func RegisterEducationServiceServer(s *grpc.Server, srv EducationServiceServer) {
	s.RegisterService(&_EducationService_serviceDesc, srv)
}

func _EducationService_CreateEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEducationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EducationServiceServer).CreateEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EducationService/CreateEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EducationServiceServer).CreateEducation(ctx, req.(*CreateEducationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EducationService_UpdateEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEducationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EducationServiceServer).UpdateEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EducationService/UpdateEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EducationServiceServer).UpdateEducation(ctx, req.(*UpdateEducationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EducationService_DeleteEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEducationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EducationServiceServer).DeleteEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EducationService/DeleteEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EducationServiceServer).DeleteEducation(ctx, req.(*DeleteEducationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EducationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EducationService",
	HandlerType: (*EducationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEducation",
			Handler:    _EducationService_CreateEducation_Handler,
		},
		{
			MethodName: "UpdateEducation",
			Handler:    _EducationService_UpdateEducation_Handler,
		},
		{
			MethodName: "DeleteEducation",
			Handler:    _EducationService_DeleteEducation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "education.proto",
}