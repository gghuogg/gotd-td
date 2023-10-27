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

// HelpGetPremiumPromoRequest represents TL type `help.getPremiumPromo#b81b93d4`.
type HelpGetPremiumPromoRequest struct {
}

// HelpGetPremiumPromoRequestTypeID is TL type id of HelpGetPremiumPromoRequest.
const HelpGetPremiumPromoRequestTypeID = 0xb81b93d4

// Ensuring interfaces in compile-time for HelpGetPremiumPromoRequest.
var (
	_ bin.Encoder     = &HelpGetPremiumPromoRequest{}
	_ bin.Decoder     = &HelpGetPremiumPromoRequest{}
	_ bin.BareEncoder = &HelpGetPremiumPromoRequest{}
	_ bin.BareDecoder = &HelpGetPremiumPromoRequest{}
)

func (g *HelpGetPremiumPromoRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetPremiumPromoRequest) String() string {
	if g == nil {
		return "HelpGetPremiumPromoRequest(nil)"
	}
	type Alias HelpGetPremiumPromoRequest
	return fmt.Sprintf("HelpGetPremiumPromoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetPremiumPromoRequest) TypeID() uint32 {
	return HelpGetPremiumPromoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetPremiumPromoRequest) TypeName() string {
	return "help.getPremiumPromo"
}

// TypeInfo returns info about TL type.
func (g *HelpGetPremiumPromoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getPremiumPromo",
		ID:   HelpGetPremiumPromoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetPremiumPromoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getPremiumPromo#b81b93d4 as nil")
	}
	b.PutID(HelpGetPremiumPromoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetPremiumPromoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getPremiumPromo#b81b93d4 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetPremiumPromoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getPremiumPromo#b81b93d4 to nil")
	}
	if err := b.ConsumeID(HelpGetPremiumPromoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getPremiumPromo#b81b93d4: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetPremiumPromoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getPremiumPromo#b81b93d4 to nil")
	}
	return nil
}

// HelpGetPremiumPromo invokes method help.getPremiumPromo#b81b93d4 returning error if any.
func (c *Client) HelpGetPremiumPromo(ctx context.Context) (*HelpPremiumPromo, error) {
	var result HelpPremiumPromo

	request := &HelpGetPremiumPromoRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
