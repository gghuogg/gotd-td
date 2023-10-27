// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// InputStickeredMediaPhoto represents TL type `inputStickeredMediaPhoto#4a992157`.
// A photo with stickers attached
//
// See https://core.telegram.org/constructor/inputStickeredMediaPhoto for reference.
type InputStickeredMediaPhoto struct {
	// The photo
	ID InputPhotoClass
}

// InputStickeredMediaPhotoTypeID is TL type id of InputStickeredMediaPhoto.
const InputStickeredMediaPhotoTypeID = 0x4a992157

// construct implements constructor of InputStickeredMediaClass.
func (i InputStickeredMediaPhoto) construct() InputStickeredMediaClass { return &i }

// Ensuring interfaces in compile-time for InputStickeredMediaPhoto.
var (
	_ bin.Encoder     = &InputStickeredMediaPhoto{}
	_ bin.Decoder     = &InputStickeredMediaPhoto{}
	_ bin.BareEncoder = &InputStickeredMediaPhoto{}
	_ bin.BareDecoder = &InputStickeredMediaPhoto{}

	_ InputStickeredMediaClass = &InputStickeredMediaPhoto{}
)

func (i *InputStickeredMediaPhoto) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickeredMediaPhoto) String() string {
	if i == nil {
		return "InputStickeredMediaPhoto(nil)"
	}
	type Alias InputStickeredMediaPhoto
	return fmt.Sprintf("InputStickeredMediaPhoto%+v", Alias(*i))
}

// FillFrom fills InputStickeredMediaPhoto from given interface.
func (i *InputStickeredMediaPhoto) FillFrom(from interface {
	GetID() (value InputPhotoClass)
}) {
	i.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickeredMediaPhoto) TypeID() uint32 {
	return InputStickeredMediaPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickeredMediaPhoto) TypeName() string {
	return "inputStickeredMediaPhoto"
}

// TypeInfo returns info about TL type.
func (i *InputStickeredMediaPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickeredMediaPhoto",
		ID:   InputStickeredMediaPhotoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStickeredMediaPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickeredMediaPhoto#4a992157 as nil")
	}
	b.PutID(InputStickeredMediaPhotoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickeredMediaPhoto) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickeredMediaPhoto#4a992157 as nil")
	}
	if i.ID == nil {
		return fmt.Errorf("unable to encode inputStickeredMediaPhoto#4a992157: field id is nil")
	}
	if err := i.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickeredMediaPhoto#4a992157: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickeredMediaPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickeredMediaPhoto#4a992157 to nil")
	}
	if err := b.ConsumeID(InputStickeredMediaPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickeredMediaPhoto#4a992157: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickeredMediaPhoto) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickeredMediaPhoto#4a992157 to nil")
	}
	{
		value, err := DecodeInputPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickeredMediaPhoto#4a992157: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputStickeredMediaPhoto) GetID() (value InputPhotoClass) {
	if i == nil {
		return
	}
	return i.ID
}

// InputStickeredMediaDocument represents TL type `inputStickeredMediaDocument#438865b`.
// A document with stickers attached
//
// See https://core.telegram.org/constructor/inputStickeredMediaDocument for reference.
type InputStickeredMediaDocument struct {
	// The document
	ID InputDocumentClass
}

// InputStickeredMediaDocumentTypeID is TL type id of InputStickeredMediaDocument.
const InputStickeredMediaDocumentTypeID = 0x438865b

// construct implements constructor of InputStickeredMediaClass.
func (i InputStickeredMediaDocument) construct() InputStickeredMediaClass { return &i }

// Ensuring interfaces in compile-time for InputStickeredMediaDocument.
var (
	_ bin.Encoder     = &InputStickeredMediaDocument{}
	_ bin.Decoder     = &InputStickeredMediaDocument{}
	_ bin.BareEncoder = &InputStickeredMediaDocument{}
	_ bin.BareDecoder = &InputStickeredMediaDocument{}

	_ InputStickeredMediaClass = &InputStickeredMediaDocument{}
)

func (i *InputStickeredMediaDocument) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickeredMediaDocument) String() string {
	if i == nil {
		return "InputStickeredMediaDocument(nil)"
	}
	type Alias InputStickeredMediaDocument
	return fmt.Sprintf("InputStickeredMediaDocument%+v", Alias(*i))
}

// FillFrom fills InputStickeredMediaDocument from given interface.
func (i *InputStickeredMediaDocument) FillFrom(from interface {
	GetID() (value InputDocumentClass)
}) {
	i.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickeredMediaDocument) TypeID() uint32 {
	return InputStickeredMediaDocumentTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickeredMediaDocument) TypeName() string {
	return "inputStickeredMediaDocument"
}

// TypeInfo returns info about TL type.
func (i *InputStickeredMediaDocument) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickeredMediaDocument",
		ID:   InputStickeredMediaDocumentTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStickeredMediaDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickeredMediaDocument#438865b as nil")
	}
	b.PutID(InputStickeredMediaDocumentTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickeredMediaDocument) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickeredMediaDocument#438865b as nil")
	}
	if i.ID == nil {
		return fmt.Errorf("unable to encode inputStickeredMediaDocument#438865b: field id is nil")
	}
	if err := i.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickeredMediaDocument#438865b: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickeredMediaDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickeredMediaDocument#438865b to nil")
	}
	if err := b.ConsumeID(InputStickeredMediaDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickeredMediaDocument#438865b: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickeredMediaDocument) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickeredMediaDocument#438865b to nil")
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickeredMediaDocument#438865b: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputStickeredMediaDocument) GetID() (value InputDocumentClass) {
	if i == nil {
		return
	}
	return i.ID
}

// InputStickeredMediaClassName is schema name of InputStickeredMediaClass.
const InputStickeredMediaClassName = "InputStickeredMedia"

// InputStickeredMediaClass represents InputStickeredMedia generic type.
//
// See https://core.telegram.org/type/InputStickeredMedia for reference.
//
// Example:
//
//	g, err := tg.DecodeInputStickeredMedia(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputStickeredMediaPhoto: // inputStickeredMediaPhoto#4a992157
//	case *tg.InputStickeredMediaDocument: // inputStickeredMediaDocument#438865b
//	default: panic(v)
//	}
type InputStickeredMediaClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputStickeredMediaClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputStickeredMedia implements binary de-serialization for InputStickeredMediaClass.
func DecodeInputStickeredMedia(buf *bin.Buffer) (InputStickeredMediaClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStickeredMediaPhotoTypeID:
		// Decoding inputStickeredMediaPhoto#4a992157.
		v := InputStickeredMediaPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickeredMediaClass: %w", err)
		}
		return &v, nil
	case InputStickeredMediaDocumentTypeID:
		// Decoding inputStickeredMediaDocument#438865b.
		v := InputStickeredMediaDocument{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickeredMediaClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStickeredMediaClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputStickeredMedia boxes the InputStickeredMediaClass providing a helper.
type InputStickeredMediaBox struct {
	InputStickeredMedia InputStickeredMediaClass
}

// Decode implements bin.Decoder for InputStickeredMediaBox.
func (b *InputStickeredMediaBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStickeredMediaBox to nil")
	}
	v, err := DecodeInputStickeredMedia(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStickeredMedia = v
	return nil
}

// Encode implements bin.Encode for InputStickeredMediaBox.
func (b *InputStickeredMediaBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputStickeredMedia == nil {
		return fmt.Errorf("unable to encode InputStickeredMediaClass as nil")
	}
	return b.InputStickeredMedia.Encode(buf)
}
