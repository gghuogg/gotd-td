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

// HelpGetTermsOfServiceUpdateRequest represents TL type `help.getTermsOfServiceUpdate#2ca51fd1`.
type HelpGetTermsOfServiceUpdateRequest struct {
}

// HelpGetTermsOfServiceUpdateRequestTypeID is TL type id of HelpGetTermsOfServiceUpdateRequest.
const HelpGetTermsOfServiceUpdateRequestTypeID = 0x2ca51fd1

// Ensuring interfaces in compile-time for HelpGetTermsOfServiceUpdateRequest.
var (
	_ bin.Encoder     = &HelpGetTermsOfServiceUpdateRequest{}
	_ bin.Decoder     = &HelpGetTermsOfServiceUpdateRequest{}
	_ bin.BareEncoder = &HelpGetTermsOfServiceUpdateRequest{}
	_ bin.BareDecoder = &HelpGetTermsOfServiceUpdateRequest{}
)

func (g *HelpGetTermsOfServiceUpdateRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetTermsOfServiceUpdateRequest) String() string {
	if g == nil {
		return "HelpGetTermsOfServiceUpdateRequest(nil)"
	}
	type Alias HelpGetTermsOfServiceUpdateRequest
	return fmt.Sprintf("HelpGetTermsOfServiceUpdateRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetTermsOfServiceUpdateRequest) TypeID() uint32 {
	return HelpGetTermsOfServiceUpdateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetTermsOfServiceUpdateRequest) TypeName() string {
	return "help.getTermsOfServiceUpdate"
}

// TypeInfo returns info about TL type.
func (g *HelpGetTermsOfServiceUpdateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getTermsOfServiceUpdate",
		ID:   HelpGetTermsOfServiceUpdateRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetTermsOfServiceUpdateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getTermsOfServiceUpdate#2ca51fd1 as nil")
	}
	b.PutID(HelpGetTermsOfServiceUpdateRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetTermsOfServiceUpdateRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getTermsOfServiceUpdate#2ca51fd1 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetTermsOfServiceUpdateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getTermsOfServiceUpdate#2ca51fd1 to nil")
	}
	if err := b.ConsumeID(HelpGetTermsOfServiceUpdateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getTermsOfServiceUpdate#2ca51fd1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetTermsOfServiceUpdateRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getTermsOfServiceUpdate#2ca51fd1 to nil")
	}
	return nil
}

// HelpGetTermsOfServiceUpdate invokes method help.getTermsOfServiceUpdate#2ca51fd1 returning error if any.
func (c *Client) HelpGetTermsOfServiceUpdate(ctx context.Context) (HelpTermsOfServiceUpdateClass, error) {
	var result HelpTermsOfServiceUpdateBox

	request := &HelpGetTermsOfServiceUpdateRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.TermsOfServiceUpdate, nil
}
