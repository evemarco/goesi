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

func easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanetsPlanetIdOkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdPlanetsPlanetIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanetsPlanetIdOk
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
func easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdOkList) {
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
func (v GetCharactersCharacterIdPlanetsPlanetIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdOk) {
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
		case "links":
			if in.IsNull() {
				in.Skip()
				out.Links = nil
			} else {
				in.Delim('[')
				if out.Links == nil {
					if !in.IsDelim(']') {
						out.Links = make([]GetCharactersCharacterIdPlanetsPlanetIdLink, 0, 2)
					} else {
						out.Links = []GetCharactersCharacterIdPlanetsPlanetIdLink{}
					}
				} else {
					out.Links = (out.Links)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdPlanetsPlanetIdLink
					easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.Links = append(out.Links, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pins":
			if in.IsNull() {
				in.Skip()
				out.Pins = nil
			} else {
				in.Delim('[')
				if out.Pins == nil {
					if !in.IsDelim(']') {
						out.Pins = make([]GetCharactersCharacterIdPlanetsPlanetIdPin, 0, 1)
					} else {
						out.Pins = []GetCharactersCharacterIdPlanetsPlanetIdPin{}
					}
				} else {
					out.Pins = (out.Pins)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetCharactersCharacterIdPlanetsPlanetIdPin
					easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi3(in, &v5)
					out.Pins = append(out.Pins, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "routes":
			if in.IsNull() {
				in.Skip()
				out.Routes = nil
			} else {
				in.Delim('[')
				if out.Routes == nil {
					if !in.IsDelim(']') {
						out.Routes = make([]GetCharactersCharacterIdPlanetsPlanetIdRoute, 0, 1)
					} else {
						out.Routes = []GetCharactersCharacterIdPlanetsPlanetIdRoute{}
					}
				} else {
					out.Routes = (out.Routes)[:0]
				}
				for !in.IsDelim(']') {
					var v6 GetCharactersCharacterIdPlanetsPlanetIdRoute
					easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi4(in, &v6)
					out.Routes = append(out.Routes, v6)
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
func easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Links) != 0 {
		const prefix string = ",\"links\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.Links {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi2(out, v8)
			}
			out.RawByte(']')
		}
	}
	if len(in.Pins) != 0 {
		const prefix string = ",\"pins\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Pins {
				if v9 > 0 {
					out.RawByte(',')
				}
				easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi3(out, v10)
			}
			out.RawByte(']')
		}
	}
	if len(in.Routes) != 0 {
		const prefix string = ",\"routes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Routes {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi4(out, v12)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi4(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdRoute) {
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
		case "route_id":
			out.RouteId = int64(in.Int64())
		case "source_pin_id":
			out.SourcePinId = int64(in.Int64())
		case "destination_pin_id":
			out.DestinationPinId = int64(in.Int64())
		case "content_type_id":
			out.ContentTypeId = int32(in.Int32())
		case "quantity":
			out.Quantity = float64(in.Float64())
		case "waypoints":
			if in.IsNull() {
				in.Skip()
				out.Waypoints = nil
			} else {
				in.Delim('[')
				if out.Waypoints == nil {
					if !in.IsDelim(']') {
						out.Waypoints = make([]int64, 0, 8)
					} else {
						out.Waypoints = []int64{}
					}
				} else {
					out.Waypoints = (out.Waypoints)[:0]
				}
				for !in.IsDelim(']') {
					var v13 int64
					v13 = int64(in.Int64())
					out.Waypoints = append(out.Waypoints, v13)
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
func easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi4(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdRoute) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RouteId != 0 {
		const prefix string = ",\"route_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RouteId))
	}
	if in.SourcePinId != 0 {
		const prefix string = ",\"source_pin_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.SourcePinId))
	}
	if in.DestinationPinId != 0 {
		const prefix string = ",\"destination_pin_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DestinationPinId))
	}
	if in.ContentTypeId != 0 {
		const prefix string = ",\"content_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ContentTypeId))
	}
	if in.Quantity != 0 {
		const prefix string = ",\"quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Quantity))
	}
	if len(in.Waypoints) != 0 {
		const prefix string = ",\"waypoints\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.Waypoints {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v15))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdPin) {
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
		case "latitude":
			out.Latitude = float64(in.Float64())
		case "longitude":
			out.Longitude = float64(in.Float64())
		case "pin_id":
			out.PinId = int64(in.Int64())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "schematic_id":
			out.SchematicId = int32(in.Int32())
		case "extractor_details":
			(out.ExtractorDetails).UnmarshalEasyJSON(in)
		case "factory_details":
			easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi5(in, &out.FactoryDetails)
		case "contents":
			if in.IsNull() {
				in.Skip()
				out.Contents = nil
			} else {
				in.Delim('[')
				if out.Contents == nil {
					if !in.IsDelim(']') {
						out.Contents = make([]GetCharactersCharacterIdPlanetsPlanetIdContent, 0, 4)
					} else {
						out.Contents = []GetCharactersCharacterIdPlanetsPlanetIdContent{}
					}
				} else {
					out.Contents = (out.Contents)[:0]
				}
				for !in.IsDelim(']') {
					var v16 GetCharactersCharacterIdPlanetsPlanetIdContent
					(v16).UnmarshalEasyJSON(in)
					out.Contents = append(out.Contents, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "install_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.InstallTime).UnmarshalJSON(data))
			}
		case "expiry_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExpiryTime).UnmarshalJSON(data))
			}
		case "last_cycle_start":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastCycleStart).UnmarshalJSON(data))
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
func easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdPin) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Latitude != 0 {
		const prefix string = ",\"latitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Latitude))
	}
	if in.Longitude != 0 {
		const prefix string = ",\"longitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Longitude))
	}
	if in.PinId != 0 {
		const prefix string = ",\"pin_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PinId))
	}
	if in.TypeId != 0 {
		const prefix string = ",\"type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TypeId))
	}
	if in.SchematicId != 0 {
		const prefix string = ",\"schematic_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SchematicId))
	}
	if true {
		const prefix string = ",\"extractor_details\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.ExtractorDetails).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"factory_details\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi5(out, in.FactoryDetails)
	}
	if len(in.Contents) != 0 {
		const prefix string = ",\"contents\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v17, v18 := range in.Contents {
				if v17 > 0 {
					out.RawByte(',')
				}
				(v18).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"install_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.InstallTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"expiry_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ExpiryTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"last_cycle_start\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastCycleStart).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi5(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) {
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
		case "schematic_id":
			out.SchematicId = int32(in.Int32())
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
func easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi5(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SchematicId != 0 {
		const prefix string = ",\"schematic_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SchematicId))
	}
	out.RawByte('}')
}
func easyjson2db8f1a3DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdLink) {
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
		case "source_pin_id":
			out.SourcePinId = int64(in.Int64())
		case "destination_pin_id":
			out.DestinationPinId = int64(in.Int64())
		case "link_level":
			out.LinkLevel = int32(in.Int32())
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
func easyjson2db8f1a3EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdLink) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SourcePinId != 0 {
		const prefix string = ",\"source_pin_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.SourcePinId))
	}
	if in.DestinationPinId != 0 {
		const prefix string = ",\"destination_pin_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DestinationPinId))
	}
	if in.LinkLevel != 0 {
		const prefix string = ",\"link_level\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.LinkLevel))
	}
	out.RawByte('}')
}