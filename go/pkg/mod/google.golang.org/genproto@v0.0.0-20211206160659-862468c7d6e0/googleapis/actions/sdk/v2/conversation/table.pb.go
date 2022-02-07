// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/actions/sdk/v2/conversation/prompt/content/table.proto

package conversation

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The alignment of the content within the cell.
type TableColumn_HorizontalAlignment int32

const (
	// Unspecified horizontal alignment.
	TableColumn_UNSPECIFIED TableColumn_HorizontalAlignment = 0
	// Leading edge of the cell. This is the default.
	TableColumn_LEADING TableColumn_HorizontalAlignment = 1
	// Content is aligned to the center of the column.
	TableColumn_CENTER TableColumn_HorizontalAlignment = 2
	// Content is aligned to the trailing edge of the column.
	TableColumn_TRAILING TableColumn_HorizontalAlignment = 3
)

// Enum value maps for TableColumn_HorizontalAlignment.
var (
	TableColumn_HorizontalAlignment_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "LEADING",
		2: "CENTER",
		3: "TRAILING",
	}
	TableColumn_HorizontalAlignment_value = map[string]int32{
		"UNSPECIFIED": 0,
		"LEADING":     1,
		"CENTER":      2,
		"TRAILING":    3,
	}
)

func (x TableColumn_HorizontalAlignment) Enum() *TableColumn_HorizontalAlignment {
	p := new(TableColumn_HorizontalAlignment)
	*p = x
	return p
}

func (x TableColumn_HorizontalAlignment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TableColumn_HorizontalAlignment) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_conversation_prompt_content_table_proto_enumTypes[0].Descriptor()
}

func (TableColumn_HorizontalAlignment) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_conversation_prompt_content_table_proto_enumTypes[0]
}

func (x TableColumn_HorizontalAlignment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TableColumn_HorizontalAlignment.Descriptor instead.
func (TableColumn_HorizontalAlignment) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescGZIP(), []int{1, 0}
}

// A table card for displaying a table of text.
type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Overall title of the table. Optional but must be set if subtitle is set.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Subtitle for the table. Optional.
	Subtitle string `protobuf:"bytes,2,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	// Image associated with the table. Optional.
	Image *Image `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	// Headers and alignment of columns.
	Columns []*TableColumn `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
	// Row data of the table. The first 3 rows are guaranteed to be shown but
	// others might be cut on certain surfaces. Please test with the simulator to
	// see which rows will be shown for a given surface. On surfaces that support
	// the WEB_BROWSER capability, you can point the user to
	// a web page with more data.
	Rows []*TableRow `protobuf:"bytes,6,rep,name=rows,proto3" json:"rows,omitempty"`
	// Button.
	Button *Link `protobuf:"bytes,7,opt,name=button,proto3" json:"button,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescGZIP(), []int{0}
}

func (x *Table) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Table) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *Table) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Table) GetColumns() []*TableColumn {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Table) GetRows() []*TableRow {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *Table) GetButton() *Link {
	if x != nil {
		return x.Button
	}
	return nil
}

// Describes a column in a table.
type TableColumn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Header text for the column.
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Horizontal alignment of content w.r.t column. If unspecified, content
	// will be aligned to the leading edge.
	Align TableColumn_HorizontalAlignment `protobuf:"varint,2,opt,name=align,proto3,enum=google.actions.sdk.v2.conversation.TableColumn_HorizontalAlignment" json:"align,omitempty"`
}

func (x *TableColumn) Reset() {
	*x = TableColumn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableColumn) ProtoMessage() {}

func (x *TableColumn) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableColumn.ProtoReflect.Descriptor instead.
func (*TableColumn) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescGZIP(), []int{1}
}

func (x *TableColumn) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *TableColumn) GetAlign() TableColumn_HorizontalAlignment {
	if x != nil {
		return x.Align
	}
	return TableColumn_UNSPECIFIED
}

// Describes a cell in a row.
type TableCell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text content of the cell.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TableCell) Reset() {
	*x = TableCell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableCell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableCell) ProtoMessage() {}

func (x *TableCell) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableCell.ProtoReflect.Descriptor instead.
func (*TableCell) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescGZIP(), []int{2}
}

