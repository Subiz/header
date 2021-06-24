// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.2
// source: error.proto

package header

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

type E int32

const (
	E_Sub                                     E = 0
	E_Email_SendRequest                       E = 10600
	E_PaymentWorkerJob                        E = 110440
	E_AccountResetPasswordEmail               E = 110533
	E_AccountPasswordChangedEmailRequested    E = 110534
	E_AccountConfirmChangeEmailEmailRequested E = 110535
	E_AccountWarnInactive                     E = 110566
	E_UnpaidInvoiceEmail                      E = 110567
	E_SubscriptionUpgradedEmail               E = 110568
	E_BillCreatedEmail                        E = 110569
	E_DowngradedToFreeEmail                   E = 110570
	E_UnpaidInvoice9DaysEmail                 E = 110572
	E_AccountCreatedEmail                     E = 110573
	E_UpdatePlanEmail                         E = 110574
	E_TrialEnding7Email                       E = 110575
	E_TrialEnding3Email                       E = 110576
	E_TrialEnding1Email                       E = 110577
	E_InvoicePaidEmail                        E = 110578
	E_AccountAgentJoinedEmail                 E = 110579
	E_UserTotalConvoUpdated                   E = 111232
	E_BqreportSynced                          E = 3000008
	E_BotSynced                               E = 4000001
	E_StartBot                                E = 4000002
	E_TeminateBot                             E = 4000003
	E_undefined                               E = 1 // unknown error
	E_internal_error                          E = 2 // typical server error
	E_missing_account_id                      E = 3
	E_invalid_proto                           E = 4 // protobuf message is broken
	E_database_error                          E = 5
	E_resource_not_found                      E = 6
	E_access_deny                             E = 7
	E_invalid_json                            E = 8 // json payload is broken
	E_account_not_found                       E = 9
	E_invalid_credential                      E = 10
	E_http_call_error                         E = 11
	E_agent_group_not_found                   E = 12
	E_agent_not_found                         E = 15
	E_facebook_page_not_found                 E = 16
	E_invalid_refresh_token                   E = 17
	E_wrong_password                          E = 18
	E_agent_is_not_active                     E = 19
	E_invalid_token                           E = 20
	E_invalid_message_size                    E = 21
	E_invalid_payload_size                    E = 22
	E_facebook_call_failed                    E = 23
	E_invalid_facebook_access_token           E = 24
	E_subiz_call_failed                       E = 25
	E_user_is_banned                          E = 26
	E_wrong_signature                         E = 27
	E_data_corrupted                          E = 28
	E_invalid_account_id                      E = 29
	E_invalid_target_id                       E = 30
	E_invalid_note_id                         E = 31
	E_invalid_user_id                         E = 32
	E_user_not_found                          E = 33
	E_automation_not_found                    E = 34
	E_invalid_attribute_value                 E = 35
	E_invalid_attribute_key                   E = 36
	E_invalid_attribute_name                  E = 37
	E_invalid_attribute_type                  E = 38
	E_too_many_attribute                      E = 39
	E_attribute_not_found                     E = 40
	E_invalid_agent_id                        E = 41
	E_invalid_note_target_id                  E = 42
	E_invalid_event_id                        E = 43
	E_invalid_content_type                    E = 44
	E_parse_query_failed                      E = 45
	E_domain_is_not_whitelisted               E = 46
	E_invalid_domain                          E = 47
	E_ip_is_blocked                           E = 48
	E_invalid_transaction_id                  E = 49
	E_invalid_amount                          E = 50
	E_inactive_promotion_referral_program     E = 51
	E_plan_not_found                          E = 52
	E_invoice_not_found                       E = 53
	E_invalid_stripe_info                     E = 54
	E_subscription_not_found                  E = 55
	E_invalid_plan                            E = 56
	E_promotion_not_found                     E = 57
	E_invalid_promotion                       E = 58
	E_invalid_email                           E = 59
	E_email_taken                             E = 60
	E_invalid_date_format                     E = 61
	E_prohibited_action                       E = 62
	E_invalid_working_day                     E = 63
	E_invalid_holiday                         E = 64
	// password_too_weak = 65;
	E_invalid_fullname                E = 66
	E_invalid_password                E = 67
	E_invalid_token_type              E = 68
	E_invalid_invoice_template        E = 69
	E_invoice_template_compile_failed E = 70
	E_invalid_google_auth_response    E = 71
	E_pdf_generate_failed             E = 72
	E_invalid_stripe_customer_id      E = 73
	E_invalid_stripe_token            E = 74
	E_stripe_call_failed              E = 75
	E_invalid_bill_id                 E = 76
	E_invalid_invoice_id              E = 77
	E_invalid_id                      E = 78
	E_payment_method_not_found        E = 79
	E_invalid_invoice_duedate         E = 80
	E_filestore_error                 E = 81
	E_invalid_country                 E = 82
	E_invalid_referrer_code           E = 83
	E_invalid_oauth_scope             E = 84
)

