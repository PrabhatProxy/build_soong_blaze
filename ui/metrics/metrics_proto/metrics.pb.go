// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics.proto

package metrics_proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MetricsBase_BUILDVARIANT int32

const (
	MetricsBase_USER      MetricsBase_BUILDVARIANT = 0
	MetricsBase_USERDEBUG MetricsBase_BUILDVARIANT = 1
	MetricsBase_ENG       MetricsBase_BUILDVARIANT = 2
)

var MetricsBase_BUILDVARIANT_name = map[int32]string{
	0: "USER",
	1: "USERDEBUG",
	2: "ENG",
}

var MetricsBase_BUILDVARIANT_value = map[string]int32{
	"USER":      0,
	"USERDEBUG": 1,
	"ENG":       2,
}

func (x MetricsBase_BUILDVARIANT) Enum() *MetricsBase_BUILDVARIANT {
	p := new(MetricsBase_BUILDVARIANT)
	*p = x
	return p
}

func (x MetricsBase_BUILDVARIANT) String() string {
	return proto.EnumName(MetricsBase_BUILDVARIANT_name, int32(x))
}

func (x *MetricsBase_BUILDVARIANT) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetricsBase_BUILDVARIANT_value, data, "MetricsBase_BUILDVARIANT")
	if err != nil {
		return err
	}
	*x = MetricsBase_BUILDVARIANT(value)
	return nil
}

func (MetricsBase_BUILDVARIANT) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0, 0}
}

type MetricsBase_ARCH int32

const (
	MetricsBase_UNKNOWN MetricsBase_ARCH = 0
	MetricsBase_ARM     MetricsBase_ARCH = 1
	MetricsBase_ARM64   MetricsBase_ARCH = 2
	MetricsBase_X86     MetricsBase_ARCH = 3
	MetricsBase_X86_64  MetricsBase_ARCH = 4
)

var MetricsBase_ARCH_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARM",
	2: "ARM64",
	3: "X86",
	4: "X86_64",
}

var MetricsBase_ARCH_value = map[string]int32{
	"UNKNOWN": 0,
	"ARM":     1,
	"ARM64":   2,
	"X86":     3,
	"X86_64":  4,
}

func (x MetricsBase_ARCH) Enum() *MetricsBase_ARCH {
	p := new(MetricsBase_ARCH)
	*p = x
	return p
}

func (x MetricsBase_ARCH) String() string {
	return proto.EnumName(MetricsBase_ARCH_name, int32(x))
}

func (x *MetricsBase_ARCH) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetricsBase_ARCH_value, data, "MetricsBase_ARCH")
	if err != nil {
		return err
	}
	*x = MetricsBase_ARCH(value)
	return nil
}

func (MetricsBase_ARCH) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0, 1}
}

type ModuleTypeInfo_BUILDSYSTEM int32

const (
	ModuleTypeInfo_UNKNOWN ModuleTypeInfo_BUILDSYSTEM = 0
	ModuleTypeInfo_SOONG   ModuleTypeInfo_BUILDSYSTEM = 1
	ModuleTypeInfo_MAKE    ModuleTypeInfo_BUILDSYSTEM = 2
)

var ModuleTypeInfo_BUILDSYSTEM_name = map[int32]string{
	0: "UNKNOWN",
	1: "SOONG",
	2: "MAKE",
}

var ModuleTypeInfo_BUILDSYSTEM_value = map[string]int32{
	"UNKNOWN": 0,
	"SOONG":   1,
	"MAKE":    2,
}

func (x ModuleTypeInfo_BUILDSYSTEM) Enum() *ModuleTypeInfo_BUILDSYSTEM {
	p := new(ModuleTypeInfo_BUILDSYSTEM)
	*p = x
	return p
}

func (x ModuleTypeInfo_BUILDSYSTEM) String() string {
	return proto.EnumName(ModuleTypeInfo_BUILDSYSTEM_name, int32(x))
}

