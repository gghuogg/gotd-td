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

// GetChatAvailableMessageSendersRequest represents TL type `getChatAvailableMessageSenders#450fe92b`.
type GetChatAvailableMessageSendersRequest struct {
	// Chat identifier
	ChatID int64
}

// GetChatAvailableMessageSendersRequestTypeID is TL type id of GetChatAvailableMessageSendersRequest.
const GetChatAvailableMessageSendersRequestTypeID = 0x450fe92b

// Ensuring interfaces in compile-time for GetChatAvailableMessageSendersRequest.
var (
	_ bin.Encoder     = &GetChatAvailableMessageSendersRequest{}
	_ bin.Decoder     = &GetChatAvailableMessageSendersRequest{}
	_ bin.BareEncoder = &GetChatAvailableMessageSendersRequest{}
	_ bin.BareDecoder = &GetChatAvailableMessageSendersRequest{}
)

func (g *GetChatAvailableMessageSendersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatAvailableMessageSendersRequest) String() string {
	if g == nil {
		return "GetChatAvailableMessageSendersRequest(nil)"
	}
	type Alias GetChatAvailableMessageSendersRequest
	return fmt.Sprintf("GetChatAvailableMessageSendersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatAvailableMessageSendersRequest) TypeID() uint32 {
	return GetChatAvailableMessageSendersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatAvailableMessageSendersRequest) TypeName() string {
	return "getChatAvailableMessageSenders"
}

// TypeInfo returns info about TL type.
func (g *GetChatAvailableMessageSendersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatAvailableMessageSenders",
		ID:   GetChatAvailableMessageSendersRequestTypeID,
	}
	if g == nil {
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
func (g *GetChatAvailableMessageSendersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatAvailableMessageSenders#450fe92b as nil")
	}
	b.PutID(GetChatAvailableMessageSendersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatAvailableMessageSendersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatAvailableMessageSenders#450fe92b as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatAvailableMessageSendersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatAvailableMessageSenders#450fe92b to nil")
	}
	if err := b.ConsumeID(GetChatAvailableMessageSendersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatAvailableMessageSenders#450fe92b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatAvailableMessageSendersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatAvailableMessageSenders#450fe92b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatAvailableMessageSenders#450fe92b: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatAvailableMessageSendersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatAvailableMessageSenders#450fe92b as nil")
	}
	b.ObjStart()
	b.PutID("getChatAvailableMessageSenders")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatAvailableMessageSendersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatAvailableMessageSenders#450fe92b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatAvailableMessageSenders"); err != nil {
				return fmt.Errorf("unable to decode getChatAvailableMessageSenders#450fe92b: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatAvailableMessageSenders#450fe92b: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatAvailableMessageSendersRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetChatAvailableMessageSenders invokes method getChatAvailableMessageSenders#450fe92b returning error if any.
func (c *Client) GetChatAvailableMessageSenders(ctx context.Context, chatid int64) (*ChatMessageSenders, error) {
	var result ChatMessageSenders

	request := &GetChatAvailableMessageSendersRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
