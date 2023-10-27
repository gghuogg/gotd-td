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

// ContactsResetSavedRequest represents TL type `contacts.resetSaved#879537f1`.
type ContactsResetSavedRequest struct {
}

// ContactsResetSavedRequestTypeID is TL type id of ContactsResetSavedRequest.
const ContactsResetSavedRequestTypeID = 0x879537f1

// Ensuring interfaces in compile-time for ContactsResetSavedRequest.
var (
	_ bin.Encoder     = &ContactsResetSavedRequest{}
	_ bin.Decoder     = &ContactsResetSavedRequest{}
	_ bin.BareEncoder = &ContactsResetSavedRequest{}
	_ bin.BareDecoder = &ContactsResetSavedRequest{}
)

func (r *ContactsResetSavedRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ContactsResetSavedRequest) String() string {
	if r == nil {
		return "ContactsResetSavedRequest(nil)"
	}
	type Alias ContactsResetSavedRequest
	return fmt.Sprintf("ContactsResetSavedRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsResetSavedRequest) TypeID() uint32 {
	return ContactsResetSavedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsResetSavedRequest) TypeName() string {
	return "contacts.resetSaved"
}

// TypeInfo returns info about TL type.
func (r *ContactsResetSavedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.resetSaved",
		ID:   ContactsResetSavedRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ContactsResetSavedRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resetSaved#879537f1 as nil")
	}
	b.PutID(ContactsResetSavedRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ContactsResetSavedRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resetSaved#879537f1 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ContactsResetSavedRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resetSaved#879537f1 to nil")
	}
	if err := b.ConsumeID(ContactsResetSavedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.resetSaved#879537f1: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ContactsResetSavedRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resetSaved#879537f1 to nil")
	}
	return nil
}

// ContactsResetSaved invokes method contacts.resetSaved#879537f1 returning error if any.
func (c *Client) ContactsResetSaved(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &ContactsResetSavedRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
