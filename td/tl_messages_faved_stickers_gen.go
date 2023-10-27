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

// MessagesFavedStickersNotModified represents TL type `messages.favedStickersNotModified#9e8fa6d3`.
type MessagesFavedStickersNotModified struct {
}

// MessagesFavedStickersNotModifiedTypeID is TL type id of MessagesFavedStickersNotModified.
const MessagesFavedStickersNotModifiedTypeID = 0x9e8fa6d3

// construct implements constructor of MessagesFavedStickersClass.
func (f MessagesFavedStickersNotModified) construct() MessagesFavedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFavedStickersNotModified.
var (
	_ bin.Encoder     = &MessagesFavedStickersNotModified{}
	_ bin.Decoder     = &MessagesFavedStickersNotModified{}
	_ bin.BareEncoder = &MessagesFavedStickersNotModified{}
	_ bin.BareDecoder = &MessagesFavedStickersNotModified{}

	_ MessagesFavedStickersClass = &MessagesFavedStickersNotModified{}
)

func (f *MessagesFavedStickersNotModified) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFavedStickersNotModified) String() string {
	if f == nil {
		return "MessagesFavedStickersNotModified(nil)"
	}
	type Alias MessagesFavedStickersNotModified
	return fmt.Sprintf("MessagesFavedStickersNotModified%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFavedStickersNotModified) TypeID() uint32 {
	return MessagesFavedStickersNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFavedStickersNotModified) TypeName() string {
	return "messages.favedStickersNotModified"
}

// TypeInfo returns info about TL type.
func (f *MessagesFavedStickersNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.favedStickersNotModified",
		ID:   MessagesFavedStickersNotModifiedTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFavedStickersNotModified) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.favedStickersNotModified#9e8fa6d3 as nil")
	}
	b.PutID(MessagesFavedStickersNotModifiedTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesFavedStickersNotModified) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.favedStickersNotModified#9e8fa6d3 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesFavedStickersNotModified) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.favedStickersNotModified#9e8fa6d3 to nil")
	}
	if err := b.ConsumeID(MessagesFavedStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.favedStickersNotModified#9e8fa6d3: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesFavedStickersNotModified) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.favedStickersNotModified#9e8fa6d3 to nil")
	}
	return nil
}

// MessagesFavedStickers represents TL type `messages.favedStickers#2cb51097`.
type MessagesFavedStickers struct {
	// Hash field of MessagesFavedStickers.
	Hash int64
	// Packs field of MessagesFavedStickers.
	Packs []StickerPack
	// Stickers field of MessagesFavedStickers.
	Stickers []DocumentClass
}

// MessagesFavedStickersTypeID is TL type id of MessagesFavedStickers.
const MessagesFavedStickersTypeID = 0x2cb51097

// construct implements constructor of MessagesFavedStickersClass.
func (f MessagesFavedStickers) construct() MessagesFavedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFavedStickers.
var (
	_ bin.Encoder     = &MessagesFavedStickers{}
	_ bin.Decoder     = &MessagesFavedStickers{}
	_ bin.BareEncoder = &MessagesFavedStickers{}
	_ bin.BareDecoder = &MessagesFavedStickers{}

	_ MessagesFavedStickersClass = &MessagesFavedStickers{}
)

