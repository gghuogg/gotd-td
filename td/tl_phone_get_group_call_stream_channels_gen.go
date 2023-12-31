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

// PhoneGetGroupCallStreamChannelsRequest represents TL type `phone.getGroupCallStreamChannels#1ab21940`.
type PhoneGetGroupCallStreamChannelsRequest struct {
	// Call field of PhoneGetGroupCallStreamChannelsRequest.
	Call InputGroupCall
}

// PhoneGetGroupCallStreamChannelsRequestTypeID is TL type id of PhoneGetGroupCallStreamChannelsRequest.
const PhoneGetGroupCallStreamChannelsRequestTypeID = 0x1ab21940

// Ensuring interfaces in compile-time for PhoneGetGroupCallStreamChannelsRequest.
var (
	_ bin.Encoder     = &PhoneGetGroupCallStreamChannelsRequest{}
	_ bin.Decoder     = &PhoneGetGroupCallStreamChannelsRequest{}
	_ bin.BareEncoder = &PhoneGetGroupCallStreamChannelsRequest{}
	_ bin.BareDecoder = &PhoneGetGroupCallStreamChannelsRequest{}
)

func (g *PhoneGetGroupCallStreamChannelsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Call.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGetGroupCallStreamChannelsRequest) String() string {
	if g == nil {
		return "PhoneGetGroupCallStreamChannelsRequest(nil)"
	}
	type Alias PhoneGetGroupCallStreamChannelsRequest
	return fmt.Sprintf("PhoneGetGroupCallStreamChannelsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneGetGroupCallStreamChannelsRequest) TypeID() uint32 {
	return PhoneGetGroupCallStreamChannelsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneGetGroupCallStreamChannelsRequest) TypeName() string {
	return "phone.getGroupCallStreamChannels"
}

// TypeInfo returns info about TL type.
func (g *PhoneGetGroupCallStreamChannelsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.getGroupCallStreamChannels",
		ID:   PhoneGetGroupCallStreamChannelsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhoneGetGroupCallStreamChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getGroupCallStreamChannels#1ab21940 as nil")
	}
	b.PutID(PhoneGetGroupCallStreamChannelsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PhoneGetGroupCallStreamChannelsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getGroupCallStreamChannels#1ab21940 as nil")
	}
	if err := g.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.getGroupCallStreamChannels#1ab21940: field call: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *PhoneGetGroupCallStreamChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getGroupCallStreamChannels#1ab21940 to nil")
	}
	if err := b.ConsumeID(PhoneGetGroupCallStreamChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.getGroupCallStreamChannels#1ab21940: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PhoneGetGroupCallStreamChannelsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getGroupCallStreamChannels#1ab21940 to nil")
	}
	{
		if err := g.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.getGroupCallStreamChannels#1ab21940: field call: %w", err)
		}
	}
	return nil
}

// GetCall returns value of Call field.
func (g *PhoneGetGroupCallStreamChannelsRequest) GetCall() (value InputGroupCall) {
	if g == nil {
		return
	}
	return g.Call
}

// PhoneGetGroupCallStreamChannels invokes method phone.getGroupCallStreamChannels#1ab21940 returning error if any.
func (c *Client) PhoneGetGroupCallStreamChannels(ctx context.Context, call InputGroupCall) (*PhoneGroupCallStreamChannels, error) {
	var result PhoneGroupCallStreamChannels

	request := &PhoneGetGroupCallStreamChannelsRequest{
		Call: call,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
