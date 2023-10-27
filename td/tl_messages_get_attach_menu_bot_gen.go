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

// MessagesGetAttachMenuBotRequest represents TL type `messages.getAttachMenuBot#77216192`.
type MessagesGetAttachMenuBotRequest struct {
	// Bot field of MessagesGetAttachMenuBotRequest.
	Bot InputUserClass
}

// MessagesGetAttachMenuBotRequestTypeID is TL type id of MessagesGetAttachMenuBotRequest.
const MessagesGetAttachMenuBotRequestTypeID = 0x77216192

// Ensuring interfaces in compile-time for MessagesGetAttachMenuBotRequest.
var (
	_ bin.Encoder     = &MessagesGetAttachMenuBotRequest{}
	_ bin.Decoder     = &MessagesGetAttachMenuBotRequest{}
	_ bin.BareEncoder = &MessagesGetAttachMenuBotRequest{}
	_ bin.BareDecoder = &MessagesGetAttachMenuBotRequest{}
)

func (g *MessagesGetAttachMenuBotRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Bot == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAttachMenuBotRequest) String() string {
	if g == nil {
		return "MessagesGetAttachMenuBotRequest(nil)"
	}
	type Alias MessagesGetAttachMenuBotRequest
	return fmt.Sprintf("MessagesGetAttachMenuBotRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetAttachMenuBotRequest) TypeID() uint32 {
	return MessagesGetAttachMenuBotRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetAttachMenuBotRequest) TypeName() string {
	return "messages.getAttachMenuBot"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetAttachMenuBotRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getAttachMenuBot",
		ID:   MessagesGetAttachMenuBotRequestTypeID,
	}
	if g == nil {
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
func (g *MessagesGetAttachMenuBotRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAttachMenuBot#77216192 as nil")
	}
	b.PutID(MessagesGetAttachMenuBotRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetAttachMenuBotRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAttachMenuBot#77216192 as nil")
	}
	if g.Bot == nil {
		return fmt.Errorf("unable to encode messages.getAttachMenuBot#77216192: field bot is nil")
	}
	if err := g.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getAttachMenuBot#77216192: field bot: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetAttachMenuBotRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAttachMenuBot#77216192 to nil")
	}
	if err := b.ConsumeID(MessagesGetAttachMenuBotRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAttachMenuBot#77216192: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetAttachMenuBotRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAttachMenuBot#77216192 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAttachMenuBot#77216192: field bot: %w", err)
		}
		g.Bot = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (g *MessagesGetAttachMenuBotRequest) GetBot() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.Bot
}

// MessagesGetAttachMenuBot invokes method messages.getAttachMenuBot#77216192 returning error if any.
func (c *Client) MessagesGetAttachMenuBot(ctx context.Context, bot InputUserClass) (*AttachMenuBotsBot, error) {
	var result AttachMenuBotsBot

	request := &MessagesGetAttachMenuBotRequest{
		Bot: bot,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
