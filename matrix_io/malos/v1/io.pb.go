// Code generated by protoc-gen-go. DO NOT EDIT.
// source: matrix_io/malos/v1/io.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EnumMatrixDeviceType int32

const (
	// Undetected device.
	EnumMatrixDeviceType_UNDETECTED EnumMatrixDeviceType = 0
	// Detected MATRIX Creator.
	EnumMatrixDeviceType_CREATOR EnumMatrixDeviceType = 1
	// Detected MATRIX Creator.
	EnumMatrixDeviceType_VOICE EnumMatrixDeviceType = 2
)

var EnumMatrixDeviceType_name = map[int32]string{
	0: "UNDETECTED",
	1: "CREATOR",
	2: "VOICE",
}

var EnumMatrixDeviceType_value = map[string]int32{
	"UNDETECTED": 0,
	"CREATOR":    1,
	"VOICE":      2,
}

func (x EnumMatrixDeviceType) String() string {
	return proto.EnumName(EnumMatrixDeviceType_name, int32(x))
}

func (EnumMatrixDeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{0}
}

// GPIO mode input/output
type GpioParams_EnumMode int32

const (
	GpioParams_INPUT  GpioParams_EnumMode = 0
	GpioParams_OUTPUT GpioParams_EnumMode = 1
)

var GpioParams_EnumMode_name = map[int32]string{
	0: "INPUT",
	1: "OUTPUT",
}

var GpioParams_EnumMode_value = map[string]int32{
	"INPUT":  0,
	"OUTPUT": 1,
}

func (x GpioParams_EnumMode) String() string {
	return proto.EnumName(GpioParams_EnumMode_name, int32(x))
}

func (GpioParams_EnumMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{3, 0}
}

// Mic channel
type WakeWordParams_MicChannel int32

const (
	WakeWordParams_channel0 WakeWordParams_MicChannel = 0
	WakeWordParams_channel1 WakeWordParams_MicChannel = 1
	WakeWordParams_channel2 WakeWordParams_MicChannel = 2
	WakeWordParams_channel3 WakeWordParams_MicChannel = 3
	WakeWordParams_channel4 WakeWordParams_MicChannel = 4
	WakeWordParams_channel5 WakeWordParams_MicChannel = 5
	WakeWordParams_channel6 WakeWordParams_MicChannel = 6
	WakeWordParams_channel7 WakeWordParams_MicChannel = 7
	WakeWordParams_channel8 WakeWordParams_MicChannel = 8
)

var WakeWordParams_MicChannel_name = map[int32]string{
	0: "channel0",
	1: "channel1",
	2: "channel2",
	3: "channel3",
	4: "channel4",
	5: "channel5",
	6: "channel6",
	7: "channel7",
	8: "channel8",
}

var WakeWordParams_MicChannel_value = map[string]int32{
	"channel0": 0,
	"channel1": 1,
	"channel2": 2,
	"channel3": 3,
	"channel4": 4,
	"channel5": 5,
	"channel6": 6,
	"channel7": 7,
	"channel8": 8,
}

func (x WakeWordParams_MicChannel) String() string {
	return proto.EnumName(WakeWordParams_MicChannel_name, int32(x))
}

func (WakeWordParams_MicChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{5, 0}
}

