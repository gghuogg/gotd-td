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

// MessagesSetTypingRequest represents TL type `messages.setTyping#58943ee2`.
type MessagesSetTypingRequest struct {
	// Flags field of MessagesSetTypingRequest.
	Flags bin.Fields
	// Peer field of MessagesSetTypingRequest.
	Peer InputPeerClass
	// TopMsgID field of MessagesSetTypingRequest.
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
	// Action field of MessagesSetTypingRequest.
	Action SendMessageActionClass
}

// MessagesSetTypingRequestTypeID is TL type id of MessagesSetTypingRequest.
const MessagesSetTypingRequestTypeID = 0x58943ee2

// Ensuring interfaces in compile-time for MessagesSetTypingRequest.
var (
	_ bin.Encoder     = &MessagesSetTypingRequest{}
	_ bin.Decoder     = &MessagesSetTypingRequest{}
	_ bin.BareEncoder = &MessagesSetTypingRequest{}
	_ bin.BareDecoder = &MessagesSetTypingRequest{}
)

func (s *MessagesSetTypingRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.TopMsgID == 0) {
		return false
	}
	if !(s.Action == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetTypingRequest) String() string {
	if s == nil {
		return "MessagesSetTypingRequest(nil)"
	}
	type Alias MessagesSetTypingRequest
	return fmt.Sprintf("MessagesSetTypingRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetTypingRequest) TypeID() uint32 {
	return MessagesSetTypingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetTypingRequest) TypeName() string {
	return "messages.setTyping"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetTypingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setTyping",
		ID:   MessagesSetTypingRequestTypeID,
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
			Name:       "TopMsgID",
			SchemaName: "top_msg_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Action",
			SchemaName: "action",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSetTypingRequest) SetFlags() {
	if !(s.TopMsgID == 0) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSetTypingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setTyping#58943ee2 as nil")
	}
	b.PutID(MessagesSetTypingRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetTypingRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setTyping#58943ee2 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.TopMsgID)
	}
	if s.Action == nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field action is nil")
	}
	if err := s.Action.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field action: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetTypingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setTyping#58943ee2 to nil")
	}
	if err := b.ConsumeID(MessagesSetTypingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setTyping#58943ee2: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetTypingRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setTyping#58943ee2 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field top_msg_id: %w", err)
		}
		s.TopMsgID = value
	}
	{
		value, err := DecodeSendMessageAction(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field action: %w", err)
		}
		s.Action = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSetTypingRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (s *MessagesSetTypingRequest) SetTopMsgID(value int) {
	s.Flags.Set(0)
	s.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSetTypingRequest) GetTopMsgID() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.TopMsgID, true
}

// GetAction returns value of Action field.
func (s *MessagesSetTypingRequest) GetAction() (value SendMessageActionClass) {
	if s == nil {
		return
	}
	return s.Action
}

// MessagesSetTyping invokes method messages.setTyping#58943ee2 returning error if any.
func (c *Client) MessagesSetTyping(ctx context.Context, request *MessagesSetTypingRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
