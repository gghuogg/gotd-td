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

// MessagesSetChatAvailableReactionsRequest represents TL type `messages.setChatAvailableReactions#feb16771`.
type MessagesSetChatAvailableReactionsRequest struct {
	// Peer field of MessagesSetChatAvailableReactionsRequest.
	Peer InputPeerClass
	// AvailableReactions field of MessagesSetChatAvailableReactionsRequest.
	AvailableReactions ChatReactionsClass
}

// MessagesSetChatAvailableReactionsRequestTypeID is TL type id of MessagesSetChatAvailableReactionsRequest.
const MessagesSetChatAvailableReactionsRequestTypeID = 0xfeb16771

// Ensuring interfaces in compile-time for MessagesSetChatAvailableReactionsRequest.
var (
	_ bin.Encoder     = &MessagesSetChatAvailableReactionsRequest{}
	_ bin.Decoder     = &MessagesSetChatAvailableReactionsRequest{}
	_ bin.BareEncoder = &MessagesSetChatAvailableReactionsRequest{}
	_ bin.BareDecoder = &MessagesSetChatAvailableReactionsRequest{}
)

func (s *MessagesSetChatAvailableReactionsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.AvailableReactions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetChatAvailableReactionsRequest) String() string {
	if s == nil {
		return "MessagesSetChatAvailableReactionsRequest(nil)"
	}
	type Alias MessagesSetChatAvailableReactionsRequest
	return fmt.Sprintf("MessagesSetChatAvailableReactionsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetChatAvailableReactionsRequest) TypeID() uint32 {
	return MessagesSetChatAvailableReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetChatAvailableReactionsRequest) TypeName() string {
	return "messages.setChatAvailableReactions"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetChatAvailableReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setChatAvailableReactions",
		ID:   MessagesSetChatAvailableReactionsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "AvailableReactions",
			SchemaName: "available_reactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetChatAvailableReactionsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setChatAvailableReactions#feb16771 as nil")
	}
	b.PutID(MessagesSetChatAvailableReactionsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetChatAvailableReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setChatAvailableReactions#feb16771 as nil")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#feb16771: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#feb16771: field peer: %w", err)
	}
	if s.AvailableReactions == nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#feb16771: field available_reactions is nil")
	}
	if err := s.AvailableReactions.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#feb16771: field available_reactions: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetChatAvailableReactionsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setChatAvailableReactions#feb16771 to nil")
	}
	if err := b.ConsumeID(MessagesSetChatAvailableReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setChatAvailableReactions#feb16771: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetChatAvailableReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setChatAvailableReactions#feb16771 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setChatAvailableReactions#feb16771: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := DecodeChatReactions(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setChatAvailableReactions#feb16771: field available_reactions: %w", err)
		}
		s.AvailableReactions = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSetChatAvailableReactionsRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetAvailableReactions returns value of AvailableReactions field.
func (s *MessagesSetChatAvailableReactionsRequest) GetAvailableReactions() (value ChatReactionsClass) {
	if s == nil {
		return
	}
	return s.AvailableReactions
}

// MessagesSetChatAvailableReactions invokes method messages.setChatAvailableReactions#feb16771 returning error if any.
func (c *Client) MessagesSetChatAvailableReactions(ctx context.Context, request *MessagesSetChatAvailableReactionsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
