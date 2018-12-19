// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package scheduler

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

func easyjsonFa65a999DecodeGithubComSubizHeaderScheduler(in *jlexer.Lexer, out *Task) {
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
			out.Id = string(in.String())
		case "callback_time":
			out.CallbackTime = int64(in.Int64())
		case "topic":
			out.Topic = string(in.String())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				out.Data = in.Bytes()
			}
		case "key":
			out.Key = string(in.String())
		case "called":
			out.Called = int64(in.Int64())
		case "sec":
			out.Sec = int64(in.Int64())
		case "par":
			out.Par = string(in.String())
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
func easyjsonFa65a999EncodeGithubComSubizHeaderScheduler(out *jwriter.Writer, in Task) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	if in.CallbackTime != 0 {
		const prefix string = ",\"callback_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.CallbackTime))
	}
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
	if in.Key != "" {
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Key))
	}
	if in.Called != 0 {
		const prefix string = ",\"called\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Called))
	}
	if in.Sec != 0 {
		const prefix string = ",\"sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Sec))
	}
	if in.Par != "" {
		const prefix string = ",\"par\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Par))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Task) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa65a999EncodeGithubComSubizHeaderScheduler(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Task) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa65a999EncodeGithubComSubizHeaderScheduler(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Task) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa65a999DecodeGithubComSubizHeaderScheduler(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Task) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa65a999DecodeGithubComSubizHeaderScheduler(l, v)
}
func easyjsonFa65a999DecodeGithubComSubizHeaderScheduler1(in *jlexer.Lexer, out *Id) {
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
			out.Id = string(in.String())
		case "callback_time":
			out.CallbackTime = int64(in.Int64())
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
func easyjsonFa65a999EncodeGithubComSubizHeaderScheduler1(out *jwriter.Writer, in Id) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	if in.CallbackTime != 0 {
		const prefix string = ",\"callback_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.CallbackTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Id) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa65a999EncodeGithubComSubizHeaderScheduler1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Id) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa65a999EncodeGithubComSubizHeaderScheduler1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Id) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa65a999DecodeGithubComSubizHeaderScheduler1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Id) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa65a999DecodeGithubComSubizHeaderScheduler1(l, v)
}
