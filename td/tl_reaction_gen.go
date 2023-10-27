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

// ReactionEmpty represents TL type `reactionEmpty#79f5d419`.
type ReactionEmpty struct {
}

// ReactionEmptyTypeID is TL type id of ReactionEmpty.
const ReactionEmptyTypeID = 0x79f5d419

// construct implements constructor of ReactionClass.
func (r ReactionEmpty) construct() ReactionClass { return &r }

// Ensuring interfaces in compile-time for ReactionEmpty.
var (
	_ bin.Encoder     = &ReactionEmpty{}
	_ bin.Decoder     = &ReactionEmpty{}
	_ bin.BareEncoder = &ReactionEmpty{}
	_ bin.BareDecoder = &ReactionEmpty{}

	_ ReactionClass = &ReactionEmpty{}
)

func (r *ReactionEmpty) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionEmpty) String() string {
	if r == nil {
		return "ReactionEmpty(nil)"
	}
	type Alias ReactionEmpty
	return fmt.Sprintf("ReactionEmpty%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionEmpty) TypeID() uint32 {
	return ReactionEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionEmpty) TypeName() string {
	return "reactionEmpty"
}

// TypeInfo returns info about TL type.
func (r *ReactionEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionEmpty",
		ID:   ReactionEmptyTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReactionEmpty) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionEmpty#79f5d419 as nil")
	}
	b.PutID(ReactionEmptyTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionEmpty) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionEmpty#79f5d419 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionEmpty) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionEmpty#79f5d419 to nil")
	}
	if err := b.ConsumeID(ReactionEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionEmpty#79f5d419: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionEmpty) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionEmpty#79f5d419 to nil")
	}
	return nil
}

// ReactionEmoji represents TL type `reactionEmoji#1b2286b8`.
type ReactionEmoji struct {
	// Emoticon field of ReactionEmoji.
	Emoticon string
}

// ReactionEmojiTypeID is TL type id of ReactionEmoji.
const ReactionEmojiTypeID = 0x1b2286b8

// construct implements constructor of ReactionClass.
func (r ReactionEmoji) construct() ReactionClass { return &r }

// Ensuring interfaces in compile-time for ReactionEmoji.
var (
	_ bin.Encoder     = &ReactionEmoji{}
	_ bin.Decoder     = &ReactionEmoji{}
	_ bin.BareEncoder = &ReactionEmoji{}
	_ bin.BareDecoder = &ReactionEmoji{}

	_ ReactionClass = &ReactionEmoji{}
)

func (r *ReactionEmoji) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Emoticon == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionEmoji) String() string {
	if r == nil {
		return "ReactionEmoji(nil)"
	}
	type Alias ReactionEmoji
	return fmt.Sprintf("ReactionEmoji%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionEmoji) TypeID() uint32 {
	return ReactionEmojiTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionEmoji) TypeName() string {
	return "reactionEmoji"
}

// TypeInfo returns info about TL type.
func (r *ReactionEmoji) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionEmoji",
		ID:   ReactionEmojiTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReactionEmoji) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionEmoji#1b2286b8 as nil")
	}
	b.PutID(ReactionEmojiTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionEmoji) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionEmoji#1b2286b8 as nil")
	}
	b.PutString(r.Emoticon)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionEmoji) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionEmoji#1b2286b8 to nil")
	}
	if err := b.ConsumeID(ReactionEmojiTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionEmoji#1b2286b8: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionEmoji) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionEmoji#1b2286b8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reactionEmoji#1b2286b8: field emoticon: %w", err)
		}
		r.Emoticon = value
	}
	return nil
}

// GetEmoticon returns value of Emoticon field.
func (r *ReactionEmoji) GetEmoticon() (value string) {
	if r == nil {
		return
	}
	return r.Emoticon
}

// ReactionCustomEmoji represents TL type `reactionCustomEmoji#8935fc73`.
type ReactionCustomEmoji struct {
	// DocumentID field of ReactionCustomEmoji.
	DocumentID int64
}

