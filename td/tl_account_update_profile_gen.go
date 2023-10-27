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

// AccountUpdateProfileRequest represents TL type `account.updateProfile#78515775`.
type AccountUpdateProfileRequest struct {
	// Flags field of AccountUpdateProfileRequest.
	Flags bin.Fields
	// FirstName field of AccountUpdateProfileRequest.
	//
	// Use SetFirstName and GetFirstName helpers.
	FirstName string
	// LastName field of AccountUpdateProfileRequest.
	//
	// Use SetLastName and GetLastName helpers.
	LastName string
	// About field of AccountUpdateProfileRequest.
	//
	// Use SetAbout and GetAbout helpers.
	About string
}

// AccountUpdateProfileRequestTypeID is TL type id of AccountUpdateProfileRequest.
const AccountUpdateProfileRequestTypeID = 0x78515775

// Ensuring interfaces in compile-time for AccountUpdateProfileRequest.
var (
	_ bin.Encoder     = &AccountUpdateProfileRequest{}
	_ bin.Decoder     = &AccountUpdateProfileRequest{}
	_ bin.BareEncoder = &AccountUpdateProfileRequest{}
	_ bin.BareDecoder = &AccountUpdateProfileRequest{}
)

func (u *AccountUpdateProfileRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.FirstName == "") {
		return false
	}
	if !(u.LastName == "") {
		return false
	}
	if !(u.About == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateProfileRequest) String() string {
	if u == nil {
		return "AccountUpdateProfileRequest(nil)"
	}
	type Alias AccountUpdateProfileRequest
	return fmt.Sprintf("AccountUpdateProfileRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateProfileRequest) TypeID() uint32 {
	return AccountUpdateProfileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateProfileRequest) TypeName() string {
	return "account.updateProfile"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateProfileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateProfile",
		ID:   AccountUpdateProfileRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FirstName",
			SchemaName: "first_name",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "About",
			SchemaName: "about",
			Null:       !u.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *AccountUpdateProfileRequest) SetFlags() {
	if !(u.FirstName == "") {
		u.Flags.Set(0)
	}
	if !(u.LastName == "") {
		u.Flags.Set(1)
	}
	if !(u.About == "") {
		u.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (u *AccountUpdateProfileRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateProfile#78515775 as nil")
	}
	b.PutID(AccountUpdateProfileRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateProfileRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateProfile#78515775 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateProfile#78515775: field flags: %w", err)
	}
	if u.Flags.Has(0) {
		b.PutString(u.FirstName)
	}
	if u.Flags.Has(1) {
		b.PutString(u.LastName)
	}
	if u.Flags.Has(2) {
		b.PutString(u.About)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateProfileRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateProfile#78515775 to nil")
	}
	if err := b.ConsumeID(AccountUpdateProfileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateProfile#78515775: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateProfileRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateProfile#78515775 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateProfile#78515775: field flags: %w", err)
		}
	}
	if u.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateProfile#78515775: field first_name: %w", err)
		}
		u.FirstName = value
	}
	if u.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateProfile#78515775: field last_name: %w", err)
		}
		u.LastName = value
	}
	if u.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateProfile#78515775: field about: %w", err)
		}
		u.About = value
	}
	return nil
}

// SetFirstName sets value of FirstName conditional field.
func (u *AccountUpdateProfileRequest) SetFirstName(value string) {
	u.Flags.Set(0)
	u.FirstName = value
}

// GetFirstName returns value of FirstName conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateProfileRequest) GetFirstName() (value string, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.FirstName, true
}

// SetLastName sets value of LastName conditional field.
func (u *AccountUpdateProfileRequest) SetLastName(value string) {
	u.Flags.Set(1)
	u.LastName = value
}

// GetLastName returns value of LastName conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateProfileRequest) GetLastName() (value string, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.LastName, true
}

// SetAbout sets value of About conditional field.
func (u *AccountUpdateProfileRequest) SetAbout(value string) {
	u.Flags.Set(2)
	u.About = value
}

// GetAbout returns value of About conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateProfileRequest) GetAbout() (value string, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.About, true
}

// AccountUpdateProfile invokes method account.updateProfile#78515775 returning error if any.
func (c *Client) AccountUpdateProfile(ctx context.Context, request *AccountUpdateProfileRequest) (UserClass, error) {
	var result UserBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.User, nil
}
