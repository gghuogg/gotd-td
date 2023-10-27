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

// MessagesUpdatePinnedMessageRequest represents TL type `messages.updatePinnedMessage#d2aaf7ec`.
type MessagesUpdatePinnedMessageRequest struct {
	// Flags field of MessagesUpdatePinnedMessageRequest.
	Flags bin.Fields
	// Silent field of MessagesUpdatePinnedMessageRequest.
	Silent bool
	// Unpin field of MessagesUpdatePinnedMessageRequest.
	Unpin bool
	// PmOneside field of MessagesUpdatePinnedMessageRequest.
	PmOneside bool
	// Peer field of MessagesUpdatePinnedMessageRequest.
	Peer InputPeerClass
	// ID field of MessagesUpdatePinnedMessageRequest.
	ID int
}

// MessagesUpdatePinnedMessageRequestTypeID is TL type id of MessagesUpdatePinnedMessageRequest.
const MessagesUpdatePinnedMessageRequestTypeID = 0xd2aaf7ec

// Ensuring interfaces in compile-time for MessagesUpdatePinnedMessageRequest.
var (
	_ bin.Encoder     = &MessagesUpdatePinnedMessageRequest{}
	_ bin.Decoder     = &MessagesUpdatePinnedMessageRequest{}
	_ bin.BareEncoder = &MessagesUpdatePinnedMessageRequest{}
	_ bin.BareDecoder = &MessagesUpdatePinnedMessageRequest{}
)

func (u *MessagesUpdatePinnedMessageRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Silent == false) {
		return false
	}
	if !(u.Unpin == false) {
		return false
	}
	if !(u.PmOneside == false) {
		return false
	}
	if !(u.Peer == nil) {
		return false
	}
	if !(u.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *MessagesUpdatePinnedMessageRequest) String() string {
	if u == nil {
		return "MessagesUpdatePinnedMessageRequest(nil)"
	}
	type Alias MessagesUpdatePinnedMessageRequest
	return fmt.Sprintf("MessagesUpdatePinnedMessageRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesUpdatePinnedMessageRequest) TypeID() uint32 {
	return MessagesUpdatePinnedMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesUpdatePinnedMessageRequest) TypeName() string {
	return "messages.updatePinnedMessage"
}

// TypeInfo returns info about TL type.
func (u *MessagesUpdatePinnedMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.updatePinnedMessage",
		ID:   MessagesUpdatePinnedMessageRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "Unpin",
			SchemaName: "unpin",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "PmOneside",
			SchemaName: "pm_oneside",
			Null:       !u.Flags.Has(2),
		},
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

// SetFlags sets flags for non-zero fields.
func (u *MessagesUpdatePinnedMessageRequest) SetFlags() {
	if !(u.Silent == false) {
		u.Flags.Set(0)
	}
	if !(u.Unpin == false) {
		u.Flags.Set(1)
	}
	if !(u.PmOneside == false) {
		u.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (u *MessagesUpdatePinnedMessageRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.updatePinnedMessage#d2aaf7ec as nil")
	}
	b.PutID(MessagesUpdatePinnedMessageRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *MessagesUpdatePinnedMessageRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.updatePinnedMessage#d2aaf7ec as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.updatePinnedMessage#d2aaf7ec: field flags: %w", err)
	}
	if u.Peer == nil {
		return fmt.Errorf("unable to encode messages.updatePinnedMessage#d2aaf7ec: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.updatePinnedMessage#d2aaf7ec: field peer: %w", err)
	}
	b.PutInt(u.ID)
	return nil
}

// Decode implements bin.Decoder.
func (u *MessagesUpdatePinnedMessageRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.updatePinnedMessage#d2aaf7ec to nil")
	}
	if err := b.ConsumeID(MessagesUpdatePinnedMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *MessagesUpdatePinnedMessageRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.updatePinnedMessage#d2aaf7ec to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: field flags: %w", err)
		}
	}
	u.Silent = u.Flags.Has(0)
	u.Unpin = u.Flags.Has(1)
	u.PmOneside = u.Flags.Has(2)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: field peer: %w", err)
		}
		u.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (u *MessagesUpdatePinnedMessageRequest) SetSilent(value bool) {
	if value {
		u.Flags.Set(0)
		u.Silent = true
	} else {
		u.Flags.Unset(0)
		u.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (u *MessagesUpdatePinnedMessageRequest) GetSilent() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// SetUnpin sets value of Unpin conditional field.
func (u *MessagesUpdatePinnedMessageRequest) SetUnpin(value bool) {
	if value {
		u.Flags.Set(1)
		u.Unpin = true
	} else {
		u.Flags.Unset(1)
		u.Unpin = false
	}
}

// GetUnpin returns value of Unpin conditional field.
func (u *MessagesUpdatePinnedMessageRequest) GetUnpin() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(1)
}

// SetPmOneside sets value of PmOneside conditional field.
func (u *MessagesUpdatePinnedMessageRequest) SetPmOneside(value bool) {
	if value {
		u.Flags.Set(2)
		u.PmOneside = true
	} else {
		u.Flags.Unset(2)
		u.PmOneside = false
	}
}

// GetPmOneside returns value of PmOneside conditional field.
func (u *MessagesUpdatePinnedMessageRequest) GetPmOneside() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(2)
}

// GetPeer returns value of Peer field.
func (u *MessagesUpdatePinnedMessageRequest) GetPeer() (value InputPeerClass) {
	if u == nil {
		return
	}
	return u.Peer
}

// GetID returns value of ID field.
func (u *MessagesUpdatePinnedMessageRequest) GetID() (value int) {
	if u == nil {
		return
	}
	return u.ID
}

// MessagesUpdatePinnedMessage invokes method messages.updatePinnedMessage#d2aaf7ec returning error if any.
func (c *Client) MessagesUpdatePinnedMessage(ctx context.Context, request *MessagesUpdatePinnedMessageRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
