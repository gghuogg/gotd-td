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

// GetCustomEmojiReactionAnimationsRequest represents TL type `getCustomEmojiReactionAnimations#49748dd2`.
type GetCustomEmojiReactionAnimationsRequest struct {
}

// GetCustomEmojiReactionAnimationsRequestTypeID is TL type id of GetCustomEmojiReactionAnimationsRequest.
const GetCustomEmojiReactionAnimationsRequestTypeID = 0x49748dd2

// Ensuring interfaces in compile-time for GetCustomEmojiReactionAnimationsRequest.
var (
	_ bin.Encoder     = &GetCustomEmojiReactionAnimationsRequest{}
	_ bin.Decoder     = &GetCustomEmojiReactionAnimationsRequest{}
	_ bin.BareEncoder = &GetCustomEmojiReactionAnimationsRequest{}
	_ bin.BareDecoder = &GetCustomEmojiReactionAnimationsRequest{}
)

func (g *GetCustomEmojiReactionAnimationsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCustomEmojiReactionAnimationsRequest) String() string {
	if g == nil {
		return "GetCustomEmojiReactionAnimationsRequest(nil)"
	}
	type Alias GetCustomEmojiReactionAnimationsRequest
	return fmt.Sprintf("GetCustomEmojiReactionAnimationsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCustomEmojiReactionAnimationsRequest) TypeID() uint32 {
	return GetCustomEmojiReactionAnimationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCustomEmojiReactionAnimationsRequest) TypeName() string {
	return "getCustomEmojiReactionAnimations"
}

// TypeInfo returns info about TL type.
func (g *GetCustomEmojiReactionAnimationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCustomEmojiReactionAnimations",
		ID:   GetCustomEmojiReactionAnimationsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCustomEmojiReactionAnimationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCustomEmojiReactionAnimations#49748dd2 as nil")
	}
	b.PutID(GetCustomEmojiReactionAnimationsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCustomEmojiReactionAnimationsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCustomEmojiReactionAnimations#49748dd2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCustomEmojiReactionAnimationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCustomEmojiReactionAnimations#49748dd2 to nil")
	}
	if err := b.ConsumeID(GetCustomEmojiReactionAnimationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCustomEmojiReactionAnimations#49748dd2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCustomEmojiReactionAnimationsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCustomEmojiReactionAnimations#49748dd2 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetCustomEmojiReactionAnimationsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getCustomEmojiReactionAnimations#49748dd2 as nil")
	}
	b.ObjStart()
	b.PutID("getCustomEmojiReactionAnimations")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetCustomEmojiReactionAnimationsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getCustomEmojiReactionAnimations#49748dd2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getCustomEmojiReactionAnimations"); err != nil {
				return fmt.Errorf("unable to decode getCustomEmojiReactionAnimations#49748dd2: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCustomEmojiReactionAnimations invokes method getCustomEmojiReactionAnimations#49748dd2 returning error if any.
func (c *Client) GetCustomEmojiReactionAnimations(ctx context.Context) (*Stickers, error) {
	var result Stickers

	request := &GetCustomEmojiReactionAnimationsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
