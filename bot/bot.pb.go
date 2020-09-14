// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: bot.proto

package bot

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/subiz/header/common"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BotCategory int32

const (
	BotCategory_users         BotCategory = 0
	BotCategory_systems       BotCategory = 1
	BotCategory_conversations BotCategory = 2
)

// Enum value maps for BotCategory.
var (
	BotCategory_name = map[int32]string{
		0: "users",
		1: "systems",
		2: "conversations",
	}
	BotCategory_value = map[string]int32{
		"users":         0,
		"systems":       1,
		"conversations": 2,
	}
)

func (x BotCategory) Enum() *BotCategory {
	p := new(BotCategory)
	*p = x
	return p
}

func (x BotCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BotCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_bot_proto_enumTypes[0].Descriptor()
}

func (BotCategory) Type() protoreflect.EnumType {
	return &file_bot_proto_enumTypes[0]
}

func (x BotCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BotCategory.Descriptor instead.
func (BotCategory) EnumDescriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{0}
}

type ActionType int32

const (
	ActionType_nil               ActionType = 0
	ActionType_condition         ActionType = 1
	ActionType_sleep             ActionType = 2
	ActionType_send_message      ActionType = 3
	ActionType_jump              ActionType = 4
	ActionType_send_email        ActionType = 5
	ActionType_tag               ActionType = 6
	ActionType_convert_to_ticket ActionType = 7
	ActionType_send_webhook      ActionType = 8
	ActionType_like_comment      ActionType = 9
	ActionType_hide_comment      ActionType = 10
	ActionType_end_chat          ActionType = 11
	ActionType_question          ActionType = 12
	ActionType_update_user       ActionType = 13
	ActionType_assign_agent      ActionType = 14
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0:  "nil",
		1:  "condition",
		2:  "sleep",
		3:  "send_message",
		4:  "jump",
		5:  "send_email",
		6:  "tag",
		7:  "convert_to_ticket",
		8:  "send_webhook",
		9:  "like_comment",
		10: "hide_comment",
		11: "end_chat",
		12: "question",
		13: "update_user",
		14: "assign_agent",
	}
	ActionType_value = map[string]int32{
		"nil":               0,
		"condition":         1,
		"sleep":             2,
		"send_message":      3,
		"jump":              4,
		"send_email":        5,
		"tag":               6,
		"convert_to_ticket": 7,
		"send_webhook":      8,
		"like_comment":      9,
		"hide_comment":      10,
		"end_chat":          11,
		"question":          12,
		"update_user":       13,
		"assign_agent":      14,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_bot_proto_enumTypes[1].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_bot_proto_enumTypes[1]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{1}
}

type ActionCondition_Group int32

const (
	ActionCondition_single ActionCondition_Group = 0
	ActionCondition_all    ActionCondition_Group = 1
	ActionCondition_any    ActionCondition_Group = 2
)

// Enum value maps for ActionCondition_Group.
var (
	ActionCondition_Group_name = map[int32]string{
		0: "single",
		1: "all",
		2: "any",
	}
	ActionCondition_Group_value = map[string]int32{
		"single": 0,
		"all":    1,
		"any":    2,
	}
)

func (x ActionCondition_Group) Enum() *ActionCondition_Group {
	p := new(ActionCondition_Group)
	*p = x
	return p
}

func (x ActionCondition_Group) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionCondition_Group) Descriptor() protoreflect.EnumDescriptor {
	return file_bot_proto_enumTypes[2].Descriptor()
}

func (ActionCondition_Group) Type() protoreflect.EnumType {
	return &file_bot_proto_enumTypes[2]
}

func (x ActionCondition_Group) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionCondition_Group.Descriptor instead.
func (ActionCondition_Group) EnumDescriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{5, 0}
}

type ActionCondition_Function int32

const (
	ActionCondition_minute_of_day ActionCondition_Function = 0
	ActionCondition_hour_of_day   ActionCondition_Function = 1
	ActionCondition_day_of_week   ActionCondition_Function = 2
	ActionCondition_day_ago       ActionCondition_Function = 3
)

