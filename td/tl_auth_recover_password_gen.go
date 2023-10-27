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

// AuthRecoverPasswordRequest represents TL type `auth.recoverPassword#37096c70`.
type AuthRecoverPasswordRequest struct {
	// Flags field of AuthRecoverPasswordRequest.
	Flags bin.Fields
	// Code field of AuthRecoverPasswordRequest.
	Code string
	// NewSettings field of AuthRecoverPasswordRequest.
	//
	// Use SetNewSettings and GetNewSettings helpers.
	NewSettings AccountPasswordInputSettings
}

// AuthRecoverPasswordRequestTypeID is TL type id of AuthRecoverPasswordRequest.
const AuthRecoverPasswordRequestTypeID = 0x37096c70

// Ensuring interfaces in compile-time for AuthRecoverPasswordRequest.
var (
	_ bin.Encoder     = &AuthRecoverPasswordRequest{}
	_ bin.Decoder     = &AuthRecoverPasswordRequest{}
	_ bin.BareEncoder = &AuthRecoverPasswordRequest{}
	_ bin.BareDecoder = &AuthRecoverPasswordRequest{}
)

func (r *AuthRecoverPasswordRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Code == "") {
		return false
	}
	if !(r.NewSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthRecoverPasswordRequest) String() string {
	if r == nil {
		return "AuthRecoverPasswordRequest(nil)"
	}
	type Alias AuthRecoverPasswordRequest
	return fmt.Sprintf("AuthRecoverPasswordRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthRecoverPasswordRequest) TypeID() uint32 {
	return AuthRecoverPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthRecoverPasswordRequest) TypeName() string {
	return "auth.recoverPassword"
}

// TypeInfo returns info about TL type.
func (r *AuthRecoverPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.recoverPassword",
		ID:   AuthRecoverPasswordRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
		{
			Name:       "NewSettings",
			SchemaName: "new_settings",
			Null:       !r.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *AuthRecoverPasswordRequest) SetFlags() {
	if !(r.NewSettings.Zero()) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *AuthRecoverPasswordRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.recoverPassword#37096c70 as nil")
	}
	b.PutID(AuthRecoverPasswordRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AuthRecoverPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.recoverPassword#37096c70 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.recoverPassword#37096c70: field flags: %w", err)
	}
	b.PutString(r.Code)
	if r.Flags.Has(0) {
		if err := r.NewSettings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode auth.recoverPassword#37096c70: field new_settings: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthRecoverPasswordRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.recoverPassword#37096c70 to nil")
	}
	if err := b.ConsumeID(AuthRecoverPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.recoverPassword#37096c70: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AuthRecoverPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.recoverPassword#37096c70 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.recoverPassword#37096c70: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.recoverPassword#37096c70: field code: %w", err)
		}
		r.Code = value
	}
	if r.Flags.Has(0) {
		if err := r.NewSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.recoverPassword#37096c70: field new_settings: %w", err)
		}
	}
	return nil
}

// GetCode returns value of Code field.
func (r *AuthRecoverPasswordRequest) GetCode() (value string) {
	if r == nil {
		return
	}
	return r.Code
}

// SetNewSettings sets value of NewSettings conditional field.
func (r *AuthRecoverPasswordRequest) SetNewSettings(value AccountPasswordInputSettings) {
	r.Flags.Set(0)
	r.NewSettings = value
}

// GetNewSettings returns value of NewSettings conditional field and
// boolean which is true if field was set.
func (r *AuthRecoverPasswordRequest) GetNewSettings() (value AccountPasswordInputSettings, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.NewSettings, true
}

// AuthRecoverPassword invokes method auth.recoverPassword#37096c70 returning error if any.
func (c *Client) AuthRecoverPassword(ctx context.Context, request *AuthRecoverPasswordRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
