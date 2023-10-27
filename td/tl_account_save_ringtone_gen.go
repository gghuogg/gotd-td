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

// AccountSaveRingtoneRequest represents TL type `account.saveRingtone#3dea5b03`.
type AccountSaveRingtoneRequest struct {
	// ID field of AccountSaveRingtoneRequest.
	ID InputDocumentClass
	// Unsave field of AccountSaveRingtoneRequest.
	Unsave bool
}

// AccountSaveRingtoneRequestTypeID is TL type id of AccountSaveRingtoneRequest.
const AccountSaveRingtoneRequestTypeID = 0x3dea5b03

// Ensuring interfaces in compile-time for AccountSaveRingtoneRequest.
var (
	_ bin.Encoder     = &AccountSaveRingtoneRequest{}
	_ bin.Decoder     = &AccountSaveRingtoneRequest{}
	_ bin.BareEncoder = &AccountSaveRingtoneRequest{}
	_ bin.BareDecoder = &AccountSaveRingtoneRequest{}
)

func (s *AccountSaveRingtoneRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == nil) {
		return false
	}
	if !(s.Unsave == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSaveRingtoneRequest) String() string {
	if s == nil {
		return "AccountSaveRingtoneRequest(nil)"
	}
	type Alias AccountSaveRingtoneRequest
	return fmt.Sprintf("AccountSaveRingtoneRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSaveRingtoneRequest) TypeID() uint32 {
	return AccountSaveRingtoneRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSaveRingtoneRequest) TypeName() string {
	return "account.saveRingtone"
}

// TypeInfo returns info about TL type.
func (s *AccountSaveRingtoneRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.saveRingtone",
		ID:   AccountSaveRingtoneRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Unsave",
			SchemaName: "unsave",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSaveRingtoneRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveRingtone#3dea5b03 as nil")
	}
	b.PutID(AccountSaveRingtoneRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSaveRingtoneRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveRingtone#3dea5b03 as nil")
	}
	if s.ID == nil {
		return fmt.Errorf("unable to encode account.saveRingtone#3dea5b03: field id is nil")
	}
	if err := s.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveRingtone#3dea5b03: field id: %w", err)
	}
	b.PutBool(s.Unsave)
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSaveRingtoneRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveRingtone#3dea5b03 to nil")
	}
	if err := b.ConsumeID(AccountSaveRingtoneRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveRingtone#3dea5b03: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSaveRingtoneRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveRingtone#3dea5b03 to nil")
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.saveRingtone#3dea5b03: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.saveRingtone#3dea5b03: field unsave: %w", err)
		}
		s.Unsave = value
	}
	return nil
}

// GetID returns value of ID field.
func (s *AccountSaveRingtoneRequest) GetID() (value InputDocumentClass) {
	if s == nil {
		return
	}
	return s.ID
}

// GetUnsave returns value of Unsave field.
func (s *AccountSaveRingtoneRequest) GetUnsave() (value bool) {
	if s == nil {
		return
	}
	return s.Unsave
}

// AccountSaveRingtone invokes method account.saveRingtone#3dea5b03 returning error if any.
func (c *Client) AccountSaveRingtone(ctx context.Context, request *AccountSaveRingtoneRequest) (AccountSavedRingtoneClass, error) {
	var result AccountSavedRingtoneBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SavedRingtone, nil
}
