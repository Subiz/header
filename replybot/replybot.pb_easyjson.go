// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package replybot

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	event "github.com/subiz/header/event"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB4e27493DecodeGithubComSubizHeaderReplybot(in *jlexer.Lexer, out *Setting) {
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
			out.AccountId = string(in.String())
		case "actions":
			if in.IsNull() {
				in.Skip()
				out.Actions = nil
			} else {
				in.Delim('[')
				if out.Actions == nil {
					if !in.IsDelim(']') {
						out.Actions = make([]*Action, 0, 8)
					} else {
						out.Actions = []*Action{}
					}
				} else {
					out.Actions = (out.Actions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Action
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Action)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Actions = append(out.Actions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "created":
			out.Created = int64(in.Int64())
		case "updated":
			out.Updated = int64(in.Int64())
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
func easyjsonB4e27493EncodeGithubComSubizHeaderReplybot(out *jwriter.Writer, in Setting) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AccountId != "" {
		const prefix string = ",\"account_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AccountId))
	}
	if len(in.Actions) != 0 {
		const prefix string = ",\"actions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Actions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Created != 0 {
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Created))
	}
	if in.Updated != 0 {
		const prefix string = ",\"updated\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Updated))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Setting) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB4e27493EncodeGithubComSubizHeaderReplybot(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Setting) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB4e27493EncodeGithubComSubizHeaderReplybot(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Setting) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB4e27493DecodeGithubComSubizHeaderReplybot(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Setting) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB4e27493DecodeGithubComSubizHeaderReplybot(l, v)
}
func easyjsonB4e27493DecodeGithubComSubizHeaderReplybot1(in *jlexer.Lexer, out *JoinedConversation) {
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
		case "conversation_id":
			out.ConversationId = string(in.String())
		case "created":
			out.Created = int64(in.Int64())
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
func easyjsonB4e27493EncodeGithubComSubizHeaderReplybot1(out *jwriter.Writer, in JoinedConversation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConversationId != "" {
		const prefix string = ",\"conversation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ConversationId))
	}
	if in.Created != 0 {
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Created))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JoinedConversation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB4e27493EncodeGithubComSubizHeaderReplybot1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JoinedConversation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB4e27493EncodeGithubComSubizHeaderReplybot1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JoinedConversation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB4e27493DecodeGithubComSubizHeaderReplybot1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JoinedConversation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB4e27493DecodeGithubComSubizHeaderReplybot1(l, v)
}
func easyjsonB4e27493DecodeGithubComSubizHeaderReplybot2(in *jlexer.Lexer, out *Action) {
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
		case "event":
			if in.IsNull() {
				in.Skip()
				out.Event = nil
			} else {
				if out.Event == nil {
					out.Event = new(event.RawEvent)
				}
				(*out.Event).UnmarshalEasyJSON(in)
			}
		case "delay":
			out.Delay = int32(in.Int32())
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
func easyjsonB4e27493EncodeGithubComSubizHeaderReplybot2(out *jwriter.Writer, in Action) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Event != nil {
		const prefix string = ",\"event\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Event).MarshalEasyJSON(out)
	}
	if in.Delay != 0 {
		const prefix string = ",\"delay\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Delay))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Action) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB4e27493EncodeGithubComSubizHeaderReplybot2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Action) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB4e27493EncodeGithubComSubizHeaderReplybot2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Action) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB4e27493DecodeGithubComSubizHeaderReplybot2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Action) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB4e27493DecodeGithubComSubizHeaderReplybot2(l, v)
}
