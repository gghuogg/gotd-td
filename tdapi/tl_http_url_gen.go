// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// HTTPURL represents TL type `httpUrl#87b775a6`.
type HTTPURL struct {
	// The URL
	URL string
}

// HTTPURLTypeID is TL type id of HTTPURL.
const HTTPURLTypeID = 0x87b775a6

// Ensuring interfaces in compile-time for HTTPURL.
var (
	_ bin.Encoder     = &HTTPURL{}
	_ bin.Decoder     = &HTTPURL{}
	_ bin.BareEncoder = &HTTPURL{}
	_ bin.BareDecoder = &HTTPURL{}
)

func (h *HTTPURL) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *HTTPURL) String() string {
	if h == nil {
		return "HTTPURL(nil)"
	}
	type Alias HTTPURL
	return fmt.Sprintf("HTTPURL%+v", Alias(*h))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HTTPURL) TypeID() uint32 {
	return HTTPURLTypeID
}

// TypeName returns name of type in TL schema.
func (*HTTPURL) TypeName() string {
	return "httpUrl"
}

// TypeInfo returns info about TL type.
func (h *HTTPURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "httpUrl",
		ID:   HTTPURLTypeID,
	}
	if h == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (h *HTTPURL) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode httpUrl#87b775a6 as nil")
	}
	b.PutID(HTTPURLTypeID)
	return h.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (h *HTTPURL) EncodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode httpUrl#87b775a6 as nil")
	}
	b.PutString(h.URL)
	return nil
}

// Decode implements bin.Decoder.
func (h *HTTPURL) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode httpUrl#87b775a6 to nil")
	}
	if err := b.ConsumeID(HTTPURLTypeID); err != nil {
		return fmt.Errorf("unable to decode httpUrl#87b775a6: %w", err)
	}
	return h.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (h *HTTPURL) DecodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode httpUrl#87b775a6 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode httpUrl#87b775a6: field url: %w", err)
		}
		h.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (h *HTTPURL) EncodeTDLibJSON(b tdjson.Encoder) error {
	if h == nil {
		return fmt.Errorf("can't encode httpUrl#87b775a6 as nil")
	}
	b.ObjStart()
	b.PutID("httpUrl")
	b.Comma()
	b.FieldStart("url")
	b.PutString(h.URL)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (h *HTTPURL) DecodeTDLibJSON(b tdjson.Decoder) error {
	if h == nil {
		return fmt.Errorf("can't decode httpUrl#87b775a6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("httpUrl"); err != nil {
				return fmt.Errorf("unable to decode httpUrl#87b775a6: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode httpUrl#87b775a6: field url: %w", err)
			}
			h.URL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (h *HTTPURL) GetURL() (value string) {
	if h == nil {
		return
	}
	return h.URL
}
