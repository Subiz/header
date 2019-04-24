// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package avatar

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	auth "github.com/subiz/header/auth"
	common "github.com/subiz/header/common"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson8e21b7afDecodeGithubComSubizHeaderAvatar(in *jlexer.Lexer, out *Avatar) {
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
				easyjson8e21b7afDecodeGithubComSubizHeaderCommon(in, &*out.Ctx)
			}
		case "key":
			out.Key = string(in.String())
		case "avatar_url":
			out.AvatarUrl = string(in.String())
		case "auto_update":
			out.AutoUpdate = bool(in.Bool())
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
func easyjson8e21b7afEncodeGithubComSubizHeaderAvatar(out *jwriter.Writer, in Avatar) {
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
		easyjson8e21b7afEncodeGithubComSubizHeaderCommon(out, *in.Ctx)
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
	if in.AvatarUrl != "" {
		const prefix string = ",\"avatar_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvatarUrl))
	}
	if in.AutoUpdate {
		const prefix string = ",\"auto_update\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.AutoUpdate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Avatar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8e21b7afEncodeGithubComSubizHeaderAvatar(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Avatar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8e21b7afEncodeGithubComSubizHeaderAvatar(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Avatar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8e21b7afDecodeGithubComSubizHeaderAvatar(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Avatar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8e21b7afDecodeGithubComSubizHeaderAvatar(l, v)
}
func easyjson8e21b7afDecodeGithubComSubizHeaderCommon(in *jlexer.Lexer, out *common.Context) {
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
					out.Tracing = new(common.Tracing)
				}
				easyjson8e21b7afDecodeGithubComSubizHeaderCommon1(in, &*out.Tracing)
			}
		case "by_device":
			if in.IsNull() {
				in.Skip()
				out.ByDevice = nil
			} else {
				if out.ByDevice == nil {
					out.ByDevice = new(common.Device)
				}
				easyjson8e21b7afDecodeGithubComSubizHeaderCommon2(in, &*out.ByDevice)
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
func easyjson8e21b7afEncodeGithubComSubizHeaderCommon(out *jwriter.Writer, in common.Context) {
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
		easyjson8e21b7afEncodeGithubComSubizHeaderCommon1(out, *in.Tracing)
	}
	if in.ByDevice != nil {
		const prefix string = ",\"by_device\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson8e21b7afEncodeGithubComSubizHeaderCommon2(out, *in.ByDevice)
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
func easyjson8e21b7afDecodeGithubComSubizHeaderCommon2(in *jlexer.Lexer, out *common.Device) {
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
					var v4 string
					v4 = string(in.String())
					out.GaTrackingIds = append(out.GaTrackingIds, v4)
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
func easyjson8e21b7afEncodeGithubComSubizHeaderCommon2(out *jwriter.Writer, in common.Device) {
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
			for v5, v6 := range in.GaTrackingIds {
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
func easyjson8e21b7afDecodeGithubComSubizHeaderCommon1(in *jlexer.Lexer, out *common.Tracing) {
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
						out.Traces = make([]*common.Trace, 0, 8)
					} else {
						out.Traces = []*common.Trace{}
					}
				} else {
					out.Traces = (out.Traces)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *common.Trace
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(common.Trace)
						}
						easyjson8e21b7afDecodeGithubComSubizHeaderCommon3(in, &*v7)
					}
					out.Traces = append(out.Traces, v7)
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
func easyjson8e21b7afEncodeGithubComSubizHeaderCommon1(out *jwriter.Writer, in common.Tracing) {
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
			for v8, v9 := range in.Traces {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					easyjson8e21b7afEncodeGithubComSubizHeaderCommon3(out, *v9)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson8e21b7afDecodeGithubComSubizHeaderCommon3(in *jlexer.Lexer, out *common.Trace) {
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
func easyjson8e21b7afEncodeGithubComSubizHeaderCommon3(out *jwriter.Writer, in common.Trace) {
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
