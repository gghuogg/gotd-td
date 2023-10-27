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

// StatsGetMessagePublicForwardsRequest represents TL type `stats.getMessagePublicForwards#5630281b`.
type StatsGetMessagePublicForwardsRequest struct {
	// Channel field of StatsGetMessagePublicForwardsRequest.
	Channel InputChannelClass
	// MsgID field of StatsGetMessagePublicForwardsRequest.
	MsgID int
	// OffsetRate field of StatsGetMessagePublicForwardsRequest.
	OffsetRate int
	// OffsetPeer field of StatsGetMessagePublicForwardsRequest.
	OffsetPeer InputPeerClass
	// OffsetID field of StatsGetMessagePublicForwardsRequest.
	OffsetID int
	// Limit field of StatsGetMessagePublicForwardsRequest.
	Limit int
}

// StatsGetMessagePublicForwardsRequestTypeID is TL type id of StatsGetMessagePublicForwardsRequest.
const StatsGetMessagePublicForwardsRequestTypeID = 0x5630281b

// Ensuring interfaces in compile-time for StatsGetMessagePublicForwardsRequest.
var (
	_ bin.Encoder     = &StatsGetMessagePublicForwardsRequest{}
	_ bin.Decoder     = &StatsGetMessagePublicForwardsRequest{}
	_ bin.BareEncoder = &StatsGetMessagePublicForwardsRequest{}
	_ bin.BareDecoder = &StatsGetMessagePublicForwardsRequest{}
)

func (g *StatsGetMessagePublicForwardsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}
	if !(g.OffsetRate == 0) {
		return false
	}
	if !(g.OffsetPeer == nil) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StatsGetMessagePublicForwardsRequest) String() string {
	if g == nil {
		return "StatsGetMessagePublicForwardsRequest(nil)"
	}
	type Alias StatsGetMessagePublicForwardsRequest
	return fmt.Sprintf("StatsGetMessagePublicForwardsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGetMessagePublicForwardsRequest) TypeID() uint32 {
	return StatsGetMessagePublicForwardsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGetMessagePublicForwardsRequest) TypeName() string {
	return "stats.getMessagePublicForwards"
}

// TypeInfo returns info about TL type.
func (g *StatsGetMessagePublicForwardsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.getMessagePublicForwards",
		ID:   StatsGetMessagePublicForwardsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "OffsetRate",
			SchemaName: "offset_rate",
		},
		{
			Name:       "OffsetPeer",
			SchemaName: "offset_peer",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *StatsGetMessagePublicForwardsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getMessagePublicForwards#5630281b as nil")
	}
	b.PutID(StatsGetMessagePublicForwardsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StatsGetMessagePublicForwardsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getMessagePublicForwards#5630281b as nil")
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5630281b: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5630281b: field channel: %w", err)
	}
	b.PutInt(g.MsgID)
	b.PutInt(g.OffsetRate)
	if g.OffsetPeer == nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5630281b: field offset_peer is nil")
	}
	if err := g.OffsetPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5630281b: field offset_peer: %w", err)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *StatsGetMessagePublicForwardsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getMessagePublicForwards#5630281b to nil")
	}
	if err := b.ConsumeID(StatsGetMessagePublicForwardsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StatsGetMessagePublicForwardsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getMessagePublicForwards#5630281b to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field offset_rate: %w", err)
		}
		g.OffsetRate = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field offset_peer: %w", err)
		}
		g.OffsetPeer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (g *StatsGetMessagePublicForwardsRequest) GetChannel() (value InputChannelClass) {
	if g == nil {
		return
	}
	return g.Channel
}

// GetMsgID returns value of MsgID field.
func (g *StatsGetMessagePublicForwardsRequest) GetMsgID() (value int) {
	if g == nil {
		return
	}
	return g.MsgID
}

// GetOffsetRate returns value of OffsetRate field.
func (g *StatsGetMessagePublicForwardsRequest) GetOffsetRate() (value int) {
	if g == nil {
		return
	}
	return g.OffsetRate
}

// GetOffsetPeer returns value of OffsetPeer field.
func (g *StatsGetMessagePublicForwardsRequest) GetOffsetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.OffsetPeer
}

// GetOffsetID returns value of OffsetID field.
func (g *StatsGetMessagePublicForwardsRequest) GetOffsetID() (value int) {
	if g == nil {
		return
	}
	return g.OffsetID
}

// GetLimit returns value of Limit field.
func (g *StatsGetMessagePublicForwardsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// StatsGetMessagePublicForwards invokes method stats.getMessagePublicForwards#5630281b returning error if any.
func (c *Client) StatsGetMessagePublicForwards(ctx context.Context, request *StatsGetMessagePublicForwardsRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
