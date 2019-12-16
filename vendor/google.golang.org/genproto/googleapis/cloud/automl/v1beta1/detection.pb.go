// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/detection.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Annotation details for image object detection.
type ImageObjectDetectionAnnotation struct {
	// Output only. The rectangle representing the object location.
	BoundingBox *BoundingPoly `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Output only. The confidence that this annotation is positive for the parent example,
	// value in [0, 1], higher means higher positivity confidence.
	Score                float32  `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageObjectDetectionAnnotation) Reset()         { *m = ImageObjectDetectionAnnotation{} }
func (m *ImageObjectDetectionAnnotation) String() string { return proto.CompactTextString(m) }
func (*ImageObjectDetectionAnnotation) ProtoMessage()    {}
func (*ImageObjectDetectionAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{0}
}

func (m *ImageObjectDetectionAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageObjectDetectionAnnotation.Unmarshal(m, b)
}
func (m *ImageObjectDetectionAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageObjectDetectionAnnotation.Marshal(b, m, deterministic)
}
func (m *ImageObjectDetectionAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageObjectDetectionAnnotation.Merge(m, src)
}
func (m *ImageObjectDetectionAnnotation) XXX_Size() int {
	return xxx_messageInfo_ImageObjectDetectionAnnotation.Size(m)
}
func (m *ImageObjectDetectionAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageObjectDetectionAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_ImageObjectDetectionAnnotation proto.InternalMessageInfo

func (m *ImageObjectDetectionAnnotation) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *ImageObjectDetectionAnnotation) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Annotation details for video object tracking.
type VideoObjectTrackingAnnotation struct {
	// Optional. The instance of the object, expressed as a positive integer. Used to tell
	// apart objects of the same type (i.e. AnnotationSpec) when multiple are
	// present on a single example.
	// NOTE: Instance ID prediction quality is not a part of model evaluation and
	// is done as best effort. Especially in cases when an entity goes
	// off-screen for a longer time (minutes), when it comes back it may be given
	// a new instance ID.
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Required. A time (frame) of a video to which this annotation pertains.
	// Represented as the duration since the video's start.
	TimeOffset *duration.Duration `protobuf:"bytes,2,opt,name=time_offset,json=timeOffset,proto3" json:"time_offset,omitempty"`
	// Required. The rectangle representing the object location on the frame (i.e.
	// at the time_offset of the video).
	BoundingBox *BoundingPoly `protobuf:"bytes,3,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Output only. The confidence that this annotation is positive for the video at
	// the time_offset, value in [0, 1], higher means higher positivity
	// confidence. For annotations created by the user the score is 1. When
	// user approves an annotation, the original float score is kept (and not
	// changed to 1).
	Score                float32  `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoObjectTrackingAnnotation) Reset()         { *m = VideoObjectTrackingAnnotation{} }
func (m *VideoObjectTrackingAnnotation) String() string { return proto.CompactTextString(m) }
func (*VideoObjectTrackingAnnotation) ProtoMessage()    {}
func (*VideoObjectTrackingAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{1}
}

func (m *VideoObjectTrackingAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoObjectTrackingAnnotation.Unmarshal(m, b)
}
func (m *VideoObjectTrackingAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoObjectTrackingAnnotation.Marshal(b, m, deterministic)
}
func (m *VideoObjectTrackingAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoObjectTrackingAnnotation.Merge(m, src)
}
func (m *VideoObjectTrackingAnnotation) XXX_Size() int {
	return xxx_messageInfo_VideoObjectTrackingAnnotation.Size(m)
}
func (m *VideoObjectTrackingAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoObjectTrackingAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_VideoObjectTrackingAnnotation proto.InternalMessageInfo

func (m *VideoObjectTrackingAnnotation) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *VideoObjectTrackingAnnotation) GetTimeOffset() *duration.Duration {
	if m != nil {
		return m.TimeOffset
	}
	return nil
}

