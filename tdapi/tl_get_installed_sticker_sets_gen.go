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

// GetInstalledStickerSetsRequest represents TL type `getInstalledStickerSets#612ef6f6`.
type GetInstalledStickerSetsRequest struct {
	// Type of the sticker sets to return
	StickerType StickerTypeClass
}

// GetInstalledStickerSetsRequestTypeID is TL type id of GetInstalledStickerSetsRequest.
const GetInstalledStickerSetsRequestTypeID = 0x612ef6f6

// Ensuring interfaces in compile-time for GetInstalledStickerSetsRequest.
var (
	_ bin.Encoder     = &GetInstalledStickerSetsRequest{}
	_ bin.Decoder     = &GetInstalledStickerSetsRequest{}
	_ bin.BareEncoder = &GetInstalledStickerSetsRequest{}
	_ bin.BareDecoder = &GetInstalledStickerSetsRequest{}
)

func (g *GetInstalledStickerSetsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.StickerType == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetInstalledStickerSetsRequest) String() string {
	if g == nil {
		return "GetInstalledStickerSetsRequest(nil)"
	}
	type Alias GetInstalledStickerSetsRequest
	return fmt.Sprintf("GetInstalledStickerSetsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetInstalledStickerSetsRequest) TypeID() uint32 {
	return GetInstalledStickerSetsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetInstalledStickerSetsRequest) TypeName() string {
	return "getInstalledStickerSets"
}

// TypeInfo returns info about TL type.
func (g *GetInstalledStickerSetsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getInstalledStickerSets",
		ID:   GetInstalledStickerSetsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StickerType",
			SchemaName: "sticker_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetInstalledStickerSetsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getInstalledStickerSets#612ef6f6 as nil")
	}
	b.PutID(GetInstalledStickerSetsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetInstalledStickerSetsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getInstalledStickerSets#612ef6f6 as nil")
	}
	if g.StickerType == nil {
		return fmt.Errorf("unable to encode getInstalledStickerSets#612ef6f6: field sticker_type is nil")
	}
	if err := g.StickerType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getInstalledStickerSets#612ef6f6: field sticker_type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetInstalledStickerSetsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getInstalledStickerSets#612ef6f6 to nil")
	}
	if err := b.ConsumeID(GetInstalledStickerSetsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getInstalledStickerSets#612ef6f6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetInstalledStickerSetsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getInstalledStickerSets#612ef6f6 to nil")
	}
	{
		value, err := DecodeStickerType(b)
		if err != nil {
			return fmt.Errorf("unable to decode getInstalledStickerSets#612ef6f6: field sticker_type: %w", err)
		}
		g.StickerType = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetInstalledStickerSetsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getInstalledStickerSets#612ef6f6 as nil")
	}
	b.ObjStart()
	b.PutID("getInstalledStickerSets")
	b.Comma()
	b.FieldStart("sticker_type")
	if g.StickerType == nil {
		return fmt.Errorf("unable to encode getInstalledStickerSets#612ef6f6: field sticker_type is nil")
	}
	if err := g.StickerType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getInstalledStickerSets#612ef6f6: field sticker_type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetInstalledStickerSetsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getInstalledStickerSets#612ef6f6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getInstalledStickerSets"); err != nil {
				return fmt.Errorf("unable to decode getInstalledStickerSets#612ef6f6: %w", err)
			}
		case "sticker_type":
			value, err := DecodeTDLibJSONStickerType(b)
			if err != nil {
				return fmt.Errorf("unable to decode getInstalledStickerSets#612ef6f6: field sticker_type: %w", err)
			}
			g.StickerType = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStickerType returns value of StickerType field.
func (g *GetInstalledStickerSetsRequest) GetStickerType() (value StickerTypeClass) {
	if g == nil {
		return
	}
	return g.StickerType
}

// GetInstalledStickerSets invokes method getInstalledStickerSets#612ef6f6 returning error if any.
func (c *Client) GetInstalledStickerSets(ctx context.Context, stickertype StickerTypeClass) (*StickerSets, error) {
	var result StickerSets

	request := &GetInstalledStickerSetsRequest{
		StickerType: stickertype,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
