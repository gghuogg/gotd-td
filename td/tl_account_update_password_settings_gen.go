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

// AccountUpdatePasswordSettingsRequest represents TL type `account.updatePasswordSettings#a59b102f`.
type AccountUpdatePasswordSettingsRequest struct {
	// Password field of AccountUpdatePasswordSettingsRequest.
	Password InputCheckPasswordSRPClass
	// NewSettings field of AccountUpdatePasswordSettingsRequest.
	NewSettings AccountPasswordInputSettings
}

// AccountUpdatePasswordSettingsRequestTypeID is TL type id of AccountUpdatePasswordSettingsRequest.
const AccountUpdatePasswordSettingsRequestTypeID = 0xa59b102f

// Ensuring interfaces in compile-time for AccountUpdatePasswordSettingsRequest.
var (
	_ bin.Encoder     = &AccountUpdatePasswordSettingsRequest{}
	_ bin.Decoder     = &AccountUpdatePasswordSettingsRequest{}
	_ bin.BareEncoder = &AccountUpdatePasswordSettingsRequest{}
	_ bin.BareDecoder = &AccountUpdatePasswordSettingsRequest{}
)

func (u *AccountUpdatePasswordSettingsRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Password == nil) {
		return false
	}
	if !(u.NewSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdatePasswordSettingsRequest) String() string {
	if u == nil {
		return "AccountUpdatePasswordSettingsRequest(nil)"
	}
	type Alias AccountUpdatePasswordSettingsRequest
	return fmt.Sprintf("AccountUpdatePasswordSettingsRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdatePasswordSettingsRequest) TypeID() uint32 {
	return AccountUpdatePasswordSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdatePasswordSettingsRequest) TypeName() string {
	return "account.updatePasswordSettings"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdatePasswordSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updatePasswordSettings",
		ID:   AccountUpdatePasswordSettingsRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
		{
			Name:       "NewSettings",
			SchemaName: "new_settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdatePasswordSettingsRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updatePasswordSettings#a59b102f as nil")
	}
	b.PutID(AccountUpdatePasswordSettingsRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdatePasswordSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updatePasswordSettings#a59b102f as nil")
	}
	if u.Password == nil {
		return fmt.Errorf("unable to encode account.updatePasswordSettings#a59b102f: field password is nil")
	}
	if err := u.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updatePasswordSettings#a59b102f: field password: %w", err)
	}
	if err := u.NewSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updatePasswordSettings#a59b102f: field new_settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdatePasswordSettingsRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updatePasswordSettings#a59b102f to nil")
	}
	if err := b.ConsumeID(AccountUpdatePasswordSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updatePasswordSettings#a59b102f: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdatePasswordSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updatePasswordSettings#a59b102f to nil")
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updatePasswordSettings#a59b102f: field password: %w", err)
		}
		u.Password = value
	}
	{
		if err := u.NewSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updatePasswordSettings#a59b102f: field new_settings: %w", err)
		}
	}
	return nil
}

// GetPassword returns value of Password field.
func (u *AccountUpdatePasswordSettingsRequest) GetPassword() (value InputCheckPasswordSRPClass) {
	if u == nil {
		return
	}
	return u.Password
}

// GetNewSettings returns value of NewSettings field.
func (u *AccountUpdatePasswordSettingsRequest) GetNewSettings() (value AccountPasswordInputSettings) {
	if u == nil {
		return
	}
	return u.NewSettings
}

// AccountUpdatePasswordSettings invokes method account.updatePasswordSettings#a59b102f returning error if any.
func (c *Client) AccountUpdatePasswordSettings(ctx context.Context, request *AccountUpdatePasswordSettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
