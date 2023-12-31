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

// HelpGetAppChangelogRequest represents TL type `help.getAppChangelog#9010ef6f`.
type HelpGetAppChangelogRequest struct {
	// PrevAppVersion field of HelpGetAppChangelogRequest.
	PrevAppVersion string
}

// HelpGetAppChangelogRequestTypeID is TL type id of HelpGetAppChangelogRequest.
const HelpGetAppChangelogRequestTypeID = 0x9010ef6f

// Ensuring interfaces in compile-time for HelpGetAppChangelogRequest.
var (
	_ bin.Encoder     = &HelpGetAppChangelogRequest{}
	_ bin.Decoder     = &HelpGetAppChangelogRequest{}
	_ bin.BareEncoder = &HelpGetAppChangelogRequest{}
	_ bin.BareDecoder = &HelpGetAppChangelogRequest{}
)

func (g *HelpGetAppChangelogRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.PrevAppVersion == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetAppChangelogRequest) String() string {
	if g == nil {
		return "HelpGetAppChangelogRequest(nil)"
	}
	type Alias HelpGetAppChangelogRequest
	return fmt.Sprintf("HelpGetAppChangelogRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetAppChangelogRequest) TypeID() uint32 {
	return HelpGetAppChangelogRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetAppChangelogRequest) TypeName() string {
	return "help.getAppChangelog"
}

// TypeInfo returns info about TL type.
func (g *HelpGetAppChangelogRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getAppChangelog",
		ID:   HelpGetAppChangelogRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PrevAppVersion",
			SchemaName: "prev_app_version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetAppChangelogRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getAppChangelog#9010ef6f as nil")
	}
	b.PutID(HelpGetAppChangelogRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetAppChangelogRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getAppChangelog#9010ef6f as nil")
	}
	b.PutString(g.PrevAppVersion)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetAppChangelogRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getAppChangelog#9010ef6f to nil")
	}
	if err := b.ConsumeID(HelpGetAppChangelogRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getAppChangelog#9010ef6f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetAppChangelogRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getAppChangelog#9010ef6f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getAppChangelog#9010ef6f: field prev_app_version: %w", err)
		}
		g.PrevAppVersion = value
	}
	return nil
}

// GetPrevAppVersion returns value of PrevAppVersion field.
func (g *HelpGetAppChangelogRequest) GetPrevAppVersion() (value string) {
	if g == nil {
		return
	}
	return g.PrevAppVersion
}

// HelpGetAppChangelog invokes method help.getAppChangelog#9010ef6f returning error if any.
func (c *Client) HelpGetAppChangelog(ctx context.Context, prevappversion string) (UpdatesClass, error) {
	var result UpdatesBox

	request := &HelpGetAppChangelogRequest{
		PrevAppVersion: prevappversion,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
