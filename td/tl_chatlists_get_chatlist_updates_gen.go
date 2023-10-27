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

// ChatlistsGetChatlistUpdatesRequest represents TL type `chatlists.getChatlistUpdates#89419521`.
type ChatlistsGetChatlistUpdatesRequest struct {
	// Chatlist field of ChatlistsGetChatlistUpdatesRequest.
	Chatlist InputChatlistDialogFilter
}

// ChatlistsGetChatlistUpdatesRequestTypeID is TL type id of ChatlistsGetChatlistUpdatesRequest.
const ChatlistsGetChatlistUpdatesRequestTypeID = 0x89419521

// Ensuring interfaces in compile-time for ChatlistsGetChatlistUpdatesRequest.
var (
	_ bin.Encoder     = &ChatlistsGetChatlistUpdatesRequest{}
	_ bin.Decoder     = &ChatlistsGetChatlistUpdatesRequest{}
	_ bin.BareEncoder = &ChatlistsGetChatlistUpdatesRequest{}
	_ bin.BareDecoder = &ChatlistsGetChatlistUpdatesRequest{}
)

func (g *ChatlistsGetChatlistUpdatesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Chatlist.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChatlistsGetChatlistUpdatesRequest) String() string {
	if g == nil {
		return "ChatlistsGetChatlistUpdatesRequest(nil)"
	}
	type Alias ChatlistsGetChatlistUpdatesRequest
	return fmt.Sprintf("ChatlistsGetChatlistUpdatesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsGetChatlistUpdatesRequest) TypeID() uint32 {
	return ChatlistsGetChatlistUpdatesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsGetChatlistUpdatesRequest) TypeName() string {
	return "chatlists.getChatlistUpdates"
}

// TypeInfo returns info about TL type.
func (g *ChatlistsGetChatlistUpdatesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.getChatlistUpdates",
		ID:   ChatlistsGetChatlistUpdatesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chatlist",
			SchemaName: "chatlist",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChatlistsGetChatlistUpdatesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode chatlists.getChatlistUpdates#89419521 as nil")
	}
	b.PutID(ChatlistsGetChatlistUpdatesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChatlistsGetChatlistUpdatesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode chatlists.getChatlistUpdates#89419521 as nil")
	}
	if err := g.Chatlist.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatlists.getChatlistUpdates#89419521: field chatlist: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChatlistsGetChatlistUpdatesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode chatlists.getChatlistUpdates#89419521 to nil")
	}
	if err := b.ConsumeID(ChatlistsGetChatlistUpdatesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.getChatlistUpdates#89419521: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChatlistsGetChatlistUpdatesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode chatlists.getChatlistUpdates#89419521 to nil")
	}
	{
		if err := g.Chatlist.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatlists.getChatlistUpdates#89419521: field chatlist: %w", err)
		}
	}
	return nil
}

// GetChatlist returns value of Chatlist field.
func (g *ChatlistsGetChatlistUpdatesRequest) GetChatlist() (value InputChatlistDialogFilter) {
	if g == nil {
		return
	}
	return g.Chatlist
}

// ChatlistsGetChatlistUpdates invokes method chatlists.getChatlistUpdates#89419521 returning error if any.
func (c *Client) ChatlistsGetChatlistUpdates(ctx context.Context, chatlist InputChatlistDialogFilter) (*ChatlistsChatlistUpdates, error) {
	var result ChatlistsChatlistUpdates

	request := &ChatlistsGetChatlistUpdatesRequest{
		Chatlist: chatlist,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