// ReactionCustomEmojiTypeID is TL type id of ReactionCustomEmoji.
const ReactionCustomEmojiTypeID = 0x8935fc73

// construct implements constructor of ReactionClass.
func (r ReactionCustomEmoji) construct() ReactionClass { return &r }

// Ensuring interfaces in compile-time for ReactionCustomEmoji.
var (
	_ bin.Encoder     = &ReactionCustomEmoji{}
	_ bin.Decoder     = &ReactionCustomEmoji{}
	_ bin.BareEncoder = &ReactionCustomEmoji{}
	_ bin.BareDecoder = &ReactionCustomEmoji{}

	_ ReactionClass = &ReactionCustomEmoji{}
)

func (r *ReactionCustomEmoji) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.DocumentID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionCustomEmoji) String() string {
	if r == nil {
		return "ReactionCustomEmoji(nil)"
	}
	type Alias ReactionCustomEmoji
	return fmt.Sprintf("ReactionCustomEmoji%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionCustomEmoji) TypeID() uint32 {
	return ReactionCustomEmojiTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionCustomEmoji) TypeName() string {
	return "reactionCustomEmoji"
}

// TypeInfo returns info about TL type.
func (r *ReactionCustomEmoji) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionCustomEmoji",
		ID:   ReactionCustomEmojiTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DocumentID",
			SchemaName: "document_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReactionCustomEmoji) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionCustomEmoji#8935fc73 as nil")
	}
	b.PutID(ReactionCustomEmojiTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionCustomEmoji) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionCustomEmoji#8935fc73 as nil")
	}
	b.PutLong(r.DocumentID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionCustomEmoji) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionCustomEmoji#8935fc73 to nil")
	}
	if err := b.ConsumeID(ReactionCustomEmojiTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionCustomEmoji#8935fc73: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionCustomEmoji) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionCustomEmoji#8935fc73 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode reactionCustomEmoji#8935fc73: field document_id: %w", err)
		}
		r.DocumentID = value
	}
	return nil
}

// GetDocumentID returns value of DocumentID field.
func (r *ReactionCustomEmoji) GetDocumentID() (value int64) {
	if r == nil {
		return
	}
	return r.DocumentID
}

// ReactionClassName is schema name of ReactionClass.
const ReactionClassName = "Reaction"

// ReactionClass represents Reaction generic type.
//
// Example:
//
//	g, err := td.DecodeReaction(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *td.ReactionEmpty: // reactionEmpty#79f5d419
//	case *td.ReactionEmoji: // reactionEmoji#1b2286b8
//	case *td.ReactionCustomEmoji: // reactionCustomEmoji#8935fc73
//	default: panic(v)
//	}
type ReactionClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ReactionClass

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

// DecodeReaction implements binary de-serialization for ReactionClass.
func DecodeReaction(buf *bin.Buffer) (ReactionClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ReactionEmptyTypeID:
		// Decoding reactionEmpty#79f5d419.
		v := ReactionEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionClass: %w", err)
		}
		return &v, nil
	case ReactionEmojiTypeID:
		// Decoding reactionEmoji#1b2286b8.
		v := ReactionEmoji{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionClass: %w", err)
		}
		return &v, nil
	case ReactionCustomEmojiTypeID:
		// Decoding reactionCustomEmoji#8935fc73.
		v := ReactionCustomEmoji{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReactionClass: %w", bin.NewUnexpectedID(id))
	}
}

// Reaction boxes the ReactionClass providing a helper.
type ReactionBox struct {
	Reaction ReactionClass
}

// Decode implements bin.Decoder for ReactionBox.
func (b *ReactionBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReactionBox to nil")
	}
	v, err := DecodeReaction(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Reaction = v
	return nil
}

// Encode implements bin.Encode for ReactionBox.
func (b *ReactionBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Reaction == nil {
		return fmt.Errorf("unable to encode ReactionClass as nil")
	}
	return b.Reaction.Encode(buf)
}