// Enum value maps for E.
var (
	E_name = map[int32]string{
		0:       "Sub",
		10600:   "Email_SendRequest",
		110440:  "PaymentWorkerJob",
		110533:  "AccountResetPasswordEmail",
		110534:  "AccountPasswordChangedEmailRequested",
		110535:  "AccountConfirmChangeEmailEmailRequested",
		110566:  "AccountWarnInactive",
		110567:  "UnpaidInvoiceEmail",
		110568:  "SubscriptionUpgradedEmail",
		110569:  "BillCreatedEmail",
		110570:  "DowngradedToFreeEmail",
		110572:  "UnpaidInvoice9DaysEmail",
		110573:  "AccountCreatedEmail",
		110574:  "UpdatePlanEmail",
		110575:  "TrialEnding7Email",
		110576:  "TrialEnding3Email",
		110577:  "TrialEnding1Email",
		110578:  "InvoicePaidEmail",
		110579:  "AccountAgentJoinedEmail",
		111232:  "UserTotalConvoUpdated",
		3000008: "BqreportSynced",
		4000001: "BotSynced",
		4000002: "StartBot",
		4000003: "TeminateBot",
		1:       "undefined",
		2:       "internal_error",
		3:       "missing_account_id",
		4:       "invalid_proto",
		5:       "database_error",
		6:       "resource_not_found",
		7:       "access_deny",
		8:       "invalid_json",
		9:       "account_not_found",
		10:      "invalid_credential",
		11:      "http_call_error",
		12:      "agent_group_not_found",
		15:      "agent_not_found",
		16:      "facebook_page_not_found",
		17:      "invalid_refresh_token",
		18:      "wrong_password",
		19:      "agent_is_not_active",
		20:      "invalid_token",
		21:      "invalid_message_size",
		22:      "invalid_payload_size",
		23:      "facebook_call_failed",
		24:      "invalid_facebook_access_token",
		25:      "subiz_call_failed",
		26:      "user_is_banned",
		27:      "wrong_signature",
		28:      "data_corrupted",
		29:      "invalid_account_id",
		30:      "invalid_target_id",
		31:      "invalid_note_id",
		32:      "invalid_user_id",
		33:      "user_not_found",
		34:      "automation_not_found",
		35:      "invalid_attribute_value",
		36:      "invalid_attribute_key",
		37:      "invalid_attribute_name",
		38:      "invalid_attribute_type",
		39:      "too_many_attribute",
		40:      "attribute_not_found",
		41:      "invalid_agent_id",
		42:      "invalid_note_target_id",
		43:      "invalid_event_id",
		44:      "invalid_content_type",
		45:      "parse_query_failed",
		46:      "domain_is_not_whitelisted",
		47:      "invalid_domain",
		48:      "ip_is_blocked",
		49:      "invalid_transaction_id",
		50:      "invalid_amount",
		51:      "inactive_promotion_referral_program",
		52:      "plan_not_found",
		53:      "invoice_not_found",
		54:      "invalid_stripe_info",
		55:      "subscription_not_found",
		56:      "invalid_plan",
		57:      "promotion_not_found",
		58:      "invalid_promotion",
		59:      "invalid_email",
		60:      "email_taken",
		61:      "invalid_date_format",
		62:      "prohibited_action",
		63:      "invalid_working_day",
		64:      "invalid_holiday",
		66:      "invalid_fullname",
		67:      "invalid_password",
		68:      "invalid_token_type",
		69:      "invalid_invoice_template",
		70:      "invoice_template_compile_failed",
		71:      "invalid_google_auth_response",
		72:      "pdf_generate_failed",
		73:      "invalid_stripe_customer_id",
		74:      "invalid_stripe_token",
		75:      "stripe_call_failed",
		76:      "invalid_bill_id",
		77:      "invalid_invoice_id",
		78:      "invalid_id",
		79:      "payment_method_not_found",
		80:      "invalid_invoice_duedate",
		81:      "filestore_error",
		82:      "invalid_country",
		83:      "invalid_referrer_code",
		84:      "invalid_oauth_scope",
	}
	E_value = map[string]int32{
		"Sub":                                     0,
		"Email_SendRequest":                       10600,
		"PaymentWorkerJob":                        110440,
		"AccountResetPasswordEmail":               110533,
		"AccountPasswordChangedEmailRequested":    110534,
		"AccountConfirmChangeEmailEmailRequested": 110535,
		"AccountWarnInactive":                     110566,
		"UnpaidInvoiceEmail":                      110567,
		"SubscriptionUpgradedEmail":               110568,
		"BillCreatedEmail":                        110569,
		"DowngradedToFreeEmail":                   110570,
		"UnpaidInvoice9DaysEmail":                 110572,
		"AccountCreatedEmail":                     110573,
		"UpdatePlanEmail":                         110574,
		"TrialEnding7Email":                       110575,
		"TrialEnding3Email":                       110576,
		"TrialEnding1Email":                       110577,
		"InvoicePaidEmail":                        110578,
		"AccountAgentJoinedEmail":                 110579,
		"UserTotalConvoUpdated":                   111232,
		"BqreportSynced":                          3000008,
		"BotSynced":                               4000001,
		"StartBot":                                4000002,
		"TeminateBot":                             4000003,
		"undefined":                               1,
		"internal_error":                          2,
		"missing_account_id":                      3,
		"invalid_proto":                           4,
		"database_error":                          5,
		"resource_not_found":                      6,
		"access_deny":                             7,
		"invalid_json":                            8,
		"account_not_found":                       9,
		"invalid_credential":                      10,
		"http_call_error":                         11,
		"agent_group_not_found":                   12,
		"agent_not_found":                         15,
		"facebook_page_not_found":                 16,
		"invalid_refresh_token":                   17,
		"wrong_password":                          18,
		"agent_is_not_active":                     19,
		"invalid_token":                           20,
		"invalid_message_size":                    21,
		"invalid_payload_size":                    22,
		"facebook_call_failed":                    23,
		"invalid_facebook_access_token":           24,
		"subiz_call_failed":                       25,
		"user_is_banned":                          26,
		"wrong_signature":                         27,
		"data_corrupted":                          28,
		"invalid_account_id":                      29,
		"invalid_target_id":                       30,
		"invalid_note_id":                         31,
		"invalid_user_id":                         32,
		"user_not_found":                          33,
		"automation_not_found":                    34,
		"invalid_attribute_value":                 35,
		"invalid_attribute_key":                   36,
		"invalid_attribute_name":                  37,
		"invalid_attribute_type":                  38,
		"too_many_attribute":                      39,
		"attribute_not_found":                     40,
		"invalid_agent_id":                        41,
		"invalid_note_target_id":                  42,
		"invalid_event_id":                        43,
		"invalid_content_type":                    44,
		"parse_query_failed":                      45,
		"domain_is_not_whitelisted":               46,
		"invalid_domain":                          47,
		"ip_is_blocked":                           48,
		"invalid_transaction_id":                  49,
		"invalid_amount":                          50,
		"inactive_promotion_referral_program":     51,
		"plan_not_found":                          52,
		"invoice_not_found":                       53,
		"invalid_stripe_info":                     54,
		"subscription_not_found":                  55,
		"invalid_plan":                            56,
		"promotion_not_found":                     57,
		"invalid_promotion":                       58,
		"invalid_email":                           59,
		"email_taken":                             60,
		"invalid_date_format":                     61,
		"prohibited_action":                       62,
		"invalid_working_day":                     63,
		"invalid_holiday":                         64,
		"invalid_fullname":                        66,
		"invalid_password":                        67,
		"invalid_token_type":                      68,
		"invalid_invoice_template":                69,
		"invoice_template_compile_failed":         70,
		"invalid_google_auth_response":            71,
		"pdf_generate_failed":                     72,
		"invalid_stripe_customer_id":              73,
		"invalid_stripe_token":                    74,
		"stripe_call_failed":                      75,
		"invalid_bill_id":                         76,
		"invalid_invoice_id":                      77,
		"invalid_id":                              78,
		"payment_method_not_found":                79,
		"invalid_invoice_duedate":                 80,
		"filestore_error":                         81,
		"invalid_country":                         82,
		"invalid_referrer_code":                   83,
		"invalid_oauth_scope":                     84,
	}
)