func (f *MessagesFavedStickers) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Hash == 0) {
		return false
	}
	if !(f.Packs == nil) {
		return false
	}
	if !(f.Stickers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFavedStickers) String() string {
	if f == nil {
		return "MessagesFavedStickers(nil)"
	}
	type Alias MessagesFavedStickers
	return fmt.Sprintf("MessagesFavedStickers%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFavedStickers) TypeID() uint32 {
	return MessagesFavedStickersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFavedStickers) TypeName() string {
	return "messages.favedStickers"
}

// TypeInfo returns info about TL type.
func (f *MessagesFavedStickers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.favedStickers",
		ID:   MessagesFavedStickersTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Packs",
			SchemaName: "packs",
		},
		{
			Name:       "Stickers",
			SchemaName: "stickers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFavedStickers) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.favedStickers#2cb51097 as nil")
	}
	b.PutID(MessagesFavedStickersTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesFavedStickers) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.favedStickers#2cb51097 as nil")
	}
	b.PutLong(f.Hash)
	b.PutVectorHeader(len(f.Packs))
	for idx, v := range f.Packs {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.favedStickers#2cb51097: field packs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Stickers))
	for idx, v := range f.Stickers {
		if v == nil {
			return fmt.Errorf("unable to encode messages.favedStickers#2cb51097: field stickers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.favedStickers#2cb51097: field stickers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesFavedStickers) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.favedStickers#2cb51097 to nil")
	}
	if err := b.ConsumeID(MessagesFavedStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.favedStickers#2cb51097: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesFavedStickers) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.favedStickers#2cb51097 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.favedStickers#2cb51097: field hash: %w", err)
		}
		f.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.favedStickers#2cb51097: field packs: %w", err)
		}

		if headerLen > 0 {
			f.Packs = make([]StickerPack, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StickerPack
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.favedStickers#2cb51097: field packs: %w", err)
			}
			f.Packs = append(f.Packs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.favedStickers#2cb51097: field stickers: %w", err)
		}

		if headerLen > 0 {
			f.Stickers = make([]DocumentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.favedStickers#2cb51097: field stickers: %w", err)
			}
			f.Stickers = append(f.Stickers, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (f *MessagesFavedStickers) GetHash() (value int64) {
	if f == nil {
		return
	}
	return f.Hash
}

// GetPacks returns value of Packs field.
func (f *MessagesFavedStickers) GetPacks() (value []StickerPack) {
	if f == nil {
		return
	}
	return f.Packs
}

// GetStickers returns value of Stickers field.
func (f *MessagesFavedStickers) GetStickers() (value []DocumentClass) {
	if f == nil {
		return
	}
	return f.Stickers
}

// MessagesFavedStickersClassName is schema name of MessagesFavedStickersClass.
const MessagesFavedStickersClassName = "messages.FavedStickers"

// MessagesFavedStickersClass represents messages.FavedStickers generic type.
//
// Example:
//
//	g, err := td.DecodeMessagesFavedStickers(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *td.MessagesFavedStickersNotModified: // messages.favedStickersNotModified#9e8fa6d3
//	case *td.MessagesFavedStickers: // messages.favedStickers#2cb51097
//	default: panic(v)
//	}
type MessagesFavedStickersClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesFavedStickersClass

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

// DecodeMessagesFavedStickers implements binary de-serialization for MessagesFavedStickersClass.
func DecodeMessagesFavedStickers(buf *bin.Buffer) (MessagesFavedStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesFavedStickersNotModifiedTypeID:
		// Decoding messages.favedStickersNotModified#9e8fa6d3.
		v := MessagesFavedStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFavedStickersClass: %w", err)
		}
		return &v, nil
	case MessagesFavedStickersTypeID:
		// Decoding messages.favedStickers#2cb51097.
		v := MessagesFavedStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFavedStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesFavedStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesFavedStickers boxes the MessagesFavedStickersClass providing a helper.
type MessagesFavedStickersBox struct {
	FavedStickers MessagesFavedStickersClass
}

// Decode implements bin.Decoder for MessagesFavedStickersBox.
func (b *MessagesFavedStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesFavedStickersBox to nil")
	}
	v, err := DecodeMessagesFavedStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FavedStickers = v
	return nil
}

// Encode implements bin.Encode for MessagesFavedStickersBox.
func (b *MessagesFavedStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FavedStickers == nil {
		return fmt.Errorf("unable to encode MessagesFavedStickersClass as nil")
	}
	return b.FavedStickers.Encode(buf)
}
