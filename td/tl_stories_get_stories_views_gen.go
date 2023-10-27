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

// StoriesGetStoriesViewsRequest represents TL type `stories.getStoriesViews#28e16cc8`.
type StoriesGetStoriesViewsRequest struct {
	// Peer field of StoriesGetStoriesViewsRequest.
	Peer InputPeerClass
	// ID field of StoriesGetStoriesViewsRequest.
	ID []int
}

// StoriesGetStoriesViewsRequestTypeID is TL type id of StoriesGetStoriesViewsRequest.
const StoriesGetStoriesViewsRequestTypeID = 0x28e16cc8

// Ensuring interfaces in compile-time for StoriesGetStoriesViewsRequest.
var (
	_ bin.Encoder     = &StoriesGetStoriesViewsRequest{}
	_ bin.Decoder     = &StoriesGetStoriesViewsRequest{}
	_ bin.BareEncoder = &StoriesGetStoriesViewsRequest{}
	_ bin.BareDecoder = &StoriesGetStoriesViewsRequest{}
)

func (g *StoriesGetStoriesViewsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StoriesGetStoriesViewsRequest) String() string {
	if g == nil {
		return "StoriesGetStoriesViewsRequest(nil)"
	}
	type Alias StoriesGetStoriesViewsRequest
	return fmt.Sprintf("StoriesGetStoriesViewsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesGetStoriesViewsRequest) TypeID() uint32 {
	return StoriesGetStoriesViewsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesGetStoriesViewsRequest) TypeName() string {
	return "stories.getStoriesViews"
}

// TypeInfo returns info about TL type.
func (g *StoriesGetStoriesViewsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.getStoriesViews",
		ID:   StoriesGetStoriesViewsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *StoriesGetStoriesViewsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getStoriesViews#28e16cc8 as nil")
	}
	b.PutID(StoriesGetStoriesViewsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StoriesGetStoriesViewsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getStoriesViews#28e16cc8 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode stories.getStoriesViews#28e16cc8: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.getStoriesViews#28e16cc8: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *StoriesGetStoriesViewsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getStoriesViews#28e16cc8 to nil")
	}
	if err := b.ConsumeID(StoriesGetStoriesViewsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.getStoriesViews#28e16cc8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StoriesGetStoriesViewsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getStoriesViews#28e16cc8 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.getStoriesViews#28e16cc8: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.getStoriesViews#28e16cc8: field id: %w", err)
		}

		if headerLen > 0 {
			g.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode stories.getStoriesViews#28e16cc8: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *StoriesGetStoriesViewsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *StoriesGetStoriesViewsRequest) GetID() (value []int) {
	if g == nil {
		return
	}
	return g.ID
}

// StoriesGetStoriesViews invokes method stories.getStoriesViews#28e16cc8 returning error if any.
func (c *Client) StoriesGetStoriesViews(ctx context.Context, request *StoriesGetStoriesViewsRequest) (*StoriesStoryViews, error) {
	var result StoriesStoryViews

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
