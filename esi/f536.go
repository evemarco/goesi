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

func easyjson3c5fedd3DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdPositionList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseSystemsSystemIdPositionList, 0, 2)
			} else {
				*out = GetUniverseSystemsSystemIdPositionList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseSystemsSystemIdPosition
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
func easyjson3c5fedd3EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseSystemsSystemIdPositionList) {
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
func (v GetUniverseSystemsSystemIdPositionList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c5fedd3EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemsSystemIdPositionList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c5fedd3EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdPositionList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c5fedd3DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdPositionList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c5fedd3DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson3c5fedd3DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdPosition) {
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
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "z":
			out.Z = float64(in.Float64())
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
func easyjson3c5fedd3EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseSystemsSystemIdPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.X))
	}
	if in.Y != 0 {
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	if in.Z != 0 {
		const prefix string = ",\"z\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Z))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseSystemsSystemIdPosition) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c5fedd3EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemsSystemIdPosition) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c5fedd3EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdPosition) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c5fedd3DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdPosition) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c5fedd3DecodeGithubComAntihaxGoesiEsi1(l, v)
}