func (m *VideoObjectTrackingAnnotation) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *VideoObjectTrackingAnnotation) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Bounding box matching model metrics for a single intersection-over-union
// threshold and multiple label match confidence thresholds.
type BoundingBoxMetricsEntry struct {
	// Output only. The intersection-over-union threshold value used to compute
	// this metrics entry.
	IouThreshold float32 `protobuf:"fixed32,1,opt,name=iou_threshold,json=iouThreshold,proto3" json:"iou_threshold,omitempty"`
	// Output only. The mean average precision, most often close to au_prc.
	MeanAveragePrecision float32 `protobuf:"fixed32,2,opt,name=mean_average_precision,json=meanAveragePrecision,proto3" json:"mean_average_precision,omitempty"`
	// Output only. Metrics for each label-match confidence_threshold from
	// 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99. Precision-recall curve is
	// derived from them.
	ConfidenceMetricsEntries []*BoundingBoxMetricsEntry_ConfidenceMetricsEntry `protobuf:"bytes,3,rep,name=confidence_metrics_entries,json=confidenceMetricsEntries,proto3" json:"confidence_metrics_entries,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                                          `json:"-"`
	XXX_unrecognized         []byte                                            `json:"-"`
	XXX_sizecache            int32                                             `json:"-"`
}

func (m *BoundingBoxMetricsEntry) Reset()         { *m = BoundingBoxMetricsEntry{} }
func (m *BoundingBoxMetricsEntry) String() string { return proto.CompactTextString(m) }
func (*BoundingBoxMetricsEntry) ProtoMessage()    {}
func (*BoundingBoxMetricsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{2}
}

func (m *BoundingBoxMetricsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingBoxMetricsEntry.Unmarshal(m, b)
}
func (m *BoundingBoxMetricsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingBoxMetricsEntry.Marshal(b, m, deterministic)
}
func (m *BoundingBoxMetricsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingBoxMetricsEntry.Merge(m, src)
}
func (m *BoundingBoxMetricsEntry) XXX_Size() int {
	return xxx_messageInfo_BoundingBoxMetricsEntry.Size(m)
}
func (m *BoundingBoxMetricsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingBoxMetricsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingBoxMetricsEntry proto.InternalMessageInfo

func (m *BoundingBoxMetricsEntry) GetIouThreshold() float32 {
	if m != nil {
		return m.IouThreshold
	}
	return 0
}

func (m *BoundingBoxMetricsEntry) GetMeanAveragePrecision() float32 {
	if m != nil {
		return m.MeanAveragePrecision
	}
	return 0
}

func (m *BoundingBoxMetricsEntry) GetConfidenceMetricsEntries() []*BoundingBoxMetricsEntry_ConfidenceMetricsEntry {
	if m != nil {
		return m.ConfidenceMetricsEntries
	}
	return nil
}

// Metrics for a single confidence threshold.
type BoundingBoxMetricsEntry_ConfidenceMetricsEntry struct {
	// Output only. The confidence threshold value used to compute the metrics.
	ConfidenceThreshold float32 `protobuf:"fixed32,1,opt,name=confidence_threshold,json=confidenceThreshold,proto3" json:"confidence_threshold,omitempty"`
	// Output only. Recall under the given confidence threshold.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. Precision under the given confidence threshold.
	Precision float32 `protobuf:"fixed32,3,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score              float32  `protobuf:"fixed32,4,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) Reset() {
	*m = BoundingBoxMetricsEntry_ConfidenceMetricsEntry{}
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) String() string {
	return proto.CompactTextString(m)
}
func (*BoundingBoxMetricsEntry_ConfidenceMetricsEntry) ProtoMessage() {}
func (*BoundingBoxMetricsEntry_ConfidenceMetricsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{2, 0}
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.Unmarshal(m, b)
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.Marshal(b, m, deterministic)
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.Merge(m, src)
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_Size() int {
	return xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.Size(m)
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry proto.InternalMessageInfo

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetConfidenceThreshold() float32 {
	if m != nil {
		return m.ConfidenceThreshold
	}
	return 0
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetRecall() float32 {
	if m != nil {
		return m.Recall
	}
	return 0
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetPrecision() float32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetF1Score() float32 {
	if m != nil {
		return m.F1Score
	}
	return 0
}

// Model evaluation metrics for image object detection problems.
// Evaluates prediction quality of labeled bounding boxes.
type ImageObjectDetectionEvaluationMetrics struct {
	// Output only. The total number of bounding boxes (i.e. summed over all
	// images) the ground truth used to create this evaluation had.
	EvaluatedBoundingBoxCount int32 `protobuf:"varint,1,opt,name=evaluated_bounding_box_count,json=evaluatedBoundingBoxCount,proto3" json:"evaluated_bounding_box_count,omitempty"`
	// Output only. The bounding boxes match metrics for each
	// Intersection-over-union threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// and each label confidence threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// pair.
	BoundingBoxMetricsEntries []*BoundingBoxMetricsEntry `protobuf:"bytes,2,rep,name=bounding_box_metrics_entries,json=boundingBoxMetricsEntries,proto3" json:"bounding_box_metrics_entries,omitempty"`
	// Output only. The single metric for bounding boxes evaluation:
	// the mean_average_precision averaged over all bounding_box_metrics_entries.
	BoundingBoxMeanAveragePrecision float32  `protobuf:"fixed32,3,opt,name=bounding_box_mean_average_precision,json=boundingBoxMeanAveragePrecision,proto3" json:"bounding_box_mean_average_precision,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *ImageObjectDetectionEvaluationMetrics) Reset()         { *m = ImageObjectDetectionEvaluationMetrics{} }
func (m *ImageObjectDetectionEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*ImageObjectDetectionEvaluationMetrics) ProtoMessage()    {}
func (*ImageObjectDetectionEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{3}
}

func (m *ImageObjectDetectionEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.Unmarshal(m, b)
}
func (m *ImageObjectDetectionEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *ImageObjectDetectionEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.Merge(m, src)
}
func (m *ImageObjectDetectionEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.Size(m)
}
func (m *ImageObjectDetectionEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ImageObjectDetectionEvaluationMetrics proto.InternalMessageInfo

func (m *ImageObjectDetectionEvaluationMetrics) GetEvaluatedBoundingBoxCount() int32 {
	if m != nil {
		return m.EvaluatedBoundingBoxCount
	}
	return 0
}

func (m *ImageObjectDetectionEvaluationMetrics) GetBoundingBoxMetricsEntries() []*BoundingBoxMetricsEntry {
	if m != nil {
		return m.BoundingBoxMetricsEntries
	}
	return nil
}

func (m *ImageObjectDetectionEvaluationMetrics) GetBoundingBoxMeanAveragePrecision() float32 {
	if m != nil {
		return m.BoundingBoxMeanAveragePrecision
	}
	return 0
}

// Model evaluation metrics for video object tracking problems.
// Evaluates prediction quality of both labeled bounding boxes and labeled
// tracks (i.e. series of bounding boxes sharing same label and instance ID).
type VideoObjectTrackingEvaluationMetrics struct {
	// Output only. The number of video frames used to create this evaluation.
	EvaluatedFrameCount int32 `protobuf:"varint,1,opt,name=evaluated_frame_count,json=evaluatedFrameCount,proto3" json:"evaluated_frame_count,omitempty"`
	// Output only. The total number of bounding boxes (i.e. summed over all
	// frames) the ground truth used to create this evaluation had.
	EvaluatedBoundingBoxCount int32 `protobuf:"varint,2,opt,name=evaluated_bounding_box_count,json=evaluatedBoundingBoxCount,proto3" json:"evaluated_bounding_box_count,omitempty"`
	// Output only. The bounding boxes match metrics for each
	// Intersection-over-union threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// and each label confidence threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// pair.
	BoundingBoxMetricsEntries []*BoundingBoxMetricsEntry `protobuf:"bytes,4,rep,name=bounding_box_metrics_entries,json=boundingBoxMetricsEntries,proto3" json:"bounding_box_metrics_entries,omitempty"`
	// Output only. The single metric for bounding boxes evaluation:
	// the mean_average_precision averaged over all bounding_box_metrics_entries.
	BoundingBoxMeanAveragePrecision float32  `protobuf:"fixed32,6,opt,name=bounding_box_mean_average_precision,json=boundingBoxMeanAveragePrecision,proto3" json:"bounding_box_mean_average_precision,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *VideoObjectTrackingEvaluationMetrics) Reset()         { *m = VideoObjectTrackingEvaluationMetrics{} }
func (m *VideoObjectTrackingEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*VideoObjectTrackingEvaluationMetrics) ProtoMessage()    {}
func (*VideoObjectTrackingEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{4}
}

func (m *VideoObjectTrackingEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.Unmarshal(m, b)
}
func (m *VideoObjectTrackingEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *VideoObjectTrackingEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.Merge(m, src)
}
func (m *VideoObjectTrackingEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.Size(m)
}
func (m *VideoObjectTrackingEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_VideoObjectTrackingEvaluationMetrics proto.InternalMessageInfo

func (m *VideoObjectTrackingEvaluationMetrics) GetEvaluatedFrameCount() int32 {
	if m != nil {
		return m.EvaluatedFrameCount
	}
	return 0
}

func (m *VideoObjectTrackingEvaluationMetrics) GetEvaluatedBoundingBoxCount() int32 {
	if m != nil {
		return m.EvaluatedBoundingBoxCount
	}
	return 0
}

func (m *VideoObjectTrackingEvaluationMetrics) GetBoundingBoxMetricsEntries() []*BoundingBoxMetricsEntry {
	if m != nil {
		return m.BoundingBoxMetricsEntries
	}
	return nil
}

func (m *VideoObjectTrackingEvaluationMetrics) GetBoundingBoxMeanAveragePrecision() float32 {
	if m != nil {
		return m.BoundingBoxMeanAveragePrecision
	}
	return 0
}

func init() {
	proto.RegisterType((*ImageObjectDetectionAnnotation)(nil), "google.cloud.automl.v1beta1.ImageObjectDetectionAnnotation")
	proto.RegisterType((*VideoObjectTrackingAnnotation)(nil), "google.cloud.automl.v1beta1.VideoObjectTrackingAnnotation")
	proto.RegisterType((*BoundingBoxMetricsEntry)(nil), "google.cloud.automl.v1beta1.BoundingBoxMetricsEntry")
	proto.RegisterType((*BoundingBoxMetricsEntry_ConfidenceMetricsEntry)(nil), "google.cloud.automl.v1beta1.BoundingBoxMetricsEntry.ConfidenceMetricsEntry")
	proto.RegisterType((*ImageObjectDetectionEvaluationMetrics)(nil), "google.cloud.automl.v1beta1.ImageObjectDetectionEvaluationMetrics")
	proto.RegisterType((*VideoObjectTrackingEvaluationMetrics)(nil), "google.cloud.automl.v1beta1.VideoObjectTrackingEvaluationMetrics")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/detection.proto", fileDescriptor_85c92c72aeb2dfbd)
}

var fileDescriptor_85c92c72aeb2dfbd = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x95, 0xc1, 0x4e, 0xdb, 0x4a,
	0x14, 0x86, 0x65, 0x07, 0xb8, 0x97, 0x09, 0x77, 0x63, 0xb8, 0xdc, 0x24, 0xe4, 0x02, 0x82, 0x56,
	0xa2, 0xad, 0x64, 0x2b, 0x29, 0x2b, 0x77, 0x51, 0x25, 0x40, 0x2b, 0x54, 0x10, 0xc8, 0x45, 0x2c,
	0x2a, 0x24, 0x6b, 0x6c, 0x1f, 0x9b, 0x69, 0xed, 0x99, 0x68, 0x3c, 0x46, 0xb0, 0xef, 0xa6, 0xef,
	0x50, 0xa9, 0xab, 0x2e, 0xfb, 0x22, 0x7d, 0x0a, 0xd6, 0x7d, 0x8a, 0x2a, 0x33, 0x76, 0x6c, 0x82,
	0x45, 0x51, 0x5b, 0xa9, 0xcb, 0x39, 0xff, 0x3f, 0x33, 0xff, 0xf9, 0xce, 0x38, 0x41, 0x4f, 0x22,
	0xc6, 0xa2, 0x18, 0x2c, 0x3f, 0x66, 0x59, 0x60, 0xe1, 0x4c, 0xb0, 0x24, 0xb6, 0x2e, 0x7a, 0x1e,
	0x08, 0xdc, 0xb3, 0x02, 0x10, 0xe0, 0x0b, 0xc2, 0xa8, 0x39, 0xe2, 0x4c, 0x30, 0x63, 0x45, 0x99,
	0x4d, 0x69, 0x36, 0x95, 0xd9, 0xcc, 0xcd, 0x9d, 0xc7, 0x77, 0x9d, 0x14, 0x01, 0x4b, 0x40, 0xf0,
	0x2b, 0x75, 0x50, 0x67, 0x35, 0xf7, 0xca, 0x95, 0x97, 0x85, 0x56, 0x90, 0x71, 0x5c, 0x5e, 0xd4,
	0xe9, 0xe6, 0x3a, 0x1e, 0x11, 0x0b, 0x53, 0xca, 0x84, 0x14, 0x53, 0xa5, 0x6e, 0xbc, 0xd7, 0xd0,
	0xea, 0x7e, 0x82, 0x23, 0x38, 0xf2, 0xde, 0x82, 0x2f, 0x76, 0x8b, 0x94, 0x83, 0x89, 0xd3, 0x38,
	0x40, 0x0b, 0x1e, 0xcb, 0x68, 0x40, 0x68, 0xe4, 0x7a, 0xec, 0xb2, 0xa5, 0xad, 0x6b, 0x5b, 0xcd,
	0xfe, 0x23, 0xf3, 0x8e, 0x06, 0xcc, 0x61, 0xbe, 0xe1, 0x98, 0xc5, 0x57, 0x4e, 0xb3, 0xd8, 0x3e,
	0x64, 0x97, 0xc6, 0x12, 0x9a, 0x4d, 0x7d, 0xc6, 0xa1, 0xa5, 0xaf, 0x6b, 0x5b, 0xba, 0xa3, 0x16,
	0x1b, 0xd7, 0x1a, 0xfa, 0xff, 0x94, 0x04, 0xc0, 0x54, 0x8c, 0x13, 0x8e, 0xfd, 0x77, 0x84, 0x46,
	0x95, 0x14, 0x6b, 0xa8, 0x49, 0x68, 0x2a, 0x30, 0xf5, 0xc1, 0x25, 0x81, 0x0c, 0x31, 0xef, 0xa0,
	0xa2, 0xb4, 0x1f, 0x18, 0x36, 0x6a, 0x0a, 0x92, 0x80, 0xcb, 0xc2, 0x30, 0x05, 0x21, 0x8f, 0x6f,
	0xf6, 0xdb, 0x45, 0xca, 0x82, 0x8e, 0xb9, 0x9b, 0xd3, 0x71, 0xd0, 0xd8, 0x7d, 0x24, 0xcd, 0xb7,
	0x5a, 0x6c, 0xfc, 0x9e, 0x16, 0x67, 0xaa, 0x2d, 0x7e, 0x6a, 0xa0, 0xff, 0x86, 0xa5, 0xeb, 0x10,
	0x04, 0x27, 0x7e, 0xba, 0x47, 0x05, 0xbf, 0x32, 0x36, 0xd1, 0x3f, 0x84, 0x65, 0xae, 0x38, 0xe7,
	0x90, 0x9e, 0xb3, 0x58, 0xb5, 0xa7, 0x3b, 0x0b, 0x84, 0x65, 0x27, 0x45, 0xcd, 0xd8, 0x46, 0xcb,
	0x09, 0x60, 0xea, 0xe2, 0x0b, 0xe0, 0x38, 0x02, 0x77, 0xc4, 0xc1, 0x27, 0x29, 0x61, 0x34, 0x47,
	0xb9, 0x34, 0x56, 0x07, 0x4a, 0x3c, 0x2e, 0x34, 0xe3, 0x83, 0x86, 0x3a, 0x3e, 0xa3, 0x21, 0x09,
	0x60, 0x8c, 0x2e, 0x51, 0xd7, 0xba, 0x40, 0x05, 0x27, 0x90, 0xb6, 0x1a, 0xeb, 0x8d, 0xad, 0x66,
	0xff, 0xd5, 0xbd, 0x3a, 0x9d, 0x4a, 0x6d, 0xee, 0x4c, 0x8e, 0xad, 0x96, 0x9d, 0x96, 0x5f, 0x57,
	0x27, 0x90, 0x76, 0x3e, 0x6a, 0x68, 0xb9, 0x7e, 0x93, 0xd1, 0x43, 0x4b, 0x95, 0x94, 0xd3, 0x20,
	0x16, 0x4b, 0xad, 0xe4, 0xb1, 0x8c, 0xe6, 0x38, 0xf8, 0x38, 0x8e, 0xf3, 0xfe, 0xf3, 0x95, 0xd1,
	0x45, 0xf3, 0x25, 0x9a, 0x86, 0x94, 0xca, 0x82, 0xd1, 0x46, 0x7f, 0x87, 0x3d, 0xb7, 0x3a, 0x9f,
	0xbf, 0xc2, 0xde, 0x6b, 0x39, 0xa1, 0x2f, 0x3a, 0x7a, 0x58, 0xf7, 0x2d, 0xec, 0x5d, 0xe0, 0x38,
	0x93, 0x8f, 0x26, 0x8f, 0x6c, 0x3c, 0x47, 0x5d, 0x50, 0x45, 0x08, 0xdc, 0xea, 0xcb, 0x71, 0x7d,
	0x96, 0x51, 0x21, 0x53, 0xcf, 0x3a, 0xed, 0x89, 0xa7, 0x42, 0x70, 0x67, 0x6c, 0x30, 0x32, 0xd4,
	0xbd, 0xb1, 0x6d, 0x7a, 0x2c, 0xba, 0x1c, 0xcb, 0xf6, 0xcf, 0x8c, 0xc5, 0x69, 0x7b, 0xb5, 0x02,
	0x81, 0xd4, 0x38, 0x40, 0x9b, 0x53, 0xd7, 0xd6, 0xbe, 0x27, 0x05, 0x6d, 0xed, 0xc6, 0x39, 0xb7,
	0x9f, 0xd6, 0xc6, 0xb5, 0x8e, 0x1e, 0xd4, 0x7c, 0xb4, 0xb7, 0x71, 0xf5, 0xd1, 0xbf, 0x25, 0xae,
	0x90, 0xe3, 0x04, 0x6e, 0x70, 0x5a, 0x9c, 0x88, 0x2f, 0xc6, 0x9a, 0x22, 0xf4, 0x23, 0xc4, 0xfa,
	0xaf, 0x22, 0x9e, 0xf9, 0xa3, 0x88, 0xe7, 0xee, 0x85, 0x78, 0xf8, 0x59, 0x43, 0x6b, 0x3e, 0x4b,
	0xee, 0x0a, 0x79, 0xac, 0xbd, 0x19, 0xe4, 0x72, 0xc4, 0x62, 0x4c, 0x23, 0x93, 0xf1, 0xc8, 0x8a,
	0x80, 0xca, 0x9f, 0x3c, 0x4b, 0x49, 0x78, 0x44, 0xd2, 0xda, 0x7f, 0x93, 0x67, 0x6a, 0xf9, 0x55,
	0x5f, 0x79, 0x29, 0x8d, 0x67, 0x3b, 0x63, 0xd3, 0xd9, 0x20, 0x13, 0xec, 0x30, 0x3e, 0x3b, 0x55,
	0xa6, 0x6f, 0xfa, 0xaa, 0x52, 0x6d, 0x5b, 0xca, 0xb6, 0x2d, 0xf5, 0x03, 0xdb, 0xce, 0x0d, 0xde,
	0x9c, 0xbc, 0xec, 0xe9, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x5c, 0xae, 0x18, 0x03, 0x07,
	0x00, 0x00,
}