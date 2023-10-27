// Code generated by gotdgen, DO NOT EDIT.

package td

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

// PhoneGroupCallStreamRtmpURL represents TL type `phone.groupCallStreamRtmpUrl#2dbf3432`.
type PhoneGroupCallStreamRtmpURL struct {
	// URL field of PhoneGroupCallStreamRtmpURL.
	URL string
	// Key field of PhoneGroupCallStreamRtmpURL.
	Key string
}

// PhoneGroupCallStreamRtmpURLTypeID is TL type id of PhoneGroupCallStreamRtmpURL.
const PhoneGroupCallStreamRtmpURLTypeID = 0x2dbf3432

// Ensuring interfaces in compile-time for PhoneGroupCallStreamRtmpURL.
var (
	_ bin.Encoder     = &PhoneGroupCallStreamRtmpURL{}
	_ bin.Decoder     = &PhoneGroupCallStreamRtmpURL{}
	_ bin.BareEncoder = &PhoneGroupCallStreamRtmpURL{}
	_ bin.BareDecoder = &PhoneGroupCallStreamRtmpURL{}
)

func (g *PhoneGroupCallStreamRtmpURL) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.URL == "") {
		return false
	}
	if !(g.Key == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGroupCallStreamRtmpURL) String() string {
	if g == nil {
		return "PhoneGroupCallStreamRtmpURL(nil)"
	}
	type Alias PhoneGroupCallStreamRtmpURL
	return fmt.Sprintf("PhoneGroupCallStreamRtmpURL%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneGroupCallStreamRtmpURL) TypeID() uint32 {
	return PhoneGroupCallStreamRtmpURLTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneGroupCallStreamRtmpURL) TypeName() string {
	return "phone.groupCallStreamRtmpUrl"
}

// TypeInfo returns info about TL type.
func (g *PhoneGroupCallStreamRtmpURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.groupCallStreamRtmpUrl",
		ID:   PhoneGroupCallStreamRtmpURLTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Key",
			SchemaName: "key",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhoneGroupCallStreamRtmpURL) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.groupCallStreamRtmpUrl#2dbf3432 as nil")
	}
	b.PutID(PhoneGroupCallStreamRtmpURLTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PhoneGroupCallStreamRtmpURL) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.groupCallStreamRtmpUrl#2dbf3432 as nil")
	}
	b.PutString(g.URL)
	b.PutString(g.Key)
	return nil
}

// Decode implements bin.Decoder.
func (g *PhoneGroupCallStreamRtmpURL) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.groupCallStreamRtmpUrl#2dbf3432 to nil")
	}
	if err := b.ConsumeID(PhoneGroupCallStreamRtmpURLTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.groupCallStreamRtmpUrl#2dbf3432: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PhoneGroupCallStreamRtmpURL) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.groupCallStreamRtmpUrl#2dbf3432 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupCallStreamRtmpUrl#2dbf3432: field url: %w", err)
		}
		g.URL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupCallStreamRtmpUrl#2dbf3432: field key: %w", err)
		}
		g.Key = value
	}
	return nil
}

// GetURL returns value of URL field.
func (g *PhoneGroupCallStreamRtmpURL) GetURL() (value string) {
	if g == nil {
		return
	}
	return g.URL
}

// GetKey returns value of Key field.
func (g *PhoneGroupCallStreamRtmpURL) GetKey() (value string) {
	if g == nil {
		return
	}
	return g.Key
}
