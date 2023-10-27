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

// MessagesUnpinAllMessagesRequest represents TL type `messages.unpinAllMessages#ee22b9a8`.
// Unpin¹ all pinned messages
//
// Links:
//  1. https://core.telegram.org/api/pin
//
// See https://core.telegram.org/method/messages.unpinAllMessages for reference.
type MessagesUnpinAllMessagesRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Chat where to unpin
	Peer InputPeerClass
	// Forum topic¹ where to unpin
	//
	// Links:
	//  1) https://core.telegram.org/api/forum#forum-topics
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
}

// MessagesUnpinAllMessagesRequestTypeID is TL type id of MessagesUnpinAllMessagesRequest.
const MessagesUnpinAllMessagesRequestTypeID = 0xee22b9a8

// Ensuring interfaces in compile-time for MessagesUnpinAllMessagesRequest.
var (
	_ bin.Encoder     = &MessagesUnpinAllMessagesRequest{}
	_ bin.Decoder     = &MessagesUnpinAllMessagesRequest{}
	_ bin.BareEncoder = &MessagesUnpinAllMessagesRequest{}
	_ bin.BareDecoder = &MessagesUnpinAllMessagesRequest{}
)

func (u *MessagesUnpinAllMessagesRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Peer == nil) {
		return false
	}
	if !(u.TopMsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *MessagesUnpinAllMessagesRequest) String() string {
	if u == nil {
		return "MessagesUnpinAllMessagesRequest(nil)"
	}
	type Alias MessagesUnpinAllMessagesRequest
	return fmt.Sprintf("MessagesUnpinAllMessagesRequest%+v", Alias(*u))
}

// FillFrom fills MessagesUnpinAllMessagesRequest from given interface.
func (u *MessagesUnpinAllMessagesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetTopMsgID() (value int, ok bool)
}) {
	u.Peer = from.GetPeer()
	if val, ok := from.GetTopMsgID(); ok {
		u.TopMsgID = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesUnpinAllMessagesRequest) TypeID() uint32 {
	return MessagesUnpinAllMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesUnpinAllMessagesRequest) TypeName() string {
	return "messages.unpinAllMessages"
}

// TypeInfo returns info about TL type.
func (u *MessagesUnpinAllMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.unpinAllMessages",
		ID:   MessagesUnpinAllMessagesRequestTypeID,
	}
	if u == nil {
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
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *MessagesUnpinAllMessagesRequest) SetFlags() {
	if !(u.TopMsgID == 0) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *MessagesUnpinAllMessagesRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.unpinAllMessages#ee22b9a8 as nil")
	}
	b.PutID(MessagesUnpinAllMessagesRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *MessagesUnpinAllMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.unpinAllMessages#ee22b9a8 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.unpinAllMessages#ee22b9a8: field flags: %w", err)
	}
	if u.Peer == nil {
		return fmt.Errorf("unable to encode messages.unpinAllMessages#ee22b9a8: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.unpinAllMessages#ee22b9a8: field peer: %w", err)
	}
	if u.Flags.Has(0) {
		b.PutInt(u.TopMsgID)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *MessagesUnpinAllMessagesRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.unpinAllMessages#ee22b9a8 to nil")
	}
	if err := b.ConsumeID(MessagesUnpinAllMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.unpinAllMessages#ee22b9a8: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *MessagesUnpinAllMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.unpinAllMessages#ee22b9a8 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.unpinAllMessages#ee22b9a8: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.unpinAllMessages#ee22b9a8: field peer: %w", err)
		}
		u.Peer = value
	}
	if u.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.unpinAllMessages#ee22b9a8: field top_msg_id: %w", err)
		}
		u.TopMsgID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (u *MessagesUnpinAllMessagesRequest) GetPeer() (value InputPeerClass) {
	if u == nil {
		return
	}
	return u.Peer
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (u *MessagesUnpinAllMessagesRequest) SetTopMsgID(value int) {
	u.Flags.Set(0)
	u.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (u *MessagesUnpinAllMessagesRequest) GetTopMsgID() (value int, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.TopMsgID, true
}

// MessagesUnpinAllMessages invokes method messages.unpinAllMessages#ee22b9a8 returning error if any.
// Unpin¹ all pinned messages
//
// Links:
//  1. https://core.telegram.org/api/pin
//
// Possible errors:
//
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_NOT_MODIFIED: No changes were made to chat information because the new information you passed is identical to the current information.
//
// See https://core.telegram.org/method/messages.unpinAllMessages for reference.
// Can be used by bots.
func (c *Client) MessagesUnpinAllMessages(ctx context.Context, request *MessagesUnpinAllMessagesRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
