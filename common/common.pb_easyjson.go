// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package common

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	auth "github.com/subiz/header/auth"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson89b94fcfDecodeGithubComSubizHeaderCommon(in *jlexer.Lexer, out *Tracing) {
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
		case "traces":
			if in.IsNull() {
				in.Skip()
				out.Traces = nil
			} else {
				in.Delim('[')
				if out.Traces == nil {
					if !in.IsDelim(']') {
						out.Traces = make([]*Trace, 0, 8)
					} else {
						out.Traces = []*Trace{}
					}
				} else {
					out.Traces = (out.Traces)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Trace
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Trace)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Traces = append(out.Traces, v1)
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon(out *jwriter.Writer, in Tracing) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Traces) != 0 {
		const prefix string = ",\"traces\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Traces {
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Tracing) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Tracing) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Tracing) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Tracing) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon1(in *jlexer.Lexer, out *Trace) {
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
		case "serviceName":
			out.ServiceName = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "duration":
			out.Duration = int64(in.Int64())
		case "started":
			out.Started = int64(in.Int64())
		case "ended":
			out.Ended = int64(in.Int64())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon1(out *jwriter.Writer, in Trace) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ServiceName != "" {
		const prefix string = ",\"serviceName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ServiceName))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Duration != 0 {
		const prefix string = ",\"duration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Duration))
	}
	if in.Started != 0 {
		const prefix string = ",\"started\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Started))
	}
	if in.Ended != 0 {
		const prefix string = ",\"ended\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Ended))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Trace) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Trace) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Trace) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Trace) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon1(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon2(in *jlexer.Lexer, out *String) {
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
		case "value":
			out.Value = string(in.String())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon2(out *jwriter.Writer, in String) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Value != "" {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v String) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v String) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *String) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *String) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon2(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon3(in *jlexer.Lexer, out *Pong) {
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
		case "message":
			out.Message = string(in.String())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon3(out *jwriter.Writer, in Pong) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Pong) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Pong) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Pong) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Pong) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon3(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon4(in *jlexer.Lexer, out *PingRequest) {
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
		case "message":
			out.Message = string(in.String())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon4(out *jwriter.Writer, in PingRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PingRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PingRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PingRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PingRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon4(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon5(in *jlexer.Lexer, out *ObjectPathRequest) {
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
					out.Ctx = new(Context)
				}
				(*out.Ctx).UnmarshalEasyJSON(in)
			}
		case "id":
			out.Id = string(in.String())
		case "path":
			out.Path = string(in.String())
		case "account_id":
			out.AccountId = string(in.String())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon5(out *jwriter.Writer, in ObjectPathRequest) {
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
	if in.Path != "" {
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Path))
	}
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ObjectPathRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ObjectPathRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ObjectPathRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ObjectPathRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon5(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon6(in *jlexer.Lexer, out *Int64) {
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
		case "value":
			out.Value = int64(in.Int64())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon6(out *jwriter.Writer, in Int64) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Value != 0 {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Int64) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Int64) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Int64) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Int64) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon6(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon7(in *jlexer.Lexer, out *Ids) {
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
					out.Ctx = new(Context)
				}
				(*out.Ctx).UnmarshalEasyJSON(in)
			}
		case "account_id":
			out.AccountId = string(in.String())
		case "ids":
			if in.IsNull() {
				in.Skip()
				out.Ids = nil
			} else {
				in.Delim('[')
				if out.Ids == nil {
					if !in.IsDelim(']') {
						out.Ids = make([]string, 0, 4)
					} else {
						out.Ids = []string{}
					}
				} else {
					out.Ids = (out.Ids)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Ids = append(out.Ids, v4)
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon7(out *jwriter.Writer, in Ids) {
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
	if len(in.Ids) != 0 {
		const prefix string = ",\"ids\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Ids {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Ids) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Ids) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Ids) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Ids) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon7(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon8(in *jlexer.Lexer, out *Id) {
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
					out.Ctx = new(Context)
				}
				(*out.Ctx).UnmarshalEasyJSON(in)
			}
		case "account_id":
			out.AccountId = string(in.String())
		case "id":
			out.Id = string(in.String())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon8(out *jwriter.Writer, in Id) {
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Id) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Id) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Id) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Id) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon8(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon9(in *jlexer.Lexer, out *Empty) {
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
					out.Ctx = new(Context)
				}
				(*out.Ctx).UnmarshalEasyJSON(in)
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon9(out *jwriter.Writer, in Empty) {
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Empty) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Empty) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Empty) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Empty) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon9(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon10(in *jlexer.Lexer, out *Device) {
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
		case "ip":
			out.Ip = string(in.String())
		case "user_agent":
			out.UserAgent = string(in.String())
		case "screen_resolution":
			out.ScreenResolution = string(in.String())
		case "timezone":
			out.Timezone = string(in.String())
		case "language":
			out.Language = string(in.String())
		case "referrer":
			out.Referrer = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "platform":
			out.Platform = string(in.String())
		case "source_referrer":
			out.SourceReferrer = string(in.String())
		case "ga_tracking_ids":
			if in.IsNull() {
				in.Skip()
				out.GaTrackingIds = nil
			} else {
				in.Delim('[')
				if out.GaTrackingIds == nil {
					if !in.IsDelim(']') {
						out.GaTrackingIds = make([]string, 0, 4)
					} else {
						out.GaTrackingIds = []string{}
					}
				} else {
					out.GaTrackingIds = (out.GaTrackingIds)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.GaTrackingIds = append(out.GaTrackingIds, v7)
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon10(out *jwriter.Writer, in Device) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ip != "" {
		const prefix string = ",\"ip\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ip))
	}
	if in.UserAgent != "" {
		const prefix string = ",\"user_agent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserAgent))
	}
	if in.ScreenResolution != "" {
		const prefix string = ",\"screen_resolution\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ScreenResolution))
	}
	if in.Timezone != "" {
		const prefix string = ",\"timezone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Timezone))
	}
	if in.Language != "" {
		const prefix string = ",\"language\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Language))
	}
	if in.Referrer != "" {
		const prefix string = ",\"referrer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Referrer))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.Platform != "" {
		const prefix string = ",\"platform\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Platform))
	}
	if in.SourceReferrer != "" {
		const prefix string = ",\"source_referrer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SourceReferrer))
	}
	if len(in.GaTrackingIds) != 0 {
		const prefix string = ",\"ga_tracking_ids\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.GaTrackingIds {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Device) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Device) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Device) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Device) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon10(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon11(in *jlexer.Lexer, out *Context) {
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
			out.EventId = string(in.String())
		case "state":
			if in.IsNull() {
				in.Skip()
				out.State = nil
			} else {
				out.State = in.Bytes()
			}
		case "node":
			out.Node = string(in.String())
		case "credential":
			if in.IsNull() {
				in.Skip()
				out.Credential = nil
			} else {
				if out.Credential == nil {
					out.Credential = new(auth.Credential)
				}
				(*out.Credential).UnmarshalEasyJSON(in)
			}
		case "tracing":
			if in.IsNull() {
				in.Skip()
				out.Tracing = nil
			} else {
				if out.Tracing == nil {
					out.Tracing = new(Tracing)
				}
				(*out.Tracing).UnmarshalEasyJSON(in)
			}
		case "by_device":
			if in.IsNull() {
				in.Skip()
				out.ByDevice = nil
			} else {
				if out.ByDevice == nil {
					out.ByDevice = new(Device)
				}
				(*out.ByDevice).UnmarshalEasyJSON(in)
			}
		case "sub_topic":
			out.SubTopic = string(in.String())
		case "kafka_partition":
			out.KafkaPartition = int32(in.Int32())
		case "kafka_offset":
			out.KafkaOffset = int64(in.Int64())
		case "kafka_term":
			out.KafkaTerm = uint64(in.Uint64())
		case "idempotency_key":
			out.IdempotencyKey = string(in.String())
		case "env":
			out.Env = string(in.String())
		case "kafka_key":
			out.KafkaKey = string(in.String())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon11(out *jwriter.Writer, in Context) {
	out.RawByte('{')
	first := true
	_ = first
	if in.EventId != "" {
		const prefix string = ",\"event_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EventId))
	}
	if len(in.State) != 0 {
		const prefix string = ",\"state\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.State)
	}
	if in.Node != "" {
		const prefix string = ",\"node\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Node))
	}
	if in.Credential != nil {
		const prefix string = ",\"credential\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Credential).MarshalEasyJSON(out)
	}
	if in.Tracing != nil {
		const prefix string = ",\"tracing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Tracing).MarshalEasyJSON(out)
	}
	if in.ByDevice != nil {
		const prefix string = ",\"by_device\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ByDevice).MarshalEasyJSON(out)
	}
	if in.SubTopic != "" {
		const prefix string = ",\"sub_topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SubTopic))
	}
	if in.KafkaPartition != 0 {
		const prefix string = ",\"kafka_partition\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.KafkaPartition))
	}
	if in.KafkaOffset != 0 {
		const prefix string = ",\"kafka_offset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.KafkaOffset))
	}
	if in.KafkaTerm != 0 {
		const prefix string = ",\"kafka_term\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.KafkaTerm))
	}
	if in.IdempotencyKey != "" {
		const prefix string = ",\"idempotency_key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IdempotencyKey))
	}
	if in.Env != "" {
		const prefix string = ",\"env\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Env))
	}
	if in.KafkaKey != "" {
		const prefix string = ",\"kafka_key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.KafkaKey))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Context) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Context) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Context) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Context) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon11(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon12(in *jlexer.Lexer, out *Bool) {
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
		case "value":
			out.Value = bool(in.Bool())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon12(out *jwriter.Writer, in Bool) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Value {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bool) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bool) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bool) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bool) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon12(l, v)
}
