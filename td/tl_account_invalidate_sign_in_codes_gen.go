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

// AccountInvalidateSignInCodesRequest represents TL type `account.invalidateSignInCodes#ca8ae8ba`.
type AccountInvalidateSignInCodesRequest struct {
	// Codes field of AccountInvalidateSignInCodesRequest.
	Codes []string
}

// AccountInvalidateSignInCodesRequestTypeID is TL type id of AccountInvalidateSignInCodesRequest.
const AccountInvalidateSignInCodesRequestTypeID = 0xca8ae8ba

// Ensuring interfaces in compile-time for AccountInvalidateSignInCodesRequest.
var (
	_ bin.Encoder     = &AccountInvalidateSignInCodesRequest{}
	_ bin.Decoder     = &AccountInvalidateSignInCodesRequest{}
	_ bin.BareEncoder = &AccountInvalidateSignInCodesRequest{}
	_ bin.BareDecoder = &AccountInvalidateSignInCodesRequest{}
)

func (i *AccountInvalidateSignInCodesRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Codes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AccountInvalidateSignInCodesRequest) String() string {
	if i == nil {
		return "AccountInvalidateSignInCodesRequest(nil)"
	}
	type Alias AccountInvalidateSignInCodesRequest
	return fmt.Sprintf("AccountInvalidateSignInCodesRequest%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountInvalidateSignInCodesRequest) TypeID() uint32 {
	return AccountInvalidateSignInCodesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountInvalidateSignInCodesRequest) TypeName() string {
	return "account.invalidateSignInCodes"
}

// TypeInfo returns info about TL type.
func (i *AccountInvalidateSignInCodesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.invalidateSignInCodes",
		ID:   AccountInvalidateSignInCodesRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Codes",
			SchemaName: "codes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *AccountInvalidateSignInCodesRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.invalidateSignInCodes#ca8ae8ba as nil")
	}
	b.PutID(AccountInvalidateSignInCodesRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *AccountInvalidateSignInCodesRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.invalidateSignInCodes#ca8ae8ba as nil")
	}
	b.PutVectorHeader(len(i.Codes))
	for _, v := range i.Codes {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *AccountInvalidateSignInCodesRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.invalidateSignInCodes#ca8ae8ba to nil")
	}
	if err := b.ConsumeID(AccountInvalidateSignInCodesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.invalidateSignInCodes#ca8ae8ba: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *AccountInvalidateSignInCodesRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.invalidateSignInCodes#ca8ae8ba to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.invalidateSignInCodes#ca8ae8ba: field codes: %w", err)
		}

		if headerLen > 0 {
			i.Codes = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode account.invalidateSignInCodes#ca8ae8ba: field codes: %w", err)
			}
			i.Codes = append(i.Codes, value)
		}
	}
	return nil
}

// GetCodes returns value of Codes field.
func (i *AccountInvalidateSignInCodesRequest) GetCodes() (value []string) {
	if i == nil {
		return
	}
	return i.Codes
}

// AccountInvalidateSignInCodes invokes method account.invalidateSignInCodes#ca8ae8ba returning error if any.
func (c *Client) AccountInvalidateSignInCodes(ctx context.Context, codes []string) (bool, error) {
	var result BoolBox

	request := &AccountInvalidateSignInCodesRequest{
		Codes: codes,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
