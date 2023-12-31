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

// MessageSponsor represents TL type `messageSponsor#a0b25071`.
type MessageSponsor struct {
	// Type of the sponsor
	Type MessageSponsorTypeClass
	// Photo of the sponsor; may be null if must not be shown
	Photo ChatPhotoInfo
	// Additional optional information about the sponsor to be shown along with the message
	Info string
}

// MessageSponsorTypeID is TL type id of MessageSponsor.
const MessageSponsorTypeID = 0xa0b25071

// Ensuring interfaces in compile-time for MessageSponsor.
var (
	_ bin.Encoder     = &MessageSponsor{}
	_ bin.Decoder     = &MessageSponsor{}
	_ bin.BareEncoder = &MessageSponsor{}
	_ bin.BareDecoder = &MessageSponsor{}
)

func (m *MessageSponsor) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Type == nil) {
		return false
	}
	if !(m.Photo.Zero()) {
		return false
	}
	if !(m.Info == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSponsor) String() string {
	if m == nil {
		return "MessageSponsor(nil)"
	}
	type Alias MessageSponsor
	return fmt.Sprintf("MessageSponsor%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSponsor) TypeID() uint32 {
	return MessageSponsorTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSponsor) TypeName() string {
	return "messageSponsor"
}

// TypeInfo returns info about TL type.
func (m *MessageSponsor) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSponsor",
		ID:   MessageSponsorTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Info",
			SchemaName: "info",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSponsor) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSponsor#a0b25071 as nil")
	}
	b.PutID(MessageSponsorTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSponsor) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSponsor#a0b25071 as nil")
	}
	if m.Type == nil {
		return fmt.Errorf("unable to encode messageSponsor#a0b25071: field type is nil")
	}
	if err := m.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageSponsor#a0b25071: field type: %w", err)
	}
	if err := m.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageSponsor#a0b25071: field photo: %w", err)
	}
	b.PutString(m.Info)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSponsor) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSponsor#a0b25071 to nil")
	}
	if err := b.ConsumeID(MessageSponsorTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSponsor#a0b25071: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSponsor) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSponsor#a0b25071 to nil")
	}
	{
		value, err := DecodeMessageSponsorType(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageSponsor#a0b25071: field type: %w", err)
		}
		m.Type = value
	}
	{
		if err := m.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageSponsor#a0b25071: field photo: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageSponsor#a0b25071: field info: %w", err)
		}
		m.Info = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSponsor) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSponsor#a0b25071 as nil")
	}
	b.ObjStart()
	b.PutID("messageSponsor")
	b.Comma()
	b.FieldStart("type")
	if m.Type == nil {
		return fmt.Errorf("unable to encode messageSponsor#a0b25071: field type is nil")
	}
	if err := m.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageSponsor#a0b25071: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("photo")
	if err := m.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageSponsor#a0b25071: field photo: %w", err)
	}
	b.Comma()
	b.FieldStart("info")
	b.PutString(m.Info)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageSponsor) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSponsor#a0b25071 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSponsor"); err != nil {
				return fmt.Errorf("unable to decode messageSponsor#a0b25071: %w", err)
			}
		case "type":
			value, err := DecodeTDLibJSONMessageSponsorType(b)
			if err != nil {
				return fmt.Errorf("unable to decode messageSponsor#a0b25071: field type: %w", err)
			}
			m.Type = value
		case "photo":
			if err := m.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode messageSponsor#a0b25071: field photo: %w", err)
			}
		case "info":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode messageSponsor#a0b25071: field info: %w", err)
			}
			m.Info = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetType returns value of Type field.
func (m *MessageSponsor) GetType() (value MessageSponsorTypeClass) {
	if m == nil {
		return
	}
	return m.Type
}

// GetPhoto returns value of Photo field.
func (m *MessageSponsor) GetPhoto() (value ChatPhotoInfo) {
	if m == nil {
		return
	}
	return m.Photo
}

// GetInfo returns value of Info field.
func (m *MessageSponsor) GetInfo() (value string) {
	if m == nil {
		return
	}
	return m.Info
}
