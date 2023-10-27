// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// RevokeChatInviteLinkRequest represents TL type `revokeChatInviteLink#d1b755a9`.
type RevokeChatInviteLinkRequest struct {
	// Chat identifier
	ChatID int64
	// Invite link to be revoked
	InviteLink string
}

// RevokeChatInviteLinkRequestTypeID is TL type id of RevokeChatInviteLinkRequest.
const RevokeChatInviteLinkRequestTypeID = 0xd1b755a9

// Ensuring interfaces in compile-time for RevokeChatInviteLinkRequest.
var (
	_ bin.Encoder     = &RevokeChatInviteLinkRequest{}
	_ bin.Decoder     = &RevokeChatInviteLinkRequest{}
	_ bin.BareEncoder = &RevokeChatInviteLinkRequest{}
	_ bin.BareDecoder = &RevokeChatInviteLinkRequest{}
)

func (r *RevokeChatInviteLinkRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}
	if !(r.InviteLink == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RevokeChatInviteLinkRequest) String() string {
	if r == nil {
		return "RevokeChatInviteLinkRequest(nil)"
	}
	type Alias RevokeChatInviteLinkRequest
	return fmt.Sprintf("RevokeChatInviteLinkRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RevokeChatInviteLinkRequest) TypeID() uint32 {
	return RevokeChatInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RevokeChatInviteLinkRequest) TypeName() string {
	return "revokeChatInviteLink"
}

// TypeInfo returns info about TL type.
func (r *RevokeChatInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "revokeChatInviteLink",
		ID:   RevokeChatInviteLinkRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RevokeChatInviteLinkRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode revokeChatInviteLink#d1b755a9 as nil")
	}
	b.PutID(RevokeChatInviteLinkRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RevokeChatInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode revokeChatInviteLink#d1b755a9 as nil")
	}
	b.PutInt53(r.ChatID)
	b.PutString(r.InviteLink)
	return nil
}

// Decode implements bin.Decoder.
func (r *RevokeChatInviteLinkRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode revokeChatInviteLink#d1b755a9 to nil")
	}
	if err := b.ConsumeID(RevokeChatInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode revokeChatInviteLink#d1b755a9: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RevokeChatInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode revokeChatInviteLink#d1b755a9 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode revokeChatInviteLink#d1b755a9: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode revokeChatInviteLink#d1b755a9: field invite_link: %w", err)
		}
		r.InviteLink = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RevokeChatInviteLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode revokeChatInviteLink#d1b755a9 as nil")
	}
	b.ObjStart()
	b.PutID("revokeChatInviteLink")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(r.InviteLink)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RevokeChatInviteLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode revokeChatInviteLink#d1b755a9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("revokeChatInviteLink"); err != nil {
				return fmt.Errorf("unable to decode revokeChatInviteLink#d1b755a9: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode revokeChatInviteLink#d1b755a9: field chat_id: %w", err)
			}
			r.ChatID = value
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode revokeChatInviteLink#d1b755a9: field invite_link: %w", err)
			}
			r.InviteLink = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *RevokeChatInviteLinkRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// GetInviteLink returns value of InviteLink field.
func (r *RevokeChatInviteLinkRequest) GetInviteLink() (value string) {
	if r == nil {
		return
	}
	return r.InviteLink
}

// RevokeChatInviteLink invokes method revokeChatInviteLink#d1b755a9 returning error if any.
func (c *Client) RevokeChatInviteLink(ctx context.Context, request *RevokeChatInviteLinkRequest) (*ChatInviteLinks, error) {
	var result ChatInviteLinks

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
