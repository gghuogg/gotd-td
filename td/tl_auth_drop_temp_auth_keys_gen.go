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

// AuthDropTempAuthKeysRequest represents TL type `auth.dropTempAuthKeys#8e48a188`.
type AuthDropTempAuthKeysRequest struct {
	// ExceptAuthKeys field of AuthDropTempAuthKeysRequest.
	ExceptAuthKeys []int64
}

// AuthDropTempAuthKeysRequestTypeID is TL type id of AuthDropTempAuthKeysRequest.
const AuthDropTempAuthKeysRequestTypeID = 0x8e48a188

// Ensuring interfaces in compile-time for AuthDropTempAuthKeysRequest.
var (
	_ bin.Encoder     = &AuthDropTempAuthKeysRequest{}
	_ bin.Decoder     = &AuthDropTempAuthKeysRequest{}
	_ bin.BareEncoder = &AuthDropTempAuthKeysRequest{}
	_ bin.BareDecoder = &AuthDropTempAuthKeysRequest{}
)

func (d *AuthDropTempAuthKeysRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ExceptAuthKeys == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *AuthDropTempAuthKeysRequest) String() string {
	if d == nil {
		return "AuthDropTempAuthKeysRequest(nil)"
	}
	type Alias AuthDropTempAuthKeysRequest
	return fmt.Sprintf("AuthDropTempAuthKeysRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthDropTempAuthKeysRequest) TypeID() uint32 {
	return AuthDropTempAuthKeysRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthDropTempAuthKeysRequest) TypeName() string {
	return "auth.dropTempAuthKeys"
}

// TypeInfo returns info about TL type.
func (d *AuthDropTempAuthKeysRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.dropTempAuthKeys",
		ID:   AuthDropTempAuthKeysRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ExceptAuthKeys",
			SchemaName: "except_auth_keys",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *AuthDropTempAuthKeysRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode auth.dropTempAuthKeys#8e48a188 as nil")
	}
	b.PutID(AuthDropTempAuthKeysRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *AuthDropTempAuthKeysRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode auth.dropTempAuthKeys#8e48a188 as nil")
	}
	b.PutVectorHeader(len(d.ExceptAuthKeys))
	for _, v := range d.ExceptAuthKeys {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *AuthDropTempAuthKeysRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode auth.dropTempAuthKeys#8e48a188 to nil")
	}
	if err := b.ConsumeID(AuthDropTempAuthKeysRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.dropTempAuthKeys#8e48a188: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *AuthDropTempAuthKeysRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode auth.dropTempAuthKeys#8e48a188 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode auth.dropTempAuthKeys#8e48a188: field except_auth_keys: %w", err)
		}

		if headerLen > 0 {
			d.ExceptAuthKeys = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode auth.dropTempAuthKeys#8e48a188: field except_auth_keys: %w", err)
			}
			d.ExceptAuthKeys = append(d.ExceptAuthKeys, value)
		}
	}
	return nil
}

// GetExceptAuthKeys returns value of ExceptAuthKeys field.
func (d *AuthDropTempAuthKeysRequest) GetExceptAuthKeys() (value []int64) {
	if d == nil {
		return
	}
	return d.ExceptAuthKeys
}

// AuthDropTempAuthKeys invokes method auth.dropTempAuthKeys#8e48a188 returning error if any.
func (c *Client) AuthDropTempAuthKeys(ctx context.Context, exceptauthkeys []int64) (bool, error) {
	var result BoolBox

	request := &AuthDropTempAuthKeysRequest{
		ExceptAuthKeys: exceptauthkeys,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