// Enum value maps for ActionCondition_Function.
var (
	ActionCondition_Function_name = map[int32]string{
		0: "minute_of_day",
		1: "hour_of_day",
		2: "day_of_week",
		3: "day_ago",
	}
	ActionCondition_Function_value = map[string]int32{
		"minute_of_day": 0,
		"hour_of_day":   1,
		"day_of_week":   2,
		"day_ago":       3,
	}
)

func (x ActionCondition_Function) Enum() *ActionCondition_Function {
	p := new(ActionCondition_Function)
	*p = x
	return p
}

func (x ActionCondition_Function) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionCondition_Function) Descriptor() protoreflect.EnumDescriptor {
	return file_bot_proto_enumTypes[3].Descriptor()
}

func (ActionCondition_Function) Type() protoreflect.EnumType {
	return &file_bot_proto_enumTypes[3]
}

func (x ActionCondition_Function) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionCondition_Function.Descriptor instead.
func (ActionCondition_Function) EnumDescriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{5, 1}
}

type Bot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id        string          `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Name      string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ChannelId string          `protobuf:"bytes,5,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Category  string          `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"` // hello, operator, reply
	Status    string          `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`     // active, inactive, deleted
	Action    *Action         `protobuf:"bytes,9,opt,name=action,proto3" json:"action,omitempty"`
	Created   int64           `protobuf:"varint,10,opt,name=created,proto3" json:"created,omitempty"`
	Updated   int64           `protobuf:"varint,11,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Bot) Reset() {
	*x = Bot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bot) ProtoMessage() {}

func (x *Bot) ProtoReflect() protoreflect.Message {
	mi := &file_bot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bot.ProtoReflect.Descriptor instead.
func (*Bot) Descriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{0}
}

func (x *Bot) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *Bot) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Bot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bot) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Bot) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Bot) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Bot) GetAction() *Action {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *Bot) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Bot) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type Bots struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Bots []*Bot          `protobuf:"bytes,2,rep,name=bots,proto3" json:"bots,omitempty"`
}

func (x *Bots) Reset() {
	*x = Bots{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bots) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bots) ProtoMessage() {}

func (x *Bots) ProtoReflect() protoreflect.Message {
	mi := &file_bot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bots.ProtoReflect.Descriptor instead.
func (*Bots) Descriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{1}
}

func (x *Bots) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *Bots) GetBots() []*Bot {
	if x != nil {
		return x.Bots
	}
	return nil
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId   string             `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	BotId       string             `protobuf:"bytes,3,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	Id          string             `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Name        string             `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Nexts       []*Action          `protobuf:"bytes,8,rep,name=nexts,proto3" json:"nexts,omitempty"` // pointed to the next action or used for branching to multiple actions
	Type        string             `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`   // send_message, send_email, tag, convert_to_ticket, condition, sleep, send_webhook
	Condition   *ActionCondition   `protobuf:"bytes,10,opt,name=condition,proto3" json:"condition,omitempty"`
	SendMessage *ActionSendMessage `protobuf:"bytes,11,opt,name=send_message,json=sendMessage,proto3" json:"send_message,omitempty"`
	SendEmail   *ActionSendMessage `protobuf:"bytes,12,opt,name=send_email,json=sendEmail,proto3" json:"send_email,omitempty"`
	Sleep       *ActionSleep       `protobuf:"bytes,13,opt,name=sleep,proto3" json:"sleep,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_bot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{2}
}

func (x *Action) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Action) GetBotId() string {
	if x != nil {
		return x.BotId
	}
	return ""
}

func (x *Action) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action) GetNexts() []*Action {
	if x != nil {
		return x.Nexts
	}
	return nil
}

func (x *Action) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Action) GetCondition() *ActionCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *Action) GetSendMessage() *ActionSendMessage {
	if x != nil {
		return x.SendMessage
	}
	return nil
}

func (x *Action) GetSendEmail() *ActionSendMessage {
	if x != nil {
		return x.SendEmail
	}
	return nil
}

func (x *Action) GetSleep() *ActionSleep {
	if x != nil {
		return x.Sleep
	}
	return nil
}

type ActionSendMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string   `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	Attachments  []string `protobuf:"bytes,10,rep,name=attachments,proto3" json:"attachments,omitempty"`
	QuickReplies []string `protobuf:"bytes,11,rep,name=quick_replies,json=quickReplies,proto3" json:"quick_replies,omitempty"`
}

