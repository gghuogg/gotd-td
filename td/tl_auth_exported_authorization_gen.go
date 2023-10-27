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

// AuthExportedAuthorization represents TL type `auth.exportedAuthorization#b434e2b8`.
type AuthExportedAuthorization struct {
	// ID field of AuthExportedAuthorization.
	ID int64
	// Bytes field of AuthExportedAuthorization.
	Bytes []byte
}

// AuthExportedAuthorizationTypeID is TL type id of AuthExportedAuthorization.
const AuthExportedAuthorizationTypeID = 0xb434e2b8

// Ensuring interfaces in compile-time for AuthExportedAuthorization.
var (
	_ bin.Encoder     = &AuthExportedAuthorization{}
	_ bin.Decoder     = &AuthExportedAuthorization{}
	_ bin.BareEncoder = &AuthExportedAuthorization{}
	_ bin.BareDecoder = &AuthExportedAuthorization{}
)

func (e *AuthExportedAuthorization) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *AuthExportedAuthorization) String() string {
	if e == nil {
		return "AuthExportedAuthorization(nil)"
	}
	type Alias AuthExportedAuthorization
	return fmt.Sprintf("AuthExportedAuthorization%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthExportedAuthorization) TypeID() uint32 {
	return AuthExportedAuthorizationTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthExportedAuthorization) TypeName() string {
	return "auth.exportedAuthorization"
}

// TypeInfo returns info about TL type.
func (e *AuthExportedAuthorization) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.exportedAuthorization",
		ID:   AuthExportedAuthorizationTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *AuthExportedAuthorization) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportedAuthorization#b434e2b8 as nil")
	}
	b.PutID(AuthExportedAuthorizationTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *AuthExportedAuthorization) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportedAuthorization#b434e2b8 as nil")
	}
	b.PutLong(e.ID)
	b.PutBytes(e.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (e *AuthExportedAuthorization) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportedAuthorization#b434e2b8 to nil")
	}
	if err := b.ConsumeID(AuthExportedAuthorizationTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.exportedAuthorization#b434e2b8: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *AuthExportedAuthorization) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportedAuthorization#b434e2b8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportedAuthorization#b434e2b8: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportedAuthorization#b434e2b8: field bytes: %w", err)
		}
		e.Bytes = value
	}
	return nil
}

// GetID returns value of ID field.
func (e *AuthExportedAuthorization) GetID() (value int64) {
	if e == nil {
		return
	}
	return e.ID
}

// GetBytes returns value of Bytes field.
func (e *AuthExportedAuthorization) GetBytes() (value []byte) {
	if e == nil {
		return
	}
	return e.Bytes
}
