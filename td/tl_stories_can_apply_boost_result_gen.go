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

// StoriesCanApplyBoostOk represents TL type `stories.canApplyBoostOk#c3173587`.
type StoriesCanApplyBoostOk struct {
}

// StoriesCanApplyBoostOkTypeID is TL type id of StoriesCanApplyBoostOk.
const StoriesCanApplyBoostOkTypeID = 0xc3173587

// construct implements constructor of StoriesCanApplyBoostResultClass.
func (c StoriesCanApplyBoostOk) construct() StoriesCanApplyBoostResultClass { return &c }

// Ensuring interfaces in compile-time for StoriesCanApplyBoostOk.
var (
	_ bin.Encoder     = &StoriesCanApplyBoostOk{}
	_ bin.Decoder     = &StoriesCanApplyBoostOk{}
	_ bin.BareEncoder = &StoriesCanApplyBoostOk{}
	_ bin.BareDecoder = &StoriesCanApplyBoostOk{}

	_ StoriesCanApplyBoostResultClass = &StoriesCanApplyBoostOk{}
)

func (c *StoriesCanApplyBoostOk) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *StoriesCanApplyBoostOk) String() string {
	if c == nil {
		return "StoriesCanApplyBoostOk(nil)"
	}
	type Alias StoriesCanApplyBoostOk
	return fmt.Sprintf("StoriesCanApplyBoostOk%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesCanApplyBoostOk) TypeID() uint32 {
	return StoriesCanApplyBoostOkTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesCanApplyBoostOk) TypeName() string {
	return "stories.canApplyBoostOk"
}

// TypeInfo returns info about TL type.
func (c *StoriesCanApplyBoostOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.canApplyBoostOk",
		ID:   StoriesCanApplyBoostOkTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *StoriesCanApplyBoostOk) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stories.canApplyBoostOk#c3173587 as nil")
	}
	b.PutID(StoriesCanApplyBoostOkTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *StoriesCanApplyBoostOk) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stories.canApplyBoostOk#c3173587 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *StoriesCanApplyBoostOk) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stories.canApplyBoostOk#c3173587 to nil")
	}
	if err := b.ConsumeID(StoriesCanApplyBoostOkTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.canApplyBoostOk#c3173587: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *StoriesCanApplyBoostOk) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stories.canApplyBoostOk#c3173587 to nil")
	}
	return nil
}

// StoriesCanApplyBoostReplace represents TL type `stories.canApplyBoostReplace#712c4655`.
type StoriesCanApplyBoostReplace struct {
	// CurrentBoost field of StoriesCanApplyBoostReplace.
	CurrentBoost PeerClass
	// Chats field of StoriesCanApplyBoostReplace.
	Chats []ChatClass
}

// StoriesCanApplyBoostReplaceTypeID is TL type id of StoriesCanApplyBoostReplace.
const StoriesCanApplyBoostReplaceTypeID = 0x712c4655

// construct implements constructor of StoriesCanApplyBoostResultClass.
func (c StoriesCanApplyBoostReplace) construct() StoriesCanApplyBoostResultClass { return &c }

// Ensuring interfaces in compile-time for StoriesCanApplyBoostReplace.
var (
	_ bin.Encoder     = &StoriesCanApplyBoostReplace{}
	_ bin.Decoder     = &StoriesCanApplyBoostReplace{}
	_ bin.BareEncoder = &StoriesCanApplyBoostReplace{}
	_ bin.BareDecoder = &StoriesCanApplyBoostReplace{}

	_ StoriesCanApplyBoostResultClass = &StoriesCanApplyBoostReplace{}
)

