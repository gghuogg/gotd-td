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

// MessagesSendScreenshotNotificationRequest represents TL type `messages.sendScreenshotNotification#a1405817`.
type MessagesSendScreenshotNotificationRequest struct {
	// Peer field of MessagesSendScreenshotNotificationRequest.
	Peer InputPeerClass
	// ReplyTo field of MessagesSendScreenshotNotificationRequest.
	ReplyTo InputReplyToClass
	// RandomID field of MessagesSendScreenshotNotificationRequest.
	RandomID int64
}

// MessagesSendScreenshotNotificationRequestTypeID is TL type id of MessagesSendScreenshotNotificationRequest.
const MessagesSendScreenshotNotificationRequestTypeID = 0xa1405817

// Ensuring interfaces in compile-time for MessagesSendScreenshotNotificationRequest.
var (
	_ bin.Encoder     = &MessagesSendScreenshotNotificationRequest{}
	_ bin.Decoder     = &MessagesSendScreenshotNotificationRequest{}
	_ bin.BareEncoder = &MessagesSendScreenshotNotificationRequest{}
	_ bin.BareDecoder = &MessagesSendScreenshotNotificationRequest{}
)

func (s *MessagesSendScreenshotNotificationRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyTo == nil) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendScreenshotNotificationRequest) String() string {
	if s == nil {
		return "MessagesSendScreenshotNotificationRequest(nil)"
	}
	type Alias MessagesSendScreenshotNotificationRequest
	return fmt.Sprintf("MessagesSendScreenshotNotificationRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendScreenshotNotificationRequest) TypeID() uint32 {
	return MessagesSendScreenshotNotificationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendScreenshotNotificationRequest) TypeName() string {
	return "messages.sendScreenshotNotification"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendScreenshotNotificationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendScreenshotNotification",
		ID:   MessagesSendScreenshotNotificationRequestTypeID,
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
			Name:       "ReplyTo",
			SchemaName: "reply_to",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSendScreenshotNotificationRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendScreenshotNotification#a1405817 as nil")
	}
	b.PutID(MessagesSendScreenshotNotificationRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendScreenshotNotificationRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendScreenshotNotification#a1405817 as nil")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendScreenshotNotification#a1405817: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendScreenshotNotification#a1405817: field peer: %w", err)
	}
	if s.ReplyTo == nil {
		return fmt.Errorf("unable to encode messages.sendScreenshotNotification#a1405817: field reply_to is nil")
	}
	if err := s.ReplyTo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendScreenshotNotification#a1405817: field reply_to: %w", err)
	}
	b.PutLong(s.RandomID)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendScreenshotNotificationRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendScreenshotNotification#a1405817 to nil")
	}
	if err := b.ConsumeID(MessagesSendScreenshotNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendScreenshotNotification#a1405817: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendScreenshotNotificationRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendScreenshotNotification#a1405817 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendScreenshotNotification#a1405817: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := DecodeInputReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendScreenshotNotification#a1405817: field reply_to: %w", err)
		}
		s.ReplyTo = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendScreenshotNotification#a1405817: field random_id: %w", err)
		}
		s.RandomID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSendScreenshotNotificationRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetReplyTo returns value of ReplyTo field.
func (s *MessagesSendScreenshotNotificationRequest) GetReplyTo() (value InputReplyToClass) {
	if s == nil {
		return
	}
	return s.ReplyTo
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendScreenshotNotificationRequest) GetRandomID() (value int64) {
	if s == nil {
		return
	}
	return s.RandomID
}

// MessagesSendScreenshotNotification invokes method messages.sendScreenshotNotification#a1405817 returning error if any.
func (c *Client) MessagesSendScreenshotNotification(ctx context.Context, request *MessagesSendScreenshotNotificationRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