func (x *ActionSendMessage) Reset() {
	*x = ActionSendMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionSendMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionSendMessage) ProtoMessage() {}

func (x *ActionSendMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionSendMessage.ProtoReflect.Descriptor instead.
func (*ActionSendMessage) Descriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{3}
}

func (x *ActionSendMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ActionSendMessage) GetAttachments() []string {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *ActionSendMessage) GetQuickReplies() []string {
	if x != nil {
		return x.QuickReplies
	}
	return nil
}

type ActionSleep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration int64 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *ActionSleep) Reset() {
	*x = ActionSleep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionSleep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionSleep) ProtoMessage() {}

func (x *ActionSleep) ProtoReflect() protoreflect.Message {
	mi := &file_bot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionSleep.ProtoReflect.Descriptor instead.
func (*ActionSleep) Descriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{4}
}

func (x *ActionSleep) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type ActionCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group      string             `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Conditions []*ActionCondition `protobuf:"bytes,3,rep,name=conditions,proto3" json:"conditions,omitempty"` // only availabe if group is all or any
	// only avaiable if group is single
	Key      string               `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`   // unique, path to properties, for example data.conversation.text, global.agent.avaliablitiy...
	Type     string               `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"` // string, number, boolean
	Number   *common.NumberParams `protobuf:"bytes,7,opt,name=number,proto3" json:"number,omitempty"`
	String_  *common.StringParams `protobuf:"bytes,8,opt,name=string,proto3" json:"string,omitempty"`
	Function string               `protobuf:"bytes,19,opt,name=function,proto3" json:"function,omitempty"` // used to transform value of left side before evaluate expression
	Name     string               `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`         // front end used only
}

func (x *ActionCondition) Reset() {
	*x = ActionCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionCondition) ProtoMessage() {}

func (x *ActionCondition) ProtoReflect() protoreflect.Message {
	mi := &file_bot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionCondition.ProtoReflect.Descriptor instead.
func (*ActionCondition) Descriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{5}
}

func (x *ActionCondition) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ActionCondition) GetConditions() []*ActionCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *ActionCondition) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ActionCondition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ActionCondition) GetNumber() *common.NumberParams {
	if x != nil {
		return x.Number
	}
	return nil
}

func (x *ActionCondition) GetString_() *common.StringParams {
	if x != nil {
		return x.String_
	}
	return nil
}

func (x *ActionCondition) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *ActionCondition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_bot_proto protoreflect.FileDescriptor

var file_bot_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x62, 0x6f, 0x74,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97,
	0x02, 0x0a, 0x03, 0x42, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x23, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x47, 0x0a, 0x04, 0x42, 0x6f, 0x74, 0x73,
	0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03,
	0x63, 0x74, 0x78, 0x12, 0x1c, 0x0a, 0x04, 0x62, 0x6f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x42, 0x6f, 0x74, 0x52, 0x04, 0x62, 0x6f, 0x74,
	0x73, 0x22, 0xe7, 0x02, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x62,
	0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x74,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x6e, 0x65, 0x78, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x6e, 0x65, 0x78, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x39, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0b, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x0a,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x6c, 0x65, 0x65, 0x70, 0x52, 0x05, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x22, 0x74, 0x0a, 0x11, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x71, 0x75, 0x69, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65,
	0x73, 0x22, 0x29, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6c, 0x65, 0x65, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x84, 0x03, 0x0a,
	0x0f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x74,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25,
	0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x61, 0x6e, 0x79, 0x10, 0x02, 0x22, 0x4c, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x64,
	0x61, 0x79, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x68, 0x6f, 0x75, 0x72, 0x5f, 0x6f, 0x66, 0x5f,
	0x64, 0x61, 0x79, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x64, 0x61, 0x79, 0x5f, 0x6f, 0x66, 0x5f,
	0x77, 0x65, 0x65, 0x6b, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x64, 0x61, 0x79, 0x5f, 0x61, 0x67,
	0x6f, 0x10, 0x03, 0x2a, 0x38, 0x0a, 0x0b, 0x42, 0x6f, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x10, 0x02, 0x2a, 0xf0, 0x01,
	0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x6e, 0x69, 0x6c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x6a, 0x75, 0x6d, 0x70, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f,
	0x74, 0x6f, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x10, 0x08, 0x12, 0x10, 0x0a,
	0x0c, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x09, 0x12,
	0x10, 0x0a, 0x0c, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10,
	0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x10, 0x0b, 0x12,
	0x0c, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x0c, 0x12, 0x0f, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x10, 0x0d, 0x12, 0x10,
	0x0a, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x10, 0x0e,
	0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x75, 0x62, 0x69, 0x7a, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x62, 0x6f, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bot_proto_rawDescOnce sync.Once
	file_bot_proto_rawDescData = file_bot_proto_rawDesc
)