func (x *TableCell) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Describes a row in the table.
type TableRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cells in this row. The first 3 cells are guaranteed to be shown but
	// others might be cut on certain surfaces. Please test with the simulator
	// to see which cells will be shown for a given surface.
	Cells []*TableCell `protobuf:"bytes,1,rep,name=cells,proto3" json:"cells,omitempty"`
	// Indicates whether there should be a divider after each row.
	Divider bool `protobuf:"varint,2,opt,name=divider,proto3" json:"divider,omitempty"`
}

func (x *TableRow) Reset() {
	*x = TableRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableRow) ProtoMessage() {}

func (x *TableRow) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableRow.ProtoReflect.Descriptor instead.
func (*TableRow) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescGZIP(), []int{3}
}

func (x *TableRow) GetCells() []*TableCell {
	if x != nil {
		return x.Cells
	}
	return nil
}

func (x *TableRow) GetDivider() bool {
	if x != nil {
		return x.Divider
	}
	return false
}

var File_google_actions_sdk_v2_conversation_prompt_content_table_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc9, 0x02, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x49, 0x0a,
	0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52,
	0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x62, 0x75,
	0x74, 0x74, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x06, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x22, 0xcf, 0x01, 0x0a,
	0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41,
	0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x22,
	0x4d, 0x0a, 0x13, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x45, 0x41, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x41, 0x49, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x22, 0x1f,
	0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22,
	0x69, 0x0a, 0x08, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x77, 0x12, 0x43, 0x0a, 0x05, 0x63,
	0x65, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x86, 0x01, 0x0a, 0x26, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescData = file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDesc
)

func file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDescData
}

var file_google_actions_sdk_v2_conversation_prompt_content_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_actions_sdk_v2_conversation_prompt_content_table_proto_goTypes = []interface{}{
	(TableColumn_HorizontalAlignment)(0), // 0: google.actions.sdk.v2.conversation.TableColumn.HorizontalAlignment
	(*Table)(nil),                        // 1: google.actions.sdk.v2.conversation.Table
	(*TableColumn)(nil),                  // 2: google.actions.sdk.v2.conversation.TableColumn
	(*TableCell)(nil),                    // 3: google.actions.sdk.v2.conversation.TableCell
	(*TableRow)(nil),                     // 4: google.actions.sdk.v2.conversation.TableRow
	(*Image)(nil),                        // 5: google.actions.sdk.v2.conversation.Image
	(*Link)(nil),                         // 6: google.actions.sdk.v2.conversation.Link
}
var file_google_actions_sdk_v2_conversation_prompt_content_table_proto_depIdxs = []int32{
	5, // 0: google.actions.sdk.v2.conversation.Table.image:type_name -> google.actions.sdk.v2.conversation.Image
	2, // 1: google.actions.sdk.v2.conversation.Table.columns:type_name -> google.actions.sdk.v2.conversation.TableColumn
	4, // 2: google.actions.sdk.v2.conversation.Table.rows:type_name -> google.actions.sdk.v2.conversation.TableRow
	6, // 3: google.actions.sdk.v2.conversation.Table.button:type_name -> google.actions.sdk.v2.conversation.Link
	0, // 4: google.actions.sdk.v2.conversation.TableColumn.align:type_name -> google.actions.sdk.v2.conversation.TableColumn.HorizontalAlignment
	3, // 5: google.actions.sdk.v2.conversation.TableRow.cells:type_name -> google.actions.sdk.v2.conversation.TableCell
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_conversation_prompt_content_table_proto_init() }
func file_google_actions_sdk_v2_conversation_prompt_content_table_proto_init() {
	if File_google_actions_sdk_v2_conversation_prompt_content_table_proto != nil {
		return
	}
	file_google_actions_sdk_v2_conversation_prompt_content_image_proto_init()
	file_google_actions_sdk_v2_conversation_prompt_content_link_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableColumn); i {
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
		file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableCell); i {
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
		file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableRow); i {
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
			RawDescriptor: file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_conversation_prompt_content_table_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_conversation_prompt_content_table_proto_depIdxs,
		EnumInfos:         file_google_actions_sdk_v2_conversation_prompt_content_table_proto_enumTypes,
		MessageInfos:      file_google_actions_sdk_v2_conversation_prompt_content_table_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_conversation_prompt_content_table_proto = out.File
	file_google_actions_sdk_v2_conversation_prompt_content_table_proto_rawDesc = nil
	file_google_actions_sdk_v2_conversation_prompt_content_table_proto_goTypes = nil
	file_google_actions_sdk_v2_conversation_prompt_content_table_proto_depIdxs = nil
}