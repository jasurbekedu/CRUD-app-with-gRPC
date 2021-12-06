// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto_files/expenditure/expenditure.proto

package expenditure

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Expenditure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpenditureId int64   `protobuf:"varint,1,opt,name=ExpenditureId,proto3" json:"ExpenditureId,omitempty"`
	Amount        float64 `protobuf:"fixed64,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Summary       string  `protobuf:"bytes,3,opt,name=Summary,proto3" json:"Summary,omitempty"`
	SpentAt       string  `protobuf:"bytes,4,opt,name=SpentAt,proto3" json:"SpentAt,omitempty"`
}

func (x *Expenditure) Reset() {
	*x = Expenditure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_expenditure_expenditure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expenditure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expenditure) ProtoMessage() {}

func (x *Expenditure) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_expenditure_expenditure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expenditure.ProtoReflect.Descriptor instead.
func (*Expenditure) Descriptor() ([]byte, []int) {
	return file_proto_files_expenditure_expenditure_proto_rawDescGZIP(), []int{0}
}

func (x *Expenditure) GetExpenditureId() int64 {
	if x != nil {
		return x.ExpenditureId
	}
	return 0
}

func (x *Expenditure) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Expenditure) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Expenditure) GetSpentAt() string {
	if x != nil {
		return x.SpentAt
	}
	return ""
}

type Identificator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *Identificator) Reset() {
	*x = Identificator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_expenditure_expenditure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identificator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identificator) ProtoMessage() {}

func (x *Identificator) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_expenditure_expenditure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identificator.ProtoReflect.Descriptor instead.
func (*Identificator) Descriptor() ([]byte, []int) {
	return file_proto_files_expenditure_expenditure_proto_rawDescGZIP(), []int{1}
}

func (x *Identificator) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Identificator) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type PostExpenditure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identificator *Identificator `protobuf:"bytes,1,opt,name=Identificator,proto3" json:"Identificator,omitempty"`
	Amount        float64        `protobuf:"fixed64,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Summary       string         `protobuf:"bytes,3,opt,name=Summary,proto3" json:"Summary,omitempty"`
}

func (x *PostExpenditure) Reset() {
	*x = PostExpenditure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_expenditure_expenditure_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostExpenditure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostExpenditure) ProtoMessage() {}

func (x *PostExpenditure) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_expenditure_expenditure_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostExpenditure.ProtoReflect.Descriptor instead.
func (*PostExpenditure) Descriptor() ([]byte, []int) {
	return file_proto_files_expenditure_expenditure_proto_rawDescGZIP(), []int{2}
}

func (x *PostExpenditure) GetIdentificator() *Identificator {
	if x != nil {
		return x.Identificator
	}
	return nil
}

func (x *PostExpenditure) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PostExpenditure) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type ListOfExpenditure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expenditures []*Expenditure `protobuf:"bytes,1,rep,name=Expenditures,proto3" json:"Expenditures,omitempty"`
}

func (x *ListOfExpenditure) Reset() {
	*x = ListOfExpenditure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_expenditure_expenditure_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfExpenditure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfExpenditure) ProtoMessage() {}

func (x *ListOfExpenditure) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_expenditure_expenditure_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfExpenditure.ProtoReflect.Descriptor instead.
func (*ListOfExpenditure) Descriptor() ([]byte, []int) {
	return file_proto_files_expenditure_expenditure_proto_rawDescGZIP(), []int{3}
}

func (x *ListOfExpenditure) GetExpenditures() []*Expenditure {
	if x != nil {
		return x.Expenditures
	}
	return nil
}

var File_proto_files_expenditure_expenditure_proto protoreflect.FileDescriptor

var file_proto_files_expenditure_expenditure_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x78, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x22, 0x7f, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x45, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0x47, 0x0a, 0x0d, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x51, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x66, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x3c, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x0c, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x73, 0x32, 0xbe, 0x01,
	0x0a, 0x12, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x2e, 0x65, 0x78, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x78, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x1c, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x66, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x1a, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1e, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x42, 0x20,
	0x5a, 0x1e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x75, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_files_expenditure_expenditure_proto_rawDescOnce sync.Once
	file_proto_files_expenditure_expenditure_proto_rawDescData = file_proto_files_expenditure_expenditure_proto_rawDesc
)

func file_proto_files_expenditure_expenditure_proto_rawDescGZIP() []byte {
	file_proto_files_expenditure_expenditure_proto_rawDescOnce.Do(func() {
		file_proto_files_expenditure_expenditure_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_expenditure_expenditure_proto_rawDescData)
	})
	return file_proto_files_expenditure_expenditure_proto_rawDescData
}

var file_proto_files_expenditure_expenditure_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_files_expenditure_expenditure_proto_goTypes = []interface{}{
	(*Expenditure)(nil),       // 0: expenditure.Expenditure
	(*Identificator)(nil),     // 1: expenditure.Identificator
	(*PostExpenditure)(nil),   // 2: expenditure.PostExpenditure
	(*ListOfExpenditure)(nil), // 3: expenditure.ListOfExpenditure
}
var file_proto_files_expenditure_expenditure_proto_depIdxs = []int32{
	1, // 0: expenditure.PostExpenditure.Identificator:type_name -> expenditure.Identificator
	0, // 1: expenditure.ListOfExpenditure.Expenditures:type_name -> expenditure.Expenditure
	2, // 2: expenditure.ExpenditureService.PostNewExpenditure:input_type -> expenditure.PostExpenditure
	1, // 3: expenditure.ExpenditureService.GetListOfExpenditure:input_type -> expenditure.Identificator
	2, // 4: expenditure.ExpenditureService.PostNewExpenditure:output_type -> expenditure.PostExpenditure
	3, // 5: expenditure.ExpenditureService.GetListOfExpenditure:output_type -> expenditure.ListOfExpenditure
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_files_expenditure_expenditure_proto_init() }
func file_proto_files_expenditure_expenditure_proto_init() {
	if File_proto_files_expenditure_expenditure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_files_expenditure_expenditure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expenditure); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_files_expenditure_expenditure_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identificator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_files_expenditure_expenditure_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostExpenditure); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_files_expenditure_expenditure_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfExpenditure); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_files_expenditure_expenditure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_files_expenditure_expenditure_proto_goTypes,
		DependencyIndexes: file_proto_files_expenditure_expenditure_proto_depIdxs,
		MessageInfos:      file_proto_files_expenditure_expenditure_proto_msgTypes,
	}.Build()
	File_proto_files_expenditure_expenditure_proto = out.File
	file_proto_files_expenditure_expenditure_proto_rawDesc = nil
	file_proto_files_expenditure_expenditure_proto_goTypes = nil
	file_proto_files_expenditure_expenditure_proto_depIdxs = nil
}
