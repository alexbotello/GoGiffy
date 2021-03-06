package backgroundservice

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// ServiceName the Background Service that will be associated with the
// commands/events. Every Background Service operates independently, but they
// share the same API.
type ServiceName string

// String returns the ServiceName as string value.
func (t ServiceName) String() string {
	return string(t)
}

// ServiceName values.
const (
	ServiceNameBackgroundFetch ServiceName = "backgroundFetch"
	ServiceNameBackgroundSync  ServiceName = "backgroundSync"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ServiceName) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ServiceName) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ServiceName) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ServiceName(in.String()) {
	case ServiceNameBackgroundFetch:
		*t = ServiceNameBackgroundFetch
	case ServiceNameBackgroundSync:
		*t = ServiceNameBackgroundSync

	default:
		in.AddError(errors.New("unknown ServiceName value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ServiceName) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
