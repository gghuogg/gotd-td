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

// ContactsUnblockRequest represents TL type `contacts.unblock#b550d328`.
type ContactsUnblockRequest struct {
	// Flags field of ContactsUnblockRequest.
	Flags bin.Fields
	// MyStoriesFrom field of ContactsUnblockRequest.
	MyStoriesFrom bool
	// ID field of ContactsUnblockRequest.
	ID InputPeerClass
}

// ContactsUnblockRequestTypeID is TL type id of ContactsUnblockRequest.
const ContactsUnblockRequestTypeID = 0xb550d328

// Ensuring interfaces in compile-time for ContactsUnblockRequest.
var (
	_ bin.Encoder     = &ContactsUnblockRequest{}
	_ bin.Decoder     = &ContactsUnblockRequest{}
	_ bin.BareEncoder = &ContactsUnblockRequest{}
	_ bin.BareDecoder = &ContactsUnblockRequest{}
)

func (u *ContactsUnblockRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.MyStoriesFrom == false) {
		return false
	}
	if !(u.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *ContactsUnblockRequest) String() string {
	if u == nil {
		return "ContactsUnblockRequest(nil)"
	}
	type Alias ContactsUnblockRequest
	return fmt.Sprintf("ContactsUnblockRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsUnblockRequest) TypeID() uint32 {
	return ContactsUnblockRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsUnblockRequest) TypeName() string {
	return "contacts.unblock"
}

// TypeInfo returns info about TL type.
func (u *ContactsUnblockRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.unblock",
		ID:   ContactsUnblockRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MyStoriesFrom",
			SchemaName: "my_stories_from",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *ContactsUnblockRequest) SetFlags() {
	if !(u.MyStoriesFrom == false) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *ContactsUnblockRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode contacts.unblock#b550d328 as nil")
	}
	b.PutID(ContactsUnblockRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *ContactsUnblockRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode contacts.unblock#b550d328 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.unblock#b550d328: field flags: %w", err)
	}
	if u.ID == nil {
		return fmt.Errorf("unable to encode contacts.unblock#b550d328: field id is nil")
	}
	if err := u.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.unblock#b550d328: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *ContactsUnblockRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode contacts.unblock#b550d328 to nil")
	}
	if err := b.ConsumeID(ContactsUnblockRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.unblock#b550d328: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *ContactsUnblockRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode contacts.unblock#b550d328 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.unblock#b550d328: field flags: %w", err)
		}
	}
	u.MyStoriesFrom = u.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.unblock#b550d328: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// SetMyStoriesFrom sets value of MyStoriesFrom conditional field.
func (u *ContactsUnblockRequest) SetMyStoriesFrom(value bool) {
	if value {
		u.Flags.Set(0)
		u.MyStoriesFrom = true
	} else {
		u.Flags.Unset(0)
		u.MyStoriesFrom = false
	}
}

// GetMyStoriesFrom returns value of MyStoriesFrom conditional field.
func (u *ContactsUnblockRequest) GetMyStoriesFrom() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// GetID returns value of ID field.
func (u *ContactsUnblockRequest) GetID() (value InputPeerClass) {
	if u == nil {
		return
	}
	return u.ID
}

// ContactsUnblock invokes method contacts.unblock#b550d328 returning error if any.
func (c *Client) ContactsUnblock(ctx context.Context, request *ContactsUnblockRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
