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

// ContactsAddContactRequest represents TL type `contacts.addContact#e8f463d0`.
type ContactsAddContactRequest struct {
	// Flags field of ContactsAddContactRequest.
	Flags bin.Fields
	// AddPhonePrivacyException field of ContactsAddContactRequest.
	AddPhonePrivacyException bool
	// ID field of ContactsAddContactRequest.
	ID InputUserClass
	// FirstName field of ContactsAddContactRequest.
	FirstName string
	// LastName field of ContactsAddContactRequest.
	LastName string
	// Phone field of ContactsAddContactRequest.
	Phone string
}

// ContactsAddContactRequestTypeID is TL type id of ContactsAddContactRequest.
const ContactsAddContactRequestTypeID = 0xe8f463d0

// Ensuring interfaces in compile-time for ContactsAddContactRequest.
var (
	_ bin.Encoder     = &ContactsAddContactRequest{}
	_ bin.Decoder     = &ContactsAddContactRequest{}
	_ bin.BareEncoder = &ContactsAddContactRequest{}
	_ bin.BareDecoder = &ContactsAddContactRequest{}
)

func (a *ContactsAddContactRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.AddPhonePrivacyException == false) {
		return false
	}
	if !(a.ID == nil) {
		return false
	}
	if !(a.FirstName == "") {
		return false
	}
	if !(a.LastName == "") {
		return false
	}
	if !(a.Phone == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *ContactsAddContactRequest) String() string {
	if a == nil {
		return "ContactsAddContactRequest(nil)"
	}
	type Alias ContactsAddContactRequest
	return fmt.Sprintf("ContactsAddContactRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsAddContactRequest) TypeID() uint32 {
	return ContactsAddContactRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsAddContactRequest) TypeName() string {
	return "contacts.addContact"
}

// TypeInfo returns info about TL type.
func (a *ContactsAddContactRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.addContact",
		ID:   ContactsAddContactRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AddPhonePrivacyException",
			SchemaName: "add_phone_privacy_exception",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "FirstName",
			SchemaName: "first_name",
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
		},
		{
			Name:       "Phone",
			SchemaName: "phone",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *ContactsAddContactRequest) SetFlags() {
	if !(a.AddPhonePrivacyException == false) {
		a.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (a *ContactsAddContactRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode contacts.addContact#e8f463d0 as nil")
	}
	b.PutID(ContactsAddContactRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *ContactsAddContactRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode contacts.addContact#e8f463d0 as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.addContact#e8f463d0: field flags: %w", err)
	}
	if a.ID == nil {
		return fmt.Errorf("unable to encode contacts.addContact#e8f463d0: field id is nil")
	}
	if err := a.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.addContact#e8f463d0: field id: %w", err)
	}
	b.PutString(a.FirstName)
	b.PutString(a.LastName)
	b.PutString(a.Phone)
	return nil
}

// Decode implements bin.Decoder.
func (a *ContactsAddContactRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode contacts.addContact#e8f463d0 to nil")
	}
	if err := b.ConsumeID(ContactsAddContactRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *ContactsAddContactRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode contacts.addContact#e8f463d0 to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field flags: %w", err)
		}
	}
	a.AddPhonePrivacyException = a.Flags.Has(0)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field id: %w", err)
		}
		a.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field first_name: %w", err)
		}
		a.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field last_name: %w", err)
		}
		a.LastName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field phone: %w", err)
		}
		a.Phone = value
	}
	return nil
}

// SetAddPhonePrivacyException sets value of AddPhonePrivacyException conditional field.
func (a *ContactsAddContactRequest) SetAddPhonePrivacyException(value bool) {
	if value {
		a.Flags.Set(0)
		a.AddPhonePrivacyException = true
	} else {
		a.Flags.Unset(0)
		a.AddPhonePrivacyException = false
	}
}

// GetAddPhonePrivacyException returns value of AddPhonePrivacyException conditional field.
func (a *ContactsAddContactRequest) GetAddPhonePrivacyException() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// GetID returns value of ID field.
func (a *ContactsAddContactRequest) GetID() (value InputUserClass) {
	if a == nil {
		return
	}
	return a.ID
}

// GetFirstName returns value of FirstName field.
func (a *ContactsAddContactRequest) GetFirstName() (value string) {
	if a == nil {
		return
	}
	return a.FirstName
}

// GetLastName returns value of LastName field.
func (a *ContactsAddContactRequest) GetLastName() (value string) {
	if a == nil {
		return
	}
	return a.LastName
}

// GetPhone returns value of Phone field.
func (a *ContactsAddContactRequest) GetPhone() (value string) {
	if a == nil {
		return
	}
	return a.Phone
}

// ContactsAddContact invokes method contacts.addContact#e8f463d0 returning error if any.
func (c *Client) ContactsAddContact(ctx context.Context, request *ContactsAddContactRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
