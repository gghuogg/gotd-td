// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// BotsToggleUsernameRequest represents TL type `bots.toggleUsername#53ca973`.
// Activate or deactivate a purchased fragment.com¹ username associated to a bot we own.
//
// Links:
//  1. https://fragment.com
//
// See https://core.telegram.org/method/bots.toggleUsername for reference.
type BotsToggleUsernameRequest struct {
	// The bot
	Bot InputUserClass
	// Username
	Username string
	// Whether to activate or deactivate it
	Active bool
}

// BotsToggleUsernameRequestTypeID is TL type id of BotsToggleUsernameRequest.
const BotsToggleUsernameRequestTypeID = 0x53ca973

// Ensuring interfaces in compile-time for BotsToggleUsernameRequest.
var (
	_ bin.Encoder     = &BotsToggleUsernameRequest{}
	_ bin.Decoder     = &BotsToggleUsernameRequest{}
	_ bin.BareEncoder = &BotsToggleUsernameRequest{}
	_ bin.BareDecoder = &BotsToggleUsernameRequest{}
)

func (t *BotsToggleUsernameRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Bot == nil) {
		return false
	}
	if !(t.Username == "") {
		return false
	}
	if !(t.Active == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *BotsToggleUsernameRequest) String() string {
	if t == nil {
		return "BotsToggleUsernameRequest(nil)"
	}
	type Alias BotsToggleUsernameRequest
	return fmt.Sprintf("BotsToggleUsernameRequest%+v", Alias(*t))
}

// FillFrom fills BotsToggleUsernameRequest from given interface.
func (t *BotsToggleUsernameRequest) FillFrom(from interface {
	GetBot() (value InputUserClass)
	GetUsername() (value string)
	GetActive() (value bool)
}) {
	t.Bot = from.GetBot()
	t.Username = from.GetUsername()
	t.Active = from.GetActive()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsToggleUsernameRequest) TypeID() uint32 {
	return BotsToggleUsernameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsToggleUsernameRequest) TypeName() string {
	return "bots.toggleUsername"
}

// TypeInfo returns info about TL type.
func (t *BotsToggleUsernameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.toggleUsername",
		ID:   BotsToggleUsernameRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Active",
			SchemaName: "active",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *BotsToggleUsernameRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode bots.toggleUsername#53ca973 as nil")
	}
	b.PutID(BotsToggleUsernameRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *BotsToggleUsernameRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode bots.toggleUsername#53ca973 as nil")
	}
	if t.Bot == nil {
		return fmt.Errorf("unable to encode bots.toggleUsername#53ca973: field bot is nil")
	}
	if err := t.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.toggleUsername#53ca973: field bot: %w", err)
	}
	b.PutString(t.Username)
	b.PutBool(t.Active)
	return nil
}

// Decode implements bin.Decoder.
func (t *BotsToggleUsernameRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode bots.toggleUsername#53ca973 to nil")
	}
	if err := b.ConsumeID(BotsToggleUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.toggleUsername#53ca973: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *BotsToggleUsernameRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode bots.toggleUsername#53ca973 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.toggleUsername#53ca973: field bot: %w", err)
		}
		t.Bot = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.toggleUsername#53ca973: field username: %w", err)
		}
		t.Username = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode bots.toggleUsername#53ca973: field active: %w", err)
		}
		t.Active = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (t *BotsToggleUsernameRequest) GetBot() (value InputUserClass) {
	if t == nil {
		return
	}
	return t.Bot
}

// GetUsername returns value of Username field.
func (t *BotsToggleUsernameRequest) GetUsername() (value string) {
	if t == nil {
		return
	}
	return t.Username
}

// GetActive returns value of Active field.
func (t *BotsToggleUsernameRequest) GetActive() (value bool) {
	if t == nil {
		return
	}
	return t.Active
}

// BotsToggleUsername invokes method bots.toggleUsername#53ca973 returning error if any.
// Activate or deactivate a purchased fragment.com¹ username associated to a bot we own.
//
// Links:
//  1. https://fragment.com
//
// Possible errors:
//
//	400 BOT_INVALID: This is not a valid bot.
//
// See https://core.telegram.org/method/bots.toggleUsername for reference.
// Can be used by bots.
func (c *Client) BotsToggleUsername(ctx context.Context, request *BotsToggleUsernameRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
