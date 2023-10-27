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

// CDNPublicKey represents TL type `cdnPublicKey#c982eaba`.
type CDNPublicKey struct {
	// DCID field of CDNPublicKey.
	DCID int
	// PublicKey field of CDNPublicKey.
	PublicKey string
}

// CDNPublicKeyTypeID is TL type id of CDNPublicKey.
const CDNPublicKeyTypeID = 0xc982eaba

// Ensuring interfaces in compile-time for CDNPublicKey.
var (
	_ bin.Encoder     = &CDNPublicKey{}
	_ bin.Decoder     = &CDNPublicKey{}
	_ bin.BareEncoder = &CDNPublicKey{}
	_ bin.BareDecoder = &CDNPublicKey{}
)

func (c *CDNPublicKey) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.DCID == 0) {
		return false
	}
	if !(c.PublicKey == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CDNPublicKey) String() string {
	if c == nil {
		return "CDNPublicKey(nil)"
	}
	type Alias CDNPublicKey
	return fmt.Sprintf("CDNPublicKey%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CDNPublicKey) TypeID() uint32 {
	return CDNPublicKeyTypeID
}

// TypeName returns name of type in TL schema.
func (*CDNPublicKey) TypeName() string {
	return "cdnPublicKey"
}

// TypeInfo returns info about TL type.
func (c *CDNPublicKey) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "cdnPublicKey",
		ID:   CDNPublicKeyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "PublicKey",
			SchemaName: "public_key",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CDNPublicKey) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode cdnPublicKey#c982eaba as nil")
	}
	b.PutID(CDNPublicKeyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CDNPublicKey) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode cdnPublicKey#c982eaba as nil")
	}
	b.PutInt(c.DCID)
	b.PutString(c.PublicKey)
	return nil
}

// Decode implements bin.Decoder.
func (c *CDNPublicKey) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode cdnPublicKey#c982eaba to nil")
	}
	if err := b.ConsumeID(CDNPublicKeyTypeID); err != nil {
		return fmt.Errorf("unable to decode cdnPublicKey#c982eaba: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CDNPublicKey) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode cdnPublicKey#c982eaba to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode cdnPublicKey#c982eaba: field dc_id: %w", err)
		}
		c.DCID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode cdnPublicKey#c982eaba: field public_key: %w", err)
		}
		c.PublicKey = value
	}
	return nil
}

// GetDCID returns value of DCID field.
func (c *CDNPublicKey) GetDCID() (value int) {
	if c == nil {
		return
	}
	return c.DCID
}

// GetPublicKey returns value of PublicKey field.
func (c *CDNPublicKey) GetPublicKey() (value string) {
	if c == nil {
		return
	}
	return c.PublicKey
}
