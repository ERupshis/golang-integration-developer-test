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

func easyjsonD2b7633eDecodeGithubComErupshisGolangIntegrationDeveloperTestInternalModels(in *jlexer.Lexer, out *Game) {
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
			out.ID = int64(in.Int64())
		case "title":
			out.Title = string(in.String())
		case "thumbnail":
			out.Thumbnail = string(in.String())
		case "short_description":
			out.ShortDescription = string(in.String())
		case "game_url":
			out.GameURL = string(in.String())
		case "genre":
			out.Genre = string(in.String())
		case "platform":
			out.Platform = string(in.String())
		case "publisher":
			out.Publisher = string(in.String())
		case "developer":
			out.Developer = string(in.String())
		case "release_date":
			out.ReleaseDate = string(in.String())
		case "freetogame_profile_url":
			out.FreeToGameProfileURL = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComErupshisGolangIntegrationDeveloperTestInternalModels(out *jwriter.Writer, in Game) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"thumbnail\":"
		out.RawString(prefix)
		out.String(string(in.Thumbnail))
	}
	{
		const prefix string = ",\"short_description\":"
		out.RawString(prefix)
		out.String(string(in.ShortDescription))
	}
	{
		const prefix string = ",\"game_url\":"
		out.RawString(prefix)
		out.String(string(in.GameURL))
	}
	{
		const prefix string = ",\"genre\":"
		out.RawString(prefix)
		out.String(string(in.Genre))
	}
	{
		const prefix string = ",\"platform\":"
		out.RawString(prefix)
		out.String(string(in.Platform))
	}
	{
		const prefix string = ",\"publisher\":"
		out.RawString(prefix)
		out.String(string(in.Publisher))
	}
	{
		const prefix string = ",\"developer\":"
		out.RawString(prefix)
		out.String(string(in.Developer))
	}
	{
		const prefix string = ",\"release_date\":"
		out.RawString(prefix)
		out.String(string(in.ReleaseDate))
	}
	{
		const prefix string = ",\"freetogame_profile_url\":"
		out.RawString(prefix)
		out.String(string(in.FreeToGameProfileURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Game) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComErupshisGolangIntegrationDeveloperTestInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Game) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComErupshisGolangIntegrationDeveloperTestInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Game) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComErupshisGolangIntegrationDeveloperTestInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Game) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComErupshisGolangIntegrationDeveloperTestInternalModels(l, v)
}
