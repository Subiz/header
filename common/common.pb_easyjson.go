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

func easyjson89b94fcfDecodeGithubComSubizHeaderCommon(in *jlexer.Lexer, out *agentClient) {
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon(out *jwriter.Writer, in agentClient) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v agentClient) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v agentClient) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *agentClient) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *agentClient) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon1(in *jlexer.Lexer, out *String) {
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon1(out *jwriter.Writer, in String) {
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
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v String) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *String) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *String) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon1(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon2(in *jlexer.Lexer, out *Pong) {
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon2(out *jwriter.Writer, in Pong) {
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
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Pong) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Pong) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Pong) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon2(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon3(in *jlexer.Lexer, out *PingRequest) {
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon3(out *jwriter.Writer, in PingRequest) {
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
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PingRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PingRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PingRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon3(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon4(in *jlexer.Lexer, out *ObjectPathRequest) {
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon4(out *jwriter.Writer, in ObjectPathRequest) {
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
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ObjectPathRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ObjectPathRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ObjectPathRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon4(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon5(in *jlexer.Lexer, out *Int64) {
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon5(out *jwriter.Writer, in Int64) {
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
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Int64) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Int64) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Int64) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon5(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon6(in *jlexer.Lexer, out *Ids) {
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
					var v1 string
					v1 = string(in.String())
					out.Ids = append(out.Ids, v1)
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon6(out *jwriter.Writer, in Ids) {
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
			for v2, v3 := range in.Ids {
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
func (v Ids) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Ids) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Ids) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Ids) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon6(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon7(in *jlexer.Lexer, out *Id) {
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon7(out *jwriter.Writer, in Id) {
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
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Id) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Id) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Id) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon7(l, v)
}
func easyjson89b94fcfDecodeGithubComSubizHeaderCommon8(in *jlexer.Lexer, out *Error) {
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
		case "description":
			out.Description = string(in.String())
		case "debug":
			out.Debug = string(in.String())
		case "hash":
			out.Hash = string(in.String())
		case "class":
			out.Class = int32(in.Int32())
		case "stack":
			out.Stack = string(in.String())
		case "created":
			out.Created = int64(in.Int64())
		case "code":
			out.Code = string(in.String())
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
func easyjson89b94fcfEncodeGithubComSubizHeaderCommon8(out *jwriter.Writer, in Error) {
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
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Debug != "" {
		const prefix string = ",\"debug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Debug))
	}
	if in.Hash != "" {
		const prefix string = ",\"hash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Hash))
	}
	if in.Class != 0 {
		const prefix string = ",\"class\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Class))
	}
	if in.Stack != "" {
		const prefix string = ",\"stack\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Stack))
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
	if in.Code != "" {
		const prefix string = ",\"code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Code))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Error) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Error) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGithubComSubizHeaderCommon8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Error) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89b94fcfDecodeGithubComSubizHeaderCommon8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Error) UnmarshalEasyJSON(l *jlexer.Lexer) {
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
		case "reply_topic":
			out.ReplyTopic = string(in.String())
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
				out.Tracing = in.Bytes()
			}
		case "reply_key":
			out.ReplyKey = string(in.String())
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
		case "topic":
			out.Topic = string(in.String())
		case "partition":
			out.Partition = int32(in.Int32())
		case "offset":
			out.Offset = int64(in.Int64())
		case "term":
			out.Term = uint64(in.Uint64())
		case "router_topic":
			out.RouterTopic = string(in.String())
		case "idempotency_key":
			out.IdempotencyKey = string(in.String())
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
	if in.ReplyTopic != "" {
		const prefix string = ",\"reply_topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ReplyTopic))
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
	if len(in.Tracing) != 0 {
		const prefix string = ",\"tracing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Tracing)
	}
	if in.ReplyKey != "" {
		const prefix string = ",\"reply_key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ReplyKey))
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
	if in.Term != 0 {
		const prefix string = ",\"term\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Term))
	}
	if in.RouterTopic != "" {
		const prefix string = ",\"router_topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RouterTopic))
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
