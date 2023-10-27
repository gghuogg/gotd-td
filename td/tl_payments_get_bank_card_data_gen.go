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

// PaymentsGetBankCardDataRequest represents TL type `payments.getBankCardData#2e79d779`.
type PaymentsGetBankCardDataRequest struct {
	// Number field of PaymentsGetBankCardDataRequest.
	Number string
}

// PaymentsGetBankCardDataRequestTypeID is TL type id of PaymentsGetBankCardDataRequest.
const PaymentsGetBankCardDataRequestTypeID = 0x2e79d779

// Ensuring interfaces in compile-time for PaymentsGetBankCardDataRequest.
var (
	_ bin.Encoder     = &PaymentsGetBankCardDataRequest{}
	_ bin.Decoder     = &PaymentsGetBankCardDataRequest{}
	_ bin.BareEncoder = &PaymentsGetBankCardDataRequest{}
	_ bin.BareDecoder = &PaymentsGetBankCardDataRequest{}
)

func (g *PaymentsGetBankCardDataRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Number == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetBankCardDataRequest) String() string {
	if g == nil {
		return "PaymentsGetBankCardDataRequest(nil)"
	}
	type Alias PaymentsGetBankCardDataRequest
	return fmt.Sprintf("PaymentsGetBankCardDataRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetBankCardDataRequest) TypeID() uint32 {
	return PaymentsGetBankCardDataRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetBankCardDataRequest) TypeName() string {
	return "payments.getBankCardData"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetBankCardDataRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getBankCardData",
		ID:   PaymentsGetBankCardDataRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Number",
			SchemaName: "number",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PaymentsGetBankCardDataRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getBankCardData#2e79d779 as nil")
	}
	b.PutID(PaymentsGetBankCardDataRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetBankCardDataRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getBankCardData#2e79d779 as nil")
	}
	b.PutString(g.Number)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetBankCardDataRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getBankCardData#2e79d779 to nil")
	}
	if err := b.ConsumeID(PaymentsGetBankCardDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getBankCardData#2e79d779: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetBankCardDataRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getBankCardData#2e79d779 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getBankCardData#2e79d779: field number: %w", err)
		}
		g.Number = value
	}
	return nil
}

// GetNumber returns value of Number field.
func (g *PaymentsGetBankCardDataRequest) GetNumber() (value string) {
	if g == nil {
		return
	}
	return g.Number
}

// PaymentsGetBankCardData invokes method payments.getBankCardData#2e79d779 returning error if any.
func (c *Client) PaymentsGetBankCardData(ctx context.Context, number string) (*PaymentsBankCardData, error) {
	var result PaymentsBankCardData

	request := &PaymentsGetBankCardDataRequest{
		Number: number,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
