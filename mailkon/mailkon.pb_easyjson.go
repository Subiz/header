// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package mailkon

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon(in *jlexer.Lexer, out *UserAvail) {
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
		case "email_address":
			if in.IsNull() {
				in.Skip()
				out.EmailAddress = nil
			} else {
				if out.EmailAddress == nil {
					out.EmailAddress = new(string)
				}
				*out.EmailAddress = string(in.String())
			}
		case "availability":
			if in.IsNull() {
				in.Skip()
				out.Availability = nil
			} else {
				if out.Availability == nil {
					out.Availability = new(bool)
				}
				*out.Availability = bool(in.Bool())
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
func easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon(out *jwriter.Writer, in UserAvail) {
	out.RawByte('{')
	first := true
	_ = first
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
	if in.EmailAddress != nil {
		const prefix string = ",\"email_address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.EmailAddress))
	}
	if in.Availability != nil {
		const prefix string = ",\"availability\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.Availability))
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserAvail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserAvail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserAvail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserAvail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon(l, v)
}
func easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon1(in *jlexer.Lexer, out *User) {
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
		case "email_address":
			if in.IsNull() {
				in.Skip()
				out.EmailAddress = nil
			} else {
				if out.EmailAddress == nil {
					out.EmailAddress = new(string)
				}
				*out.EmailAddress = string(in.String())
			}
		case "seen":
			if in.IsNull() {
				in.Skip()
				out.Seen = nil
			} else {
				if out.Seen == nil {
					out.Seen = new(int64)
				}
				*out.Seen = int64(in.Int64())
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
func easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
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
	if in.EmailAddress != nil {
		const prefix string = ",\"email_address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.EmailAddress))
	}
	if in.Seen != nil {
		const prefix string = ",\"seen\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.Seen))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon1(l, v)
}
func easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon2(in *jlexer.Lexer, out *SendgridEvent) {
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
		case "event_id":
			if in.IsNull() {
				in.Skip()
				out.EventId = nil
			} else {
				if out.EventId == nil {
					out.EventId = new(string)
				}
				*out.EventId = string(in.String())
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
		case "message_id":
			if in.IsNull() {
				in.Skip()
				out.MessageId = nil
			} else {
				if out.MessageId == nil {
					out.MessageId = new(string)
				}
				*out.MessageId = string(in.String())
			}
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]string, 0, 4)
					} else {
						out.Data = []string{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Data = append(out.Data, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "full_message_id":
			if in.IsNull() {
				in.Skip()
				out.FullMessageId = nil
			} else {
				if out.FullMessageId == nil {
					out.FullMessageId = new(string)
				}
				*out.FullMessageId = string(in.String())
			}
		case "subject":
			if in.IsNull() {
				in.Skip()
				out.Subject = nil
			} else {
				if out.Subject == nil {
					out.Subject = new(string)
				}
				*out.Subject = string(in.String())
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
func easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon2(out *jwriter.Writer, in SendgridEvent) {
	out.RawByte('{')
	first := true
	_ = first
	if in.EventId != nil {
		const prefix string = ",\"event_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.EventId))
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
	if in.MessageId != nil {
		const prefix string = ",\"message_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.MessageId))
	}
	if len(in.Data) != 0 {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Data {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.FullMessageId != nil {
		const prefix string = ",\"full_message_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.FullMessageId))
	}
	if in.Subject != nil {
		const prefix string = ",\"subject\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Subject))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SendgridEvent) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SendgridEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SendgridEvent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SendgridEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon2(l, v)
}
func easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon3(in *jlexer.Lexer, out *Message) {
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
		case "message_id":
			if in.IsNull() {
				in.Skip()
				out.MessageId = nil
			} else {
				if out.MessageId == nil {
					out.MessageId = new(string)
				}
				*out.MessageId = string(in.String())
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
		case "from_addr":
			if in.IsNull() {
				in.Skip()
				out.FromAddr = nil
			} else {
				if out.FromAddr == nil {
					out.FromAddr = new(string)
				}
				*out.FromAddr = string(in.String())
			}
		case "to_addr":
			if in.IsNull() {
				in.Skip()
				out.ToAddr = nil
			} else {
				if out.ToAddr == nil {
					out.ToAddr = new(string)
				}
				*out.ToAddr = string(in.String())
			}
		case "subject":
			if in.IsNull() {
				in.Skip()
				out.Subject = nil
			} else {
				if out.Subject == nil {
					out.Subject = new(string)
				}
				*out.Subject = string(in.String())
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
		case "header":
			if in.IsNull() {
				in.Skip()
				out.Header = nil
			} else {
				if out.Header == nil {
					out.Header = new(string)
				}
				*out.Header = string(in.String())
			}
		case "attachments":
			if in.IsNull() {
				in.Skip()
				out.Attachments = nil
			} else {
				in.Delim('[')
				if out.Attachments == nil {
					if !in.IsDelim(']') {
						out.Attachments = make([]string, 0, 4)
					} else {
						out.Attachments = []string{}
					}
				} else {
					out.Attachments = (out.Attachments)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Attachments = append(out.Attachments, v4)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon3(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MessageId != nil {
		const prefix string = ",\"message_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.MessageId))
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
	if in.FromAddr != nil {
		const prefix string = ",\"from_addr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.FromAddr))
	}
	if in.ToAddr != nil {
		const prefix string = ",\"to_addr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.ToAddr))
	}
	if in.Subject != nil {
		const prefix string = ",\"subject\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Subject))
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
	if in.Header != nil {
		const prefix string = ",\"header\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Header))
	}
	if len(in.Attachments) != 0 {
		const prefix string = ",\"attachments\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Attachments {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon3(l, v)
}
func easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon4(in *jlexer.Lexer, out *Job) {
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
		case "topic":
			if in.IsNull() {
				in.Skip()
				out.Topic = nil
			} else {
				if out.Topic == nil {
					out.Topic = new(string)
				}
				*out.Topic = string(in.String())
			}
		case "partition":
			if in.IsNull() {
				in.Skip()
				out.Partition = nil
			} else {
				if out.Partition == nil {
					out.Partition = new(int32)
				}
				*out.Partition = int32(in.Int32())
			}
		case "offset":
			if in.IsNull() {
				in.Skip()
				out.Offset = nil
			} else {
				if out.Offset == nil {
					out.Offset = new(int64)
				}
				*out.Offset = int64(in.Int64())
			}
		case "content_type":
			if in.IsNull() {
				in.Skip()
				out.ContentType = nil
			} else {
				if out.ContentType == nil {
					out.ContentType = new(string)
				}
				*out.ContentType = string(in.String())
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
		case "request_id":
			if in.IsNull() {
				in.Skip()
				out.RequestId = nil
			} else {
				if out.RequestId == nil {
					out.RequestId = new(string)
				}
				*out.RequestId = string(in.String())
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
func easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon4(out *jwriter.Writer, in Job) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Topic != nil {
		const prefix string = ",\"topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Topic))
	}
	if in.Partition != nil {
		const prefix string = ",\"partition\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.Partition))
	}
	if in.Offset != nil {
		const prefix string = ",\"offset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.Offset))
	}
	if in.ContentType != nil {
		const prefix string = ",\"content_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.ContentType))
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
	if in.RequestId != nil {
		const prefix string = ",\"request_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.RequestId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Job) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Job) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Job) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Job) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon4(l, v)
}
func easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon5(in *jlexer.Lexer, out *HttpChunk) {
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
		case "id":
			if in.IsNull() {
				in.Skip()
				out.Id = nil
			} else {
				if out.Id == nil {
					out.Id = new(string)
				}
				*out.Id = string(in.String())
			}
		case "chunk_id":
			if in.IsNull() {
				in.Skip()
				out.ChunkId = nil
			} else {
				if out.ChunkId == nil {
					out.ChunkId = new(int32)
				}
				*out.ChunkId = int32(in.Int32())
			}
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				out.Data = in.Bytes()
			}
		case "chunk_size":
			if in.IsNull() {
				in.Skip()
				out.ChunkSize = nil
			} else {
				if out.ChunkSize == nil {
					out.ChunkSize = new(int32)
				}
				*out.ChunkSize = int32(in.Int32())
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
func easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon5(out *jwriter.Writer, in HttpChunk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != nil {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Id))
	}
	if in.ChunkId != nil {
		const prefix string = ",\"chunk_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.ChunkId))
	}
	if len(in.Data) != 0 {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Data)
	}
	if in.ChunkSize != nil {
		const prefix string = ",\"chunk_size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(*in.ChunkSize))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HttpChunk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HttpChunk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HttpChunk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HttpChunk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon5(l, v)
}
func easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon6(in *jlexer.Lexer, out *Domain) {
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
		case "domain":
			if in.IsNull() {
				in.Skip()
				out.Domain = nil
			} else {
				if out.Domain == nil {
					out.Domain = new(string)
				}
				*out.Domain = string(in.String())
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
		case "dnstype":
			if in.IsNull() {
				in.Skip()
				out.Dnstype = nil
			} else {
				if out.Dnstype == nil {
					out.Dnstype = new(string)
				}
				*out.Dnstype = string(in.String())
			}
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				if out.Data == nil {
					out.Data = new(string)
				}
				*out.Data = string(in.String())
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
func easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon6(out *jwriter.Writer, in Domain) {
	out.RawByte('{')
	first := true
	_ = first
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
	if in.Domain != nil {
		const prefix string = ",\"domain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Domain))
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
	if in.Dnstype != nil {
		const prefix string = ",\"dnstype\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Dnstype))
	}
	if in.Data != nil {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Data))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Domain) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Domain) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Domain) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Domain) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon6(l, v)
}
func easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon7(in *jlexer.Lexer, out *Address) {
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
		case "address":
			if in.IsNull() {
				in.Skip()
				out.Address = nil
			} else {
				if out.Address == nil {
					out.Address = new(string)
				}
				*out.Address = string(in.String())
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
func easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon7(out *jwriter.Writer, in Address) {
	out.RawByte('{')
	first := true
	_ = first
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
	if in.Address != nil {
		const prefix string = ",\"address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Address))
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Address) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Address) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbdc5db9EncodeGitSubizNetHeaderMailkon7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Address) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Address) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbdc5db9DecodeGitSubizNetHeaderMailkon7(l, v)
}