// Value for a led. Values range from 0 to 255.
type LedValue struct {
	Red                  uint32   `protobuf:"varint,1,opt,name=red,proto3" json:"red,omitempty"`
	Green                uint32   `protobuf:"varint,2,opt,name=green,proto3" json:"green,omitempty"`
	Blue                 uint32   `protobuf:"varint,3,opt,name=blue,proto3" json:"blue,omitempty"`
	White                uint32   `protobuf:"varint,4,opt,name=white,proto3" json:"white,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LedValue) Reset()         { *m = LedValue{} }
func (m *LedValue) String() string { return proto.CompactTextString(m) }
func (*LedValue) ProtoMessage()    {}
func (*LedValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{0}
}

func (m *LedValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedValue.Unmarshal(m, b)
}
func (m *LedValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedValue.Marshal(b, m, deterministic)
}
func (m *LedValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedValue.Merge(m, src)
}
func (m *LedValue) XXX_Size() int {
	return xxx_messageInfo_LedValue.Size(m)
}
func (m *LedValue) XXX_DiscardUnknown() {
	xxx_messageInfo_LedValue.DiscardUnknown(m)
}

var xxx_messageInfo_LedValue proto.InternalMessageInfo

func (m *LedValue) GetRed() uint32 {
	if m != nil {
		return m.Red
	}
	return 0
}

func (m *LedValue) GetGreen() uint32 {
	if m != nil {
		return m.Green
	}
	return 0
}

func (m *LedValue) GetBlue() uint32 {
	if m != nil {
		return m.Blue
	}
	return 0
}

func (m *LedValue) GetWhite() uint32 {
	if m != nil {
		return m.White
	}
	return 0
}

// The led array.
type EverloopImage struct {
	Led []*LedValue `protobuf:"bytes,1,rep,name=led,proto3" json:"led,omitempty"`
	// Number of leds in the Everloop
	EverloopLength       int32    `protobuf:"varint,2,opt,name=everloop_length,json=everloopLength,proto3" json:"everloop_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EverloopImage) Reset()         { *m = EverloopImage{} }
func (m *EverloopImage) String() string { return proto.CompactTextString(m) }
func (*EverloopImage) ProtoMessage()    {}
func (*EverloopImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{1}
}

func (m *EverloopImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EverloopImage.Unmarshal(m, b)
}
func (m *EverloopImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EverloopImage.Marshal(b, m, deterministic)
}
func (m *EverloopImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EverloopImage.Merge(m, src)
}
func (m *EverloopImage) XXX_Size() int {
	return xxx_messageInfo_EverloopImage.Size(m)
}
func (m *EverloopImage) XXX_DiscardUnknown() {
	xxx_messageInfo_EverloopImage.DiscardUnknown(m)
}

var xxx_messageInfo_EverloopImage proto.InternalMessageInfo

func (m *EverloopImage) GetLed() []*LedValue {
	if m != nil {
		return m.Led
	}
	return nil
}

func (m *EverloopImage) GetEverloopLength() int32 {
	if m != nil {
		return m.EverloopLength
	}
	return 0
}

// Servo handler params
type ServoParams struct {
	// GPIO to config
	Pin uint32 `protobuf:"varint,1,opt,name=pin,proto3" json:"pin,omitempty"`
	// Servo angle
	Angle                uint32   `protobuf:"varint,2,opt,name=angle,proto3" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServoParams) Reset()         { *m = ServoParams{} }
func (m *ServoParams) String() string { return proto.CompactTextString(m) }
func (*ServoParams) ProtoMessage()    {}
func (*ServoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{2}
}

func (m *ServoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServoParams.Unmarshal(m, b)
}
func (m *ServoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServoParams.Marshal(b, m, deterministic)
}
func (m *ServoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServoParams.Merge(m, src)
}
func (m *ServoParams) XXX_Size() int {
	return xxx_messageInfo_ServoParams.Size(m)
}
func (m *ServoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ServoParams.DiscardUnknown(m)
}

var xxx_messageInfo_ServoParams proto.InternalMessageInfo

func (m *ServoParams) GetPin() uint32 {
	if m != nil {
		return m.Pin
	}
	return 0
}

func (m *ServoParams) GetAngle() uint32 {
	if m != nil {
		return m.Angle
	}
	return 0
}

// GPIO handler params
type GpioParams struct {
	// GPIO to config
	Pin  uint32              `protobuf:"varint,1,opt,name=pin,proto3" json:"pin,omitempty"`
	Mode GpioParams_EnumMode `protobuf:"varint,2,opt,name=mode,proto3,enum=matrix_io.malos.v1.io.GpioParams_EnumMode" json:"mode,omitempty"`
	// GPIO value
	Value uint32 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	// GPIO vector value
	Values               uint32   `protobuf:"varint,4,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GpioParams) Reset()         { *m = GpioParams{} }
func (m *GpioParams) String() string { return proto.CompactTextString(m) }
func (*GpioParams) ProtoMessage()    {}
func (*GpioParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{3}
}

func (m *GpioParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GpioParams.Unmarshal(m, b)
}
func (m *GpioParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GpioParams.Marshal(b, m, deterministic)
}
func (m *GpioParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GpioParams.Merge(m, src)
}
func (m *GpioParams) XXX_Size() int {
	return xxx_messageInfo_GpioParams.Size(m)
}
func (m *GpioParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GpioParams.DiscardUnknown(m)
}

var xxx_messageInfo_GpioParams proto.InternalMessageInfo

func (m *GpioParams) GetPin() uint32 {
	if m != nil {
		return m.Pin
	}
	return 0
}

func (m *GpioParams) GetMode() GpioParams_EnumMode {
	if m != nil {
		return m.Mode
	}
	return GpioParams_INPUT
}

func (m *GpioParams) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *GpioParams) GetValues() uint32 {
	if m != nil {
		return m.Values
	}
	return 0
}

// Microphone Array params
type MicArrayParams struct {
	// gain for all microphones
	Gain int32 `protobuf:"varint,1,opt,name=gain,proto3" json:"gain,omitempty"`
	// Spherical coordinates for beamforming (azimutal_angle,
	// polar_angle, radial_distance_mm).
	//
	// azimutal angle (must lie between -pi/2 an pi/2)
	AzimutalAngle float32 `protobuf:"fixed32,2,opt,name=azimutal_angle,json=azimutalAngle,proto3" json:"azimutal_angle,omitempty"`
	// polar angle (must lie between -pi/2 an pi/2)
	PolarAngle float32 `protobuf:"fixed32,3,opt,name=polar_angle,json=polarAngle,proto3" json:"polar_angle,omitempty"`
	// distance from the center of the MATRIX Creator to the sound source.
	RadialDistanceMm float32 `protobuf:"fixed32,4,opt,name=radial_distance_mm,json=radialDistanceMm,proto3" json:"radial_distance_mm,omitempty"`
	// sound speed in mm/seg
	SoundSpeedMmseg float32 `protobuf:"fixed32,5,opt,name=sound_speed_mmseg,json=soundSpeedMmseg,proto3" json:"sound_speed_mmseg,omitempty"`
	// sampling frequency in Hz
	SamplingFrequencyHz  uint32   `protobuf:"varint,6,opt,name=sampling_frequency_hz,json=samplingFrequencyHz,proto3" json:"sampling_frequency_hz,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MicArrayParams) Reset()         { *m = MicArrayParams{} }
func (m *MicArrayParams) String() string { return proto.CompactTextString(m) }
func (*MicArrayParams) ProtoMessage()    {}
func (*MicArrayParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{4}
}

func (m *MicArrayParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MicArrayParams.Unmarshal(m, b)
}
func (m *MicArrayParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MicArrayParams.Marshal(b, m, deterministic)
}
func (m *MicArrayParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicArrayParams.Merge(m, src)
}
func (m *MicArrayParams) XXX_Size() int {
	return xxx_messageInfo_MicArrayParams.Size(m)
}
func (m *MicArrayParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MicArrayParams.DiscardUnknown(m)
}

var xxx_messageInfo_MicArrayParams proto.InternalMessageInfo

func (m *MicArrayParams) GetGain() int32 {
	if m != nil {
		return m.Gain
	}
	return 0
}

func (m *MicArrayParams) GetAzimutalAngle() float32 {
	if m != nil {
		return m.AzimutalAngle
	}
	return 0
}

func (m *MicArrayParams) GetPolarAngle() float32 {
	if m != nil {
		return m.PolarAngle
	}
	return 0
}

func (m *MicArrayParams) GetRadialDistanceMm() float32 {
	if m != nil {
		return m.RadialDistanceMm
	}
	return 0
}

func (m *MicArrayParams) GetSoundSpeedMmseg() float32 {
	if m != nil {
		return m.SoundSpeedMmseg
	}
	return 0
}

func (m *MicArrayParams) GetSamplingFrequencyHz() uint32 {
	if m != nil {
		return m.SamplingFrequencyHz
	}
	return 0
}

type WakeWordParams struct {
	// Wake Word
	WakeWord string                    `protobuf:"bytes,1,opt,name=wake_word,json=wakeWord,proto3" json:"wake_word,omitempty"`
	Channel  WakeWordParams_MicChannel `protobuf:"varint,2,opt,name=channel,proto3,enum=matrix_io.malos.v1.io.WakeWordParams_MicChannel" json:"channel,omitempty"`
	// lenguaje model path from lmtool or similar alternative:
	// http://www.speech.cs.cmu.edu/tools/lmtool-new.html
	LmPath string `protobuf:"bytes,3,opt,name=lm_path,json=lmPath,proto3" json:"lm_path,omitempty"`
	// dictionary path from lmtool
	DicPath string `protobuf:"bytes,4,opt,name=dic_path,json=dicPath,proto3" json:"dic_path,omitempty"`
	// enable pocketsphinx verbose mode
	EnableVerbose bool `protobuf:"varint,5,opt,name=enable_verbose,json=enableVerbose,proto3" json:"enable_verbose,omitempty"`
	// stop recognition service
	StopRecognition      bool     `protobuf:"varint,6,opt,name=stop_recognition,json=stopRecognition,proto3" json:"stop_recognition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WakeWordParams) Reset()         { *m = WakeWordParams{} }
func (m *WakeWordParams) String() string { return proto.CompactTextString(m) }
func (*WakeWordParams) ProtoMessage()    {}
func (*WakeWordParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{5}
}

func (m *WakeWordParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WakeWordParams.Unmarshal(m, b)
}
func (m *WakeWordParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WakeWordParams.Marshal(b, m, deterministic)
}
func (m *WakeWordParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WakeWordParams.Merge(m, src)
}
func (m *WakeWordParams) XXX_Size() int {
	return xxx_messageInfo_WakeWordParams.Size(m)
}
func (m *WakeWordParams) XXX_DiscardUnknown() {
	xxx_messageInfo_WakeWordParams.DiscardUnknown(m)
}

var xxx_messageInfo_WakeWordParams proto.InternalMessageInfo

func (m *WakeWordParams) GetWakeWord() string {
	if m != nil {
		return m.WakeWord
	}
	return ""
}

func (m *WakeWordParams) GetChannel() WakeWordParams_MicChannel {
	if m != nil {
		return m.Channel
	}
	return WakeWordParams_channel0
}

func (m *WakeWordParams) GetLmPath() string {
	if m != nil {
		return m.LmPath
	}
	return ""
}

func (m *WakeWordParams) GetDicPath() string {
	if m != nil {
		return m.DicPath
	}
	return ""
}

func (m *WakeWordParams) GetEnableVerbose() bool {
	if m != nil {
		return m.EnableVerbose
	}
	return false
}

func (m *WakeWordParams) GetStopRecognition() bool {
	if m != nil {
		return m.StopRecognition
	}
	return false
}

type MatrixDeviceParams struct {
	// Matrix Device Type (1: CREATOR - 2:VOICE)
	DeviceType           EnumMatrixDeviceType `protobuf:"varint,1,opt,name=device_type,json=deviceType,proto3,enum=matrix_io.malos.v1.io.EnumMatrixDeviceType" json:"device_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MatrixDeviceParams) Reset()         { *m = MatrixDeviceParams{} }
func (m *MatrixDeviceParams) String() string { return proto.CompactTextString(m) }
func (*MatrixDeviceParams) ProtoMessage()    {}
func (*MatrixDeviceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961acb3747a91b5, []int{6}
}

func (m *MatrixDeviceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatrixDeviceParams.Unmarshal(m, b)
}
func (m *MatrixDeviceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatrixDeviceParams.Marshal(b, m, deterministic)
}
func (m *MatrixDeviceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatrixDeviceParams.Merge(m, src)
}
func (m *MatrixDeviceParams) XXX_Size() int {
	return xxx_messageInfo_MatrixDeviceParams.Size(m)
}
func (m *MatrixDeviceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MatrixDeviceParams.DiscardUnknown(m)
}

var xxx_messageInfo_MatrixDeviceParams proto.InternalMessageInfo

func (m *MatrixDeviceParams) GetDeviceType() EnumMatrixDeviceType {
	if m != nil {
		return m.DeviceType
	}
	return EnumMatrixDeviceType_UNDETECTED
}

func init() {
	proto.RegisterEnum("matrix_io.malos.v1.io.EnumMatrixDeviceType", EnumMatrixDeviceType_name, EnumMatrixDeviceType_value)
	proto.RegisterEnum("matrix_io.malos.v1.io.GpioParams_EnumMode", GpioParams_EnumMode_name, GpioParams_EnumMode_value)
	proto.RegisterEnum("matrix_io.malos.v1.io.WakeWordParams_MicChannel", WakeWordParams_MicChannel_name, WakeWordParams_MicChannel_value)
	proto.RegisterType((*LedValue)(nil), "matrix_io.malos.v1.io.LedValue")
	proto.RegisterType((*EverloopImage)(nil), "matrix_io.malos.v1.io.EverloopImage")
	proto.RegisterType((*ServoParams)(nil), "matrix_io.malos.v1.io.ServoParams")
	proto.RegisterType((*GpioParams)(nil), "matrix_io.malos.v1.io.GpioParams")
	proto.RegisterType((*MicArrayParams)(nil), "matrix_io.malos.v1.io.MicArrayParams")
	proto.RegisterType((*WakeWordParams)(nil), "matrix_io.malos.v1.io.WakeWordParams")
	proto.RegisterType((*MatrixDeviceParams)(nil), "matrix_io.malos.v1.io.MatrixDeviceParams")
}

func init() { proto.RegisterFile("matrix_io/malos/v1/io.proto", fileDescriptor_8961acb3747a91b5) }

var fileDescriptor_8961acb3747a91b5 = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdf, 0x6f, 0xe3, 0x44,
	0x10, 0xae, 0xf3, 0xbb, 0x13, 0x9a, 0x9a, 0xa5, 0x07, 0x41, 0x7d, 0x38, 0xb0, 0x74, 0xe2, 0x28,
	0x9c, 0x7b, 0xc9, 0x71, 0x70, 0x4f, 0x27, 0xf5, 0xda, 0x00, 0x41, 0xcd, 0x25, 0xf2, 0xa5, 0x3d,
	0x09, 0x21, 0x59, 0x1b, 0x7b, 0x70, 0x56, 0xb5, 0x77, 0xcd, 0xda, 0x49, 0x49, 0xff, 0x04, 0xc4,
	0x5f, 0xc2, 0x0b, 0xff, 0x1f, 0x4f, 0xc8, 0x63, 0xbb, 0xf1, 0x49, 0xcd, 0xdb, 0x7c, 0xdf, 0x7c,
	0x33, 0xfb, 0xcd, 0xac, 0xbd, 0x70, 0x1c, 0xf1, 0x54, 0x8b, 0x3f, 0x5d, 0xa1, 0x4e, 0x23, 0x1e,
	0xaa, 0xe4, 0x74, 0x3d, 0x38, 0x15, 0xca, 0x8e, 0xb5, 0x4a, 0x15, 0x7b, 0x74, 0x9f, 0xb4, 0x29,
	0x69, 0xaf, 0x07, 0xb6, 0x50, 0xd6, 0x6f, 0xd0, 0xb9, 0x44, 0xff, 0x9a, 0x87, 0x2b, 0x64, 0x26,
	0xd4, 0x35, 0xfa, 0x7d, 0xe3, 0x0b, 0xe3, 0xe9, 0x81, 0x93, 0x85, 0xec, 0x08, 0x9a, 0x81, 0x46,
	0x94, 0xfd, 0x1a, 0x71, 0x39, 0x60, 0x0c, 0x1a, 0x8b, 0x70, 0x85, 0xfd, 0x3a, 0x91, 0x14, 0x67,
	0xca, 0xdb, 0xa5, 0x48, 0xb1, 0xdf, 0xc8, 0x95, 0x04, 0xac, 0x1b, 0x38, 0x18, 0xad, 0x51, 0x87,
	0x4a, 0xc5, 0xe3, 0x88, 0x07, 0xc8, 0x06, 0x50, 0x0f, 0xe9, 0x88, 0xfa, 0xd3, 0xee, 0xf0, 0xb1,
	0xfd, 0xa0, 0x27, 0xbb, 0x34, 0xe4, 0x64, 0x5a, 0xf6, 0x15, 0x1c, 0x62, 0xd1, 0xc3, 0x0d, 0x51,
	0x06, 0xe9, 0x92, 0xdc, 0x34, 0x9d, 0x5e, 0x49, 0x5f, 0x12, 0x6b, 0xbd, 0x84, 0xee, 0x3b, 0xd4,
	0x6b, 0x35, 0xe3, 0x9a, 0x47, 0x49, 0x36, 0x4d, 0x2c, 0x64, 0x39, 0x4d, 0x2c, 0x64, 0xe6, 0x91,
	0xcb, 0x20, 0xc4, 0x72, 0x1a, 0x02, 0xd6, 0xbf, 0x06, 0xc0, 0x4f, 0xb1, 0xd8, 0x5d, 0xf6, 0x1a,
	0x1a, 0x91, 0xf2, 0xf3, 0xaa, 0xde, 0xf0, 0x64, 0x87, 0xe9, 0x6d, 0x0b, 0x7b, 0x24, 0x57, 0xd1,
	0x44, 0xf9, 0xe8, 0x50, 0x5d, 0x76, 0xec, 0x9a, 0x6f, 0xf7, 0x95, 0x03, 0xf6, 0x29, 0xb4, 0x28,
	0x48, 0x8a, 0x8d, 0x15, 0xc8, 0xfa, 0x12, 0x3a, 0x65, 0x3d, 0xdb, 0x87, 0xe6, 0xf8, 0xed, 0xec,
	0x6a, 0x6e, 0xee, 0x31, 0x80, 0xd6, 0xf4, 0x6a, 0x9e, 0xc5, 0x86, 0xf5, 0x9f, 0x01, 0xbd, 0x89,
	0xf0, 0xce, 0xb4, 0xe6, 0x9b, 0xc2, 0x35, 0x83, 0x46, 0xc0, 0x0b, 0xdb, 0x4d, 0x87, 0x62, 0xf6,
	0x04, 0x7a, 0xfc, 0x4e, 0x44, 0xab, 0x94, 0x87, 0xee, 0x76, 0xee, 0x9a, 0x73, 0x50, 0xb2, 0x67,
	0x19, 0xc9, 0x1e, 0x43, 0x37, 0x56, 0x21, 0xd7, 0x85, 0xa6, 0x4e, 0x1a, 0x20, 0x2a, 0x17, 0x7c,
	0x0b, 0x4c, 0x73, 0x5f, 0xf0, 0xd0, 0xf5, 0x45, 0x92, 0x72, 0xe9, 0xa1, 0x1b, 0x45, 0xe4, 0xba,
	0xe6, 0x98, 0x79, 0xe6, 0xa2, 0x48, 0x4c, 0x22, 0x76, 0x02, 0x1f, 0x27, 0x6a, 0x25, 0x7d, 0x37,
	0x89, 0x11, 0x7d, 0x37, 0x8a, 0x12, 0x0c, 0xfa, 0x4d, 0x12, 0x1f, 0x52, 0xe2, 0x5d, 0xc6, 0x4f,
	0x32, 0x9a, 0x0d, 0xe1, 0x51, 0xc2, 0xa3, 0x38, 0x14, 0x32, 0x70, 0x7f, 0xd7, 0xf8, 0xc7, 0x0a,
	0xa5, 0xb7, 0x71, 0x97, 0x77, 0xfd, 0x16, 0xad, 0xe4, 0x93, 0x32, 0xf9, 0x63, 0x99, 0xfb, 0xf9,
	0xce, 0xfa, 0xbb, 0x0e, 0xbd, 0xf7, 0xfc, 0x06, 0xdf, 0x2b, 0xed, 0x17, 0xc3, 0x1f, 0xc3, 0xfe,
	0x2d, 0xbf, 0x41, 0xf7, 0x56, 0xe9, 0xfc, 0xeb, 0xdd, 0x77, 0x3a, 0xb7, 0x85, 0x84, 0xfd, 0x02,
	0x6d, 0x6f, 0xc9, 0xa5, 0xc4, 0xb0, 0xb8, 0xc0, 0xe7, 0x3b, 0x2e, 0xf0, 0xc3, 0xa6, 0xf6, 0x44,
	0x78, 0xe7, 0x79, 0x9d, 0x53, 0x36, 0x60, 0x9f, 0x41, 0x3b, 0x8c, 0xdc, 0x98, 0xa7, 0x4b, 0x5a,
	0xd3, 0xbe, 0xd3, 0x0a, 0xa3, 0x19, 0x4f, 0x97, 0xec, 0x73, 0xe8, 0xf8, 0xc2, 0xcb, 0x33, 0x0d,
	0xca, 0xb4, 0x7d, 0xe1, 0x51, 0xea, 0x09, 0xf4, 0x50, 0xf2, 0x45, 0x88, 0xee, 0x1a, 0xf5, 0x42,
	0x25, 0x48, 0xcb, 0xe8, 0x38, 0x07, 0x39, 0x7b, 0x9d, 0x93, 0xec, 0x6b, 0x30, 0x93, 0x54, 0xc5,
	0xae, 0x46, 0x4f, 0x05, 0x52, 0xa4, 0x42, 0x49, 0xda, 0x42, 0xc7, 0x39, 0xcc, 0x78, 0x67, 0x4b,
	0x5b, 0x7f, 0x19, 0x00, 0x5b, 0x77, 0xec, 0x23, 0xe8, 0x14, 0xfe, 0x9e, 0x9b, 0x7b, 0x15, 0x34,
	0x30, 0x8d, 0x0a, 0x1a, 0x9a, 0xb5, 0x0a, 0x7a, 0x61, 0xd6, 0x2b, 0xe8, 0x3b, 0xb3, 0x51, 0x41,
	0x2f, 0xcd, 0x66, 0x05, 0x7d, 0x6f, 0xb6, 0x2a, 0xe8, 0x07, 0xb3, 0x5d, 0x41, 0xaf, 0xcc, 0x8e,
	0xb5, 0x00, 0x36, 0xa1, 0x75, 0x5e, 0xe0, 0x5a, 0x78, 0x58, 0xdc, 0xc8, 0x25, 0x74, 0x7d, 0xc2,
	0x6e, 0xba, 0x89, 0x91, 0xee, 0xa4, 0x37, 0xfc, 0x66, 0xc7, 0xe2, 0xe9, 0x73, 0xaf, 0xf4, 0x98,
	0x6f, 0x62, 0x74, 0xc0, 0xbf, 0x8f, 0x4f, 0x5e, 0xc3, 0xd1, 0x43, 0x1a, 0xd6, 0x03, 0xb8, 0x7a,
	0x7b, 0x31, 0x9a, 0x8f, 0xce, 0xe7, 0xa3, 0x0b, 0x73, 0x8f, 0x75, 0xa1, 0x7d, 0xee, 0x8c, 0xce,
	0xe6, 0x53, 0xc7, 0x34, 0xb2, 0x7f, 0xe7, 0x7a, 0x3a, 0x3e, 0x1f, 0x99, 0xb5, 0x37, 0x1b, 0x38,
	0x56, 0x12, 0x8b, 0xd3, 0xcb, 0x07, 0xf1, 0xde, 0xc2, 0x9b, 0xf6, 0x78, 0x3a, 0xcb, 0x98, 0x99,
	0xf1, 0xeb, 0xab, 0x40, 0xa4, 0xcb, 0xd5, 0xc2, 0xf6, 0x54, 0x74, 0x9a, 0xcb, 0x9f, 0xd1, 0x63,
	0x4a, 0x11, 0x95, 0x25, 0xcf, 0x82, 0x92, 0xa8, 0xbe, 0xb3, 0xff, 0xd4, 0x8e, 0x72, 0x7b, 0xe3,
	0xa9, 0x3d, 0xa1, 0xc6, 0xe3, 0xa9, 0x7d, 0x3d, 0x58, 0xb4, 0xa8, 0xe8, 0xc5, 0xff, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x08, 0x2a, 0x75, 0x0f, 0x9b, 0x05, 0x00, 0x00,
}
