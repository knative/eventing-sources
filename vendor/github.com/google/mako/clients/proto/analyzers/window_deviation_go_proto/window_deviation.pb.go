// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clients/proto/analyzers/window_deviation.proto

package mako_window_deviation

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	mako_go_proto "github.com/google/mako/spec/proto/mako_go_proto"
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

type ToleranceCheck_DirectionBias int32

const (
	ToleranceCheck_NO_BIAS         ToleranceCheck_DirectionBias = 0
	ToleranceCheck_IGNORE_INCREASE ToleranceCheck_DirectionBias = 1
	ToleranceCheck_IGNORE_DECREASE ToleranceCheck_DirectionBias = 2
)

var ToleranceCheck_DirectionBias_name = map[int32]string{
	0: "NO_BIAS",
	1: "IGNORE_INCREASE",
	2: "IGNORE_DECREASE",
}

var ToleranceCheck_DirectionBias_value = map[string]int32{
	"NO_BIAS":         0,
	"IGNORE_INCREASE": 1,
	"IGNORE_DECREASE": 2,
}

func (x ToleranceCheck_DirectionBias) Enum() *ToleranceCheck_DirectionBias {
	p := new(ToleranceCheck_DirectionBias)
	*p = x
	return p
}

func (x ToleranceCheck_DirectionBias) String() string {
	return proto.EnumName(ToleranceCheck_DirectionBias_name, int32(x))
}

func (x *ToleranceCheck_DirectionBias) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ToleranceCheck_DirectionBias_value, data, "ToleranceCheck_DirectionBias")
	if err != nil {
		return err
	}
	*x = ToleranceCheck_DirectionBias(value)
	return nil
}

func (ToleranceCheck_DirectionBias) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{1, 0}
}

type WindowDeviationOutput_ToleranceCheckOutput_CheckResult int32

const (
	WindowDeviationOutput_ToleranceCheckOutput_UNDEFINED WindowDeviationOutput_ToleranceCheckOutput_CheckResult = 0
	WindowDeviationOutput_ToleranceCheckOutput_REGRESSED WindowDeviationOutput_ToleranceCheckOutput_CheckResult = 1
	WindowDeviationOutput_ToleranceCheckOutput_SKIPPED   WindowDeviationOutput_ToleranceCheckOutput_CheckResult = 2
	WindowDeviationOutput_ToleranceCheckOutput_PASSED    WindowDeviationOutput_ToleranceCheckOutput_CheckResult = 3
)

var WindowDeviationOutput_ToleranceCheckOutput_CheckResult_name = map[int32]string{
	0: "UNDEFINED",
	1: "REGRESSED",
	2: "SKIPPED",
	3: "PASSED",
}

var WindowDeviationOutput_ToleranceCheckOutput_CheckResult_value = map[string]int32{
	"UNDEFINED": 0,
	"REGRESSED": 1,
	"SKIPPED":   2,
	"PASSED":    3,
}

func (x WindowDeviationOutput_ToleranceCheckOutput_CheckResult) Enum() *WindowDeviationOutput_ToleranceCheckOutput_CheckResult {
	p := new(WindowDeviationOutput_ToleranceCheckOutput_CheckResult)
	*p = x
	return p
}

func (x WindowDeviationOutput_ToleranceCheckOutput_CheckResult) String() string {
	return proto.EnumName(WindowDeviationOutput_ToleranceCheckOutput_CheckResult_name, int32(x))
}

func (x *WindowDeviationOutput_ToleranceCheckOutput_CheckResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WindowDeviationOutput_ToleranceCheckOutput_CheckResult_value, data, "WindowDeviationOutput_ToleranceCheckOutput_CheckResult")
	if err != nil {
		return err
	}
	*x = WindowDeviationOutput_ToleranceCheckOutput_CheckResult(value)
	return nil
}

func (WindowDeviationOutput_ToleranceCheckOutput_CheckResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{4, 0, 0}
}

type WindowDeviationInput struct {
	RunInfoQueryList     []*mako_go_proto.RunInfoQuery `protobuf:"bytes,1,rep,name=run_info_query_list,json=runInfoQueryList" json:"run_info_query_list,omitempty"`
	ToleranceCheckList   []*ToleranceCheck             `protobuf:"bytes,2,rep,name=tolerance_check_list,json=toleranceCheckList" json:"tolerance_check_list,omitempty"`
	Name                 *string                       `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *WindowDeviationInput) Reset()         { *m = WindowDeviationInput{} }
func (m *WindowDeviationInput) String() string { return proto.CompactTextString(m) }
func (*WindowDeviationInput) ProtoMessage()    {}
func (*WindowDeviationInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{0}
}

func (m *WindowDeviationInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WindowDeviationInput.Unmarshal(m, b)
}
func (m *WindowDeviationInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WindowDeviationInput.Marshal(b, m, deterministic)
}
func (m *WindowDeviationInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WindowDeviationInput.Merge(m, src)
}
func (m *WindowDeviationInput) XXX_Size() int {
	return xxx_messageInfo_WindowDeviationInput.Size(m)
}
func (m *WindowDeviationInput) XXX_DiscardUnknown() {
	xxx_messageInfo_WindowDeviationInput.DiscardUnknown(m)
}

var xxx_messageInfo_WindowDeviationInput proto.InternalMessageInfo

func (m *WindowDeviationInput) GetRunInfoQueryList() []*mako_go_proto.RunInfoQuery {
	if m != nil {
		return m.RunInfoQueryList
	}
	return nil
}

func (m *WindowDeviationInput) GetToleranceCheckList() []*ToleranceCheck {
	if m != nil {
		return m.ToleranceCheckList
	}
	return nil
}

func (m *WindowDeviationInput) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type ToleranceCheck struct {
	DataFilter                  *mako_go_proto.DataFilter     `protobuf:"bytes,1,opt,name=data_filter,json=dataFilter" json:"data_filter,omitempty"`
	RecentWindowSize            *int32                        `protobuf:"varint,2,opt,name=recent_window_size,json=recentWindowSize,def=1" json:"recent_window_size,omitempty"`
	MinimumHistoricalWindowSize *int32                        `protobuf:"varint,6,opt,name=minimum_historical_window_size,json=minimumHistoricalWindowSize,def=3" json:"minimum_historical_window_size,omitempty"`
	DirectionBias               *ToleranceCheck_DirectionBias `protobuf:"varint,3,opt,name=direction_bias,json=directionBias,enum=mako.window_deviation.ToleranceCheck_DirectionBias,def=0" json:"direction_bias,omitempty"`
	MeanToleranceParamsList     []*MeanToleranceParams        `protobuf:"bytes,4,rep,name=mean_tolerance_params_list,json=meanToleranceParamsList" json:"mean_tolerance_params_list,omitempty"`
	MedianToleranceParamsList   []*MedianToleranceParams      `protobuf:"bytes,5,rep,name=median_tolerance_params_list,json=medianToleranceParamsList" json:"median_tolerance_params_list,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                      `json:"-"`
	XXX_unrecognized            []byte                        `json:"-"`
	XXX_sizecache               int32                         `json:"-"`
}

