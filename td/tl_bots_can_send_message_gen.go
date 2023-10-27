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

// BotsCanSendMessageRequest represents TL type `bots.canSendMessage#1359f4e6`.
type BotsCanSendMessageRequest struct {
	// Bot field of BotsCanSendMessageRequest.
	Bot InputUserClass
}

// BotsCanSendMessageRequestTypeID is TL type id of BotsCanSendMessageRequest.
const BotsCanSendMessageRequestTypeID = 0x1359f4e6

// Ensuring interfaces in compile-time for BotsCanSendMessageRequest.
var (
	_ bin.Encoder     = &BotsCanSendMessageRequest{}
	_ bin.Decoder     = &BotsCanSendMessageRequest{}
	_ bin.BareEncoder = &BotsCanSendMessageRequest{}
	_ bin.BareDecoder = &BotsCanSendMessageRequest{}
)

func (c *BotsCanSendMessageRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Bot == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *BotsCanSendMessageRequest) String() string {
	if c == nil {
		return "BotsCanSendMessageRequest(nil)"
	}
	type Alias BotsCanSendMessageRequest
	return fmt.Sprintf("BotsCanSendMessageRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsCanSendMessageRequest) TypeID() uint32 {
	return BotsCanSendMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsCanSendMessageRequest) TypeName() string {
	return "bots.canSendMessage"
}

// TypeInfo returns info about TL type.
func (c *BotsCanSendMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.canSendMessage",
		ID:   BotsCanSendMessageRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *BotsCanSendMessageRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode bots.canSendMessage#1359f4e6 as nil")
	}
	b.PutID(BotsCanSendMessageRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *BotsCanSendMessageRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode bots.canSendMessage#1359f4e6 as nil")
	}
	if c.Bot == nil {
		return fmt.Errorf("unable to encode bots.canSendMessage#1359f4e6: field bot is nil")
	}
	if err := c.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.canSendMessage#1359f4e6: field bot: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *BotsCanSendMessageRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode bots.canSendMessage#1359f4e6 to nil")
	}
	if err := b.ConsumeID(BotsCanSendMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.canSendMessage#1359f4e6: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *BotsCanSendMessageRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode bots.canSendMessage#1359f4e6 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.canSendMessage#1359f4e6: field bot: %w", err)
		}
		c.Bot = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (c *BotsCanSendMessageRequest) GetBot() (value InputUserClass) {
	if c == nil {
		return
	}
	return c.Bot
}

// BotsCanSendMessage invokes method bots.canSendMessage#1359f4e6 returning error if any.
func (c *Client) BotsCanSendMessage(ctx context.Context, bot InputUserClass) (bool, error) {
	var result BoolBox

	request := &BotsCanSendMessageRequest{
		Bot: bot,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
