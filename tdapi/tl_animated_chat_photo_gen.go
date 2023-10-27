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

// AnimatedChatPhoto represents TL type `animatedChatPhoto#b719c2e`.
type AnimatedChatPhoto struct {
	// Animation width and height
	Length int32
	// Information about the animation file
	File File
	// Timestamp of the frame, used as a static chat photo
	MainFrameTimestamp float64
}

// AnimatedChatPhotoTypeID is TL type id of AnimatedChatPhoto.
const AnimatedChatPhotoTypeID = 0xb719c2e

// Ensuring interfaces in compile-time for AnimatedChatPhoto.
var (
	_ bin.Encoder     = &AnimatedChatPhoto{}
	_ bin.Decoder     = &AnimatedChatPhoto{}
	_ bin.BareEncoder = &AnimatedChatPhoto{}
	_ bin.BareDecoder = &AnimatedChatPhoto{}
)

func (a *AnimatedChatPhoto) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Length == 0) {
		return false
	}
	if !(a.File.Zero()) {
		return false
	}
	if !(a.MainFrameTimestamp == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AnimatedChatPhoto) String() string {
	if a == nil {
		return "AnimatedChatPhoto(nil)"
	}
	type Alias AnimatedChatPhoto
	return fmt.Sprintf("AnimatedChatPhoto%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AnimatedChatPhoto) TypeID() uint32 {
	return AnimatedChatPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*AnimatedChatPhoto) TypeName() string {
	return "animatedChatPhoto"
}

// TypeInfo returns info about TL type.
func (a *AnimatedChatPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "animatedChatPhoto",
		ID:   AnimatedChatPhotoTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Length",
			SchemaName: "length",
		},
		{
			Name:       "File",
			SchemaName: "file",
		},
		{
			Name:       "MainFrameTimestamp",
			SchemaName: "main_frame_timestamp",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AnimatedChatPhoto) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode animatedChatPhoto#b719c2e as nil")
	}
	b.PutID(AnimatedChatPhotoTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AnimatedChatPhoto) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode animatedChatPhoto#b719c2e as nil")
	}
	b.PutInt32(a.Length)
	if err := a.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode animatedChatPhoto#b719c2e: field file: %w", err)
	}
	b.PutDouble(a.MainFrameTimestamp)
	return nil
}

// Decode implements bin.Decoder.
func (a *AnimatedChatPhoto) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode animatedChatPhoto#b719c2e to nil")
	}
	if err := b.ConsumeID(AnimatedChatPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode animatedChatPhoto#b719c2e: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AnimatedChatPhoto) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode animatedChatPhoto#b719c2e to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode animatedChatPhoto#b719c2e: field length: %w", err)
		}
		a.Length = value
	}
	{
		if err := a.File.Decode(b); err != nil {
			return fmt.Errorf("unable to decode animatedChatPhoto#b719c2e: field file: %w", err)
		}
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode animatedChatPhoto#b719c2e: field main_frame_timestamp: %w", err)
		}
		a.MainFrameTimestamp = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AnimatedChatPhoto) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode animatedChatPhoto#b719c2e as nil")
	}
	b.ObjStart()
	b.PutID("animatedChatPhoto")
	b.Comma()
	b.FieldStart("length")
	b.PutInt32(a.Length)
	b.Comma()
	b.FieldStart("file")
	if err := a.File.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode animatedChatPhoto#b719c2e: field file: %w", err)
	}
	b.Comma()
	b.FieldStart("main_frame_timestamp")
	b.PutDouble(a.MainFrameTimestamp)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AnimatedChatPhoto) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode animatedChatPhoto#b719c2e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("animatedChatPhoto"); err != nil {
				return fmt.Errorf("unable to decode animatedChatPhoto#b719c2e: %w", err)
			}
		case "length":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode animatedChatPhoto#b719c2e: field length: %w", err)
			}
			a.Length = value
		case "file":
			if err := a.File.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode animatedChatPhoto#b719c2e: field file: %w", err)
			}
		case "main_frame_timestamp":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode animatedChatPhoto#b719c2e: field main_frame_timestamp: %w", err)
			}
			a.MainFrameTimestamp = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLength returns value of Length field.
func (a *AnimatedChatPhoto) GetLength() (value int32) {
	if a == nil {
		return
	}
	return a.Length
}

// GetFile returns value of File field.
func (a *AnimatedChatPhoto) GetFile() (value File) {
	if a == nil {
		return
	}
	return a.File
}

// GetMainFrameTimestamp returns value of MainFrameTimestamp field.
func (a *AnimatedChatPhoto) GetMainFrameTimestamp() (value float64) {
	if a == nil {
		return
	}
	return a.MainFrameTimestamp
}
