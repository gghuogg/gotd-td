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

// MessagesGetRecentReactionsRequest represents TL type `messages.getRecentReactions#39461db2`.
type MessagesGetRecentReactionsRequest struct {
	// Limit field of MessagesGetRecentReactionsRequest.
	Limit int
	// Hash field of MessagesGetRecentReactionsRequest.
	Hash int64
}

// MessagesGetRecentReactionsRequestTypeID is TL type id of MessagesGetRecentReactionsRequest.
const MessagesGetRecentReactionsRequestTypeID = 0x39461db2

// Ensuring interfaces in compile-time for MessagesGetRecentReactionsRequest.
var (
	_ bin.Encoder     = &MessagesGetRecentReactionsRequest{}
	_ bin.Decoder     = &MessagesGetRecentReactionsRequest{}
	_ bin.BareEncoder = &MessagesGetRecentReactionsRequest{}
	_ bin.BareDecoder = &MessagesGetRecentReactionsRequest{}
)

func (g *MessagesGetRecentReactionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetRecentReactionsRequest) String() string {
	if g == nil {
		return "MessagesGetRecentReactionsRequest(nil)"
	}
	type Alias MessagesGetRecentReactionsRequest
	return fmt.Sprintf("MessagesGetRecentReactionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetRecentReactionsRequest) TypeID() uint32 {
	return MessagesGetRecentReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetRecentReactionsRequest) TypeName() string {
	return "messages.getRecentReactions"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetRecentReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getRecentReactions",
		ID:   MessagesGetRecentReactionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetRecentReactionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getRecentReactions#39461db2 as nil")
	}
	b.PutID(MessagesGetRecentReactionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetRecentReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getRecentReactions#39461db2 as nil")
	}
	b.PutInt(g.Limit)
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetRecentReactionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getRecentReactions#39461db2 to nil")
	}
	if err := b.ConsumeID(MessagesGetRecentReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getRecentReactions#39461db2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetRecentReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getRecentReactions#39461db2 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getRecentReactions#39461db2: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getRecentReactions#39461db2: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetLimit returns value of Limit field.
func (g *MessagesGetRecentReactionsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *MessagesGetRecentReactionsRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetRecentReactions invokes method messages.getRecentReactions#39461db2 returning error if any.
func (c *Client) MessagesGetRecentReactions(ctx context.Context, request *MessagesGetRecentReactionsRequest) (MessagesReactionsClass, error) {
	var result MessagesReactionsBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Reactions, nil
}
