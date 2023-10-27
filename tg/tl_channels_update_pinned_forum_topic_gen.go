// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// ChannelsUpdatePinnedForumTopicRequest represents TL type `channels.updatePinnedForumTopic#6c2d9026`.
// Pin or unpin forum topics¹
//
// Links:
//  1. https://core.telegram.org/api/forum
//
// See https://core.telegram.org/method/channels.updatePinnedForumTopic for reference.
type ChannelsUpdatePinnedForumTopicRequest struct {
	// Supergroup ID
	Channel InputChannelClass
	// Forum topic ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/forum
	TopicID int
	// Whether to pin or unpin the topic
	Pinned bool
}

// ChannelsUpdatePinnedForumTopicRequestTypeID is TL type id of ChannelsUpdatePinnedForumTopicRequest.
const ChannelsUpdatePinnedForumTopicRequestTypeID = 0x6c2d9026

// Ensuring interfaces in compile-time for ChannelsUpdatePinnedForumTopicRequest.
var (
	_ bin.Encoder     = &ChannelsUpdatePinnedForumTopicRequest{}
	_ bin.Decoder     = &ChannelsUpdatePinnedForumTopicRequest{}
	_ bin.BareEncoder = &ChannelsUpdatePinnedForumTopicRequest{}
	_ bin.BareDecoder = &ChannelsUpdatePinnedForumTopicRequest{}
)

func (u *ChannelsUpdatePinnedForumTopicRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Channel == nil) {
		return false
	}
	if !(u.TopicID == 0) {
		return false
	}
	if !(u.Pinned == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *ChannelsUpdatePinnedForumTopicRequest) String() string {
	if u == nil {
		return "ChannelsUpdatePinnedForumTopicRequest(nil)"
	}
	type Alias ChannelsUpdatePinnedForumTopicRequest
	return fmt.Sprintf("ChannelsUpdatePinnedForumTopicRequest%+v", Alias(*u))
}

// FillFrom fills ChannelsUpdatePinnedForumTopicRequest from given interface.
func (u *ChannelsUpdatePinnedForumTopicRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetTopicID() (value int)
	GetPinned() (value bool)
}) {
	u.Channel = from.GetChannel()
	u.TopicID = from.GetTopicID()
	u.Pinned = from.GetPinned()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsUpdatePinnedForumTopicRequest) TypeID() uint32 {
	return ChannelsUpdatePinnedForumTopicRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsUpdatePinnedForumTopicRequest) TypeName() string {
	return "channels.updatePinnedForumTopic"
}

// TypeInfo returns info about TL type.
func (u *ChannelsUpdatePinnedForumTopicRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.updatePinnedForumTopic",
		ID:   ChannelsUpdatePinnedForumTopicRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "TopicID",
			SchemaName: "topic_id",
		},
		{
			Name:       "Pinned",
			SchemaName: "pinned",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *ChannelsUpdatePinnedForumTopicRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updatePinnedForumTopic#6c2d9026 as nil")
	}
	b.PutID(ChannelsUpdatePinnedForumTopicRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *ChannelsUpdatePinnedForumTopicRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updatePinnedForumTopic#6c2d9026 as nil")
	}
	if u.Channel == nil {
		return fmt.Errorf("unable to encode channels.updatePinnedForumTopic#6c2d9026: field channel is nil")
	}
	if err := u.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.updatePinnedForumTopic#6c2d9026: field channel: %w", err)
	}
	b.PutInt(u.TopicID)
	b.PutBool(u.Pinned)
	return nil
}

// Decode implements bin.Decoder.
func (u *ChannelsUpdatePinnedForumTopicRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updatePinnedForumTopic#6c2d9026 to nil")
	}
	if err := b.ConsumeID(ChannelsUpdatePinnedForumTopicRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.updatePinnedForumTopic#6c2d9026: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *ChannelsUpdatePinnedForumTopicRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updatePinnedForumTopic#6c2d9026 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.updatePinnedForumTopic#6c2d9026: field channel: %w", err)
		}
		u.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.updatePinnedForumTopic#6c2d9026: field topic_id: %w", err)
		}
		u.TopicID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.updatePinnedForumTopic#6c2d9026: field pinned: %w", err)
		}
		u.Pinned = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (u *ChannelsUpdatePinnedForumTopicRequest) GetChannel() (value InputChannelClass) {
	if u == nil {
		return
	}
	return u.Channel
}

// GetTopicID returns value of TopicID field.
func (u *ChannelsUpdatePinnedForumTopicRequest) GetTopicID() (value int) {
	if u == nil {
		return
	}
	return u.TopicID
}

// GetPinned returns value of Pinned field.
func (u *ChannelsUpdatePinnedForumTopicRequest) GetPinned() (value bool) {
	if u == nil {
		return
	}
	return u.Pinned
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (u *ChannelsUpdatePinnedForumTopicRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return u.Channel.AsNotEmpty()
}

// ChannelsUpdatePinnedForumTopic invokes method channels.updatePinnedForumTopic#6c2d9026 returning error if any.
// Pin or unpin forum topics¹
//
// Links:
//  1. https://core.telegram.org/api/forum
//
// See https://core.telegram.org/method/channels.updatePinnedForumTopic for reference.
// Can be used by bots.
func (c *Client) ChannelsUpdatePinnedForumTopic(ctx context.Context, request *ChannelsUpdatePinnedForumTopicRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
