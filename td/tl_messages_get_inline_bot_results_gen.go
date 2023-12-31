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

// MessagesGetInlineBotResultsRequest represents TL type `messages.getInlineBotResults#514e999d`.
type MessagesGetInlineBotResultsRequest struct {
	// Flags field of MessagesGetInlineBotResultsRequest.
	Flags bin.Fields
	// Bot field of MessagesGetInlineBotResultsRequest.
	Bot InputUserClass
	// Peer field of MessagesGetInlineBotResultsRequest.
	Peer InputPeerClass
	// GeoPoint field of MessagesGetInlineBotResultsRequest.
	//
	// Use SetGeoPoint and GetGeoPoint helpers.
	GeoPoint InputGeoPointClass
	// Query field of MessagesGetInlineBotResultsRequest.
	Query string
	// Offset field of MessagesGetInlineBotResultsRequest.
	Offset string
}

// MessagesGetInlineBotResultsRequestTypeID is TL type id of MessagesGetInlineBotResultsRequest.
const MessagesGetInlineBotResultsRequestTypeID = 0x514e999d

// Ensuring interfaces in compile-time for MessagesGetInlineBotResultsRequest.
var (
	_ bin.Encoder     = &MessagesGetInlineBotResultsRequest{}
	_ bin.Decoder     = &MessagesGetInlineBotResultsRequest{}
	_ bin.BareEncoder = &MessagesGetInlineBotResultsRequest{}
	_ bin.BareDecoder = &MessagesGetInlineBotResultsRequest{}
)

func (g *MessagesGetInlineBotResultsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Bot == nil) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.GeoPoint == nil) {
		return false
	}
	if !(g.Query == "") {
		return false
	}
	if !(g.Offset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetInlineBotResultsRequest) String() string {
	if g == nil {
		return "MessagesGetInlineBotResultsRequest(nil)"
	}
	type Alias MessagesGetInlineBotResultsRequest
	return fmt.Sprintf("MessagesGetInlineBotResultsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetInlineBotResultsRequest) TypeID() uint32 {
	return MessagesGetInlineBotResultsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetInlineBotResultsRequest) TypeName() string {
	return "messages.getInlineBotResults"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetInlineBotResultsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getInlineBotResults",
		ID:   MessagesGetInlineBotResultsRequestTypeID,
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
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "GeoPoint",
			SchemaName: "geo_point",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *MessagesGetInlineBotResultsRequest) SetFlags() {
	if !(g.GeoPoint == nil) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *MessagesGetInlineBotResultsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getInlineBotResults#514e999d as nil")
	}
	b.PutID(MessagesGetInlineBotResultsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetInlineBotResultsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getInlineBotResults#514e999d as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getInlineBotResults#514e999d: field flags: %w", err)
	}
	if g.Bot == nil {
		return fmt.Errorf("unable to encode messages.getInlineBotResults#514e999d: field bot is nil")
	}
	if err := g.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getInlineBotResults#514e999d: field bot: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getInlineBotResults#514e999d: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getInlineBotResults#514e999d: field peer: %w", err)
	}
	if g.Flags.Has(0) {
		if g.GeoPoint == nil {
			return fmt.Errorf("unable to encode messages.getInlineBotResults#514e999d: field geo_point is nil")
		}
		if err := g.GeoPoint.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getInlineBotResults#514e999d: field geo_point: %w", err)
		}
	}
	b.PutString(g.Query)
	b.PutString(g.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetInlineBotResultsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getInlineBotResults#514e999d to nil")
	}
	if err := b.ConsumeID(MessagesGetInlineBotResultsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getInlineBotResults#514e999d: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetInlineBotResultsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getInlineBotResults#514e999d to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getInlineBotResults#514e999d: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getInlineBotResults#514e999d: field bot: %w", err)
		}
		g.Bot = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getInlineBotResults#514e999d: field peer: %w", err)
		}
		g.Peer = value
	}
	if g.Flags.Has(0) {
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getInlineBotResults#514e999d: field geo_point: %w", err)
		}
		g.GeoPoint = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getInlineBotResults#514e999d: field query: %w", err)
		}
		g.Query = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getInlineBotResults#514e999d: field offset: %w", err)
		}
		g.Offset = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (g *MessagesGetInlineBotResultsRequest) GetBot() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.Bot
}

// GetPeer returns value of Peer field.
func (g *MessagesGetInlineBotResultsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// SetGeoPoint sets value of GeoPoint conditional field.
func (g *MessagesGetInlineBotResultsRequest) SetGeoPoint(value InputGeoPointClass) {
	g.Flags.Set(0)
	g.GeoPoint = value
}

// GetGeoPoint returns value of GeoPoint conditional field and
// boolean which is true if field was set.
func (g *MessagesGetInlineBotResultsRequest) GetGeoPoint() (value InputGeoPointClass, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.GeoPoint, true
}

// GetQuery returns value of Query field.
func (g *MessagesGetInlineBotResultsRequest) GetQuery() (value string) {
	if g == nil {
		return
	}
	return g.Query
}

// GetOffset returns value of Offset field.
func (g *MessagesGetInlineBotResultsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// MessagesGetInlineBotResults invokes method messages.getInlineBotResults#514e999d returning error if any.
func (c *Client) MessagesGetInlineBotResults(ctx context.Context, request *MessagesGetInlineBotResultsRequest) (*MessagesBotResults, error) {
	var result MessagesBotResults

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
