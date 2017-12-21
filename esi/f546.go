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

func easyjsonBceff015DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetWarsWarIdAggressorList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetWarsWarIdAggressorList, 0, 4)
			} else {
				*out = GetWarsWarIdAggressorList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetWarsWarIdAggressor
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
func easyjsonBceff015EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetWarsWarIdAggressorList) {
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
func (v GetWarsWarIdAggressorList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBceff015EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWarsWarIdAggressorList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBceff015EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWarsWarIdAggressorList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBceff015DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWarsWarIdAggressorList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBceff015DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonBceff015DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetWarsWarIdAggressor) {
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
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "ships_killed":
			out.ShipsKilled = int32(in.Int32())
		case "isk_destroyed":
			out.IskDestroyed = float64(in.Float64())
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
func easyjsonBceff015EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetWarsWarIdAggressor) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CorporationId != 0 {
		const prefix string = ",\"corporation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CorporationId))
	}
	if in.AllianceId != 0 {
		const prefix string = ",\"alliance_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.AllianceId))
	}
	if in.ShipsKilled != 0 {
		const prefix string = ",\"ships_killed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ShipsKilled))
	}
	if in.IskDestroyed != 0 {
		const prefix string = ",\"isk_destroyed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.IskDestroyed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetWarsWarIdAggressor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBceff015EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWarsWarIdAggressor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBceff015EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWarsWarIdAggressor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBceff015DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWarsWarIdAggressor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBceff015DecodeGithubComAntihaxGoesiEsi1(l, v)
}
