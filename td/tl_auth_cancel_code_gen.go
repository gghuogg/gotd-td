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

// AuthCancelCodeRequest represents TL type `auth.cancelCode#1f040578`.
type AuthCancelCodeRequest struct {
	// PhoneNumber field of AuthCancelCodeRequest.
	PhoneNumber string
	// PhoneCodeHash field of AuthCancelCodeRequest.
	PhoneCodeHash string
}

// AuthCancelCodeRequestTypeID is TL type id of AuthCancelCodeRequest.
const AuthCancelCodeRequestTypeID = 0x1f040578

// Ensuring interfaces in compile-time for AuthCancelCodeRequest.
var (
	_ bin.Encoder     = &AuthCancelCodeRequest{}
	_ bin.Decoder     = &AuthCancelCodeRequest{}
	_ bin.BareEncoder = &AuthCancelCodeRequest{}
	_ bin.BareDecoder = &AuthCancelCodeRequest{}
)

func (c *AuthCancelCodeRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.PhoneNumber == "") {
		return false
	}
	if !(c.PhoneCodeHash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AuthCancelCodeRequest) String() string {
	if c == nil {
		return "AuthCancelCodeRequest(nil)"
	}
	type Alias AuthCancelCodeRequest
	return fmt.Sprintf("AuthCancelCodeRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthCancelCodeRequest) TypeID() uint32 {
	return AuthCancelCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthCancelCodeRequest) TypeName() string {
	return "auth.cancelCode"
}

// TypeInfo returns info about TL type.
func (c *AuthCancelCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.cancelCode",
		ID:   AuthCancelCodeRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "PhoneCodeHash",
			SchemaName: "phone_code_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *AuthCancelCodeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.cancelCode#1f040578 as nil")
	}
	b.PutID(AuthCancelCodeRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AuthCancelCodeRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.cancelCode#1f040578 as nil")
	}
	b.PutString(c.PhoneNumber)
	b.PutString(c.PhoneCodeHash)
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCancelCodeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.cancelCode#1f040578 to nil")
	}
	if err := b.ConsumeID(AuthCancelCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.cancelCode#1f040578: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AuthCancelCodeRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.cancelCode#1f040578 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.cancelCode#1f040578: field phone_number: %w", err)
		}
		c.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.cancelCode#1f040578: field phone_code_hash: %w", err)
		}
		c.PhoneCodeHash = value
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (c *AuthCancelCodeRequest) GetPhoneNumber() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (c *AuthCancelCodeRequest) GetPhoneCodeHash() (value string) {
	if c == nil {
		return
	}
	return c.PhoneCodeHash
}

// AuthCancelCode invokes method auth.cancelCode#1f040578 returning error if any.
func (c *Client) AuthCancelCode(ctx context.Context, request *AuthCancelCodeRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