func (x *ModuleTypeInfo_BUILDSYSTEM) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ModuleTypeInfo_BUILDSYSTEM_value, data, "ModuleTypeInfo_BUILDSYSTEM")
	if err != nil {
		return err
	}
	*x = ModuleTypeInfo_BUILDSYSTEM(value)
	return nil
}

func (ModuleTypeInfo_BUILDSYSTEM) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{3, 0}
}

type MetricsBase struct {
	// Timestamp generated when the build starts.
	BuildDateTimestamp *int64 `protobuf:"varint,1,opt,name=build_date_timestamp,json=buildDateTimestamp" json:"build_date_timestamp,omitempty"`
	// It is usually used to specify the branch name [and release candidate].
	BuildId *string `protobuf:"bytes,2,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	// The platform version codename, eg. P, Q, REL.
	PlatformVersionCodename *string `protobuf:"bytes,3,opt,name=platform_version_codename,json=platformVersionCodename" json:"platform_version_codename,omitempty"`
	// The target product information, eg. aosp_arm.
	TargetProduct *string `protobuf:"bytes,4,opt,name=target_product,json=targetProduct" json:"target_product,omitempty"`
	// The target build variant information, eg. eng.
	TargetBuildVariant *MetricsBase_BUILDVARIANT `protobuf:"varint,5,opt,name=target_build_variant,json=targetBuildVariant,enum=build_metrics.MetricsBase_BUILDVARIANT,def=2" json:"target_build_variant,omitempty"`
	// The target arch information, eg. arm.
	TargetArch *MetricsBase_ARCH `protobuf:"varint,6,opt,name=target_arch,json=targetArch,enum=build_metrics.MetricsBase_ARCH,def=0" json:"target_arch,omitempty"`
	// The target arch variant information, eg. armv7-a-neon.
	TargetArchVariant *string `protobuf:"bytes,7,opt,name=target_arch_variant,json=targetArchVariant" json:"target_arch_variant,omitempty"`
	// The target cpu variant information, eg. generic.
	TargetCpuVariant *string `protobuf:"bytes,8,opt,name=target_cpu_variant,json=targetCpuVariant" json:"target_cpu_variant,omitempty"`
	// The host arch information, eg. x86_64.
	HostArch *MetricsBase_ARCH `protobuf:"varint,9,opt,name=host_arch,json=hostArch,enum=build_metrics.MetricsBase_ARCH,def=0" json:"host_arch,omitempty"`
	// The host 2nd arch information, eg. x86.
	Host_2NdArch *MetricsBase_ARCH `protobuf:"varint,10,opt,name=host_2nd_arch,json=host2ndArch,enum=build_metrics.MetricsBase_ARCH,def=0" json:"host_2nd_arch,omitempty"`
	// The host os information, eg. linux.
	HostOs *string `protobuf:"bytes,11,opt,name=host_os,json=hostOs" json:"host_os,omitempty"`
	// The host os extra information, eg. Linux-4.17.0-3rodete2-amd64-x86_64-Debian-GNU.
	HostOsExtra *string `protobuf:"bytes,12,opt,name=host_os_extra,json=hostOsExtra" json:"host_os_extra,omitempty"`
	// The host cross os information, eg. windows.
	HostCrossOs *string `protobuf:"bytes,13,opt,name=host_cross_os,json=hostCrossOs" json:"host_cross_os,omitempty"`
	// The host cross arch information, eg. x86.
	HostCrossArch *string `protobuf:"bytes,14,opt,name=host_cross_arch,json=hostCrossArch" json:"host_cross_arch,omitempty"`
	// The host cross 2nd arch information, eg. x86_64.
	HostCross_2NdArch *string `protobuf:"bytes,15,opt,name=host_cross_2nd_arch,json=hostCross2ndArch" json:"host_cross_2nd_arch,omitempty"`
	// The directory for generated built artifacts installation, eg. out.
	OutDir *string `protobuf:"bytes,16,opt,name=out_dir,json=outDir" json:"out_dir,omitempty"`
	// The metrics for calling various tools (microfactory) before Soong_UI starts.
	SetupTools []*PerfInfo `protobuf:"bytes,17,rep,name=setup_tools,json=setupTools" json:"setup_tools,omitempty"`
	// The metrics for calling Kati by multiple times.
	KatiRuns []*PerfInfo `protobuf:"bytes,18,rep,name=kati_runs,json=katiRuns" json:"kati_runs,omitempty"`
	// The metrics for calling Soong.
	SoongRuns []*PerfInfo `protobuf:"bytes,19,rep,name=soong_runs,json=soongRuns" json:"soong_runs,omitempty"`
	// The metrics for calling Ninja.
	NinjaRuns   []*PerfInfo  `protobuf:"bytes,20,rep,name=ninja_runs,json=ninjaRuns" json:"ninja_runs,omitempty"`
	BuildConfig *BuildConfig `protobuf:"bytes,23,opt,name=build_config,json=buildConfig" json:"build_config,omitempty"`
	// The hostname of the machine.
	Hostname             *string  `protobuf:"bytes,24,opt,name=hostname" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsBase) Reset()         { *m = MetricsBase{} }
