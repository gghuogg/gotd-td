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

// MessagesGetDialogUnreadMarksRequest represents TL type `messages.getDialogUnreadMarks#22e24e22`.
type MessagesGetDialogUnreadMarksRequest struct {
}

// MessagesGetDialogUnreadMarksRequestTypeID is TL type id of MessagesGetDialogUnreadMarksRequest.
const MessagesGetDialogUnreadMarksRequestTypeID = 0x22e24e22

// Ensuring interfaces in compile-time for MessagesGetDialogUnreadMarksRequest.
var (
	_ bin.Encoder     = &MessagesGetDialogUnreadMarksRequest{}
	_ bin.Decoder     = &MessagesGetDialogUnreadMarksRequest{}
	_ bin.BareEncoder = &MessagesGetDialogUnreadMarksRequest{}
	_ bin.BareDecoder = &MessagesGetDialogUnreadMarksRequest{}
)

func (g *MessagesGetDialogUnreadMarksRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDialogUnreadMarksRequest) String() string {
	if g == nil {
		return "MessagesGetDialogUnreadMarksRequest(nil)"
	}
	type Alias MessagesGetDialogUnreadMarksRequest
	return fmt.Sprintf("MessagesGetDialogUnreadMarksRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetDialogUnreadMarksRequest) TypeID() uint32 {
	return MessagesGetDialogUnreadMarksRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetDialogUnreadMarksRequest) TypeName() string {
	return "messages.getDialogUnreadMarks"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetDialogUnreadMarksRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getDialogUnreadMarks",
		ID:   MessagesGetDialogUnreadMarksRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetDialogUnreadMarksRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogUnreadMarks#22e24e22 as nil")
	}
	b.PutID(MessagesGetDialogUnreadMarksRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetDialogUnreadMarksRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogUnreadMarks#22e24e22 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDialogUnreadMarksRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogUnreadMarks#22e24e22 to nil")
	}
	if err := b.ConsumeID(MessagesGetDialogUnreadMarksRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDialogUnreadMarks#22e24e22: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetDialogUnreadMarksRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogUnreadMarks#22e24e22 to nil")
	}
	return nil
}

// MessagesGetDialogUnreadMarks invokes method messages.getDialogUnreadMarks#22e24e22 returning error if any.
func (c *Client) MessagesGetDialogUnreadMarks(ctx context.Context) ([]DialogPeerClass, error) {
	var result DialogPeerClassVector

	request := &MessagesGetDialogUnreadMarksRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []DialogPeerClass(result.Elems), nil
}
