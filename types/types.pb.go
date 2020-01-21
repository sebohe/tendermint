// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/types.proto

package proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BlockIdFlag indicates which BlcokID the signature is for
type BlockIDFlag int32

const (
	BlockIDFlag_BLOCK_ID_FLAG_UNSPECIFIED BlockIDFlag = 0
	BlockIDFlag_BLOCK_ID_FLAG_ABSENT      BlockIDFlag = 1
	BlockIDFlag_BLOCK_ID_FLAG_COMMIT      BlockIDFlag = 2
	BlockIDFlag_BLOCK_ID_FLAG_NIL         BlockIDFlag = 3
)

var BlockIDFlag_name = map[int32]string{
	0: "BLOCK_ID_FLAG_UNSPECIFIED",
	1: "BLOCK_ID_FLAG_ABSENT",
	2: "BLOCK_ID_FLAG_COMMIT",
	3: "BLOCK_ID_FLAG_NIL",
}

var BlockIDFlag_value = map[string]int32{
	"BLOCK_ID_FLAG_UNSPECIFIED": 0,
	"BLOCK_ID_FLAG_ABSENT":      1,
	"BLOCK_ID_FLAG_COMMIT":      2,
	"BLOCK_ID_FLAG_NIL":         3,
}

func (x BlockIDFlag) String() string {
	return proto.EnumName(BlockIDFlag_name, int32(x))
}

func (BlockIDFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{0}
}

// SignedMsgType is a type of signed message in the consensus.
type SignedMsgType int32

const (
	SignedMsgType_SIGNED_MSG_TYPE_UNSPECIFIED    SignedMsgType = 0
	SignedMsgType_SIGNED_MSG_TYPE_PREVOTE_TYPE   SignedMsgType = 1
	SignedMsgType_SIGNED_MSG_TYPE_PRECOMMIT_TYPE SignedMsgType = 2
	SignedMsgType_SIGNED_MSG_TYPE_PRPOSAL_TYPE   SignedMsgType = 3
)

var SignedMsgType_name = map[int32]string{
	0: "SIGNED_MSG_TYPE_UNSPECIFIED",
	1: "SIGNED_MSG_TYPE_PREVOTE_TYPE",
	2: "SIGNED_MSG_TYPE_PRECOMMIT_TYPE",
	3: "SIGNED_MSG_TYPE_PRPOSAL_TYPE",
}

var SignedMsgType_value = map[string]int32{
	"SIGNED_MSG_TYPE_UNSPECIFIED":    0,
	"SIGNED_MSG_TYPE_PREVOTE_TYPE":   1,
	"SIGNED_MSG_TYPE_PRECOMMIT_TYPE": 2,
	"SIGNED_MSG_TYPE_PRPOSAL_TYPE":   3,
}

func (x SignedMsgType) String() string {
	return proto.EnumName(SignedMsgType_name, int32(x))
}

func (SignedMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{1}
}

func init() {
	proto.RegisterEnum("types.proto.BlockIDFlag", BlockIDFlag_name, BlockIDFlag_value)
	proto.RegisterEnum("types.proto.SignedMsgType", SignedMsgType_name, SignedMsgType_value)
}

func init() { proto.RegisterFile("types/types.proto", fileDescriptor_2c0f90c600ad7e2e) }

var fileDescriptor_2c0f90c600ad7e2e = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc, 0x48, 0x1c, 0xad, 0x72,
	0x2e, 0x6e, 0xa7, 0x9c, 0xfc, 0xe4, 0x6c, 0x4f, 0x17, 0xb7, 0x9c, 0xc4, 0x74, 0x21, 0x59, 0x2e,
	0x49, 0x27, 0x1f, 0x7f, 0x67, 0xef, 0x78, 0x4f, 0x97, 0x78, 0x37, 0x1f, 0x47, 0xf7, 0xf8, 0x50,
	0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x09, 0x2e, 0x11,
	0x54, 0x69, 0x47, 0xa7, 0x60, 0x57, 0xbf, 0x10, 0x01, 0x46, 0x4c, 0x19, 0x67, 0x7f, 0x5f, 0x5f,
	0xcf, 0x10, 0x01, 0x26, 0x21, 0x51, 0x2e, 0x41, 0x54, 0x19, 0x3f, 0x4f, 0x1f, 0x01, 0x66, 0xad,
	0x19, 0x8c, 0x5c, 0xbc, 0xc1, 0x99, 0xe9, 0x79, 0xa9, 0x29, 0xbe, 0xc5, 0xe9, 0x21, 0x95, 0x05,
	0xa9, 0x42, 0xf2, 0x5c, 0xd2, 0xc1, 0x9e, 0xee, 0x7e, 0xae, 0x2e, 0xf1, 0xbe, 0xc1, 0xee, 0xf1,
	0x21, 0x91, 0x01, 0xae, 0x68, 0xb6, 0x2b, 0x70, 0xc9, 0xa0, 0x2b, 0x08, 0x08, 0x72, 0x0d, 0xf3,
	0x0f, 0x71, 0x05, 0x73, 0x04, 0x18, 0x85, 0x94, 0xb8, 0xe4, 0xb0, 0xa8, 0x80, 0x38, 0x05, 0xa2,
	0x86, 0x09, 0xbb, 0x29, 0x01, 0xfe, 0xc1, 0x8e, 0x3e, 0x10, 0x15, 0xcc, 0x4e, 0x7a, 0x51, 0x3a,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x25, 0xa9, 0x79, 0x29, 0xa9,
	0x45, 0xb9, 0x99, 0x79, 0x25, 0x28, 0x4c, 0x70, 0xb0, 0x82, 0xc3, 0x30, 0x89, 0x0d, 0x4c, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xa1, 0x79, 0x43, 0x6c, 0x01, 0x00, 0x00,
}