func (m *ToleranceCheck) Reset()         { *m = ToleranceCheck{} }
func (m *ToleranceCheck) String() string { return proto.CompactTextString(m) }
func (*ToleranceCheck) ProtoMessage()    {}
func (*ToleranceCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{1}
}

func (m *ToleranceCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToleranceCheck.Unmarshal(m, b)
}
func (m *ToleranceCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToleranceCheck.Marshal(b, m, deterministic)
}
func (m *ToleranceCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToleranceCheck.Merge(m, src)
}
func (m *ToleranceCheck) XXX_Size() int {
	return xxx_messageInfo_ToleranceCheck.Size(m)
}
func (m *ToleranceCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ToleranceCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ToleranceCheck proto.InternalMessageInfo

const Default_ToleranceCheck_RecentWindowSize int32 = 1
const Default_ToleranceCheck_MinimumHistoricalWindowSize int32 = 3
const Default_ToleranceCheck_DirectionBias ToleranceCheck_DirectionBias = ToleranceCheck_NO_BIAS

func (m *ToleranceCheck) GetDataFilter() *mako_go_proto.DataFilter {
	if m != nil {
		return m.DataFilter
	}
	return nil
}

func (m *ToleranceCheck) GetRecentWindowSize() int32 {
	if m != nil && m.RecentWindowSize != nil {
		return *m.RecentWindowSize
	}
	return Default_ToleranceCheck_RecentWindowSize
}

func (m *ToleranceCheck) GetMinimumHistoricalWindowSize() int32 {
	if m != nil && m.MinimumHistoricalWindowSize != nil {
		return *m.MinimumHistoricalWindowSize
	}
	return Default_ToleranceCheck_MinimumHistoricalWindowSize
}

func (m *ToleranceCheck) GetDirectionBias() ToleranceCheck_DirectionBias {
	if m != nil && m.DirectionBias != nil {
		return *m.DirectionBias
	}
	return Default_ToleranceCheck_DirectionBias
}

func (m *ToleranceCheck) GetMeanToleranceParamsList() []*MeanToleranceParams {
	if m != nil {
		return m.MeanToleranceParamsList
	}
	return nil
}

func (m *ToleranceCheck) GetMedianToleranceParamsList() []*MedianToleranceParams {
	if m != nil {
		return m.MedianToleranceParamsList
	}
	return nil
}

type MeanToleranceParams struct {
	ConstTerm            *float64 `protobuf:"fixed64,1,opt,name=const_term,json=constTerm,def=0" json:"const_term,omitempty"`
	MeanCoeff            *float64 `protobuf:"fixed64,2,opt,name=mean_coeff,json=meanCoeff,def=0" json:"mean_coeff,omitempty"`
	StddevCoeff          *float64 `protobuf:"fixed64,3,opt,name=stddev_coeff,json=stddevCoeff,def=0" json:"stddev_coeff,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeanToleranceParams) Reset()         { *m = MeanToleranceParams{} }
func (m *MeanToleranceParams) String() string { return proto.CompactTextString(m) }
func (*MeanToleranceParams) ProtoMessage()    {}
func (*MeanToleranceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{2}
}

func (m *MeanToleranceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeanToleranceParams.Unmarshal(m, b)
}
func (m *MeanToleranceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeanToleranceParams.Marshal(b, m, deterministic)
}
func (m *MeanToleranceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeanToleranceParams.Merge(m, src)
}
func (m *MeanToleranceParams) XXX_Size() int {
	return xxx_messageInfo_MeanToleranceParams.Size(m)
}
func (m *MeanToleranceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MeanToleranceParams.DiscardUnknown(m)
}

var xxx_messageInfo_MeanToleranceParams proto.InternalMessageInfo

const Default_MeanToleranceParams_ConstTerm float64 = 0
const Default_MeanToleranceParams_MeanCoeff float64 = 0
const Default_MeanToleranceParams_StddevCoeff float64 = 0

func (m *MeanToleranceParams) GetConstTerm() float64 {
	if m != nil && m.ConstTerm != nil {
		return *m.ConstTerm
	}
	return Default_MeanToleranceParams_ConstTerm
}

func (m *MeanToleranceParams) GetMeanCoeff() float64 {
	if m != nil && m.MeanCoeff != nil {
		return *m.MeanCoeff
	}
	return Default_MeanToleranceParams_MeanCoeff
}

func (m *MeanToleranceParams) GetStddevCoeff() float64 {
	if m != nil && m.StddevCoeff != nil {
		return *m.StddevCoeff
	}
	return Default_MeanToleranceParams_StddevCoeff
}

type MedianToleranceParams struct {
	ConstTerm            *float64 `protobuf:"fixed64,1,opt,name=const_term,json=constTerm,def=0" json:"const_term,omitempty"`
	MedianCoeff          *float64 `protobuf:"fixed64,2,opt,name=median_coeff,json=medianCoeff,def=0" json:"median_coeff,omitempty"`
	MadCoeff             *float64 `protobuf:"fixed64,3,opt,name=mad_coeff,json=madCoeff,def=0" json:"mad_coeff,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MedianToleranceParams) Reset()         { *m = MedianToleranceParams{} }
func (m *MedianToleranceParams) String() string { return proto.CompactTextString(m) }
func (*MedianToleranceParams) ProtoMessage()    {}
func (*MedianToleranceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{3}
}

func (m *MedianToleranceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MedianToleranceParams.Unmarshal(m, b)
}
func (m *MedianToleranceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MedianToleranceParams.Marshal(b, m, deterministic)
}
func (m *MedianToleranceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MedianToleranceParams.Merge(m, src)
}
func (m *MedianToleranceParams) XXX_Size() int {
	return xxx_messageInfo_MedianToleranceParams.Size(m)
}
func (m *MedianToleranceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MedianToleranceParams.DiscardUnknown(m)
}

var xxx_messageInfo_MedianToleranceParams proto.InternalMessageInfo

const Default_MedianToleranceParams_ConstTerm float64 = 0
const Default_MedianToleranceParams_MedianCoeff float64 = 0
const Default_MedianToleranceParams_MadCoeff float64 = 0

func (m *MedianToleranceParams) GetConstTerm() float64 {
	if m != nil && m.ConstTerm != nil {
		return *m.ConstTerm
	}
	return Default_MedianToleranceParams_ConstTerm
}

func (m *MedianToleranceParams) GetMedianCoeff() float64 {
	if m != nil && m.MedianCoeff != nil {
		return *m.MedianCoeff
	}
	return Default_MedianToleranceParams_MedianCoeff
}

func (m *MedianToleranceParams) GetMadCoeff() float64 {
	if m != nil && m.MadCoeff != nil {
		return *m.MadCoeff
	}
	return Default_MedianToleranceParams_MadCoeff
}

type WindowDeviationOutput struct {
	OutputMessage               *string                                       `protobuf:"bytes,1,opt,name=output_message,json=outputMessage" json:"output_message,omitempty"`
	Checks                      []*WindowDeviationOutput_ToleranceCheckOutput `protobuf:"bytes,3,rep,name=checks" json:"checks,omitempty"`
	ChecksSkippedForMissingData []*ToleranceCheck                             `protobuf:"bytes,2,rep,name=checks_skipped_for_missing_data,json=checksSkippedForMissingData" json:"checks_skipped_for_missing_data,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                                      `json:"-"`
	XXX_unrecognized            []byte                                        `json:"-"`
	XXX_sizecache               int32                                         `json:"-"`
}

func (m *WindowDeviationOutput) Reset()         { *m = WindowDeviationOutput{} }
func (m *WindowDeviationOutput) String() string { return proto.CompactTextString(m) }
func (*WindowDeviationOutput) ProtoMessage()    {}
func (*WindowDeviationOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{4}
}

func (m *WindowDeviationOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WindowDeviationOutput.Unmarshal(m, b)
}
func (m *WindowDeviationOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WindowDeviationOutput.Marshal(b, m, deterministic)
}
func (m *WindowDeviationOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WindowDeviationOutput.Merge(m, src)
}
func (m *WindowDeviationOutput) XXX_Size() int {
	return xxx_messageInfo_WindowDeviationOutput.Size(m)
}
func (m *WindowDeviationOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_WindowDeviationOutput.DiscardUnknown(m)
}

var xxx_messageInfo_WindowDeviationOutput proto.InternalMessageInfo

func (m *WindowDeviationOutput) GetOutputMessage() string {
	if m != nil && m.OutputMessage != nil {
		return *m.OutputMessage
	}
	return ""
}

func (m *WindowDeviationOutput) GetChecks() []*WindowDeviationOutput_ToleranceCheckOutput {
	if m != nil {
		return m.Checks
	}
	return nil
}

func (m *WindowDeviationOutput) GetChecksSkippedForMissingData() []*ToleranceCheck {
	if m != nil {
		return m.ChecksSkippedForMissingData
	}
	return nil
}

type WindowDeviationOutput_ToleranceCheckOutput struct {
	ToleranceCheck                 *ToleranceCheck                                         `protobuf:"bytes,1,opt,name=tolerance_check,json=toleranceCheck" json:"tolerance_check,omitempty"`
	Result                         *WindowDeviationOutput_ToleranceCheckOutput_CheckResult `protobuf:"varint,2,opt,name=result,enum=mako.window_deviation.WindowDeviationOutput_ToleranceCheckOutput_CheckResult,def=0" json:"result,omitempty"`
	MetricLabel                    *string                                                 `protobuf:"bytes,3,opt,name=metric_label,json=metricLabel" json:"metric_label,omitempty"`
	Stats                          *ToleranceCheckStats                                    `protobuf:"bytes,4,opt,name=stats" json:"stats,omitempty"`
	HistoricalWindowMinTimestampMs *float64                                                `protobuf:"fixed64,5,opt,name=historical_window_min_timestamp_ms,json=historicalWindowMinTimestampMs" json:"historical_window_min_timestamp_ms,omitempty"`
	HistoricalWindowMaxTimestampMs *float64                                                `protobuf:"fixed64,6,opt,name=historical_window_max_timestamp_ms,json=historicalWindowMaxTimestampMs" json:"historical_window_max_timestamp_ms,omitempty"`
	HistoricalWindowMinBuildId     *float64                                                `protobuf:"fixed64,7,opt,name=historical_window_min_build_id,json=historicalWindowMinBuildId" json:"historical_window_min_build_id,omitempty"`
	HistoricalWindowMaxBuildId     *float64                                                `protobuf:"fixed64,8,opt,name=historical_window_max_build_id,json=historicalWindowMaxBuildId" json:"historical_window_max_build_id,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}                                                `json:"-"`
	XXX_unrecognized               []byte                                                  `json:"-"`
	XXX_sizecache                  int32                                                   `json:"-"`
}

func (m *WindowDeviationOutput_ToleranceCheckOutput) Reset() {
	*m = WindowDeviationOutput_ToleranceCheckOutput{}
}
func (m *WindowDeviationOutput_ToleranceCheckOutput) String() string {
	return proto.CompactTextString(m)
}
func (*WindowDeviationOutput_ToleranceCheckOutput) ProtoMessage() {}
func (*WindowDeviationOutput_ToleranceCheckOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{4, 0}
}

func (m *WindowDeviationOutput_ToleranceCheckOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WindowDeviationOutput_ToleranceCheckOutput.Unmarshal(m, b)
}
func (m *WindowDeviationOutput_ToleranceCheckOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WindowDeviationOutput_ToleranceCheckOutput.Marshal(b, m, deterministic)
}
func (m *WindowDeviationOutput_ToleranceCheckOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WindowDeviationOutput_ToleranceCheckOutput.Merge(m, src)
}
func (m *WindowDeviationOutput_ToleranceCheckOutput) XXX_Size() int {
	return xxx_messageInfo_WindowDeviationOutput_ToleranceCheckOutput.Size(m)
}
func (m *WindowDeviationOutput_ToleranceCheckOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_WindowDeviationOutput_ToleranceCheckOutput.DiscardUnknown(m)
}

var xxx_messageInfo_WindowDeviationOutput_ToleranceCheckOutput proto.InternalMessageInfo

const Default_WindowDeviationOutput_ToleranceCheckOutput_Result WindowDeviationOutput_ToleranceCheckOutput_CheckResult = WindowDeviationOutput_ToleranceCheckOutput_UNDEFINED

func (m *WindowDeviationOutput_ToleranceCheckOutput) GetToleranceCheck() *ToleranceCheck {
	if m != nil {
		return m.ToleranceCheck
	}
	return nil
}

func (m *WindowDeviationOutput_ToleranceCheckOutput) GetResult() WindowDeviationOutput_ToleranceCheckOutput_CheckResult {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return Default_WindowDeviationOutput_ToleranceCheckOutput_Result
}

func (m *WindowDeviationOutput_ToleranceCheckOutput) GetMetricLabel() string {
	if m != nil && m.MetricLabel != nil {
		return *m.MetricLabel
	}
	return ""
}

func (m *WindowDeviationOutput_ToleranceCheckOutput) GetStats() *ToleranceCheckStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *WindowDeviationOutput_ToleranceCheckOutput) GetHistoricalWindowMinTimestampMs() float64 {
	if m != nil && m.HistoricalWindowMinTimestampMs != nil {
		return *m.HistoricalWindowMinTimestampMs
	}
	return 0
}

func (m *WindowDeviationOutput_ToleranceCheckOutput) GetHistoricalWindowMaxTimestampMs() float64 {
	if m != nil && m.HistoricalWindowMaxTimestampMs != nil {
		return *m.HistoricalWindowMaxTimestampMs
	}
	return 0
}

func (m *WindowDeviationOutput_ToleranceCheckOutput) GetHistoricalWindowMinBuildId() float64 {
	if m != nil && m.HistoricalWindowMinBuildId != nil {
		return *m.HistoricalWindowMinBuildId
	}
	return 0
}

func (m *WindowDeviationOutput_ToleranceCheckOutput) GetHistoricalWindowMaxBuildId() float64 {
	if m != nil && m.HistoricalWindowMaxBuildId != nil {
		return *m.HistoricalWindowMaxBuildId
	}
	return 0
}

type ToleranceCheckStats struct {
	HistoricDataLength         *int32                        `protobuf:"varint,1,opt,name=historic_data_length,json=historicDataLength" json:"historic_data_length,omitempty"`
	RecentDataLength           *int32                        `protobuf:"varint,2,opt,name=recent_data_length,json=recentDataLength" json:"recent_data_length,omitempty"`
	HistoricMean               *float64                      `protobuf:"fixed64,3,opt,name=historic_mean,json=historicMean" json:"historic_mean,omitempty"`
	HistoricMedian             *float64                      `protobuf:"fixed64,4,opt,name=historic_median,json=historicMedian" json:"historic_median,omitempty"`
	HistoricStddev             *float64                      `protobuf:"fixed64,5,opt,name=historic_stddev,json=historicStddev" json:"historic_stddev,omitempty"`
	HistoricMad                *float64                      `protobuf:"fixed64,6,opt,name=historic_mad,json=historicMad" json:"historic_mad,omitempty"`
	RecentMean                 *float64                      `protobuf:"fixed64,7,opt,name=recent_mean,json=recentMean" json:"recent_mean,omitempty"`
	RecentMedian               *float64                      `protobuf:"fixed64,8,opt,name=recent_median,json=recentMedian" json:"recent_median,omitempty"`
	DeltaMean                  *float64                      `protobuf:"fixed64,9,opt,name=delta_mean,json=deltaMean" json:"delta_mean,omitempty"`
	DeltaMedian                *float64                      `protobuf:"fixed64,10,opt,name=delta_median,json=deltaMedian" json:"delta_median,omitempty"`
	MeanToleranceCheckResult   []*MeanToleranceCheckResult   `protobuf:"bytes,11,rep,name=mean_tolerance_check_result,json=meanToleranceCheckResult" json:"mean_tolerance_check_result,omitempty"`
	MedianToleranceCheckResult []*MedianToleranceCheckResult `protobuf:"bytes,12,rep,name=median_tolerance_check_result,json=medianToleranceCheckResult" json:"median_tolerance_check_result,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                      `json:"-"`
	XXX_unrecognized           []byte                        `json:"-"`
	XXX_sizecache              int32                         `json:"-"`
}

func (m *ToleranceCheckStats) Reset()         { *m = ToleranceCheckStats{} }
func (m *ToleranceCheckStats) String() string { return proto.CompactTextString(m) }
func (*ToleranceCheckStats) ProtoMessage()    {}
func (*ToleranceCheckStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{5}
}

func (m *ToleranceCheckStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToleranceCheckStats.Unmarshal(m, b)
}
func (m *ToleranceCheckStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToleranceCheckStats.Marshal(b, m, deterministic)
}
func (m *ToleranceCheckStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToleranceCheckStats.Merge(m, src)
}
func (m *ToleranceCheckStats) XXX_Size() int {
	return xxx_messageInfo_ToleranceCheckStats.Size(m)
}
func (m *ToleranceCheckStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ToleranceCheckStats.DiscardUnknown(m)
}

var xxx_messageInfo_ToleranceCheckStats proto.InternalMessageInfo

func (m *ToleranceCheckStats) GetHistoricDataLength() int32 {
	if m != nil && m.HistoricDataLength != nil {
		return *m.HistoricDataLength
	}
	return 0
}

func (m *ToleranceCheckStats) GetRecentDataLength() int32 {
	if m != nil && m.RecentDataLength != nil {
		return *m.RecentDataLength
	}
	return 0
}

func (m *ToleranceCheckStats) GetHistoricMean() float64 {
	if m != nil && m.HistoricMean != nil {
		return *m.HistoricMean
	}
	return 0
}

func (m *ToleranceCheckStats) GetHistoricMedian() float64 {
	if m != nil && m.HistoricMedian != nil {
		return *m.HistoricMedian
	}
	return 0
}

func (m *ToleranceCheckStats) GetHistoricStddev() float64 {
	if m != nil && m.HistoricStddev != nil {
		return *m.HistoricStddev
	}
	return 0
}

func (m *ToleranceCheckStats) GetHistoricMad() float64 {
	if m != nil && m.HistoricMad != nil {
		return *m.HistoricMad
	}
	return 0
}

func (m *ToleranceCheckStats) GetRecentMean() float64 {
	if m != nil && m.RecentMean != nil {
		return *m.RecentMean
	}
	return 0
}

func (m *ToleranceCheckStats) GetRecentMedian() float64 {
	if m != nil && m.RecentMedian != nil {
		return *m.RecentMedian
	}
	return 0
}

func (m *ToleranceCheckStats) GetDeltaMean() float64 {
	if m != nil && m.DeltaMean != nil {
		return *m.DeltaMean
	}
	return 0
}

func (m *ToleranceCheckStats) GetDeltaMedian() float64 {
	if m != nil && m.DeltaMedian != nil {
		return *m.DeltaMedian
	}
	return 0
}

func (m *ToleranceCheckStats) GetMeanToleranceCheckResult() []*MeanToleranceCheckResult {
	if m != nil {
		return m.MeanToleranceCheckResult
	}
	return nil
}

func (m *ToleranceCheckStats) GetMedianToleranceCheckResult() []*MedianToleranceCheckResult {
	if m != nil {
		return m.MedianToleranceCheckResult
	}
	return nil
}

type MeanToleranceCheckResult struct {
	Params               *MeanToleranceParams `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	Tolerance            *float64             `protobuf:"fixed64,2,opt,name=tolerance" json:"tolerance,omitempty"`
	IsRegressed          *bool                `protobuf:"varint,3,opt,name=is_regressed,json=isRegressed" json:"is_regressed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MeanToleranceCheckResult) Reset()         { *m = MeanToleranceCheckResult{} }
func (m *MeanToleranceCheckResult) String() string { return proto.CompactTextString(m) }
func (*MeanToleranceCheckResult) ProtoMessage()    {}
func (*MeanToleranceCheckResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{6}
}

func (m *MeanToleranceCheckResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeanToleranceCheckResult.Unmarshal(m, b)
}
func (m *MeanToleranceCheckResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeanToleranceCheckResult.Marshal(b, m, deterministic)
}
func (m *MeanToleranceCheckResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeanToleranceCheckResult.Merge(m, src)
}
func (m *MeanToleranceCheckResult) XXX_Size() int {
	return xxx_messageInfo_MeanToleranceCheckResult.Size(m)
}
func (m *MeanToleranceCheckResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MeanToleranceCheckResult.DiscardUnknown(m)
}

var xxx_messageInfo_MeanToleranceCheckResult proto.InternalMessageInfo

func (m *MeanToleranceCheckResult) GetParams() *MeanToleranceParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *MeanToleranceCheckResult) GetTolerance() float64 {
	if m != nil && m.Tolerance != nil {
		return *m.Tolerance
	}
	return 0
}

func (m *MeanToleranceCheckResult) GetIsRegressed() bool {
	if m != nil && m.IsRegressed != nil {
		return *m.IsRegressed
	}
	return false
}

type MedianToleranceCheckResult struct {
	Params               *MedianToleranceParams `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	Tolerance            *float64               `protobuf:"fixed64,2,opt,name=tolerance" json:"tolerance,omitempty"`
	IsRegressed          *bool                  `protobuf:"varint,3,opt,name=is_regressed,json=isRegressed" json:"is_regressed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MedianToleranceCheckResult) Reset()         { *m = MedianToleranceCheckResult{} }
func (m *MedianToleranceCheckResult) String() string { return proto.CompactTextString(m) }
func (*MedianToleranceCheckResult) ProtoMessage()    {}
func (*MedianToleranceCheckResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20e6734acd2d200, []int{7}
}

func (m *MedianToleranceCheckResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MedianToleranceCheckResult.Unmarshal(m, b)
}
func (m *MedianToleranceCheckResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MedianToleranceCheckResult.Marshal(b, m, deterministic)
}
func (m *MedianToleranceCheckResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MedianToleranceCheckResult.Merge(m, src)
}
func (m *MedianToleranceCheckResult) XXX_Size() int {
	return xxx_messageInfo_MedianToleranceCheckResult.Size(m)
}
func (m *MedianToleranceCheckResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MedianToleranceCheckResult.DiscardUnknown(m)
}

var xxx_messageInfo_MedianToleranceCheckResult proto.InternalMessageInfo

func (m *MedianToleranceCheckResult) GetParams() *MedianToleranceParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *MedianToleranceCheckResult) GetTolerance() float64 {
	if m != nil && m.Tolerance != nil {
		return *m.Tolerance
	}
	return 0
}

func (m *MedianToleranceCheckResult) GetIsRegressed() bool {
	if m != nil && m.IsRegressed != nil {
		return *m.IsRegressed
	}
	return false
}

func init() {
	proto.RegisterEnum("mako.window_deviation.ToleranceCheck_DirectionBias", ToleranceCheck_DirectionBias_name, ToleranceCheck_DirectionBias_value)
	proto.RegisterEnum("mako.window_deviation.WindowDeviationOutput_ToleranceCheckOutput_CheckResult", WindowDeviationOutput_ToleranceCheckOutput_CheckResult_name, WindowDeviationOutput_ToleranceCheckOutput_CheckResult_value)
	proto.RegisterType((*WindowDeviationInput)(nil), "mako.window_deviation.WindowDeviationInput")
	proto.RegisterType((*ToleranceCheck)(nil), "mako.window_deviation.ToleranceCheck")
	proto.RegisterType((*MeanToleranceParams)(nil), "mako.window_deviation.MeanToleranceParams")
	proto.RegisterType((*MedianToleranceParams)(nil), "mako.window_deviation.MedianToleranceParams")
	proto.RegisterType((*WindowDeviationOutput)(nil), "mako.window_deviation.WindowDeviationOutput")
	proto.RegisterType((*WindowDeviationOutput_ToleranceCheckOutput)(nil), "mako.window_deviation.WindowDeviationOutput.ToleranceCheckOutput")
	proto.RegisterType((*ToleranceCheckStats)(nil), "mako.window_deviation.ToleranceCheckStats")
	proto.RegisterType((*MeanToleranceCheckResult)(nil), "mako.window_deviation.MeanToleranceCheckResult")
	proto.RegisterType((*MedianToleranceCheckResult)(nil), "mako.window_deviation.MedianToleranceCheckResult")
}

func init() {
	proto.RegisterFile("clients/proto/analyzers/window_deviation.proto", fileDescriptor_b20e6734acd2d200)
}

var fileDescriptor_b20e6734acd2d200 = []byte{
	// 1160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xcf, 0x6f, 0xdb, 0x36,
	0x14, 0xae, 0x92, 0xd8, 0xad, 0x9f, 0x6c, 0xd7, 0x60, 0x12, 0xcc, 0x73, 0xdb, 0x34, 0x75, 0x5b,
	0x2c, 0x28, 0x0a, 0xa5, 0x4d, 0xb1, 0x4b, 0x4e, 0x8b, 0x6b, 0xa7, 0xf3, 0x56, 0xbb, 0x19, 0xdd,
	0xa1, 0xd8, 0x89, 0x60, 0x25, 0xda, 0x21, 0x22, 0x4a, 0x9e, 0x48, 0xb7, 0x69, 0x2f, 0xfd, 0x53,
	0x76, 0xe8, 0x75, 0xff, 0xc4, 0x2e, 0xfb, 0xaf, 0x06, 0x0c, 0x24, 0x25, 0xff, 0x8a, 0x05, 0x38,
	0xd8, 0x6e, 0xf2, 0xc7, 0xef, 0x7d, 0xef, 0x23, 0xdf, 0xd3, 0xa3, 0x0c, 0x9e, 0x1f, 0x72, 0x16,
	0x29, 0x79, 0x38, 0x4e, 0x62, 0x15, 0x1f, 0xd2, 0x88, 0x86, 0x9f, 0x3e, 0xb3, 0x44, 0x1e, 0x7e,
	0xe4, 0x51, 0x10, 0x7f, 0x24, 0x01, 0xfb, 0xc0, 0xa9, 0xe2, 0x71, 0xe4, 0x19, 0x02, 0xda, 0x15,
	0xf4, 0x22, 0xf6, 0x96, 0x17, 0x1b, 0xbb, 0x72, 0xcc, 0xfc, 0x54, 0xc3, 0x30, 0xcc, 0x63, 0xf3,
	0x6f, 0x07, 0x76, 0xde, 0x19, 0x6e, 0x3b, 0xa3, 0x76, 0xa3, 0xf1, 0x44, 0xa1, 0x13, 0xd8, 0x4e,
	0x26, 0x11, 0xe1, 0xd1, 0x30, 0x26, 0xbf, 0x4f, 0x58, 0xf2, 0x89, 0x84, 0x5c, 0xaa, 0xba, 0xb3,
	0xbf, 0x79, 0xe0, 0x1e, 0x21, 0xcf, 0x48, 0xe0, 0x49, 0xd4, 0x8d, 0x86, 0xf1, 0x2f, 0x7a, 0x19,
	0xd7, 0x92, 0xb9, 0x5f, 0xaf, 0xb9, 0x54, 0xe8, 0x1d, 0xec, 0xa8, 0x38, 0x64, 0x09, 0x8d, 0x7c,
	0x46, 0xfc, 0x73, 0xe6, 0x5f, 0x58, 0x8d, 0x0d, 0xa3, 0xf1, 0xd8, 0x5b, 0x69, 0xd4, 0x7b, 0x9b,
	0x85, 0xbc, 0xd4, 0x11, 0x18, 0xa9, 0x85, 0xdf, 0x46, 0x18, 0xc1, 0x56, 0x44, 0x05, 0xab, 0x6f,
	0xee, 0x3b, 0x07, 0x25, 0x6c, 0x9e, 0x9b, 0x7f, 0x6d, 0x41, 0x75, 0x31, 0x14, 0x3d, 0x07, 0x37,
	0xa0, 0x8a, 0x92, 0x21, 0x0f, 0x15, 0x4b, 0xea, 0xce, 0xbe, 0x73, 0xe0, 0x1e, 0xd5, 0x6c, 0xda,
	0x36, 0x55, 0xf4, 0xd4, 0xe0, 0x18, 0x82, 0xe9, 0x33, 0x3a, 0x04, 0x94, 0x30, 0x9f, 0x45, 0x8a,
	0xa4, 0xbe, 0x24, 0xff, 0xcc, 0xea, 0x1b, 0xfb, 0xce, 0x41, 0xe1, 0xd8, 0x79, 0x8e, 0x6b, 0x76,
	0xd1, 0x1e, 0xd8, 0x80, 0x7f, 0x66, 0xe8, 0x14, 0xf6, 0x04, 0x8f, 0xb8, 0x98, 0x08, 0x72, 0xce,
	0xa5, 0x8a, 0x13, 0xee, 0xd3, 0x70, 0x21, 0xb8, 0x68, 0x83, 0x5f, 0xe0, 0x3b, 0x29, 0xf1, 0xc7,
	0x29, 0x6f, 0x4e, 0xc7, 0x87, 0x6a, 0xc0, 0x13, 0xe6, 0xeb, 0x23, 0x20, 0xef, 0x39, 0x95, 0x66,
	0x73, 0xd5, 0xa3, 0x17, 0x6b, 0x9d, 0x92, 0xd7, 0xce, 0x62, 0x5b, 0x9c, 0xca, 0xe3, 0x9b, 0xfd,
	0x37, 0xa4, 0xd5, 0x3d, 0x19, 0xe0, 0x4a, 0x30, 0x8f, 0xa3, 0x11, 0x34, 0x04, 0xa3, 0x11, 0x99,
	0x55, 0x65, 0x4c, 0x13, 0x2a, 0xa4, 0x2d, 0xcb, 0x96, 0x29, 0xcb, 0x93, 0x9c, 0x84, 0x3d, 0x46,
	0xa3, 0x69, 0xd2, 0x33, 0x13, 0x86, 0xbf, 0x11, 0x57, 0x41, 0x53, 0x20, 0x01, 0x77, 0x05, 0x0b,
	0x78, 0x6e, 0xaa, 0x82, 0x49, 0xf5, 0x34, 0x37, 0x95, 0x0e, 0x5d, 0x4e, 0xf6, 0xad, 0x58, 0x05,
	0xeb, 0x74, 0xcd, 0x53, 0xa8, 0x2c, 0x1c, 0x00, 0x72, 0x21, 0x3b, 0x82, 0xda, 0x0d, 0xb4, 0x0d,
	0xb7, 0xbb, 0xaf, 0xfa, 0x6f, 0x70, 0x87, 0x74, 0xfb, 0x2f, 0x71, 0xe7, 0x64, 0xd0, 0xa9, 0x39,
	0x73, 0x60, 0xbb, 0x93, 0x82, 0x1b, 0xcd, 0x2f, 0xb0, 0xbd, 0x62, 0x9b, 0x68, 0x1f, 0xc0, 0x8f,
	0x23, 0xa9, 0x88, 0x62, 0x89, 0x30, 0x6d, 0xe4, 0x1c, 0x3b, 0xcf, 0x70, 0xc9, 0x80, 0x6f, 0x59,
	0x22, 0x34, 0xc3, 0x1c, 0xac, 0x1f, 0xb3, 0xe1, 0xd0, 0xb4, 0x8b, 0x65, 0x68, 0xf0, 0xa5, 0xc6,
	0xd0, 0x23, 0x28, 0x4b, 0x15, 0x04, 0xec, 0x43, 0xca, 0xd9, 0xcc, 0x38, 0xae, 0x85, 0x0d, 0xab,
	0xf9, 0x05, 0x76, 0x57, 0x6e, 0x7e, 0x0d, 0x0b, 0x8f, 0xa0, 0x9c, 0x1e, 0xf9, 0x92, 0x09, 0xd7,
	0xc2, 0xd6, 0xc6, 0x1e, 0x94, 0x04, 0x0d, 0x96, 0x3d, 0xdc, 0x12, 0x34, 0xb0, 0x06, 0xfe, 0xbc,
	0x09, 0xbb, 0x4b, 0xe3, 0xe0, 0xcd, 0x44, 0xe9, 0x79, 0xf0, 0x18, 0xaa, 0xb1, 0x79, 0x22, 0x82,
	0x49, 0x49, 0x47, 0xcc, 0xb8, 0x28, 0xe1, 0x8a, 0x45, 0x7b, 0x16, 0x44, 0xbf, 0x41, 0xd1, 0xbc,
	0xe9, 0xba, 0x7f, 0x75, 0x8d, 0x4f, 0x72, 0x6a, 0xbc, 0x32, 0xc9, 0x52, 0x57, 0x5b, 0x10, 0xa7,
	0x82, 0xe8, 0x02, 0xee, 0xdb, 0x27, 0x22, 0x2f, 0xf8, 0x78, 0xcc, 0x02, 0x32, 0x8c, 0x13, 0x22,
	0xb8, 0x94, 0x3c, 0x1a, 0x11, 0xfd, 0x12, 0x5f, 0x6f, 0xb2, 0xdc, 0xb1, 0x6a, 0x03, 0x2b, 0x76,
	0x1a, 0x27, 0x3d, 0x2b, 0xa5, 0x47, 0x43, 0xe3, 0x6b, 0x01, 0x76, 0x56, 0xb9, 0x41, 0x7d, 0xb8,
	0xbd, 0x34, 0xd4, 0xd2, 0xc1, 0xb2, 0x66, 0xd6, 0xea, 0xe2, 0x3c, 0x43, 0x63, 0x28, 0x26, 0x4c,
	0x4e, 0x42, 0x65, 0x2a, 0x56, 0x3d, 0xea, 0xfd, 0xe7, 0x03, 0xf3, 0x6c, 0x22, 0x23, 0x7a, 0x5c,
	0xfa, 0xb5, 0xdf, 0xee, 0x9c, 0x76, 0xfb, 0x9d, 0x36, 0x4e, 0xf3, 0xa0, 0x07, 0xba, 0x53, 0x54,
	0xc2, 0x7d, 0x12, 0xd2, 0xf7, 0x2c, 0x4c, 0xa7, 0xa8, 0x6b, 0xb1, 0xd7, 0x1a, 0x42, 0x3f, 0x40,
	0x41, 0x2a, 0xaa, 0x64, 0x7d, 0xcb, 0x6c, 0xed, 0xc9, 0x5a, 0x5b, 0x1b, 0xe8, 0x08, 0x6c, 0x03,
	0xd1, 0x4f, 0xd0, 0xbc, 0x3a, 0x0f, 0x05, 0x8f, 0x88, 0xe2, 0x82, 0x49, 0x45, 0xc5, 0x98, 0x08,
	0x59, 0x2f, 0xe8, 0x0e, 0xc4, 0x7b, 0xe7, 0x4b, 0x13, 0xb1, 0xc7, 0xa3, 0xb7, 0x19, 0xad, 0x97,
	0xa7, 0x45, 0x2f, 0x17, 0xb5, 0x8a, 0x39, 0x5a, 0xf4, 0x72, 0x5e, 0xab, 0x05, 0x7b, 0xab, 0x7d,
	0xbd, 0x9f, 0xf0, 0x30, 0x20, 0x3c, 0xa8, 0xdf, 0x34, 0x3a, 0x8d, 0x15, 0x9e, 0x5a, 0x9a, 0xd2,
	0x0d, 0x72, 0x34, 0xe8, 0xe5, 0x4c, 0xe3, 0x56, 0x8e, 0x06, 0xbd, 0x4c, 0x35, 0x9a, 0x6d, 0x70,
	0xe7, 0xca, 0x84, 0x2a, 0x30, 0x2b, 0x54, 0xed, 0x86, 0xfe, 0x89, 0x3b, 0xaf, 0x70, 0x67, 0x30,
	0xe8, 0xb4, 0x6b, 0x8e, 0x1e, 0x67, 0x83, 0x9f, 0xbb, 0x67, 0x67, 0x9d, 0x76, 0x6d, 0x03, 0x01,
	0x14, 0xcf, 0x4e, 0xcc, 0xc2, 0x66, 0xf3, 0x9f, 0x2d, 0xd8, 0x5e, 0x51, 0x04, 0xf4, 0x0c, 0x76,
	0xb2, 0xdc, 0xe6, 0xc5, 0x20, 0x21, 0x8b, 0x46, 0xea, 0xdc, 0x74, 0x6a, 0x01, 0xa3, 0x6c, 0x4d,
	0x77, 0xfa, 0x6b, 0xb3, 0x82, 0x9e, 0x4e, 0x2f, 0xbe, 0x79, 0xbe, 0xb9, 0xf8, 0xb2, 0x5b, 0x6f,
	0x8e, 0xfd, 0x10, 0x2a, 0x53, 0x7d, 0x3d, 0xe3, 0xec, 0x28, 0xc1, 0xe5, 0x0c, 0xd4, 0x53, 0x14,
	0x7d, 0x07, 0xb7, 0xe7, 0x48, 0x7a, 0x06, 0x99, 0x76, 0x72, 0x70, 0x75, 0x46, 0xd3, 0xe8, 0x02,
	0xd1, 0x4e, 0xc3, 0xb4, 0x31, 0xa6, 0xc4, 0x81, 0x41, 0x75, 0xe7, 0xce, 0x14, 0x69, 0x90, 0x96,
	0xdc, 0x9d, 0xca, 0xd1, 0x00, 0xdd, 0x07, 0x37, 0xdd, 0x87, 0xf1, 0x65, 0x8b, 0x09, 0x16, 0x32,
	0xae, 0x1e, 0x42, 0x65, 0x4a, 0x30, 0x9e, 0x6c, 0xad, 0xca, 0x19, 0xc5, 0x38, 0xba, 0x07, 0x10,
	0xb0, 0x50, 0x51, 0x2b, 0x52, 0x32, 0x8c, 0x92, 0x41, 0x8c, 0xc6, 0x03, 0x28, 0x67, 0xcb, 0x46,
	0x02, 0xac, 0x8f, 0x94, 0x60, 0x14, 0x22, 0xb8, 0xb3, 0x74, 0xd5, 0xda, 0x0f, 0xa0, 0xf4, 0x5d,
	0x77, 0xcd, 0xa0, 0x3a, 0x5c, 0xe7, 0xae, 0x9d, 0x6b, 0x13, 0x5c, 0x17, 0x39, 0x2b, 0x48, 0xc1,
	0xbd, 0x2b, 0x37, 0xee, 0x42, 0xc6, 0xb2, 0xc9, 0xf8, 0x7c, 0xbd, 0x2b, 0x77, 0x3e, 0x67, 0x43,
	0xe4, 0xae, 0x35, 0xff, 0x70, 0xa0, 0x9e, 0x67, 0x16, 0xb5, 0xa0, 0x68, 0xef, 0xfc, 0x74, 0x40,
	0x5e, 0xe7, 0xcb, 0x22, 0x8d, 0x44, 0x77, 0xa1, 0x34, 0xdd, 0x8f, 0xbd, 0xd2, 0xf0, 0x0c, 0xd0,
	0x75, 0xe0, 0x92, 0x24, 0x6c, 0x94, 0x30, 0x29, 0x59, 0x60, 0xba, 0xf0, 0x16, 0x76, 0xb9, 0xc4,
	0x19, 0xd4, 0xfc, 0xea, 0x40, 0x23, 0x7f, 0x73, 0xa8, 0xbd, 0xe4, 0xf1, 0x7a, 0x9f, 0x24, 0xff,
	0x97, 0xcb, 0x56, 0x0f, 0xbe, 0xf7, 0x63, 0xe1, 0x8d, 0xe2, 0x78, 0x14, 0x32, 0x4f, 0x31, 0xa9,
	0x78, 0x34, 0xf2, 0xc6, 0x2c, 0x19, 0xc6, 0x89, 0xd0, 0x22, 0xd6, 0x56, 0xfa, 0x4f, 0xc0, 0x9b,
	0xfe, 0x07, 0x68, 0x2d, 0x7f, 0xbb, 0x9f, 0xe9, 0x8f, 0xfa, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x42, 0x5b, 0xd8, 0x0f, 0x33, 0x0c, 0x00, 0x00,
}
