// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjsonA818f49aDecodeGithubComGoParkMailRu20222VDonateInternalModels(in *jlexer.Lexer, out *Cookie) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "value":
			out.Value = string(in.String())
		case "userID":
			out.UserID = uint64(in.Uint64())
		case "expire_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Expires).UnmarshalJSON(data))
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
func easyjsonA818f49aEncodeGithubComGoParkMailRu20222VDonateInternalModels(out *jwriter.Writer, in Cookie) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix[1:])
		out.String(string(in.Value))
	}
	{
		const prefix string = ",\"userID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.UserID))
	}
	{
		const prefix string = ",\"expire_date\":"
		out.RawString(prefix)
		out.Raw((in.Expires).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Cookie) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA818f49aEncodeGithubComGoParkMailRu20222VDonateInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Cookie) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA818f49aEncodeGithubComGoParkMailRu20222VDonateInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Cookie) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA818f49aDecodeGithubComGoParkMailRu20222VDonateInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Cookie) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA818f49aDecodeGithubComGoParkMailRu20222VDonateInternalModels(l, v)
}