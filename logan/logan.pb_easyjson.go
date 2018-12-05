// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package logan

import (
	json "encoding/json"
	common "git.subiz.net/header/common"
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

func easyjson1241f361DecodeGitSubizNetHeaderLogan(in *jlexer.Lexer, out *Log) {
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
		case "trace_id":
			out.TraceId = string(in.String())
		case "created":
			out.Created = int64(in.Int64())
		case "level":
			out.Level = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Tags = append(out.Tags, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "debug":
			if in.IsNull() {
				in.Skip()
				out.Debug = nil
			} else {
				if out.Debug == nil {
					out.Debug = new(Debug)
				}
				(*out.Debug).UnmarshalEasyJSON(in)
			}
		case "message":
			if in.IsNull() {
				in.Skip()
				out.Message = nil
			} else {
				out.Message = in.Bytes()
			}
		case "service_name":
			out.ServiceName = string(in.String())
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
func easyjson1241f361EncodeGitSubizNetHeaderLogan(out *jwriter.Writer, in Log) {
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
	if in.TraceId != "" {
		const prefix string = ",\"trace_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TraceId))
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
	if in.Level != "" {
		const prefix string = ",\"level\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Level))
	}
	if len(in.Tags) != 0 {
		const prefix string = ",\"tags\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.Tags {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if in.Debug != nil {
		const prefix string = ",\"debug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Debug).MarshalEasyJSON(out)
	}
	if len(in.Message) != 0 {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Message)
	}
	if in.ServiceName != "" {
		const prefix string = ",\"service_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ServiceName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Log) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1241f361EncodeGitSubizNetHeaderLogan(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Log) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1241f361EncodeGitSubizNetHeaderLogan(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Log) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1241f361DecodeGitSubizNetHeaderLogan(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Log) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1241f361DecodeGitSubizNetHeaderLogan(l, v)
}
func easyjson1241f361DecodeGitSubizNetHeaderLogan1(in *jlexer.Lexer, out *KafkaInfo) {
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
			out.Topic = string(in.String())
		case "partition":
			out.Partition = int32(in.Int32())
		case "offset":
			out.Offset = int64(in.Int64())
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
func easyjson1241f361EncodeGitSubizNetHeaderLogan1(out *jwriter.Writer, in KafkaInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Topic != "" {
		const prefix string = ",\"topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Topic))
	}
	if in.Partition != 0 {
		const prefix string = ",\"partition\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Partition))
	}
	if in.Offset != 0 {
		const prefix string = ",\"offset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Offset))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v KafkaInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1241f361EncodeGitSubizNetHeaderLogan1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v KafkaInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1241f361EncodeGitSubizNetHeaderLogan1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *KafkaInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1241f361DecodeGitSubizNetHeaderLogan1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *KafkaInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1241f361DecodeGitSubizNetHeaderLogan1(l, v)
}
func easyjson1241f361DecodeGitSubizNetHeaderLogan2(in *jlexer.Lexer, out *Debug) {
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
		case "stack_trace":
			if in.IsNull() {
				in.Skip()
				out.StackTrace = nil
			} else {
				out.StackTrace = in.Bytes()
			}
		case "hostname":
			out.Hostname = string(in.String())
		case "kafka":
			if in.IsNull() {
				in.Skip()
				out.Kafka = nil
			} else {
				if out.Kafka == nil {
					out.Kafka = new(KafkaInfo)
				}
				(*out.Kafka).UnmarshalEasyJSON(in)
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
func easyjson1241f361EncodeGitSubizNetHeaderLogan2(out *jwriter.Writer, in Debug) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.StackTrace) != 0 {
		const prefix string = ",\"stack_trace\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.StackTrace)
	}
	if in.Hostname != "" {
		const prefix string = ",\"hostname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Hostname))
	}
	if in.Kafka != nil {
		const prefix string = ",\"kafka\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Kafka).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Debug) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1241f361EncodeGitSubizNetHeaderLogan2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Debug) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1241f361EncodeGitSubizNetHeaderLogan2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Debug) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1241f361DecodeGitSubizNetHeaderLogan2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Debug) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1241f361DecodeGitSubizNetHeaderLogan2(l, v)
}
