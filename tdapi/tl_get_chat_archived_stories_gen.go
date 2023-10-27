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

// GetChatArchivedStoriesRequest represents TL type `getChatArchivedStories#af1e9488`.
type GetChatArchivedStoriesRequest struct {
	// Chat identifier
	ChatID int64
	// Identifier of the story starting from which stories must be returned; use 0 to get
	// results from the last story
	FromStoryID int32
	// The maximum number of stories to be returned
	Limit int32
}

// GetChatArchivedStoriesRequestTypeID is TL type id of GetChatArchivedStoriesRequest.
const GetChatArchivedStoriesRequestTypeID = 0xaf1e9488

// Ensuring interfaces in compile-time for GetChatArchivedStoriesRequest.
var (
	_ bin.Encoder     = &GetChatArchivedStoriesRequest{}
	_ bin.Decoder     = &GetChatArchivedStoriesRequest{}
	_ bin.BareEncoder = &GetChatArchivedStoriesRequest{}
	_ bin.BareDecoder = &GetChatArchivedStoriesRequest{}
)

func (g *GetChatArchivedStoriesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.FromStoryID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatArchivedStoriesRequest) String() string {
	if g == nil {
		return "GetChatArchivedStoriesRequest(nil)"
	}
	type Alias GetChatArchivedStoriesRequest
	return fmt.Sprintf("GetChatArchivedStoriesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatArchivedStoriesRequest) TypeID() uint32 {
	return GetChatArchivedStoriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatArchivedStoriesRequest) TypeName() string {
	return "getChatArchivedStories"
}

// TypeInfo returns info about TL type.
func (g *GetChatArchivedStoriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatArchivedStories",
		ID:   GetChatArchivedStoriesRequestTypeID,
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
		{
			Name:       "FromStoryID",
			SchemaName: "from_story_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatArchivedStoriesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatArchivedStories#af1e9488 as nil")
	}
	b.PutID(GetChatArchivedStoriesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatArchivedStoriesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatArchivedStories#af1e9488 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt32(g.FromStoryID)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatArchivedStoriesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatArchivedStories#af1e9488 to nil")
	}
	if err := b.ConsumeID(GetChatArchivedStoriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatArchivedStories#af1e9488: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatArchivedStoriesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatArchivedStories#af1e9488 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatArchivedStories#af1e9488: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatArchivedStories#af1e9488: field from_story_id: %w", err)
		}
		g.FromStoryID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatArchivedStories#af1e9488: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatArchivedStoriesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatArchivedStories#af1e9488 as nil")
	}
	b.ObjStart()
	b.PutID("getChatArchivedStories")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("from_story_id")
	b.PutInt32(g.FromStoryID)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatArchivedStoriesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatArchivedStories#af1e9488 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatArchivedStories"); err != nil {
				return fmt.Errorf("unable to decode getChatArchivedStories#af1e9488: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatArchivedStories#af1e9488: field chat_id: %w", err)
			}
			g.ChatID = value
		case "from_story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatArchivedStories#af1e9488: field from_story_id: %w", err)
			}
			g.FromStoryID = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatArchivedStories#af1e9488: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatArchivedStoriesRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetFromStoryID returns value of FromStoryID field.
func (g *GetChatArchivedStoriesRequest) GetFromStoryID() (value int32) {
	if g == nil {
		return
	}
	return g.FromStoryID
}

// GetLimit returns value of Limit field.
func (g *GetChatArchivedStoriesRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetChatArchivedStories invokes method getChatArchivedStories#af1e9488 returning error if any.
func (c *Client) GetChatArchivedStories(ctx context.Context, request *GetChatArchivedStoriesRequest) (*Stories, error) {
	var result Stories

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
