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

// MessagesGetAvailableReactionsRequest represents TL type `messages.getAvailableReactions#18dea0ac`.
type MessagesGetAvailableReactionsRequest struct {
	// Hash field of MessagesGetAvailableReactionsRequest.
	Hash int
}

// MessagesGetAvailableReactionsRequestTypeID is TL type id of MessagesGetAvailableReactionsRequest.
const MessagesGetAvailableReactionsRequestTypeID = 0x18dea0ac

// Ensuring interfaces in compile-time for MessagesGetAvailableReactionsRequest.
var (
	_ bin.Encoder     = &MessagesGetAvailableReactionsRequest{}
	_ bin.Decoder     = &MessagesGetAvailableReactionsRequest{}
	_ bin.BareEncoder = &MessagesGetAvailableReactionsRequest{}
	_ bin.BareDecoder = &MessagesGetAvailableReactionsRequest{}
)

func (g *MessagesGetAvailableReactionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAvailableReactionsRequest) String() string {
	if g == nil {
		return "MessagesGetAvailableReactionsRequest(nil)"
	}
	type Alias MessagesGetAvailableReactionsRequest
	return fmt.Sprintf("MessagesGetAvailableReactionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetAvailableReactionsRequest) TypeID() uint32 {
	return MessagesGetAvailableReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetAvailableReactionsRequest) TypeName() string {
	return "messages.getAvailableReactions"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetAvailableReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getAvailableReactions",
		ID:   MessagesGetAvailableReactionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetAvailableReactionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAvailableReactions#18dea0ac as nil")
	}
	b.PutID(MessagesGetAvailableReactionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetAvailableReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAvailableReactions#18dea0ac as nil")
	}
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetAvailableReactionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAvailableReactions#18dea0ac to nil")
	}
	if err := b.ConsumeID(MessagesGetAvailableReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAvailableReactions#18dea0ac: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetAvailableReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAvailableReactions#18dea0ac to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAvailableReactions#18dea0ac: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetAvailableReactionsRequest) GetHash() (value int) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetAvailableReactions invokes method messages.getAvailableReactions#18dea0ac returning error if any.
func (c *Client) MessagesGetAvailableReactions(ctx context.Context, hash int) (MessagesAvailableReactionsClass, error) {
	var result MessagesAvailableReactionsBox

	request := &MessagesGetAvailableReactionsRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AvailableReactions, nil
}