func (m *MetricsBase) String() string { return proto.CompactTextString(m) }
func (*MetricsBase) ProtoMessage()    {}
func (*MetricsBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}

func (m *MetricsBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsBase.Unmarshal(m, b)
}
func (m *MetricsBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsBase.Marshal(b, m, deterministic)
}
func (m *MetricsBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsBase.Merge(m, src)
}
func (m *MetricsBase) XXX_Size() int {
	return xxx_messageInfo_MetricsBase.Size(m)
}
func (m *MetricsBase) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsBase.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsBase proto.InternalMessageInfo

const Default_MetricsBase_TargetBuildVariant MetricsBase_BUILDVARIANT = MetricsBase_ENG
const Default_MetricsBase_TargetArch MetricsBase_ARCH = MetricsBase_UNKNOWN
const Default_MetricsBase_HostArch MetricsBase_ARCH = MetricsBase_UNKNOWN
const Default_MetricsBase_Host_2NdArch MetricsBase_ARCH = MetricsBase_UNKNOWN

func (m *MetricsBase) GetBuildDateTimestamp() int64 {
	if m != nil && m.BuildDateTimestamp != nil {
		return *m.BuildDateTimestamp
	}
	return 0
}

func (m *MetricsBase) GetBuildId() string {
	if m != nil && m.BuildId != nil {
		return *m.BuildId
	}
	return ""
}

func (m *MetricsBase) GetPlatformVersionCodename() string {
	if m != nil && m.PlatformVersionCodename != nil {
		return *m.PlatformVersionCodename
	}
	return ""
}

func (m *MetricsBase) GetTargetProduct() string {
	if m != nil && m.TargetProduct != nil {
		return *m.TargetProduct
	}
	return ""
}

func (m *MetricsBase) GetTargetBuildVariant() MetricsBase_BUILDVARIANT {
	if m != nil && m.TargetBuildVariant != nil {
		return *m.TargetBuildVariant
	}
	return Default_MetricsBase_TargetBuildVariant
}

func (m *MetricsBase) GetTargetArch() MetricsBase_ARCH {
	if m != nil && m.TargetArch != nil {
		return *m.TargetArch
	}
	return Default_MetricsBase_TargetArch
}

func (m *MetricsBase) GetTargetArchVariant() string {
	if m != nil && m.TargetArchVariant != nil {
		return *m.TargetArchVariant
	}
	return ""
}

func (m *MetricsBase) GetTargetCpuVariant() string {
	if m != nil && m.TargetCpuVariant != nil {
		return *m.TargetCpuVariant
	}
	return ""
}

func (m *MetricsBase) GetHostArch() MetricsBase_ARCH {
	if m != nil && m.HostArch != nil {
		return *m.HostArch
	}
	return Default_MetricsBase_HostArch
}

func (m *MetricsBase) GetHost_2NdArch() MetricsBase_ARCH {
	if m != nil && m.Host_2NdArch != nil {
		return *m.Host_2NdArch
	}
	return Default_MetricsBase_Host_2NdArch
}

func (m *MetricsBase) GetHostOs() string {
	if m != nil && m.HostOs != nil {
		return *m.HostOs
	}
	return ""
}

