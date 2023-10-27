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

// BoostChatRequest represents TL type `boostChat#4e314c2d`.
type BoostChatRequest struct {
	// Identifier of the chat
	ChatID int64
}

// BoostChatRequestTypeID is TL type id of BoostChatRequest.
const BoostChatRequestTypeID = 0x4e314c2d

// Ensuring interfaces in compile-time for BoostChatRequest.
var (
	_ bin.Encoder     = &BoostChatRequest{}
	_ bin.Decoder     = &BoostChatRequest{}
	_ bin.BareEncoder = &BoostChatRequest{}
	_ bin.BareDecoder = &BoostChatRequest{}
)

func (b *BoostChatRequest) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BoostChatRequest) String() string {
	if b == nil {
		return "BoostChatRequest(nil)"
	}
	type Alias BoostChatRequest
	return fmt.Sprintf("BoostChatRequest%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BoostChatRequest) TypeID() uint32 {
	return BoostChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BoostChatRequest) TypeName() string {
	return "boostChat"
}

// TypeInfo returns info about TL type.
func (b *BoostChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "boostChat",
		ID:   BoostChatRequestTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BoostChatRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boostChat#4e314c2d as nil")
	}
	buf.PutID(BoostChatRequestTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BoostChatRequest) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boostChat#4e314c2d as nil")
	}
	buf.PutInt53(b.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BoostChatRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boostChat#4e314c2d to nil")
	}
	if err := buf.ConsumeID(BoostChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode boostChat#4e314c2d: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BoostChatRequest) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boostChat#4e314c2d to nil")
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode boostChat#4e314c2d: field chat_id: %w", err)
		}
		b.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BoostChatRequest) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode boostChat#4e314c2d as nil")
	}
	buf.ObjStart()
	buf.PutID("boostChat")
	buf.Comma()
	buf.FieldStart("chat_id")
	buf.PutInt53(b.ChatID)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BoostChatRequest) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode boostChat#4e314c2d to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("boostChat"); err != nil {
				return fmt.Errorf("unable to decode boostChat#4e314c2d: %w", err)
			}
		case "chat_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode boostChat#4e314c2d: field chat_id: %w", err)
			}
			b.ChatID = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (b *BoostChatRequest) GetChatID() (value int64) {
	if b == nil {
		return
	}
	return b.ChatID
}

// BoostChat invokes method boostChat#4e314c2d returning error if any.
func (c *Client) BoostChat(ctx context.Context, chatid int64) error {
	var ok Ok

	request := &BoostChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
