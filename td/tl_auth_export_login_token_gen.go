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

// AuthExportLoginTokenRequest represents TL type `auth.exportLoginToken#b7e085fe`.
type AuthExportLoginTokenRequest struct {
	// APIID field of AuthExportLoginTokenRequest.
	APIID int
	// APIHash field of AuthExportLoginTokenRequest.
	APIHash string
	// ExceptIDs field of AuthExportLoginTokenRequest.
	ExceptIDs []int64
}

// AuthExportLoginTokenRequestTypeID is TL type id of AuthExportLoginTokenRequest.
const AuthExportLoginTokenRequestTypeID = 0xb7e085fe

// Ensuring interfaces in compile-time for AuthExportLoginTokenRequest.
var (
	_ bin.Encoder     = &AuthExportLoginTokenRequest{}
	_ bin.Decoder     = &AuthExportLoginTokenRequest{}
	_ bin.BareEncoder = &AuthExportLoginTokenRequest{}
	_ bin.BareDecoder = &AuthExportLoginTokenRequest{}
)

func (e *AuthExportLoginTokenRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.APIID == 0) {
		return false
	}
	if !(e.APIHash == "") {
		return false
	}
	if !(e.ExceptIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *AuthExportLoginTokenRequest) String() string {
	if e == nil {
		return "AuthExportLoginTokenRequest(nil)"
	}
	type Alias AuthExportLoginTokenRequest
	return fmt.Sprintf("AuthExportLoginTokenRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthExportLoginTokenRequest) TypeID() uint32 {
	return AuthExportLoginTokenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthExportLoginTokenRequest) TypeName() string {
	return "auth.exportLoginToken"
}

// TypeInfo returns info about TL type.
func (e *AuthExportLoginTokenRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.exportLoginToken",
		ID:   AuthExportLoginTokenRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "APIID",
			SchemaName: "api_id",
		},
		{
			Name:       "APIHash",
			SchemaName: "api_hash",
		},
		{
			Name:       "ExceptIDs",
			SchemaName: "except_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *AuthExportLoginTokenRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportLoginToken#b7e085fe as nil")
	}
	b.PutID(AuthExportLoginTokenRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *AuthExportLoginTokenRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportLoginToken#b7e085fe as nil")
	}
	b.PutInt(e.APIID)
	b.PutString(e.APIHash)
	b.PutVectorHeader(len(e.ExceptIDs))
	for _, v := range e.ExceptIDs {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *AuthExportLoginTokenRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportLoginToken#b7e085fe to nil")
	}
	if err := b.ConsumeID(AuthExportLoginTokenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.exportLoginToken#b7e085fe: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *AuthExportLoginTokenRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportLoginToken#b7e085fe to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportLoginToken#b7e085fe: field api_id: %w", err)
		}
		e.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportLoginToken#b7e085fe: field api_hash: %w", err)
		}
		e.APIHash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportLoginToken#b7e085fe: field except_ids: %w", err)
		}

		if headerLen > 0 {
			e.ExceptIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode auth.exportLoginToken#b7e085fe: field except_ids: %w", err)
			}
			e.ExceptIDs = append(e.ExceptIDs, value)
		}
	}
	return nil
}

// GetAPIID returns value of APIID field.
func (e *AuthExportLoginTokenRequest) GetAPIID() (value int) {
	if e == nil {
		return
	}
	return e.APIID
}

// GetAPIHash returns value of APIHash field.
func (e *AuthExportLoginTokenRequest) GetAPIHash() (value string) {
	if e == nil {
		return
	}
	return e.APIHash
}

// GetExceptIDs returns value of ExceptIDs field.
func (e *AuthExportLoginTokenRequest) GetExceptIDs() (value []int64) {
	if e == nil {
		return
	}
	return e.ExceptIDs
}

// AuthExportLoginToken invokes method auth.exportLoginToken#b7e085fe returning error if any.
func (c *Client) AuthExportLoginToken(ctx context.Context, request *AuthExportLoginTokenRequest) (AuthLoginTokenClass, error) {
	var result AuthLoginTokenBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.LoginToken, nil
}
