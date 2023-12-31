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

// AccountConfirmPasswordEmailRequest represents TL type `account.confirmPasswordEmail#8fdf1920`.
type AccountConfirmPasswordEmailRequest struct {
	// Code field of AccountConfirmPasswordEmailRequest.
	Code string
}

// AccountConfirmPasswordEmailRequestTypeID is TL type id of AccountConfirmPasswordEmailRequest.
const AccountConfirmPasswordEmailRequestTypeID = 0x8fdf1920

// Ensuring interfaces in compile-time for AccountConfirmPasswordEmailRequest.
var (
	_ bin.Encoder     = &AccountConfirmPasswordEmailRequest{}
	_ bin.Decoder     = &AccountConfirmPasswordEmailRequest{}
	_ bin.BareEncoder = &AccountConfirmPasswordEmailRequest{}
	_ bin.BareDecoder = &AccountConfirmPasswordEmailRequest{}
)

func (c *AccountConfirmPasswordEmailRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Code == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountConfirmPasswordEmailRequest) String() string {
	if c == nil {
		return "AccountConfirmPasswordEmailRequest(nil)"
	}
	type Alias AccountConfirmPasswordEmailRequest
	return fmt.Sprintf("AccountConfirmPasswordEmailRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountConfirmPasswordEmailRequest) TypeID() uint32 {
	return AccountConfirmPasswordEmailRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountConfirmPasswordEmailRequest) TypeName() string {
	return "account.confirmPasswordEmail"
}

// TypeInfo returns info about TL type.
func (c *AccountConfirmPasswordEmailRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.confirmPasswordEmail",
		ID:   AccountConfirmPasswordEmailRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *AccountConfirmPasswordEmailRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.confirmPasswordEmail#8fdf1920 as nil")
	}
	b.PutID(AccountConfirmPasswordEmailRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AccountConfirmPasswordEmailRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.confirmPasswordEmail#8fdf1920 as nil")
	}
	b.PutString(c.Code)
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountConfirmPasswordEmailRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.confirmPasswordEmail#8fdf1920 to nil")
	}
	if err := b.ConsumeID(AccountConfirmPasswordEmailRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.confirmPasswordEmail#8fdf1920: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AccountConfirmPasswordEmailRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.confirmPasswordEmail#8fdf1920 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.confirmPasswordEmail#8fdf1920: field code: %w", err)
		}
		c.Code = value
	}
	return nil
}

// GetCode returns value of Code field.
func (c *AccountConfirmPasswordEmailRequest) GetCode() (value string) {
	if c == nil {
		return
	}
	return c.Code
}

// AccountConfirmPasswordEmail invokes method account.confirmPasswordEmail#8fdf1920 returning error if any.
func (c *Client) AccountConfirmPasswordEmail(ctx context.Context, code string) (bool, error) {
	var result BoolBox

	request := &AccountConfirmPasswordEmailRequest{
		Code: code,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
