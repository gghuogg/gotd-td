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

// PaymentsBankCardData represents TL type `payments.bankCardData#3e24e573`.
type PaymentsBankCardData struct {
	// Title field of PaymentsBankCardData.
	Title string
	// OpenURLs field of PaymentsBankCardData.
	OpenURLs []BankCardOpenURL
}

// PaymentsBankCardDataTypeID is TL type id of PaymentsBankCardData.
const PaymentsBankCardDataTypeID = 0x3e24e573

// Ensuring interfaces in compile-time for PaymentsBankCardData.
var (
	_ bin.Encoder     = &PaymentsBankCardData{}
	_ bin.Decoder     = &PaymentsBankCardData{}
	_ bin.BareEncoder = &PaymentsBankCardData{}
	_ bin.BareDecoder = &PaymentsBankCardData{}
)

func (b *PaymentsBankCardData) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Title == "") {
		return false
	}
	if !(b.OpenURLs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *PaymentsBankCardData) String() string {
	if b == nil {
		return "PaymentsBankCardData(nil)"
	}
	type Alias PaymentsBankCardData
	return fmt.Sprintf("PaymentsBankCardData%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsBankCardData) TypeID() uint32 {
	return PaymentsBankCardDataTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsBankCardData) TypeName() string {
	return "payments.bankCardData"
}

// TypeInfo returns info about TL type.
func (b *PaymentsBankCardData) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.bankCardData",
		ID:   PaymentsBankCardDataTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "OpenURLs",
			SchemaName: "open_urls",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *PaymentsBankCardData) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode payments.bankCardData#3e24e573 as nil")
	}
	buf.PutID(PaymentsBankCardDataTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *PaymentsBankCardData) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode payments.bankCardData#3e24e573 as nil")
	}
	buf.PutString(b.Title)
	buf.PutVectorHeader(len(b.OpenURLs))
	for idx, v := range b.OpenURLs {
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode payments.bankCardData#3e24e573: field open_urls element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *PaymentsBankCardData) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode payments.bankCardData#3e24e573 to nil")
	}
	if err := buf.ConsumeID(PaymentsBankCardDataTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *PaymentsBankCardData) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode payments.bankCardData#3e24e573 to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: field title: %w", err)
		}
		b.Title = value
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: field open_urls: %w", err)
		}

		if headerLen > 0 {
			b.OpenURLs = make([]BankCardOpenURL, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BankCardOpenURL
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: field open_urls: %w", err)
			}
			b.OpenURLs = append(b.OpenURLs, value)
		}
	}
	return nil
}

// GetTitle returns value of Title field.
func (b *PaymentsBankCardData) GetTitle() (value string) {
	if b == nil {
		return
	}
	return b.Title
}

// GetOpenURLs returns value of OpenURLs field.
func (b *PaymentsBankCardData) GetOpenURLs() (value []BankCardOpenURL) {
	if b == nil {
		return
	}
	return b.OpenURLs
}
