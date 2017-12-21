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

func easyjson8404fadcDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetIncursions200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetIncursions200OkList, 0, 1)
			} else {
				*out = GetIncursions200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetIncursions200Ok
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
func easyjson8404fadcEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetIncursions200OkList) {
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
func (v GetIncursions200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8404fadcEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetIncursions200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8404fadcEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetIncursions200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8404fadcDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetIncursions200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8404fadcDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson8404fadcDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetIncursions200Ok) {
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
		case "type":
			out.Type_ = string(in.String())
		case "state":
			out.State = string(in.String())
		case "influence":
			out.Influence = float64(in.Float64())
		case "has_boss":
			out.HasBoss = bool(in.Bool())
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "constellation_id":
			out.ConstellationId = int32(in.Int32())
		case "staging_solar_system_id":
			out.StagingSolarSystemId = int32(in.Int32())
		case "infested_solar_systems":
			if in.IsNull() {
				in.Skip()
				out.InfestedSolarSystems = nil
			} else {
				in.Delim('[')
				if out.InfestedSolarSystems == nil {
					if !in.IsDelim(']') {
						out.InfestedSolarSystems = make([]int32, 0, 16)
					} else {
						out.InfestedSolarSystems = []int32{}
					}
				} else {
					out.InfestedSolarSystems = (out.InfestedSolarSystems)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.InfestedSolarSystems = append(out.InfestedSolarSystems, v4)
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
func easyjson8404fadcEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetIncursions200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Type_ != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type_))
	}
	if in.State != "" {
		const prefix string = ",\"state\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.State))
	}
	if in.Influence != 0 {
		const prefix string = ",\"influence\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Influence))
	}
	if in.HasBoss {
		const prefix string = ",\"has_boss\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HasBoss))
	}
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	if in.ConstellationId != 0 {
		const prefix string = ",\"constellation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ConstellationId))
	}
	if in.StagingSolarSystemId != 0 {
		const prefix string = ",\"staging_solar_system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.StagingSolarSystemId))
	}
	if len(in.InfestedSolarSystems) != 0 {
		const prefix string = ",\"infested_solar_systems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.InfestedSolarSystems {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetIncursions200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8404fadcEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetIncursions200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8404fadcEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetIncursions200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8404fadcDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetIncursions200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8404fadcDecodeGithubComAntihaxGoesiEsi1(l, v)
}
