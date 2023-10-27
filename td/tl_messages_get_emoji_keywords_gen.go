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

// MessagesGetEmojiKeywordsRequest represents TL type `messages.getEmojiKeywords#35a0e062`.
type MessagesGetEmojiKeywordsRequest struct {
	// LangCode field of MessagesGetEmojiKeywordsRequest.
	LangCode string
}

// MessagesGetEmojiKeywordsRequestTypeID is TL type id of MessagesGetEmojiKeywordsRequest.
const MessagesGetEmojiKeywordsRequestTypeID = 0x35a0e062

// Ensuring interfaces in compile-time for MessagesGetEmojiKeywordsRequest.
var (
	_ bin.Encoder     = &MessagesGetEmojiKeywordsRequest{}
	_ bin.Decoder     = &MessagesGetEmojiKeywordsRequest{}
	_ bin.BareEncoder = &MessagesGetEmojiKeywordsRequest{}
	_ bin.BareDecoder = &MessagesGetEmojiKeywordsRequest{}
)

func (g *MessagesGetEmojiKeywordsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetEmojiKeywordsRequest) String() string {
	if g == nil {
		return "MessagesGetEmojiKeywordsRequest(nil)"
	}
	type Alias MessagesGetEmojiKeywordsRequest
	return fmt.Sprintf("MessagesGetEmojiKeywordsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetEmojiKeywordsRequest) TypeID() uint32 {
	return MessagesGetEmojiKeywordsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetEmojiKeywordsRequest) TypeName() string {
	return "messages.getEmojiKeywords"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetEmojiKeywordsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getEmojiKeywords",
		ID:   MessagesGetEmojiKeywordsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiKeywordsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywords#35a0e062 as nil")
	}
	b.PutID(MessagesGetEmojiKeywordsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetEmojiKeywordsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywords#35a0e062 as nil")
	}
	b.PutString(g.LangCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiKeywordsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywords#35a0e062 to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiKeywordsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiKeywords#35a0e062: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetEmojiKeywordsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywords#35a0e062 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywords#35a0e062: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// GetLangCode returns value of LangCode field.
func (g *MessagesGetEmojiKeywordsRequest) GetLangCode() (value string) {
	if g == nil {
		return
	}
	return g.LangCode
}

// MessagesGetEmojiKeywords invokes method messages.getEmojiKeywords#35a0e062 returning error if any.
func (c *Client) MessagesGetEmojiKeywords(ctx context.Context, langcode string) (*EmojiKeywordsDifference, error) {
	var result EmojiKeywordsDifference

	request := &MessagesGetEmojiKeywordsRequest{
		LangCode: langcode,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