func (m *MetricsBase) GetHostOsExtra() string {
	if m != nil && m.HostOsExtra != nil {
		return *m.HostOsExtra
	}
	return ""
}

func (m *MetricsBase) GetHostCrossOs() string {
	if m != nil && m.HostCrossOs != nil {
		return *m.HostCrossOs
	}
	return ""
}

func (m *MetricsBase) GetHostCrossArch() string {
	if m != nil && m.HostCrossArch != nil {
		return *m.HostCrossArch
	}
	return ""
}

func (m *MetricsBase) GetHostCross_2NdArch() string {
	if m != nil && m.HostCross_2NdArch != nil {
		return *m.HostCross_2NdArch
	}
	return ""
}

func (m *MetricsBase) GetOutDir() string {
	if m != nil && m.OutDir != nil {
		return *m.OutDir
	}
	return ""
}

func (m *MetricsBase) GetSetupTools() []*PerfInfo {
	if m != nil {
		return m.SetupTools
	}
	return nil
}

func (m *MetricsBase) GetKatiRuns() []*PerfInfo {
	if m != nil {
		return m.KatiRuns
	}
	return nil
}

func (m *MetricsBase) GetSoongRuns() []*PerfInfo {
	if m != nil {
		return m.SoongRuns
	}
	return nil
}

func (m *MetricsBase) GetNinjaRuns() []*PerfInfo {
	if m != nil {
		return m.NinjaRuns
	}
	return nil
}

func (m *MetricsBase) GetBuildConfig() *BuildConfig {
	if m != nil {
		return m.BuildConfig
	}
	return nil
}

func (m *MetricsBase) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

