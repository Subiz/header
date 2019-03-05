// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package noti5

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	common "github.com/subiz/header/common"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7e984e63DecodeGithubComSubizHeaderNoti5(in *jlexer.Lexer, out *Token) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ctx":
			if in.IsNull() {
				in.Skip()
				out.Ctx = nil
			} else {
				if out.Ctx == nil {
					out.Ctx = new(common.Context)
				}
				(*out.Ctx).UnmarshalEasyJSON(in)
			}
		case "account_id":
			if in.IsNull() {
				in.Skip()
				out.AccountId = nil
			} else {
				if out.AccountId == nil {
					out.AccountId = new(string)
				}
				*out.AccountId = string(in.String())
			}
		case "user_id":
			if in.IsNull() {
				in.Skip()
				out.UserId = nil
			} else {
				if out.UserId == nil {
					out.UserId = new(string)
				}
				*out.UserId = string(in.String())
			}
		case "user_type":
			if in.IsNull() {
				in.Skip()
				out.UserType = nil
			} else {
				if out.UserType == nil {
					out.UserType = new(string)
				}
				*out.UserType = string(in.String())
			}
		case "fcm_token":
			if in.IsNull() {
				in.Skip()
				out.FcmToken = nil
			} else {
				if out.FcmToken == nil {
					out.FcmToken = new(string)
				}
				*out.FcmToken = string(in.String())
			}
		case "device_id":
			if in.IsNull() {
				in.Skip()
				out.DeviceId = nil
			} else {
				if out.DeviceId == nil {
					out.DeviceId = new(string)
				}
				*out.DeviceId = string(in.String())
			}
		case "platform":
			if in.IsNull() {
				in.Skip()
				out.Platform = nil
			} else {
				if out.Platform == nil {
					out.Platform = new(string)
				}
				*out.Platform = string(in.String())
			}
		case "created":
			if in.IsNull() {
				in.Skip()
				out.Created = nil
			} else {
				if out.Created == nil {
					out.Created = new(int64)
				}
				*out.Created = int64(in.Int64())
			}
		case "topics":
			if in.IsNull() {
				in.Skip()
				out.Topics = nil
			} else {
				in.Delim('[')
				if out.Topics == nil {
					if !in.IsDelim(']') {
						out.Topics = make([]string, 0, 4)
					} else {
						out.Topics = []string{}
					}
				} else {
					out.Topics = (out.Topics)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Topics = append(out.Topics, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7e984e63EncodeGithubComSubizHeaderNoti5(out *jwriter.Writer, in Token) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ctx != nil {
		const prefix string = ",\"ctx\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Ctx).MarshalEasyJSON(out)
	}
	if in.AccountId != nil {
		const prefix string = ",\"account_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AccountId))
	}
	if in.UserId != nil {
		const prefix string = ",\"user_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.UserId))
	}
	if in.UserType != nil {
		const prefix string = ",\"user_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.UserType))
	}
	if in.FcmToken != nil {
		const prefix string = ",\"fcm_token\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.FcmToken))
	}
	if in.DeviceId != nil {
		const prefix string = ",\"device_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.DeviceId))
	}
	if in.Platform != nil {
		const prefix string = ",\"platform\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Platform))
	}
	if in.Created != nil {
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.Created))
	}
	if len(in.Topics) != 0 {
		const prefix string = ",\"topics\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Topics {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Token) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7e984e63EncodeGithubComSubizHeaderNoti5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Token) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7e984e63EncodeGithubComSubizHeaderNoti5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Token) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7e984e63DecodeGithubComSubizHeaderNoti5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Token) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7e984e63DecodeGithubComSubizHeaderNoti5(l, v)
}
func easyjson7e984e63DecodeGithubComSubizHeaderNoti51(in *jlexer.Lexer, out *Subscription) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "new_message":
			if in.IsNull() {
				in.Skip()
				out.NewMessage = nil
			} else {
				if out.NewMessage == nil {
					out.NewMessage = new(bool)
				}
				*out.NewMessage = bool(in.Bool())
			}
		case "unassigned_conversation":
			if in.IsNull() {
				in.Skip()
				out.UnassignedConversation = nil
			} else {
				if out.UnassignedConversation == nil {
					out.UnassignedConversation = new(bool)
				}
				*out.UnassignedConversation = bool(in.Bool())
			}
		case "delay":
			if in.IsNull() {
				in.Skip()
				out.Delay = nil
			} else {
				if out.Delay == nil {
					out.Delay = new(int32)
				}
				*out.Delay = int32(in.Int32())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7e984e63EncodeGithubComSubizHeaderNoti51(out *jwriter.Writer, in Subscription) {
	out.RawByte('{')
	first := true
	_ = first
	if in.NewMessage != nil {
		const prefix string = ",\"new_message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.NewMessage))
	}
	if in.UnassignedConversation != nil {
		const prefix string = ",\"unassigned_conversation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.UnassignedConversation))
	}
	if in.Delay != nil {
		const prefix string = ",\"delay\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.Delay))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Subscription) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7e984e63EncodeGithubComSubizHeaderNoti51(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Subscription) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7e984e63EncodeGithubComSubizHeaderNoti51(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Subscription) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7e984e63DecodeGithubComSubizHeaderNoti51(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Subscription) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7e984e63DecodeGithubComSubizHeaderNoti51(l, v)
}
func easyjson7e984e63DecodeGithubComSubizHeaderNoti52(in *jlexer.Lexer, out *SubscribeStatus) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ctx":
			if in.IsNull() {
				in.Skip()
				out.Ctx = nil
			} else {
				if out.Ctx == nil {
					out.Ctx = new(common.Context)
				}
				(*out.Ctx).UnmarshalEasyJSON(in)
			}
		case "account_id":
			if in.IsNull() {
				in.Skip()
				out.AccountId = nil
			} else {
				if out.AccountId == nil {
					out.AccountId = new(string)
				}
				*out.AccountId = string(in.String())
			}
		case "user_id":
			if in.IsNull() {
				in.Skip()
				out.UserId = nil
			} else {
				if out.UserId == nil {
					out.UserId = new(string)
				}
				*out.UserId = string(in.String())
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(bool)
				}
				*out.Status = bool(in.Bool())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7e984e63EncodeGithubComSubizHeaderNoti52(out *jwriter.Writer, in SubscribeStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ctx != nil {
		const prefix string = ",\"ctx\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Ctx).MarshalEasyJSON(out)
	}
	if in.AccountId != nil {
		const prefix string = ",\"account_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AccountId))
	}
	if in.UserId != nil {
		const prefix string = ",\"user_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.UserId))
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SubscribeStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7e984e63EncodeGithubComSubizHeaderNoti52(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SubscribeStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7e984e63EncodeGithubComSubizHeaderNoti52(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SubscribeStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7e984e63DecodeGithubComSubizHeaderNoti52(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SubscribeStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7e984e63DecodeGithubComSubizHeaderNoti52(l, v)
}
func easyjson7e984e63DecodeGithubComSubizHeaderNoti53(in *jlexer.Lexer, out *Setting) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ctx":
			if in.IsNull() {
				in.Skip()
				out.Ctx = nil
			} else {
				if out.Ctx == nil {
					out.Ctx = new(common.Context)
				}
				(*out.Ctx).UnmarshalEasyJSON(in)
			}
		case "account_id":
			if in.IsNull() {
				in.Skip()
				out.AccountId = nil
			} else {
				if out.AccountId == nil {
					out.AccountId = new(string)
				}
				*out.AccountId = string(in.String())
			}
		case "agent_id":
			if in.IsNull() {
				in.Skip()
				out.AgentId = nil
			} else {
				if out.AgentId == nil {
					out.AgentId = new(string)
				}
				*out.AgentId = string(in.String())
			}
		case "web_type":
			if in.IsNull() {
				in.Skip()
				out.WebType = nil
			} else {
				if out.WebType == nil {
					out.WebType = new(string)
				}
				*out.WebType = string(in.String())
			}
		case "mobile_type":
			if in.IsNull() {
				in.Skip()
				out.MobileType = nil
			} else {
				if out.MobileType == nil {
					out.MobileType = new(string)
				}
				*out.MobileType = string(in.String())
			}
		case "email_type":
			if in.IsNull() {
				in.Skip()
				out.EmailType = nil
			} else {
				if out.EmailType == nil {
					out.EmailType = new(string)
				}
				*out.EmailType = string(in.String())
			}
		case "email_after":
			if in.IsNull() {
				in.Skip()
				out.EmailAfter = nil
			} else {
				if out.EmailAfter == nil {
					out.EmailAfter = new(int32)
				}
				*out.EmailAfter = int32(in.Int32())
			}
		case "updated":
			if in.IsNull() {
				in.Skip()
				out.Updated = nil
			} else {
				if out.Updated == nil {
					out.Updated = new(int64)
				}
				*out.Updated = int64(in.Int64())
			}
		case "web":
			if in.IsNull() {
				in.Skip()
				out.Web = nil
			} else {
				if out.Web == nil {
					out.Web = new(Subscription)
				}
				(*out.Web).UnmarshalEasyJSON(in)
			}
		case "mobile":
			if in.IsNull() {
				in.Skip()
				out.Mobile = nil
			} else {
				if out.Mobile == nil {
					out.Mobile = new(Subscription)
				}
				(*out.Mobile).UnmarshalEasyJSON(in)
			}
		case "email":
			if in.IsNull() {
				in.Skip()
				out.Email = nil
			} else {
				if out.Email == nil {
					out.Email = new(Subscription)
				}
				(*out.Email).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7e984e63EncodeGithubComSubizHeaderNoti53(out *jwriter.Writer, in Setting) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ctx != nil {
		const prefix string = ",\"ctx\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Ctx).MarshalEasyJSON(out)
	}
	if in.AccountId != nil {
		const prefix string = ",\"account_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AccountId))
	}
	if in.AgentId != nil {
		const prefix string = ",\"agent_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AgentId))
	}
	if in.WebType != nil {
		const prefix string = ",\"web_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.WebType))
	}
	if in.MobileType != nil {
		const prefix string = ",\"mobile_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.MobileType))
	}
	if in.EmailType != nil {
		const prefix string = ",\"email_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.EmailType))
	}
	if in.EmailAfter != nil {
		const prefix string = ",\"email_after\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.EmailAfter))
	}
	if in.Updated != nil {
		const prefix string = ",\"updated\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.Updated))
	}
	if in.Web != nil {
		const prefix string = ",\"web\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Web).MarshalEasyJSON(out)
	}
	if in.Mobile != nil {
		const prefix string = ",\"mobile\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Mobile).MarshalEasyJSON(out)
	}
	if in.Email != nil {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Email).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Setting) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7e984e63EncodeGithubComSubizHeaderNoti53(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Setting) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7e984e63EncodeGithubComSubizHeaderNoti53(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Setting) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7e984e63DecodeGithubComSubizHeaderNoti53(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Setting) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7e984e63DecodeGithubComSubizHeaderNoti53(l, v)
}
func easyjson7e984e63DecodeGithubComSubizHeaderNoti54(in *jlexer.Lexer, out *PushNoti) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ctx":
			if in.IsNull() {
				in.Skip()
				out.Ctx = nil
			} else {
				if out.Ctx == nil {
					out.Ctx = new(common.Context)
				}
				(*out.Ctx).UnmarshalEasyJSON(in)
			}
		case "account_id":
			if in.IsNull() {
				in.Skip()
				out.AccountId = nil
			} else {
				if out.AccountId == nil {
					out.AccountId = new(string)
				}
				*out.AccountId = string(in.String())
			}
		case "user_id":
			if in.IsNull() {
				in.Skip()
				out.UserId = nil
			} else {
				if out.UserId == nil {
					out.UserId = new(string)
				}
				*out.UserId = string(in.String())
			}
		case "platform":
			if in.IsNull() {
				in.Skip()
				out.Platform = nil
			} else {
				if out.Platform == nil {
					out.Platform = new(string)
				}
				*out.Platform = string(in.String())
			}
		case "title":
			if in.IsNull() {
				in.Skip()
				out.Title = nil
			} else {
				if out.Title == nil {
					out.Title = new(string)
				}
				*out.Title = string(in.String())
			}
		case "body":
			if in.IsNull() {
				in.Skip()
				out.Body = nil
			} else {
				if out.Body == nil {
					out.Body = new(string)
				}
				*out.Body = string(in.String())
			}
		case "conversation_id":
			if in.IsNull() {
				in.Skip()
				out.ConversationId = nil
			} else {
				if out.ConversationId == nil {
					out.ConversationId = new(string)
				}
				*out.ConversationId = string(in.String())
			}
		case "sender_id":
			if in.IsNull() {
				in.Skip()
				out.SenderId = nil
			} else {
				if out.SenderId == nil {
					out.SenderId = new(string)
				}
				*out.SenderId = string(in.String())
			}
		case "sender_type":
			if in.IsNull() {
				in.Skip()
				out.SenderType = nil
			} else {
				if out.SenderType == nil {
					out.SenderType = new(string)
				}
				*out.SenderType = string(in.String())
			}
		case "icon_url":
			if in.IsNull() {
				in.Skip()
				out.IconUrl = nil
			} else {
				if out.IconUrl == nil {
					out.IconUrl = new(string)
				}
				*out.IconUrl = string(in.String())
			}
		case "last_page_view_url":
			if in.IsNull() {
				in.Skip()
				out.LastPageViewUrl = nil
			} else {
				if out.LastPageViewUrl == nil {
					out.LastPageViewUrl = new(string)
				}
				*out.LastPageViewUrl = string(in.String())
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
			}
		case "payload":
			if in.IsNull() {
				in.Skip()
				out.Payload = nil
			} else {
				out.Payload = in.Bytes()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7e984e63EncodeGithubComSubizHeaderNoti54(out *jwriter.Writer, in PushNoti) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ctx != nil {
		const prefix string = ",\"ctx\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Ctx).MarshalEasyJSON(out)
	}
	if in.AccountId != nil {
		const prefix string = ",\"account_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AccountId))
	}
	if in.UserId != nil {
		const prefix string = ",\"user_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.UserId))
	}
	if in.Platform != nil {
		const prefix string = ",\"platform\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Platform))
	}
	if in.Title != nil {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Title))
	}
	if in.Body != nil {
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Body))
	}
	if in.ConversationId != nil {
		const prefix string = ",\"conversation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.ConversationId))
	}
	if in.SenderId != nil {
		const prefix string = ",\"sender_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.SenderId))
	}
	if in.SenderType != nil {
		const prefix string = ",\"sender_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.SenderType))
	}
	if in.IconUrl != nil {
		const prefix string = ",\"icon_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.IconUrl))
	}
	if in.LastPageViewUrl != nil {
		const prefix string = ",\"last_page_view_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.LastPageViewUrl))
	}
	if in.Type != nil {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Type))
	}
	if len(in.Payload) != 0 {
		const prefix string = ",\"payload\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Payload)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PushNoti) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7e984e63EncodeGithubComSubizHeaderNoti54(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PushNoti) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7e984e63EncodeGithubComSubizHeaderNoti54(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PushNoti) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7e984e63DecodeGithubComSubizHeaderNoti54(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PushNoti) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7e984e63DecodeGithubComSubizHeaderNoti54(l, v)
}
func easyjson7e984e63DecodeGithubComSubizHeaderNoti55(in *jlexer.Lexer, out *Message) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fcm_token":
			if in.IsNull() {
				in.Skip()
				out.FcmToken = nil
			} else {
				if out.FcmToken == nil {
					out.FcmToken = new(string)
				}
				*out.FcmToken = string(in.String())
			}
		case "title":
			if in.IsNull() {
				in.Skip()
				out.Title = nil
			} else {
				if out.Title == nil {
					out.Title = new(string)
				}
				*out.Title = string(in.String())
			}
		case "body":
			if in.IsNull() {
				in.Skip()
				out.Body = nil
			} else {
				if out.Body == nil {
					out.Body = new(string)
				}
				*out.Body = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7e984e63EncodeGithubComSubizHeaderNoti55(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FcmToken != nil {
		const prefix string = ",\"fcm_token\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.FcmToken))
	}
	if in.Title != nil {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Title))
	}
	if in.Body != nil {
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Body))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7e984e63EncodeGithubComSubizHeaderNoti55(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7e984e63EncodeGithubComSubizHeaderNoti55(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7e984e63DecodeGithubComSubizHeaderNoti55(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7e984e63DecodeGithubComSubizHeaderNoti55(l, v)
}
