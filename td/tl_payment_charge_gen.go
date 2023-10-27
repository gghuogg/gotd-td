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

// PaymentCharge represents TL type `paymentCharge#ea02c27e`.
type PaymentCharge struct {
	// ID field of PaymentCharge.
	ID string
	// ProviderChargeID field of PaymentCharge.
	ProviderChargeID string
}

// PaymentChargeTypeID is TL type id of PaymentCharge.
const PaymentChargeTypeID = 0xea02c27e

// Ensuring interfaces in compile-time for PaymentCharge.
var (
	_ bin.Encoder     = &PaymentCharge{}
	_ bin.Decoder     = &PaymentCharge{}
	_ bin.BareEncoder = &PaymentCharge{}
	_ bin.BareDecoder = &PaymentCharge{}
)

func (p *PaymentCharge) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == "") {
		return false
	}
	if !(p.ProviderChargeID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentCharge) String() string {
	if p == nil {
		return "PaymentCharge(nil)"
	}
	type Alias PaymentCharge
	return fmt.Sprintf("PaymentCharge%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentCharge) TypeID() uint32 {
	return PaymentChargeTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentCharge) TypeName() string {
	return "paymentCharge"
}

// TypeInfo returns info about TL type.
func (p *PaymentCharge) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentCharge",
		ID:   PaymentChargeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "ProviderChargeID",
			SchemaName: "provider_charge_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentCharge) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentCharge#ea02c27e as nil")
	}
	b.PutID(PaymentChargeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentCharge) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentCharge#ea02c27e as nil")
	}
	b.PutString(p.ID)
	b.PutString(p.ProviderChargeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentCharge) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentCharge#ea02c27e to nil")
	}
	if err := b.ConsumeID(PaymentChargeTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentCharge#ea02c27e: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentCharge) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentCharge#ea02c27e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentCharge#ea02c27e: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentCharge#ea02c27e: field provider_charge_id: %w", err)
		}
		p.ProviderChargeID = value
	}
	return nil
}

// GetID returns value of ID field.
func (p *PaymentCharge) GetID() (value string) {
	if p == nil {
		return
	}
	return p.ID
}

// GetProviderChargeID returns value of ProviderChargeID field.
func (p *PaymentCharge) GetProviderChargeID() (value string) {
	if p == nil {
		return
	}
	return p.ProviderChargeID
}