type BuildConfig struct {
	UseGoma              *bool    `protobuf:"varint,1,opt,name=use_goma,json=useGoma" json:"use_goma,omitempty"`
	UseRbe               *bool    `protobuf:"varint,2,opt,name=use_rbe,json=useRbe" json:"use_rbe,omitempty"`
	ForceUseGoma         *bool    `protobuf:"varint,3,opt,name=force_use_goma,json=forceUseGoma" json:"force_use_goma,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildConfig) Reset()         { *m = BuildConfig{} }
func (m *BuildConfig) String() string { return proto.CompactTextString(m) }
func (*BuildConfig) ProtoMessage()    {}
func (*BuildConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{1}
}

func (m *BuildConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildConfig.Unmarshal(m, b)
}
func (m *BuildConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildConfig.Marshal(b, m, deterministic)
}
func (m *BuildConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildConfig.Merge(m, src)
}
func (m *BuildConfig) XXX_Size() int {
	return xxx_messageInfo_BuildConfig.Size(m)
}
func (m *BuildConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BuildConfig proto.InternalMessageInfo

func (m *BuildConfig) GetUseGoma() bool {
	if m != nil && m.UseGoma != nil {
		return *m.UseGoma
	}
	return false
}

func (m *BuildConfig) GetUseRbe() bool {
	if m != nil && m.UseRbe != nil {
		return *m.UseRbe
	}
	return false
}

func (m *BuildConfig) GetForceUseGoma() bool {
	if m != nil && m.ForceUseGoma != nil {
		return *m.ForceUseGoma
	}
	return false
}

type PerfInfo struct {
	// The description for the phase/action/part while the tool running.
	Desc *string `protobuf:"bytes,1,opt,name=desc" json:"desc,omitempty"`
	// The name for the running phase/action/part.
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// The absolute start time.
	// The number of nanoseconds elapsed since January 1, 1970 UTC.
	StartTime *uint64 `protobuf:"varint,3,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The real running time.
	// The number of nanoseconds elapsed since start_time.
	RealTime *uint64 `protobuf:"varint,4,opt,name=real_time,json=realTime" json:"real_time,omitempty"`
	// The number of MB for memory use.
	MemoryUse            *uint64  `protobuf:"varint,5,opt,name=memory_use,json=memoryUse" json:"memory_use,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PerfInfo) Reset()         { *m = PerfInfo{} }
func (m *PerfInfo) String() string { return proto.CompactTextString(m) }
func (*PerfInfo) ProtoMessage()    {}
func (*PerfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{2}
}

func (m *PerfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerfInfo.Unmarshal(m, b)
}
func (m *PerfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerfInfo.Marshal(b, m, deterministic)
}
func (m *PerfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerfInfo.Merge(m, src)
}
func (m *PerfInfo) XXX_Size() int {
	return xxx_messageInfo_PerfInfo.Size(m)
}
func (m *PerfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PerfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PerfInfo proto.InternalMessageInfo

func (m *PerfInfo) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

func (m *PerfInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PerfInfo) GetStartTime() uint64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *PerfInfo) GetRealTime() uint64 {
	if m != nil && m.RealTime != nil {
		return *m.RealTime
	}
	return 0
}

func (m *PerfInfo) GetMemoryUse() uint64 {
	if m != nil && m.MemoryUse != nil {
		return *m.MemoryUse
	}
	return 0
}

type ModuleTypeInfo struct {
	// The build system, eg. Soong or Make.
	BuildSystem *ModuleTypeInfo_BUILDSYSTEM `protobuf:"varint,1,opt,name=build_system,json=buildSystem,enum=build_metrics.ModuleTypeInfo_BUILDSYSTEM,def=0" json:"build_system,omitempty"`
	// The module type, eg. java_library, cc_binary, and etc.
	ModuleType *string `protobuf:"bytes,2,opt,name=module_type,json=moduleType" json:"module_type,omitempty"`
	// The number of logical modules.
	NumOfModules         *uint32  `protobuf:"varint,3,opt,name=num_of_modules,json=numOfModules" json:"num_of_modules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModuleTypeInfo) Reset()         { *m = ModuleTypeInfo{} }
func (m *ModuleTypeInfo) String() string { return proto.CompactTextString(m) }
func (*ModuleTypeInfo) ProtoMessage()    {}
func (*ModuleTypeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{3}
}

func (m *ModuleTypeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModuleTypeInfo.Unmarshal(m, b)
}
func (m *ModuleTypeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModuleTypeInfo.Marshal(b, m, deterministic)
}
func (m *ModuleTypeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleTypeInfo.Merge(m, src)
}
func (m *ModuleTypeInfo) XXX_Size() int {
	return xxx_messageInfo_ModuleTypeInfo.Size(m)
}
func (m *ModuleTypeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleTypeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleTypeInfo proto.InternalMessageInfo

const Default_ModuleTypeInfo_BuildSystem ModuleTypeInfo_BUILDSYSTEM = ModuleTypeInfo_UNKNOWN

func (m *ModuleTypeInfo) GetBuildSystem() ModuleTypeInfo_BUILDSYSTEM {
	if m != nil && m.BuildSystem != nil {
		return *m.BuildSystem
	}
	return Default_ModuleTypeInfo_BuildSystem
}

func (m *ModuleTypeInfo) GetModuleType() string {
	if m != nil && m.ModuleType != nil {
		return *m.ModuleType
	}
	return ""
}

func (m *ModuleTypeInfo) GetNumOfModules() uint32 {
	if m != nil && m.NumOfModules != nil {
		return *m.NumOfModules
	}
	return 0
}

func init() {
	proto.RegisterEnum("build_metrics.MetricsBase_BUILDVARIANT", MetricsBase_BUILDVARIANT_name, MetricsBase_BUILDVARIANT_value)
	proto.RegisterEnum("build_metrics.MetricsBase_ARCH", MetricsBase_ARCH_name, MetricsBase_ARCH_value)
	proto.RegisterEnum("build_metrics.ModuleTypeInfo_BUILDSYSTEM", ModuleTypeInfo_BUILDSYSTEM_name, ModuleTypeInfo_BUILDSYSTEM_value)
	proto.RegisterType((*MetricsBase)(nil), "build_metrics.MetricsBase")
	proto.RegisterType((*BuildConfig)(nil), "build_metrics.BuildConfig")
	proto.RegisterType((*PerfInfo)(nil), "build_metrics.PerfInfo")
	proto.RegisterType((*ModuleTypeInfo)(nil), "build_metrics.ModuleTypeInfo")
}

func init() { proto.RegisterFile("metrics.proto", fileDescriptor_6039342a2ba47b72) }

var fileDescriptor_6039342a2ba47b72 = []byte{
	// 871 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xae, 0x63, 0x27, 0x96, 0x8e, 0x62, 0x57, 0x61, 0x02, 0x44, 0xed, 0x50, 0xd4, 0x30, 0xf6,
	0xe3, 0x01, 0x9b, 0x57, 0x18, 0x81, 0x11, 0x04, 0xdb, 0x85, 0xed, 0x18, 0xa9, 0xd1, 0xda, 0x2e,
	0x68, 0x3b, 0xeb, 0x76, 0x31, 0x41, 0x91, 0x68, 0x47, 0x9b, 0x25, 0x1a, 0x24, 0x55, 0x2c, 0x0f,
	0xb1, 0xe7, 0xdb, 0xe5, 0x5e, 0x65, 0xe0, 0xa1, 0xe5, 0x28, 0xb9, 0x48, 0xd1, 0x3b, 0xf2, 0x7c,
	0x3f, 0xfc, 0x48, 0x91, 0x47, 0x50, 0x4b, 0x98, 0x12, 0x71, 0x28, 0xdb, 0x1b, 0xc1, 0x15, 0x27,
	0xb5, 0x9b, 0x2c, 0x5e, 0x47, 0xfe, 0xb6, 0xd8, 0xfc, 0xd7, 0x06, 0x67, 0x6c, 0xc6, 0xfd, 0x40,
	0x32, 0xf2, 0x06, 0x4e, 0x0c, 0x21, 0x0a, 0x14, 0xf3, 0x55, 0x9c, 0x30, 0xa9, 0x82, 0x64, 0xe3,
	0x95, 0x1a, 0xa5, 0x56, 0x99, 0x12, 0xc4, 0x2e, 0x03, 0xc5, 0xe6, 0x39, 0x42, 0x5e, 0x80, 0x65,
	0x14, 0x71, 0xe4, 0xed, 0x35, 0x4a, 0x2d, 0x9b, 0x56, 0x71, 0x3e, 0x8a, 0xc8, 0x05, 0xbc, 0xd8,
	0xac, 0x03, 0xb5, 0xe4, 0x22, 0xf1, 0x3f, 0x31, 0x21, 0x63, 0x9e, 0xfa, 0x21, 0x8f, 0x58, 0x1a,
	0x24, 0xcc, 0x2b, 0x23, 0xf7, 0x34, 0x27, 0x5c, 0x1b, 0x7c, 0xb0, 0x85, 0xc9, 0x37, 0x50, 0x57,
	0x81, 0x58, 0x31, 0xe5, 0x6f, 0x04, 0x8f, 0xb2, 0x50, 0x79, 0x15, 0x14, 0xd4, 0x4c, 0xf5, 0x83,
	0x29, 0x92, 0x3f, 0xe0, 0x64, 0x4b, 0x33, 0x21, 0x3e, 0x05, 0x22, 0x0e, 0x52, 0xe5, 0xed, 0x37,
	0x4a, 0xad, 0x7a, 0xe7, 0xbb, 0xf6, 0x83, 0xdd, 0xb6, 0x0b, 0x3b, 0x6d, 0xf7, 0x17, 0xa3, 0xf7,
	0x97, 0xd7, 0x3d, 0x3a, 0xea, 0x4d, 0xe6, 0x17, 0xe5, 0xe1, 0xe4, 0x8a, 0x12, 0xe3, 0xd4, 0xd7,
	0x92, 0x6b, 0xe3, 0x43, 0x46, 0xe0, 0x6c, 0xfd, 0x03, 0x11, 0xde, 0x7a, 0x07, 0x68, 0xfb, 0xfa,
	0x09, 0xdb, 0x1e, 0x1d, 0xbc, 0xbd, 0xa8, 0x2e, 0x26, 0xef, 0x26, 0xd3, 0x5f, 0x27, 0x14, 0x8c,
	0xb8, 0x27, 0xc2, 0x5b, 0xd2, 0x86, 0xe3, 0x82, 0xd5, 0x2e, 0x69, 0x15, 0xb7, 0x75, 0x74, 0x4f,
	0xcc, 0x97, 0xfe, 0x01, 0xb6, 0x81, 0xfc, 0x70, 0x93, 0xed, 0xe8, 0x16, 0xd2, 0x5d, 0x83, 0x0c,
	0x36, 0x59, 0xce, 0x1e, 0x82, 0x7d, 0xcb, 0xe5, 0x36, 0xa6, 0xfd, 0x85, 0x31, 0x2d, 0x2d, 0xc5,
	0x90, 0xef, 0xa1, 0x86, 0x36, 0x9d, 0x34, 0x32, 0x56, 0xf0, 0x85, 0x56, 0x8e, 0x96, 0x77, 0xd2,
	0x08, 0xdd, 0x4e, 0xa1, 0x8a, 0x6e, 0x5c, 0x7a, 0x0e, 0xe6, 0x3e, 0xd0, 0xd3, 0xa9, 0x24, 0xcd,
	0xed, 0x32, 0x5c, 0xfa, 0xec, 0x6f, 0x25, 0x02, 0xef, 0x10, 0x61, 0xc7, 0xc0, 0x43, 0x5d, 0xda,
	0x71, 0x42, 0xc1, 0xa5, 0xd4, 0x16, 0xb5, 0x7b, 0xce, 0x40, 0xd7, 0xa6, 0x92, 0x7c, 0x0b, 0xcf,
	0x0b, 0x1c, 0x0c, 0x5c, 0x37, 0xd7, 0x64, 0xc7, 0xc2, 0x20, 0x3f, 0xc2, 0x71, 0x81, 0xb7, 0xdb,
	0xdc, 0x73, 0x73, 0x98, 0x3b, 0x6e, 0x21, 0x37, 0xcf, 0x94, 0x1f, 0xc5, 0xc2, 0x73, 0x4d, 0x6e,
	0x9e, 0xa9, 0xcb, 0x58, 0x90, 0x73, 0x70, 0x24, 0x53, 0xd9, 0xc6, 0x57, 0x9c, 0xaf, 0xa5, 0x77,
	0xd4, 0x28, 0xb7, 0x9c, 0xce, 0xe9, 0xa3, 0xc3, 0xf9, 0xc0, 0xc4, 0x72, 0x94, 0x2e, 0x39, 0x05,
	0xe4, 0xce, 0x35, 0x95, 0x9c, 0x81, 0xfd, 0x57, 0xa0, 0x62, 0x5f, 0x64, 0xa9, 0xf4, 0xc8, 0xd3,
	0x3a, 0x4b, 0x33, 0x69, 0x96, 0x4a, 0xd2, 0x05, 0x90, 0x9c, 0xa7, 0x2b, 0x23, 0x3b, 0x7e, 0x5a,
	0x66, 0x23, 0x35, 0xd7, 0xa5, 0x71, 0xfa, 0x67, 0x60, 0x74, 0x27, 0x9f, 0xd1, 0x21, 0x15, 0x75,
	0xbf, 0xc0, 0xa1, 0x21, 0x85, 0x3c, 0x5d, 0xc6, 0x2b, 0xef, 0xb4, 0x51, 0x6a, 0x39, 0x9d, 0x97,
	0x8f, 0x94, 0xf8, 0x42, 0x06, 0xc8, 0xa0, 0xce, 0xcd, 0xfd, 0x84, 0xbc, 0x04, 0xbc, 0x49, 0xf8,
	0xbe, 0x3d, 0x3c, 0xb8, 0xdd, 0xbc, 0xf9, 0x06, 0x0e, 0x8b, 0x4f, 0x8e, 0x58, 0x50, 0x59, 0xcc,
	0x86, 0xd4, 0x7d, 0x46, 0x6a, 0x60, 0xeb, 0xd1, 0xe5, 0xb0, 0xbf, 0xb8, 0x72, 0x4b, 0xa4, 0x0a,
	0xfa, 0x35, 0xba, 0x7b, 0xcd, 0x9f, 0xa1, 0xa2, 0xef, 0x16, 0x71, 0x20, 0xbf, 0x5d, 0xee, 0x33,
	0x8d, 0xf6, 0xe8, 0xd8, 0x2d, 0x11, 0x1b, 0xf6, 0x7b, 0x74, 0xdc, 0x3d, 0x73, 0xf7, 0x74, 0xed,
	0xe3, 0x79, 0xd7, 0x2d, 0x13, 0x80, 0x83, 0x8f, 0xe7, 0x5d, 0xbf, 0x7b, 0xe6, 0x56, 0x9a, 0x2b,
	0x70, 0x0a, 0x39, 0x75, 0x9b, 0xca, 0x24, 0xf3, 0x57, 0x3c, 0x09, 0xb0, 0x99, 0x59, 0xb4, 0x9a,
	0x49, 0x76, 0xc5, 0x93, 0x40, 0x7f, 0x6d, 0x0d, 0x89, 0x1b, 0x86, 0x0d, 0xcc, 0xa2, 0x07, 0x99,
	0x64, 0xf4, 0x86, 0x91, 0xaf, 0xa1, 0xbe, 0xe4, 0x22, 0x64, 0xfe, 0x4e, 0x59, 0x46, 0xfc, 0x10,
	0xab, 0x0b, 0x23, 0x6f, 0xfe, 0x53, 0x02, 0x2b, 0x3f, 0x4b, 0x42, 0xa0, 0x12, 0x31, 0x19, 0xe2,
	0x12, 0x36, 0xc5, 0xb1, 0xae, 0xe1, 0x89, 0x98, 0xee, 0x88, 0x63, 0xf2, 0x0a, 0x40, 0xaa, 0x40,
	0x28, 0x6c, 0xb1, 0x68, 0x5b, 0xa1, 0x36, 0x56, 0x74, 0x67, 0x25, 0x5f, 0x81, 0x2d, 0x58, 0xb0,
	0x36, 0x68, 0x05, 0x51, 0x4b, 0x17, 0x10, 0x7c, 0x05, 0x90, 0xb0, 0x84, 0x8b, 0x3b, 0x9d, 0x0b,
	0x3b, 0x5d, 0x85, 0xda, 0xa6, 0xb2, 0x90, 0xac, 0xf9, 0x5f, 0x09, 0xea, 0x63, 0x1e, 0x65, 0x6b,
	0x36, 0xbf, 0xdb, 0x30, 0x4c, 0xb5, 0xc8, 0x3f, 0xab, 0xbc, 0x93, 0x8a, 0x25, 0x98, 0xae, 0xde,
	0xf9, 0xfe, 0xf1, 0xa3, 0x7e, 0x20, 0x32, 0x0d, 0x72, 0xf6, 0xdb, 0x6c, 0x3e, 0x1c, 0x17, 0x9e,
	0x37, 0x4a, 0x66, 0x68, 0x43, 0x5e, 0x83, 0x93, 0xa0, 0xc6, 0x57, 0x77, 0x9b, 0x7c, 0x7f, 0x90,
	0xec, 0x6c, 0xf4, 0x01, 0xa6, 0x59, 0xe2, 0xf3, 0xa5, 0x6f, 0x8a, 0x12, 0x77, 0x5a, 0xa3, 0x87,
	0x69, 0x96, 0x4c, 0x97, 0x66, 0x3d, 0xd9, 0xfc, 0x09, 0x9c, 0xc2, 0x5a, 0x0f, 0x3f, 0xb7, 0x0d,
	0xfb, 0xb3, 0xe9, 0x74, 0xa2, 0xef, 0x85, 0x05, 0x95, 0x71, 0xef, 0xdd, 0xd0, 0xdd, 0xeb, 0x1f,
	0xbd, 0x2d, 0xff, 0x9e, 0xff, 0xd6, 0x7c, 0xfc, 0xad, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x66,
	0x62, 0xba, 0xc0, 0xe6, 0x06, 0x00, 0x00,
}