func file_bot_proto_rawDescGZIP() []byte {
	file_bot_proto_rawDescOnce.Do(func() {
		file_bot_proto_rawDescData = protoimpl.X.CompressGZIP(file_bot_proto_rawDescData)
	})
	return file_bot_proto_rawDescData
}

var file_bot_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_bot_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bot_proto_goTypes = []interface{}{
	(BotCategory)(0),              // 0: bot.BotCategory
	(ActionType)(0),               // 1: bot.ActionType
	(ActionCondition_Group)(0),    // 2: bot.ActionCondition.Group
	(ActionCondition_Function)(0), // 3: bot.ActionCondition.Function
	(*Bot)(nil),                   // 4: bot.Bot
	(*Bots)(nil),                  // 5: bot.Bots
	(*Action)(nil),                // 6: bot.Action
	(*ActionSendMessage)(nil),     // 7: bot.ActionSendMessage
	(*ActionSleep)(nil),           // 8: bot.ActionSleep
	(*ActionCondition)(nil),       // 9: bot.ActionCondition
	(*common.Context)(nil),        // 10: common.Context
	(*common.NumberParams)(nil),   // 11: common.NumberParams
	(*common.StringParams)(nil),   // 12: common.StringParams
}
var file_bot_proto_depIdxs = []int32{
	10, // 0: bot.Bot.ctx:type_name -> common.Context
	6,  // 1: bot.Bot.action:type_name -> bot.Action
	10, // 2: bot.Bots.ctx:type_name -> common.Context
	4,  // 3: bot.Bots.bots:type_name -> bot.Bot
	6,  // 4: bot.Action.nexts:type_name -> bot.Action
	9,  // 5: bot.Action.condition:type_name -> bot.ActionCondition
	7,  // 6: bot.Action.send_message:type_name -> bot.ActionSendMessage
	7,  // 7: bot.Action.send_email:type_name -> bot.ActionSendMessage
	8,  // 8: bot.Action.sleep:type_name -> bot.ActionSleep
	9,  // 9: bot.ActionCondition.conditions:type_name -> bot.ActionCondition
	11, // 10: bot.ActionCondition.number:type_name -> common.NumberParams
	12, // 11: bot.ActionCondition.string:type_name -> common.StringParams
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_bot_proto_init() }
func file_bot_proto_init() {
	if File_bot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bot); i {
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
		file_bot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bots); i {
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
		file_bot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_bot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionSendMessage); i {
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
		file_bot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionSleep); i {
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
		file_bot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionCondition); i {
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
			RawDescriptor: file_bot_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bot_proto_goTypes,
		DependencyIndexes: file_bot_proto_depIdxs,
		EnumInfos:         file_bot_proto_enumTypes,
		MessageInfos:      file_bot_proto_msgTypes,
	}.Build()
	File_bot_proto = out.File
	file_bot_proto_rawDesc = nil
	file_bot_proto_goTypes = nil
	file_bot_proto_depIdxs = nil
}
