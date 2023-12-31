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

// StoriesApplyBoostRequest represents TL type `stories.applyBoost#f29d7c2b`.
type StoriesApplyBoostRequest struct {
	// Peer field of StoriesApplyBoostRequest.
	Peer InputPeerClass
}

// StoriesApplyBoostRequestTypeID is TL type id of StoriesApplyBoostRequest.
const StoriesApplyBoostRequestTypeID = 0xf29d7c2b

// Ensuring interfaces in compile-time for StoriesApplyBoostRequest.
var (
	_ bin.Encoder     = &StoriesApplyBoostRequest{}
	_ bin.Decoder     = &StoriesApplyBoostRequest{}
	_ bin.BareEncoder = &StoriesApplyBoostRequest{}
	_ bin.BareDecoder = &StoriesApplyBoostRequest{}
)

func (a *StoriesApplyBoostRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *StoriesApplyBoostRequest) String() string {
	if a == nil {
		return "StoriesApplyBoostRequest(nil)"
	}
	type Alias StoriesApplyBoostRequest
	return fmt.Sprintf("StoriesApplyBoostRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesApplyBoostRequest) TypeID() uint32 {
	return StoriesApplyBoostRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesApplyBoostRequest) TypeName() string {
	return "stories.applyBoost"
}

// TypeInfo returns info about TL type.
func (a *StoriesApplyBoostRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.applyBoost",
		ID:   StoriesApplyBoostRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *StoriesApplyBoostRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stories.applyBoost#f29d7c2b as nil")
	}
	b.PutID(StoriesApplyBoostRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *StoriesApplyBoostRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stories.applyBoost#f29d7c2b as nil")
	}
	if a.Peer == nil {
		return fmt.Errorf("unable to encode stories.applyBoost#f29d7c2b: field peer is nil")
	}
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.applyBoost#f29d7c2b: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *StoriesApplyBoostRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stories.applyBoost#f29d7c2b to nil")
	}
	if err := b.ConsumeID(StoriesApplyBoostRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.applyBoost#f29d7c2b: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *StoriesApplyBoostRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stories.applyBoost#f29d7c2b to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.applyBoost#f29d7c2b: field peer: %w", err)
		}
		a.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (a *StoriesApplyBoostRequest) GetPeer() (value InputPeerClass) {
	if a == nil {
		return
	}
	return a.Peer
}

// StoriesApplyBoost invokes method stories.applyBoost#f29d7c2b returning error if any.
func (c *Client) StoriesApplyBoost(ctx context.Context, peer InputPeerClass) (bool, error) {
	var result BoolBox

	request := &StoriesApplyBoostRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
