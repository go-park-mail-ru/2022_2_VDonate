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

func easyjson6fbf8f0cDecodeGithubComGoParkMailRu20222VDonateInternalModels(in *jlexer.Lexer, out *Subscription) {
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
		case "authorID":
			out.AuthorID = uint64(in.Uint64())
		case "subscriberID":
			out.SubscriberID = uint64(in.Uint64())
		case "authorSubscriptionID":
			out.AuthorSubscriptionID = uint64(in.Uint64())
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
func easyjson6fbf8f0cEncodeGithubComGoParkMailRu20222VDonateInternalModels(out *jwriter.Writer, in Subscription) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"authorID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.AuthorID))
	}
	{
		const prefix string = ",\"subscriberID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.SubscriberID))
	}
	{
		const prefix string = ",\"authorSubscriptionID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.AuthorSubscriptionID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Subscription) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6fbf8f0cEncodeGithubComGoParkMailRu20222VDonateInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Subscription) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6fbf8f0cEncodeGithubComGoParkMailRu20222VDonateInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Subscription) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6fbf8f0cDecodeGithubComGoParkMailRu20222VDonateInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Subscription) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6fbf8f0cDecodeGithubComGoParkMailRu20222VDonateInternalModels(l, v)
}
func easyjson6fbf8f0cDecodeGithubComGoParkMailRu20222VDonateInternalModels1(in *jlexer.Lexer, out *AuthorSubscription) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "authorID":
			out.AuthorID = uint64(in.Uint64())
		case "img":
			out.Img = string(in.String())
		case "tier":
			out.Tier = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "price":
			out.Price = uint64(in.Uint64())
		case "authorName":
			out.AuthorName = string(in.String())
		case "authorAvatar":
			out.AuthorAvatar = string(in.String())
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
func easyjson6fbf8f0cEncodeGithubComGoParkMailRu20222VDonateInternalModels1(out *jwriter.Writer, in AuthorSubscription) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"authorID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.AuthorID))
	}
	{
		const prefix string = ",\"img\":"
		out.RawString(prefix)
		out.String(string(in.Img))
	}
	{
		const prefix string = ",\"tier\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Tier))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Price))
	}
	if in.AuthorName != "" {
		const prefix string = ",\"authorName\":"
		out.RawString(prefix)
		out.String(string(in.AuthorName))
	}
	if in.AuthorAvatar != "" {
		const prefix string = ",\"authorAvatar\":"
		out.RawString(prefix)
		out.String(string(in.AuthorAvatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthorSubscription) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6fbf8f0cEncodeGithubComGoParkMailRu20222VDonateInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthorSubscription) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6fbf8f0cEncodeGithubComGoParkMailRu20222VDonateInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthorSubscription) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6fbf8f0cDecodeGithubComGoParkMailRu20222VDonateInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthorSubscription) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6fbf8f0cDecodeGithubComGoParkMailRu20222VDonateInternalModels1(l, v)
}