func (x E) Enum() *E {
	p := new(E)
	*p = x
	return p
}

func (x E) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (E) Descriptor() protoreflect.EnumDescriptor {
	return file_error_proto_enumTypes[0].Descriptor()
}

func (E) Type() protoreflect.EnumType {
	return &file_error_proto_enumTypes[0]
}

func (x E) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use E.Descriptor instead.
func (E) EnumDescriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Debug       string `protobuf:"bytes,3,opt,name=debug,proto3" json:"debug,omitempty"`
	Class       int32  `protobuf:"varint,4,opt,name=class,proto3" json:"class,omitempty"`
	Stack       string `protobuf:"bytes,5,opt,name=stack,proto3" json:"stack,omitempty"`
	Created     int64  `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"` // ms
	Code        string `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	Root        string `protobuf:"bytes,8,opt,name=root,proto3" json:"root,omitempty"`
	RequestId   string `protobuf:"bytes,9,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Error) GetDebug() string {
	if x != nil {
		return x.Debug
	}
	return ""
}

func (x *Error) GetClass() int32 {
	if x != nil {
		return x.Class
	}
	return 0
}

func (x *Error) GetStack() string {
	if x != nil {
		return x.Stack
	}
	return ""
}

func (x *Error) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *Error) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xcc, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x2a, 0x9e, 0x14, 0x0a, 0x01, 0x45, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x75,
	0x62, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x11, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0xe8, 0x52, 0x12, 0x16, 0x0a, 0x10, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x10,
	0xe8, 0xde, 0x06, 0x12, 0x1f, 0x0a, 0x19, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x10, 0xc5, 0xdf, 0x06, 0x12, 0x2a, 0x0a, 0x24, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10, 0xc6, 0xdf, 0x06,
	0x12, 0x2d, 0x0a, 0x27, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10, 0xc7, 0xdf, 0x06, 0x12,
	0x19, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x61, 0x72, 0x6e, 0x49, 0x6e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0xe6, 0xdf, 0x06, 0x12, 0x18, 0x0a, 0x12, 0x55, 0x6e,
	0x70, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x10, 0xe7, 0xdf, 0x06, 0x12, 0x1f, 0x0a, 0x19, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x10, 0xe8, 0xdf, 0x06, 0x12, 0x16, 0x0a, 0x10, 0x42, 0x69, 0x6c, 0x6c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0xe9, 0xdf, 0x06, 0x12, 0x1b, 0x0a,
	0x15, 0x44, 0x6f, 0x77, 0x6e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x54, 0x6f, 0x46, 0x72, 0x65,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0xea, 0xdf, 0x06, 0x12, 0x1d, 0x0a, 0x17, 0x55, 0x6e,
	0x70, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x39, 0x44, 0x61, 0x79, 0x73,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0xec, 0xdf, 0x06, 0x12, 0x19, 0x0a, 0x13, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x10, 0xed, 0xdf, 0x06, 0x12, 0x15, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0xee, 0xdf, 0x06, 0x12, 0x17, 0x0a, 0x11, 0x54,
	0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x37, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x10, 0xef, 0xdf, 0x06, 0x12, 0x17, 0x0a, 0x11, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x33, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0xf0, 0xdf, 0x06, 0x12, 0x17, 0x0a,
	0x11, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x31, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x10, 0xf1, 0xdf, 0x06, 0x12, 0x16, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x50, 0x61, 0x69, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0xf2, 0xdf, 0x06, 0x12, 0x1d,
	0x0a, 0x17, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4a, 0x6f,
	0x69, 0x6e, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0xf3, 0xdf, 0x06, 0x12, 0x1b, 0x0a,
	0x15, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x10, 0x80, 0xe5, 0x06, 0x12, 0x15, 0x0a, 0x0e, 0x42, 0x71,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x10, 0xc8, 0x8d, 0xb7,
	0x01, 0x12, 0x10, 0x0a, 0x09, 0x42, 0x6f, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x10, 0x81,
	0x92, 0xf4, 0x01, 0x12, 0x0f, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6f, 0x74, 0x10,
	0x82, 0x92, 0xf4, 0x01, 0x12, 0x12, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x74, 0x10, 0x83, 0x92, 0xf4, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x75, 0x6e, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x6e,
	0x79, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6a,
	0x73, 0x6f, 0x6e, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12,
	0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0b, 0x12, 0x19, 0x0a, 0x15, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f,
	0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x0f, 0x12, 0x1b, 0x0a, 0x17, 0x66, 0x61, 0x63,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x10, 0x10, 0x12, 0x19, 0x0a, 0x15, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x10,
	0x11, 0x12, 0x12, 0x0a, 0x0e, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x10, 0x12, 0x12, 0x17, 0x0a, 0x13, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x13, 0x12, 0x11,
	0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x10,
	0x14, 0x12, 0x18, 0x0a, 0x14, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x10, 0x15, 0x12, 0x18, 0x0a, 0x14, 0x69,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x10, 0x16, 0x12, 0x18, 0x0a, 0x14, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x17, 0x12,
	0x21, 0x0a, 0x1d, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x61, 0x63, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x10, 0x18, 0x12, 0x15, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x69, 0x7a, 0x5f, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x19, 0x12, 0x12, 0x0a, 0x0e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x73, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x10, 0x1a, 0x12, 0x13, 0x0a,
	0x0f, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x10, 0x1b, 0x12, 0x12, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x75,
	0x70, 0x74, 0x65, 0x64, 0x10, 0x1c, 0x12, 0x16, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x10, 0x1d, 0x12, 0x15,
	0x0a, 0x11, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x10, 0x1e, 0x12, 0x13, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x10, 0x1f, 0x12, 0x13, 0x0a, 0x0f, 0x69, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x10, 0x20, 0x12,
	0x12, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x10, 0x21, 0x12, 0x18, 0x0a, 0x14, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x22, 0x12, 0x1b, 0x0a,
	0x17, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x10, 0x23, 0x12, 0x19, 0x0a, 0x15, 0x69, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x10, 0x24, 0x12, 0x1a, 0x0a, 0x16, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x10,
	0x25, 0x12, 0x1a, 0x0a, 0x16, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x10, 0x26, 0x12, 0x16, 0x0a,
	0x12, 0x74, 0x6f, 0x6f, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x10, 0x27, 0x12, 0x17, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x28, 0x12, 0x14,
	0x0a, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x10, 0x29, 0x12, 0x1a, 0x0a, 0x16, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x10, 0x2a,
	0x12, 0x14, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x10, 0x2b, 0x12, 0x18, 0x0a, 0x14, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x10, 0x2c,
	0x12, 0x16, 0x0a, 0x12, 0x70, 0x61, 0x72, 0x73, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x2d, 0x12, 0x1d, 0x0a, 0x19, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x64, 0x10, 0x2e, 0x12, 0x12, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x10, 0x2f, 0x12, 0x11, 0x0a, 0x0d, 0x69,
	0x70, 0x5f, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x10, 0x30, 0x12, 0x1a,
	0x0a, 0x16, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x10, 0x31, 0x12, 0x12, 0x0a, 0x0e, 0x69, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x32, 0x12, 0x27,
	0x0a, 0x23, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x5f, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x10, 0x33, 0x12, 0x12, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x6e, 0x5f,
	0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x34, 0x12, 0x15, 0x0a, 0x11, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x10, 0x35, 0x12, 0x17, 0x0a, 0x13, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x10, 0x36, 0x12, 0x1a, 0x0a, 0x16, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x37, 0x12, 0x10, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x10, 0x38, 0x12, 0x17, 0x0a, 0x13, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x10, 0x39, 0x12, 0x15, 0x0a, 0x11, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x3a, 0x12, 0x11, 0x0a, 0x0d, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x3b, 0x12, 0x0f, 0x0a, 0x0b,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x10, 0x3c, 0x12, 0x17, 0x0a,
	0x13, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x10, 0x3d, 0x12, 0x15, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x68, 0x69, 0x62,
	0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x3e, 0x12, 0x17, 0x0a,
	0x13, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x64, 0x61, 0x79, 0x10, 0x3f, 0x12, 0x13, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x10, 0x40, 0x12, 0x14, 0x0a, 0x10, 0x69,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x10,
	0x42, 0x12, 0x14, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x43, 0x12, 0x16, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x10, 0x44, 0x12,
	0x1c, 0x0a, 0x18, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x10, 0x45, 0x12, 0x23, 0x0a,
	0x1f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x46, 0x12, 0x20, 0x0a, 0x1c, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x10, 0x47, 0x12, 0x17, 0x0a, 0x13, 0x70, 0x64, 0x66, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x48, 0x12, 0x1e, 0x0a,
	0x1a, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x10, 0x49, 0x12, 0x18, 0x0a,
	0x14, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x4a, 0x12, 0x16, 0x0a, 0x12, 0x73, 0x74, 0x72, 0x69, 0x70,
	0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x4b, 0x12,
	0x13, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x10, 0x4c, 0x12, 0x16, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x10, 0x4d, 0x12, 0x0e, 0x0a, 0x0a,
	0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x10, 0x4e, 0x12, 0x1c, 0x0a, 0x18,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e,
	0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x4f, 0x12, 0x1b, 0x0a, 0x17, 0x69, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x75,
	0x65, 0x64, 0x61, 0x74, 0x65, 0x10, 0x50, 0x12, 0x13, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x51, 0x12, 0x13, 0x0a, 0x0f,
	0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x10,
	0x52, 0x12, 0x19, 0x0a, 0x15, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x10, 0x53, 0x12, 0x17, 0x0a, 0x13,
	0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x10, 0x54, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62, 0x69, 0x7a, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func file_error_proto_rawDescGZIP() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_error_proto_goTypes = []interface{}{
	(E)(0),        // 0: header.E
	(*Error)(nil), // 1: header.Error
}
var file_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		EnumInfos:         file_error_proto_enumTypes,
		MessageInfos:      file_error_proto_msgTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}
