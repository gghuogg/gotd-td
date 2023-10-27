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

// Booster represents TL type `booster#e9e6380`.
type Booster struct {
	// UserID field of Booster.
	UserID int64
	// Expires field of Booster.
	Expires int
}

// BoosterTypeID is TL type id of Booster.
const BoosterTypeID = 0xe9e6380

// Ensuring interfaces in compile-time for Booster.
var (
	_ bin.Encoder     = &Booster{}
	_ bin.Decoder     = &Booster{}
	_ bin.BareEncoder = &Booster{}
	_ bin.BareDecoder = &Booster{}
)

func (b *Booster) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.UserID == 0) {
		return false
	}
	if !(b.Expires == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *Booster) String() string {
	if b == nil {
		return "Booster(nil)"
	}
	type Alias Booster
	return fmt.Sprintf("Booster%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Booster) TypeID() uint32 {
	return BoosterTypeID
}

// TypeName returns name of type in TL schema.
func (*Booster) TypeName() string {
	return "booster"
}

// TypeInfo returns info about TL type.
func (b *Booster) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "booster",
		ID:   BoosterTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *Booster) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode booster#e9e6380 as nil")
	}
	buf.PutID(BoosterTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *Booster) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode booster#e9e6380 as nil")
	}
	buf.PutLong(b.UserID)
	buf.PutInt(b.Expires)
	return nil
}

// Decode implements bin.Decoder.
func (b *Booster) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode booster#e9e6380 to nil")
	}
	if err := buf.ConsumeID(BoosterTypeID); err != nil {
		return fmt.Errorf("unable to decode booster#e9e6380: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *Booster) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode booster#e9e6380 to nil")
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode booster#e9e6380: field user_id: %w", err)
		}
		b.UserID = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode booster#e9e6380: field expires: %w", err)
		}
		b.Expires = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (b *Booster) GetUserID() (value int64) {
	if b == nil {
		return
	}
	return b.UserID
}

// GetExpires returns value of Expires field.
func (b *Booster) GetExpires() (value int) {
	if b == nil {
		return
	}
	return b.Expires
}
