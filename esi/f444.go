// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson191b3d72DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetOpportunitiesTasksTaskIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetOpportunitiesTasksTaskIdOkList, 0, 1)
			} else {
				*out = GetOpportunitiesTasksTaskIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetOpportunitiesTasksTaskIdOk
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson191b3d72EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetOpportunitiesTasksTaskIdOkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetOpportunitiesTasksTaskIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191b3d72EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOpportunitiesTasksTaskIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191b3d72EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOpportunitiesTasksTaskIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191b3d72DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOpportunitiesTasksTaskIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191b3d72DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson191b3d72DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetOpportunitiesTasksTaskIdOk) {
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
		case "task_id":
			out.TaskId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "notification":
			out.Notification = string(in.String())
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
func easyjson191b3d72EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetOpportunitiesTasksTaskIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.TaskId != 0 {
		const prefix string = ",\"task_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TaskId))
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
	if in.Notification != "" {
		const prefix string = ",\"notification\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Notification))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetOpportunitiesTasksTaskIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191b3d72EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOpportunitiesTasksTaskIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191b3d72EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOpportunitiesTasksTaskIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191b3d72DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOpportunitiesTasksTaskIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191b3d72DecodeGithubComAntihaxGoesiEsi1(l, v)
}