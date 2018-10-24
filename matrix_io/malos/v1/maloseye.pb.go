// Code generated by protoc-gen-go. DO NOT EDIT.
// source: matrix_io/malos/v1/maloseye.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/matrix-io/matrix-protos-go/matrix_io/common"
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

type EnumMalosAction int32

const (
	// Start the detection.
	EnumMalosAction_START_DETECTION EnumMalosAction = 0
	// Stop the detection.
	EnumMalosAction_STOP_DETECTION EnumMalosAction = 1
	// Restart the detection.
	EnumMalosAction_RESTART_DETECTION EnumMalosAction = 2
	// Capture preview image.
	EnumMalosAction_CAPTURE_PREVIEW EnumMalosAction = 3
)

var EnumMalosAction_name = map[int32]string{
	0: "START_DETECTION",
	1: "STOP_DETECTION",
	2: "RESTART_DETECTION",
	3: "CAPTURE_PREVIEW",
}

var EnumMalosAction_value = map[string]int32{
	"START_DETECTION":   0,
	"STOP_DETECTION":    1,
	"RESTART_DETECTION": 2,
	"CAPTURE_PREVIEW":   3,
}

func (x EnumMalosAction) String() string {
	return proto.EnumName(EnumMalosAction_name, int32(x))
}

func (EnumMalosAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bd20f3a2ecc2cac, []int{0}
}

type EnumMalosEyeDetectionType int32

const (
	// Stop. Don't do detections.
	// This only makes sense when this is the only command sent to MalosEye.
	EnumMalosEyeDetectionType_STOP EnumMalosEyeDetectionType = 0
	// Detect frontal faces.
	EnumMalosEyeDetectionType_FACE EnumMalosEyeDetectionType = 20
	// Detect faces and enhance the detection with demographics.
	// This might involve a connection to a remote detection server.
	EnumMalosEyeDetectionType_FACE_DEMOGRAPHICS EnumMalosEyeDetectionType = 21
	// Get face descriptors for detected faces.
	// The descriptors are used for face recognition,
	// that is, tell people apart.
	EnumMalosEyeDetectionType_FACE_DESCRIPTOR EnumMalosEyeDetectionType = 30
	// Detect thum-up gesture.
	EnumMalosEyeDetectionType_HAND_THUMB_UP EnumMalosEyeDetectionType = 40
	// Detect palm gesture.
	EnumMalosEyeDetectionType_HAND_PALM EnumMalosEyeDetectionType = 41
	// Detect pinch gesture.
	EnumMalosEyeDetectionType_HAND_PINCH EnumMalosEyeDetectionType = 42
	// Detect fist gesture.
	EnumMalosEyeDetectionType_HAND_FIST EnumMalosEyeDetectionType = 43
)

var EnumMalosEyeDetectionType_name = map[int32]string{
	0:  "STOP",
	20: "FACE",
	21: "FACE_DEMOGRAPHICS",
	30: "FACE_DESCRIPTOR",
	40: "HAND_THUMB_UP",
	41: "HAND_PALM",
	42: "HAND_PINCH",
	43: "HAND_FIST",
}

var EnumMalosEyeDetectionType_value = map[string]int32{
	"STOP":              0,
	"FACE":              20,
	"FACE_DEMOGRAPHICS": 21,
	"FACE_DESCRIPTOR":   30,
	"HAND_THUMB_UP":     40,
	"HAND_PALM":         41,
	"HAND_PINCH":        42,
	"HAND_FIST":         43,
}

func (x EnumMalosEyeDetectionType) String() string {
	return proto.EnumName(EnumMalosEyeDetectionType_name, int32(x))
}

func (EnumMalosEyeDetectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bd20f3a2ecc2cac, []int{1}
}