func (c *StoriesCanApplyBoostReplace) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.CurrentBoost == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *StoriesCanApplyBoostReplace) String() string {
	if c == nil {
		return "StoriesCanApplyBoostReplace(nil)"
	}
	type Alias StoriesCanApplyBoostReplace
	return fmt.Sprintf("StoriesCanApplyBoostReplace%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesCanApplyBoostReplace) TypeID() uint32 {
	return StoriesCanApplyBoostReplaceTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesCanApplyBoostReplace) TypeName() string {
	return "stories.canApplyBoostReplace"
}

// TypeInfo returns info about TL type.
func (c *StoriesCanApplyBoostReplace) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.canApplyBoostReplace",
		ID:   StoriesCanApplyBoostReplaceTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CurrentBoost",
			SchemaName: "current_boost",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *StoriesCanApplyBoostReplace) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stories.canApplyBoostReplace#712c4655 as nil")
	}
	b.PutID(StoriesCanApplyBoostReplaceTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *StoriesCanApplyBoostReplace) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stories.canApplyBoostReplace#712c4655 as nil")
	}
	if c.CurrentBoost == nil {
		return fmt.Errorf("unable to encode stories.canApplyBoostReplace#712c4655: field current_boost is nil")
	}
	if err := c.CurrentBoost.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.canApplyBoostReplace#712c4655: field current_boost: %w", err)
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode stories.canApplyBoostReplace#712c4655: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.canApplyBoostReplace#712c4655: field chats element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *StoriesCanApplyBoostReplace) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stories.canApplyBoostReplace#712c4655 to nil")
	}
	if err := b.ConsumeID(StoriesCanApplyBoostReplaceTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.canApplyBoostReplace#712c4655: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *StoriesCanApplyBoostReplace) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stories.canApplyBoostReplace#712c4655 to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.canApplyBoostReplace#712c4655: field current_boost: %w", err)
		}
		c.CurrentBoost = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.canApplyBoostReplace#712c4655: field chats: %w", err)
		}

		if headerLen > 0 {
			c.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.canApplyBoostReplace#712c4655: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	return nil
}

// GetCurrentBoost returns value of CurrentBoost field.
func (c *StoriesCanApplyBoostReplace) GetCurrentBoost() (value PeerClass) {
	if c == nil {
		return
	}
	return c.CurrentBoost
}

// GetChats returns value of Chats field.
func (c *StoriesCanApplyBoostReplace) GetChats() (value []ChatClass) {
	if c == nil {
		return
	}
	return c.Chats
}

// StoriesCanApplyBoostResultClassName is schema name of StoriesCanApplyBoostResultClass.
const StoriesCanApplyBoostResultClassName = "stories.CanApplyBoostResult"

// StoriesCanApplyBoostResultClass represents stories.CanApplyBoostResult generic type.
//
// Example:
//
//	g, err := td.DecodeStoriesCanApplyBoostResult(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *td.StoriesCanApplyBoostOk: // stories.canApplyBoostOk#c3173587
//	case *td.StoriesCanApplyBoostReplace: // stories.canApplyBoostReplace#712c4655
//	default: panic(v)
//	}
type StoriesCanApplyBoostResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StoriesCanApplyBoostResultClass

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

// DecodeStoriesCanApplyBoostResult implements binary de-serialization for StoriesCanApplyBoostResultClass.
func DecodeStoriesCanApplyBoostResult(buf *bin.Buffer) (StoriesCanApplyBoostResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StoriesCanApplyBoostOkTypeID:
		// Decoding stories.canApplyBoostOk#c3173587.
		v := StoriesCanApplyBoostOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoriesCanApplyBoostResultClass: %w", err)
		}
		return &v, nil
	case StoriesCanApplyBoostReplaceTypeID:
		// Decoding stories.canApplyBoostReplace#712c4655.
		v := StoriesCanApplyBoostReplace{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoriesCanApplyBoostResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoriesCanApplyBoostResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// StoriesCanApplyBoostResult boxes the StoriesCanApplyBoostResultClass providing a helper.
type StoriesCanApplyBoostResultBox struct {
	CanApplyBoostResult StoriesCanApplyBoostResultClass
}

// Decode implements bin.Decoder for StoriesCanApplyBoostResultBox.
func (b *StoriesCanApplyBoostResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoriesCanApplyBoostResultBox to nil")
	}
	v, err := DecodeStoriesCanApplyBoostResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CanApplyBoostResult = v
	return nil
}

// Encode implements bin.Encode for StoriesCanApplyBoostResultBox.
func (b *StoriesCanApplyBoostResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CanApplyBoostResult == nil {
		return fmt.Errorf("unable to encode StoriesCanApplyBoostResultClass as nil")
	}
	return b.CanApplyBoostResult.Encode(buf)
}
