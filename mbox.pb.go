// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mbox.proto

package mbox

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	From    string   `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To      []string `protobuf:"bytes,2,rep,name=to" json:"to,omitempty"`
	Subject string   `protobuf:"bytes,3,opt,name=subject" json:"subject,omitempty"`
	// Types that are valid to be assigned to Body:
	//	*Message_Text
	//	*Message_Html
	Body                 isMessage_Body `protobuf_oneof:"body"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_mbox_b89e99d1b1d7f985, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Body interface {
	isMessage_Body()
}

type Message_Text struct {
	Text string `protobuf:"bytes,4,opt,name=text,oneof"`
}
type Message_Html struct {
	Html string `protobuf:"bytes,5,opt,name=html,oneof"`
}

func (*Message_Text) isMessage_Body() {}
func (*Message_Html) isMessage_Body() {}

func (m *Message) GetBody() isMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Message) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Message) GetText() string {
	if x, ok := m.GetBody().(*Message_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Message) GetHtml() string {
	if x, ok := m.GetBody().(*Message_Html); ok {
		return x.Html
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Text)(nil),
		(*Message_Html)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// body
	switch x := m.Body.(type) {
	case *Message_Text:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case *Message_Html:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Html)
	case nil:
	default:
		return fmt.Errorf("Message.Body has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 4: // body.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Body = &Message_Text{x}
		return true, err
	case 5: // body.html
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Body = &Message_Html{x}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// body
	switch x := m.Body.(type) {
	case *Message_Text:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case *Message_Html:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Html)))
		n += len(x.Html)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Result struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_mbox_b89e99d1b1d7f985, []int{1}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

type Folder struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Folder) Reset()         { *m = Folder{} }
func (m *Folder) String() string { return proto.CompactTextString(m) }
func (*Folder) ProtoMessage()    {}
func (*Folder) Descriptor() ([]byte, []int) {
	return fileDescriptor_mbox_b89e99d1b1d7f985, []int{2}
}
func (m *Folder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Folder.Unmarshal(m, b)
}
func (m *Folder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Folder.Marshal(b, m, deterministic)
}
func (dst *Folder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Folder.Merge(dst, src)
}
func (m *Folder) XXX_Size() int {
	return xxx_messageInfo_Folder.Size(m)
}
func (m *Folder) XXX_DiscardUnknown() {
	xxx_messageInfo_Folder.DiscardUnknown(m)
}

var xxx_messageInfo_Folder proto.InternalMessageInfo

func (m *Folder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "mbox.Message")
	proto.RegisterType((*Result)(nil), "mbox.Result")
	proto.RegisterType((*Folder)(nil), "mbox.Folder")
}

func init() { proto.RegisterFile("mbox.proto", fileDescriptor_mbox_b89e99d1b1d7f985) }

var fileDescriptor_mbox_b89e99d1b1d7f985 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x49, 0x6b, 0x92, 0x72, 0xa2, 0x08, 0x1d, 0x0c, 0x56, 0xd4, 0xa1, 0xca, 0x84, 0x18,
	0x1a, 0x04, 0x0b, 0xea, 0xc8, 0x80, 0x58, 0xba, 0x84, 0x27, 0x70, 0x9a, 0x23, 0x04, 0x25, 0xbe,
	0x2a, 0x76, 0x51, 0x11, 0x4c, 0xbc, 0x02, 0x8f, 0xc6, 0x2b, 0xf0, 0x20, 0xd4, 0x4e, 0x82, 0xc4,
	0xf6, 0xff, 0x9f, 0xff, 0x3b, 0xff, 0x36, 0x40, 0x93, 0xf3, 0x6e, 0xb1, 0x69, 0xd9, 0x32, 0x0a,
	0xa7, 0xe3, 0x59, 0xc9, 0x5c, 0xd6, 0x94, 0xaa, 0x4d, 0x95, 0x2a, 0xad, 0xd9, 0x2a, 0x5b, 0xb1,
	0x36, 0x5d, 0x26, 0x79, 0x87, 0x68, 0x45, 0xc6, 0xa8, 0x92, 0x10, 0x41, 0x3c, 0xb5, 0xdc, 0xc8,
	0x60, 0x1e, 0x5c, 0x1c, 0x65, 0x5e, 0xe3, 0x09, 0x8c, 0x2c, 0xcb, 0xd1, 0x7c, 0xbc, 0x27, 0x7b,
	0x85, 0x12, 0x22, 0xb3, 0xcd, 0x5f, 0x68, 0x6d, 0xe5, 0xd8, 0xc7, 0x06, 0x8b, 0xe7, 0x20, 0x2c,
	0xed, 0xac, 0x14, 0x0e, 0x3f, 0x1c, 0x64, 0xde, 0x39, 0xfa, 0x6c, 0x9b, 0x5a, 0x1e, 0x0e, 0xd4,
	0xb9, 0xbb, 0x10, 0x44, 0xce, 0xc5, 0x5b, 0x32, 0x81, 0x30, 0x23, 0xb3, 0xad, 0x6d, 0x32, 0x83,
	0xf0, 0x9e, 0xeb, 0x82, 0x5a, 0xd7, 0x42, 0xab, 0x86, 0x86, 0x16, 0x4e, 0x5f, 0x7f, 0x80, 0x58,
	0xa9, 0xaa, 0xc6, 0x5b, 0x10, 0x8f, 0xa4, 0x0b, 0x9c, 0x2e, 0xfc, 0x2b, 0xfb, 0xe2, 0xf1, 0x71,
	0x67, 0xfb, 0x55, 0x67, 0x9f, 0xdf, 0x3f, 0x5f, 0xa3, 0x69, 0x32, 0x49, 0x9b, 0xee, 0x7c, 0x19,
	0x5c, 0xe2, 0x12, 0xa2, 0x8c, 0xd6, 0x54, 0xbd, 0x12, 0xf6, 0xe9, 0xee, 0xba, 0xf8, 0xff, 0xaa,
	0xe4, 0xd4, 0x0f, 0x03, 0xfe, 0x0d, 0x5f, 0x05, 0x79, 0xe8, 0x7f, 0xea, 0xe6, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0x68, 0xe5, 0xb4, 0x5b, 0x01, 0x00, 0x00,
}
