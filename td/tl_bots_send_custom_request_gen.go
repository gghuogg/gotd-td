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

// BotsSendCustomRequestRequest represents TL type `bots.sendCustomRequest#aa2769ed`.
type BotsSendCustomRequestRequest struct {
	// CustomMethod field of BotsSendCustomRequestRequest.
	CustomMethod string
	// Params field of BotsSendCustomRequestRequest.
	Params DataJSON
}

// BotsSendCustomRequestRequestTypeID is TL type id of BotsSendCustomRequestRequest.
const BotsSendCustomRequestRequestTypeID = 0xaa2769ed

// Ensuring interfaces in compile-time for BotsSendCustomRequestRequest.
var (
	_ bin.Encoder     = &BotsSendCustomRequestRequest{}
	_ bin.Decoder     = &BotsSendCustomRequestRequest{}
	_ bin.BareEncoder = &BotsSendCustomRequestRequest{}
	_ bin.BareDecoder = &BotsSendCustomRequestRequest{}
)

func (s *BotsSendCustomRequestRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.CustomMethod == "") {
		return false
	}
	if !(s.Params.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *BotsSendCustomRequestRequest) String() string {
	if s == nil {
		return "BotsSendCustomRequestRequest(nil)"
	}
	type Alias BotsSendCustomRequestRequest
	return fmt.Sprintf("BotsSendCustomRequestRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsSendCustomRequestRequest) TypeID() uint32 {
	return BotsSendCustomRequestRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsSendCustomRequestRequest) TypeName() string {
	return "bots.sendCustomRequest"
}

// TypeInfo returns info about TL type.
func (s *BotsSendCustomRequestRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.sendCustomRequest",
		ID:   BotsSendCustomRequestRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CustomMethod",
			SchemaName: "custom_method",
		},
		{
			Name:       "Params",
			SchemaName: "params",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *BotsSendCustomRequestRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode bots.sendCustomRequest#aa2769ed as nil")
	}
	b.PutID(BotsSendCustomRequestRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *BotsSendCustomRequestRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode bots.sendCustomRequest#aa2769ed as nil")
	}
	b.PutString(s.CustomMethod)
	if err := s.Params.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.sendCustomRequest#aa2769ed: field params: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *BotsSendCustomRequestRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode bots.sendCustomRequest#aa2769ed to nil")
	}
	if err := b.ConsumeID(BotsSendCustomRequestRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.sendCustomRequest#aa2769ed: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *BotsSendCustomRequestRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode bots.sendCustomRequest#aa2769ed to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.sendCustomRequest#aa2769ed: field custom_method: %w", err)
		}
		s.CustomMethod = value
	}
	{
		if err := s.Params.Decode(b); err != nil {
			return fmt.Errorf("unable to decode bots.sendCustomRequest#aa2769ed: field params: %w", err)
		}
	}
	return nil
}

// GetCustomMethod returns value of CustomMethod field.
func (s *BotsSendCustomRequestRequest) GetCustomMethod() (value string) {
	if s == nil {
		return
	}
	return s.CustomMethod
}

// GetParams returns value of Params field.
func (s *BotsSendCustomRequestRequest) GetParams() (value DataJSON) {
	if s == nil {
		return
	}
	return s.Params
}

// BotsSendCustomRequest invokes method bots.sendCustomRequest#aa2769ed returning error if any.
func (c *Client) BotsSendCustomRequest(ctx context.Context, request *BotsSendCustomRequestRequest) (*DataJSON, error) {
	var result DataJSON

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
