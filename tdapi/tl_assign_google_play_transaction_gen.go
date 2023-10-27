// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// AssignGooglePlayTransactionRequest represents TL type `assignGooglePlayTransaction#8939bca4`.
type AssignGooglePlayTransactionRequest struct {
	// Application package name
	PackageName string
	// Identifier of the purchased store product
	StoreProductID string
	// Google Play purchase token
	PurchaseToken string
	// Transaction purpose
	Purpose StorePaymentPurposeClass
}

// AssignGooglePlayTransactionRequestTypeID is TL type id of AssignGooglePlayTransactionRequest.
const AssignGooglePlayTransactionRequestTypeID = 0x8939bca4

// Ensuring interfaces in compile-time for AssignGooglePlayTransactionRequest.
var (
	_ bin.Encoder     = &AssignGooglePlayTransactionRequest{}
	_ bin.Decoder     = &AssignGooglePlayTransactionRequest{}
	_ bin.BareEncoder = &AssignGooglePlayTransactionRequest{}
	_ bin.BareDecoder = &AssignGooglePlayTransactionRequest{}
)

func (a *AssignGooglePlayTransactionRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.PackageName == "") {
		return false
	}
	if !(a.StoreProductID == "") {
		return false
	}
	if !(a.PurchaseToken == "") {
		return false
	}
	if !(a.Purpose == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AssignGooglePlayTransactionRequest) String() string {
	if a == nil {
		return "AssignGooglePlayTransactionRequest(nil)"
	}
	type Alias AssignGooglePlayTransactionRequest
	return fmt.Sprintf("AssignGooglePlayTransactionRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AssignGooglePlayTransactionRequest) TypeID() uint32 {
	return AssignGooglePlayTransactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AssignGooglePlayTransactionRequest) TypeName() string {
	return "assignGooglePlayTransaction"
}

// TypeInfo returns info about TL type.
func (a *AssignGooglePlayTransactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "assignGooglePlayTransaction",
		ID:   AssignGooglePlayTransactionRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PackageName",
			SchemaName: "package_name",
		},
		{
			Name:       "StoreProductID",
			SchemaName: "store_product_id",
		},
		{
			Name:       "PurchaseToken",
			SchemaName: "purchase_token",
		},
		{
			Name:       "Purpose",
			SchemaName: "purpose",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AssignGooglePlayTransactionRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode assignGooglePlayTransaction#8939bca4 as nil")
	}
	b.PutID(AssignGooglePlayTransactionRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AssignGooglePlayTransactionRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode assignGooglePlayTransaction#8939bca4 as nil")
	}
	b.PutString(a.PackageName)
	b.PutString(a.StoreProductID)
	b.PutString(a.PurchaseToken)
	if a.Purpose == nil {
		return fmt.Errorf("unable to encode assignGooglePlayTransaction#8939bca4: field purpose is nil")
	}
	if err := a.Purpose.Encode(b); err != nil {
		return fmt.Errorf("unable to encode assignGooglePlayTransaction#8939bca4: field purpose: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AssignGooglePlayTransactionRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode assignGooglePlayTransaction#8939bca4 to nil")
	}
	if err := b.ConsumeID(AssignGooglePlayTransactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AssignGooglePlayTransactionRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode assignGooglePlayTransaction#8939bca4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: field package_name: %w", err)
		}
		a.PackageName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: field store_product_id: %w", err)
		}
		a.StoreProductID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: field purchase_token: %w", err)
		}
		a.PurchaseToken = value
	}
	{
		value, err := DecodeStorePaymentPurpose(b)
		if err != nil {
			return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: field purpose: %w", err)
		}
		a.Purpose = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AssignGooglePlayTransactionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode assignGooglePlayTransaction#8939bca4 as nil")
	}
	b.ObjStart()
	b.PutID("assignGooglePlayTransaction")
	b.Comma()
	b.FieldStart("package_name")
	b.PutString(a.PackageName)
	b.Comma()
	b.FieldStart("store_product_id")
	b.PutString(a.StoreProductID)
	b.Comma()
	b.FieldStart("purchase_token")
	b.PutString(a.PurchaseToken)
	b.Comma()
	b.FieldStart("purpose")
	if a.Purpose == nil {
		return fmt.Errorf("unable to encode assignGooglePlayTransaction#8939bca4: field purpose is nil")
	}
	if err := a.Purpose.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode assignGooglePlayTransaction#8939bca4: field purpose: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AssignGooglePlayTransactionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode assignGooglePlayTransaction#8939bca4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("assignGooglePlayTransaction"); err != nil {
				return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: %w", err)
			}
		case "package_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: field package_name: %w", err)
			}
			a.PackageName = value
		case "store_product_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: field store_product_id: %w", err)
			}
			a.StoreProductID = value
		case "purchase_token":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: field purchase_token: %w", err)
			}
			a.PurchaseToken = value
		case "purpose":
			value, err := DecodeTDLibJSONStorePaymentPurpose(b)
			if err != nil {
				return fmt.Errorf("unable to decode assignGooglePlayTransaction#8939bca4: field purpose: %w", err)
			}
			a.Purpose = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPackageName returns value of PackageName field.
func (a *AssignGooglePlayTransactionRequest) GetPackageName() (value string) {
	if a == nil {
		return
	}
	return a.PackageName
}

// GetStoreProductID returns value of StoreProductID field.
func (a *AssignGooglePlayTransactionRequest) GetStoreProductID() (value string) {
	if a == nil {
		return
	}
	return a.StoreProductID
}

// GetPurchaseToken returns value of PurchaseToken field.
func (a *AssignGooglePlayTransactionRequest) GetPurchaseToken() (value string) {
	if a == nil {
		return
	}
	return a.PurchaseToken
}

// GetPurpose returns value of Purpose field.
func (a *AssignGooglePlayTransactionRequest) GetPurpose() (value StorePaymentPurposeClass) {
	if a == nil {
		return
	}
	return a.Purpose
}

// AssignGooglePlayTransaction invokes method assignGooglePlayTransaction#8939bca4 returning error if any.
func (c *Client) AssignGooglePlayTransaction(ctx context.Context, request *AssignGooglePlayTransactionRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
