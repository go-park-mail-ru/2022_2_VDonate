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

func easyjsonB7188a36DecodeGithubComGoParkMailRu20222VDonateInternalModels(in *jlexer.Lexer, out *ResponseImageUsers) {
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
		case "userID":
			out.UserID = uint64(in.Uint64())
		case "username":
			out.Username = string(in.String())
		case "imgPath":
			out.ImgPath = string(in.String())
		case "about":
			out.About = string(in.String())
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
func easyjsonB7188a36EncodeGithubComGoParkMailRu20222VDonateInternalModels(out *jwriter.Writer, in ResponseImageUsers) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.UserID))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"imgPath\":"
		out.RawString(prefix)
		out.String(string(in.ImgPath))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseImageUsers) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB7188a36EncodeGithubComGoParkMailRu20222VDonateInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseImageUsers) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB7188a36EncodeGithubComGoParkMailRu20222VDonateInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseImageUsers) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB7188a36DecodeGithubComGoParkMailRu20222VDonateInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseImageUsers) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB7188a36DecodeGithubComGoParkMailRu20222VDonateInternalModels(l, v)
}
func easyjsonB7188a36DecodeGithubComGoParkMailRu20222VDonateInternalModels1(in *jlexer.Lexer, out *ResponseImageSubscription) {
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
		case "subscriptionID":
			out.SubscriptionID = uint64(in.Uint64())
		case "imgPath":
			out.ImgPath = string(in.String())
		case "authorName":
			out.AuthorName = string(in.String())
		case "authorImg":
			out.AuthorImg = string(in.String())
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
func easyjsonB7188a36EncodeGithubComGoParkMailRu20222VDonateInternalModels1(out *jwriter.Writer, in ResponseImageSubscription) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"subscriptionID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.SubscriptionID))
	}
	{
		const prefix string = ",\"imgPath\":"
		out.RawString(prefix)
		out.String(string(in.ImgPath))
	}
	{
		const prefix string = ",\"authorName\":"
		out.RawString(prefix)
		out.String(string(in.AuthorName))
	}
	{
		const prefix string = ",\"authorImg\":"
		out.RawString(prefix)
		out.String(string(in.AuthorImg))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseImageSubscription) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB7188a36EncodeGithubComGoParkMailRu20222VDonateInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseImageSubscription) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB7188a36EncodeGithubComGoParkMailRu20222VDonateInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseImageSubscription) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB7188a36DecodeGithubComGoParkMailRu20222VDonateInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseImageSubscription) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB7188a36DecodeGithubComGoParkMailRu20222VDonateInternalModels1(l, v)
}
func easyjsonB7188a36DecodeGithubComGoParkMailRu20222VDonateInternalModels2(in *jlexer.Lexer, out *ResponseImagePosts) {
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
		case "postID":
			out.PostID = uint64(in.Uint64())
		case "content":
			out.Content = string(in.String())
		case "dateCreated":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateCreated).UnmarshalJSON(data))
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
func easyjsonB7188a36EncodeGithubComGoParkMailRu20222VDonateInternalModels2(out *jwriter.Writer, in ResponseImagePosts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"postID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.PostID))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"dateCreated\":"
		out.RawString(prefix)
		out.Raw((in.DateCreated).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseImagePosts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB7188a36EncodeGithubComGoParkMailRu20222VDonateInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseImagePosts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB7188a36EncodeGithubComGoParkMailRu20222VDonateInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseImagePosts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB7188a36DecodeGithubComGoParkMailRu20222VDonateInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseImagePosts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB7188a36DecodeGithubComGoParkMailRu20222VDonateInternalModels2(l, v)
}
