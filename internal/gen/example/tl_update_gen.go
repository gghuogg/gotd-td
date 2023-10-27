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

// Update represents TL type `update#b03e2ef8`.
//
// See https://localhost:80/doc/constructor/update for reference.
type Update struct {
	// Msg field of Update.
	Msg AbstractMessageClass
	// Delay field of Update.
	Delay int32
}

// UpdateTypeID is TL type id of Update.
const UpdateTypeID = 0xb03e2ef8

// Ensuring interfaces in compile-time for Update.
var (
	_ bin.Encoder     = &Update{}
	_ bin.Decoder     = &Update{}
	_ bin.BareEncoder = &Update{}
	_ bin.BareDecoder = &Update{}
)

func (u *Update) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Msg == nil) {
		return false
	}
	if !(u.Delay == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *Update) String() string {
	if u == nil {
		return "Update(nil)"
	}
	type Alias Update
	return fmt.Sprintf("Update%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Update) TypeID() uint32 {
	return UpdateTypeID
}

// TypeName returns name of type in TL schema.
func (*Update) TypeName() string {
	return "update"
}

// TypeInfo returns info about TL type.
func (u *Update) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "update",
		ID:   UpdateTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Msg",
			SchemaName: "msg",
		},
		{
			Name:       "Delay",
			SchemaName: "delay",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *Update) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode update#b03e2ef8 as nil")
	}
	b.PutID(UpdateTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *Update) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode update#b03e2ef8 as nil")
	}
	if u.Msg == nil {
		return fmt.Errorf("unable to encode update#b03e2ef8: field msg is nil")
	}
	if err := u.Msg.Encode(b); err != nil {
		return fmt.Errorf("unable to encode update#b03e2ef8: field msg: %w", err)
	}
	b.PutInt32(u.Delay)
	return nil
}

// Decode implements bin.Decoder.
func (u *Update) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode update#b03e2ef8 to nil")
	}
	if err := b.ConsumeID(UpdateTypeID); err != nil {
		return fmt.Errorf("unable to decode update#b03e2ef8: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *Update) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode update#b03e2ef8 to nil")
	}
	{
		value, err := DecodeAbstractMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode update#b03e2ef8: field msg: %w", err)
		}
		u.Msg = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode update#b03e2ef8: field delay: %w", err)
		}
		u.Delay = value
	}
	return nil
}

// GetMsg returns value of Msg field.
func (u *Update) GetMsg() (value AbstractMessageClass) {
	if u == nil {
		return
	}
	return u.Msg
}

// GetDelay returns value of Delay field.
func (u *Update) GetDelay() (value int32) {
	if u == nil {
		return
	}
	return u.Delay
}