// Configuration for MALOS-eye. If object_to_detect has at least
// one element the rest of the configuration will be ignored and the
// objects to be detected will change acordingly.
type MalosEyeConfig struct {
	// Camera configuration. Camera id, width, height, and so on.
	CameraConfig *CameraConfig `protobuf:"bytes,1,opt,name=camera_config,json=cameraConfig,proto3" json:"camera_config,omitempty"`
	// Face configuration. models path, send slow, send text proto.
	FaceConfig *FaceConfig `protobuf:"bytes,2,opt,name=face_config,json=faceConfig,proto3" json:"face_config,omitempty"`
	// Configuration of the detection server.
	DetectionServerConfig *DetectionServerConfig `protobuf:"bytes,4,opt,name=detection_server_config,json=detectionServerConfig,proto3" json:"detection_server_config,omitempty"`
	// Detections to perform. When this field is set the other fields
	// are ignored.
	ObjectToDetect []EnumMalosEyeDetectionType `protobuf:"varint,21,rep,packed,name=object_to_detect,json=objectToDetect,proto3,enum=matrix_io.malos.v1.maloseye.EnumMalosEyeDetectionType" json:"object_to_detect,omitempty"`
	// Action to start, stop or restart the detection.
	Action EnumMalosAction `protobuf:"varint,5,opt,name=action,proto3,enum=matrix_io.malos.v1.maloseye.EnumMalosAction" json:"action,omitempty"`
	// Models configuration
	Models []*common.Entity `protobuf:"bytes,6,rep,name=models,proto3" json:"models,omitempty"`
	// Config Entity
	Config               *common.Entity `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MalosEyeConfig) Reset()         { *m = MalosEyeConfig{} }
func (m *MalosEyeConfig) String() string { return proto.CompactTextString(m) }
func (*MalosEyeConfig) ProtoMessage()    {}
func (*MalosEyeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd20f3a2ecc2cac, []int{0}
}

func (m *MalosEyeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MalosEyeConfig.Unmarshal(m, b)
}
func (m *MalosEyeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MalosEyeConfig.Marshal(b, m, deterministic)
}
func (m *MalosEyeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MalosEyeConfig.Merge(m, src)
}
func (m *MalosEyeConfig) XXX_Size() int {
	return xxx_messageInfo_MalosEyeConfig.Size(m)
}
func (m *MalosEyeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MalosEyeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MalosEyeConfig proto.InternalMessageInfo

func (m *MalosEyeConfig) GetCameraConfig() *CameraConfig {
	if m != nil {
		return m.CameraConfig
	}
	return nil
}

func (m *MalosEyeConfig) GetFaceConfig() *FaceConfig {
	if m != nil {
		return m.FaceConfig
	}
	return nil
}

func (m *MalosEyeConfig) GetDetectionServerConfig() *DetectionServerConfig {
	if m != nil {
		return m.DetectionServerConfig
	}
	return nil
}

func (m *MalosEyeConfig) GetObjectToDetect() []EnumMalosEyeDetectionType {
	if m != nil {
		return m.ObjectToDetect
	}
	return nil
}

func (m *MalosEyeConfig) GetAction() EnumMalosAction {
	if m != nil {
		return m.Action
	}
	return EnumMalosAction_START_DETECTION
}

func (m *MalosEyeConfig) GetModels() []*common.Entity {
	if m != nil {
		return m.Models
	}
	return nil
}

func (m *MalosEyeConfig) GetConfig() *common.Entity {
	if m != nil {
		return m.Config
	}
	return nil
}

type CameraConfig struct {
	// Source (camera URL, file). When set this has precedence over
	// camera_id.
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	// What camera to open? 0 is usually the first camera and it's the
	// default value.
	CameraId int32 `protobuf:"varint,1,opt,name=camera_id,json=cameraId,proto3" json:"camera_id,omitempty"`
	// Capture width, i.e. 640.
	Width int32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	// Capture height, i.e. 480.
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CameraConfig) Reset()         { *m = CameraConfig{} }
func (m *CameraConfig) String() string { return proto.CompactTextString(m) }
func (*CameraConfig) ProtoMessage()    {}
func (*CameraConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd20f3a2ecc2cac, []int{1}
}

func (m *CameraConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CameraConfig.Unmarshal(m, b)
}
func (m *CameraConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CameraConfig.Marshal(b, m, deterministic)
}
func (m *CameraConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CameraConfig.Merge(m, src)
}
func (m *CameraConfig) XXX_Size() int {
	return xxx_messageInfo_CameraConfig.Size(m)
}
func (m *CameraConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CameraConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CameraConfig proto.InternalMessageInfo

func (m *CameraConfig) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *CameraConfig) GetCameraId() int32 {
	if m != nil {
		return m.CameraId
	}
	return 0
}

func (m *CameraConfig) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *CameraConfig) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

// Warning! This setting is not used yet. It will probably need a token.
type DetectionServerConfig struct {
	// Detection server connection string. For instance, "127.0.0.1:32061".
	DetectionServerAddress string `protobuf:"bytes,1,opt,name=detection_server_address,json=detectionServerAddress,proto3" json:"detection_server_address,omitempty"`
	// Timeout for detection server (in milliseconds).
	DetectionServerTimeout int32    `protobuf:"varint,2,opt,name=detection_server_timeout,json=detectionServerTimeout,proto3" json:"detection_server_timeout,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *DetectionServerConfig) Reset()         { *m = DetectionServerConfig{} }
func (m *DetectionServerConfig) String() string { return proto.CompactTextString(m) }
func (*DetectionServerConfig) ProtoMessage()    {}
func (*DetectionServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd20f3a2ecc2cac, []int{2}
}

func (m *DetectionServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectionServerConfig.Unmarshal(m, b)
}
func (m *DetectionServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectionServerConfig.Marshal(b, m, deterministic)
}
func (m *DetectionServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectionServerConfig.Merge(m, src)
}
func (m *DetectionServerConfig) XXX_Size() int {
	return xxx_messageInfo_DetectionServerConfig.Size(m)
}
func (m *DetectionServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectionServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DetectionServerConfig proto.InternalMessageInfo

func (m *DetectionServerConfig) GetDetectionServerAddress() string {
	if m != nil {
		return m.DetectionServerAddress
	}
	return ""
}

func (m *DetectionServerConfig) GetDetectionServerTimeout() int32 {
	if m != nil {
		return m.DetectionServerTimeout
	}
	return 0
}

type FaceConfig struct {
	// The path with the models folder.
	ModelsPath string `protobuf:"bytes,1,opt,name=models_path,json=modelsPath,proto3" json:"models_path,omitempty"`
	// Send slow frames for debugging
	SendSlow bool `protobuf:"varint,2,opt,name=send_slow,json=sendSlow,proto3" json:"send_slow,omitempty"`
	// Send text proto for debugging
	SendTextProto        bool     `protobuf:"varint,3,opt,name=send_text_proto,json=sendTextProto,proto3" json:"send_text_proto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaceConfig) Reset()         { *m = FaceConfig{} }
func (m *FaceConfig) String() string { return proto.CompactTextString(m) }
func (*FaceConfig) ProtoMessage()    {}
func (*FaceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd20f3a2ecc2cac, []int{3}
}

func (m *FaceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceConfig.Unmarshal(m, b)
}
func (m *FaceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceConfig.Marshal(b, m, deterministic)
}
func (m *FaceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceConfig.Merge(m, src)
}
func (m *FaceConfig) XXX_Size() int {
	return xxx_messageInfo_FaceConfig.Size(m)
}
func (m *FaceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FaceConfig proto.InternalMessageInfo

func (m *FaceConfig) GetModelsPath() string {
	if m != nil {
		return m.ModelsPath
	}
	return ""
}

func (m *FaceConfig) GetSendSlow() bool {
	if m != nil {
		return m.SendSlow
	}
	return false
}

func (m *FaceConfig) GetSendTextProto() bool {
	if m != nil {
		return m.SendTextProto
	}
	return false
}

func init() {
	proto.RegisterEnum("matrix_io.malos.v1.maloseye.EnumMalosAction", EnumMalosAction_name, EnumMalosAction_value)
	proto.RegisterEnum("matrix_io.malos.v1.maloseye.EnumMalosEyeDetectionType", EnumMalosEyeDetectionType_name, EnumMalosEyeDetectionType_value)
	proto.RegisterType((*MalosEyeConfig)(nil), "matrix_io.malos.v1.maloseye.MalosEyeConfig")
	proto.RegisterType((*CameraConfig)(nil), "matrix_io.malos.v1.maloseye.CameraConfig")
	proto.RegisterType((*DetectionServerConfig)(nil), "matrix_io.malos.v1.maloseye.DetectionServerConfig")
	proto.RegisterType((*FaceConfig)(nil), "matrix_io.malos.v1.maloseye.FaceConfig")
}

func init() { proto.RegisterFile("matrix_io/malos/v1/maloseye.proto", fileDescriptor_8bd20f3a2ecc2cac) }

var fileDescriptor_8bd20f3a2ecc2cac = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x72, 0xda, 0x48,
	0x10, 0xb5, 0x8c, 0x61, 0xa1, 0x31, 0x58, 0x9e, 0x35, 0x5e, 0xd6, 0xae, 0xdd, 0x65, 0x39, 0x24,
	0xd8, 0x89, 0x45, 0x4c, 0xaa, 0x52, 0xbe, 0xca, 0x42, 0x0e, 0xaa, 0x0a, 0xa0, 0x1a, 0x64, 0xa7,
	0x2a, 0x17, 0x45, 0x96, 0x06, 0x90, 0x0b, 0x31, 0x8e, 0x34, 0xd8, 0xe6, 0x03, 0x72, 0xca, 0x2f,
	0xe4, 0x94, 0x63, 0xbe, 0x32, 0xa5, 0x19, 0x09, 0x53, 0xc4, 0xa1, 0x72, 0x9b, 0x7e, 0xdd, 0xef,
	0xf5, 0xa8, 0xfb, 0x69, 0xe0, 0xff, 0xc0, 0x61, 0xa1, 0xff, 0x60, 0xfb, 0xb4, 0x19, 0x38, 0x13,
	0x1a, 0x35, 0xef, 0x4e, 0xc5, 0x81, 0xcc, 0x89, 0x72, 0x1b, 0x52, 0x46, 0xd1, 0xe1, 0xa2, 0x44,
	0xe1, 0x19, 0xe5, 0xee, 0x54, 0x49, 0x4b, 0x0e, 0xfe, 0x79, 0xe4, 0xbb, 0x34, 0x08, 0xe8, 0xb4,
	0x49, 0xa6, 0xcc, 0x67, 0x73, 0xc1, 0xad, 0x7f, 0xdb, 0x82, 0x72, 0x37, 0xae, 0xd5, 0xe7, 0x44,
	0xa3, 0xd3, 0xa1, 0x3f, 0x42, 0x3d, 0x28, 0xb9, 0x4e, 0x40, 0x42, 0xc7, 0x76, 0x39, 0x50, 0x95,
	0x6a, 0x52, 0xa3, 0xd8, 0x3a, 0x52, 0xd6, 0xb4, 0x51, 0x34, 0xce, 0x10, 0x0a, 0x78, 0xdb, 0x5d,
	0x8a, 0x50, 0x07, 0x8a, 0x43, 0xc7, 0x25, 0xa9, 0xda, 0x26, 0x57, 0x7b, 0xbe, 0x56, 0xed, 0xc2,
	0x71, 0x93, 0xdb, 0x60, 0x18, 0x2e, 0xce, 0xe8, 0x06, 0xfe, 0xf2, 0x08, 0x23, 0x2e, 0xf3, 0xe9,
	0xd4, 0x8e, 0x48, 0x78, 0x47, 0xc2, 0x54, 0x75, 0x8b, 0xab, 0xb6, 0xd6, 0xaa, 0xb6, 0x53, 0xee,
	0x80, 0x53, 0x93, 0x06, 0x15, 0xef, 0x29, 0x18, 0x7d, 0x04, 0x99, 0x5e, 0xdf, 0x10, 0x97, 0xd9,
	0x8c, 0xda, 0xa2, 0xa4, 0x5a, 0xa9, 0x65, 0x1a, 0xe5, 0xd6, 0x9b, 0xb5, 0x4d, 0xf4, 0xe9, 0x2c,
	0x48, 0x07, 0xba, 0x68, 0x68, 0xcd, 0x6f, 0x09, 0x2e, 0x0b, 0x3d, 0x8b, 0x0a, 0x18, 0xb5, 0x21,
	0xe7, 0xf0, 0x6c, 0x35, 0x5b, 0x93, 0x1a, 0xe5, 0xd6, 0xcb, 0xdf, 0xd3, 0x55, 0x39, 0x07, 0x27,
	0x5c, 0xf4, 0x0a, 0x72, 0x01, 0xf5, 0xc8, 0x24, 0xaa, 0xe6, 0x6a, 0x99, 0x46, 0xb1, 0x55, 0x5d,
	0x52, 0x11, 0x0b, 0x57, 0x74, 0xbe, 0x70, 0x9c, 0xd4, 0xc5, 0x8c, 0x64, 0x68, 0x7f, 0xf0, 0xa1,
	0xad, 0x61, 0x88, 0xba, 0xfa, 0x27, 0xd8, 0x5e, 0xde, 0x2f, 0xda, 0x87, 0x5c, 0x44, 0x67, 0xa1,
	0x4b, 0xf8, 0xd8, 0x0b, 0x38, 0x89, 0xd0, 0x21, 0x14, 0x12, 0xe7, 0xf8, 0x1e, 0x77, 0x4d, 0x16,
	0xe7, 0x05, 0x60, 0x78, 0x68, 0x0f, 0xb2, 0xf7, 0xbe, 0xc7, 0xc6, 0xdc, 0x00, 0x59, 0x2c, 0x82,
	0x58, 0x6a, 0x4c, 0xfc, 0xd1, 0x98, 0x55, 0x33, 0x1c, 0x4e, 0xa2, 0xfa, 0x17, 0x09, 0x2a, 0x4f,
	0xee, 0x0b, 0x9d, 0x41, 0xf5, 0x27, 0x13, 0x38, 0x9e, 0x17, 0x92, 0x28, 0xe2, 0x3d, 0x0b, 0x78,
	0x7f, 0x65, 0xa3, 0xaa, 0xc8, 0x3e, 0xc9, 0x64, 0x7e, 0x40, 0xe8, 0x8c, 0x25, 0x97, 0x5a, 0x65,
	0x5a, 0x22, 0x5b, 0x0f, 0x01, 0x1e, 0x2d, 0x89, 0xfe, 0x83, 0xa2, 0x18, 0xa5, 0x7d, 0xeb, 0xb0,
	0x71, 0xd2, 0x14, 0x04, 0x64, 0x3a, 0x6c, 0x1c, 0xcf, 0x21, 0x22, 0x53, 0xcf, 0x8e, 0x26, 0xf4,
	0x9e, 0x2b, 0xe7, 0x71, 0x3e, 0x06, 0x06, 0x13, 0x7a, 0x8f, 0x9e, 0xc1, 0x0e, 0x4f, 0x32, 0xf2,
	0xc0, 0x6c, 0xfe, 0x13, 0xf2, 0x4f, 0xcf, 0xe3, 0x52, 0x0c, 0x5b, 0xe4, 0x81, 0x99, 0x31, 0x78,
	0x3c, 0x84, 0x9d, 0x95, 0x9d, 0xa3, 0x3f, 0x61, 0x67, 0x60, 0xa9, 0xd8, 0xb2, 0xdb, 0xba, 0xa5,
	0x6b, 0x96, 0xd1, 0xef, 0xc9, 0x1b, 0x08, 0x41, 0x79, 0x60, 0xf5, 0xcd, 0x25, 0x4c, 0x42, 0x15,
	0xd8, 0xc5, 0xfa, 0x6a, 0xe9, 0x66, 0xcc, 0xd7, 0x54, 0xd3, 0xba, 0xc4, 0xba, 0x6d, 0x62, 0xfd,
	0xca, 0xd0, 0xdf, 0xcb, 0x99, 0xe3, 0xaf, 0x12, 0xfc, 0xfd, 0x4b, 0xd3, 0xa2, 0x3c, 0x6c, 0xc5,
	0xea, 0xf2, 0x46, 0x7c, 0xba, 0x50, 0x35, 0x5d, 0xde, 0x8b, 0xd5, 0xe3, 0x93, 0xdd, 0xd6, 0xbb,
	0xfd, 0xb7, 0x58, 0x35, 0x3b, 0x86, 0x36, 0x90, 0x2b, 0xb1, 0x7a, 0x02, 0x0f, 0x34, 0x6c, 0x98,
	0x56, 0x1f, 0xcb, 0xff, 0xa2, 0x5d, 0x28, 0x75, 0xd4, 0x5e, 0xdb, 0xb6, 0x3a, 0x97, 0xdd, 0x73,
	0xfb, 0xd2, 0x94, 0x1b, 0xa8, 0x04, 0x05, 0x0e, 0x99, 0xea, 0xbb, 0xae, 0x7c, 0x84, 0xca, 0x00,
	0x22, 0x34, 0x7a, 0x5a, 0x47, 0x3e, 0x5e, 0xa4, 0x2f, 0x8c, 0x81, 0x25, 0xbf, 0x38, 0xff, 0x2c,
	0xc1, 0x21, 0x9d, 0x92, 0xc4, 0xa3, 0x3e, 0x15, 0xef, 0xd6, 0xe2, 0x07, 0x39, 0x2f, 0xa5, 0xf7,
	0xe6, 0x53, 0x33, 0xa5, 0x0f, 0x67, 0x23, 0x9f, 0x8d, 0x67, 0xd7, 0xb1, 0x95, 0x9b, 0x82, 0x74,
	0xc2, 0xdf, 0x4e, 0x7e, 0xe2, 0xe4, 0xe8, 0x64, 0x94, 0x02, 0xcb, 0xcf, 0xea, 0xf7, 0xcd, 0x83,
	0x2e, 0x07, 0x8d, 0xbe, 0xc2, 0x45, 0x95, 0x54, 0x5a, 0xb9, 0x3a, 0xbd, 0xce, 0x71, 0xea, 0xeb,
	0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0x5a, 0x9b, 0xf5, 0x90, 0x05, 0x00, 0x00,
}
