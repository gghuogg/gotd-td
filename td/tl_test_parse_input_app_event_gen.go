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

// TestParseInputAppEventRequest represents TL type `test.parseInputAppEvent#bb0d87f1`.
type TestParseInputAppEventRequest struct {
}

// TestParseInputAppEventRequestTypeID is TL type id of TestParseInputAppEventRequest.
const TestParseInputAppEventRequestTypeID = 0xbb0d87f1

// Ensuring interfaces in compile-time for TestParseInputAppEventRequest.
var (
	_ bin.Encoder     = &TestParseInputAppEventRequest{}
	_ bin.Decoder     = &TestParseInputAppEventRequest{}
	_ bin.BareEncoder = &TestParseInputAppEventRequest{}
	_ bin.BareDecoder = &TestParseInputAppEventRequest{}
)

func (p *TestParseInputAppEventRequest) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *TestParseInputAppEventRequest) String() string {
	if p == nil {
		return "TestParseInputAppEventRequest(nil)"
	}
	type Alias TestParseInputAppEventRequest
	return fmt.Sprintf("TestParseInputAppEventRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestParseInputAppEventRequest) TypeID() uint32 {
	return TestParseInputAppEventRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestParseInputAppEventRequest) TypeName() string {
	return "test.parseInputAppEvent"
}

// TypeInfo returns info about TL type.
func (p *TestParseInputAppEventRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "test.parseInputAppEvent",
		ID:   TestParseInputAppEventRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *TestParseInputAppEventRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode test.parseInputAppEvent#bb0d87f1 as nil")
	}
	b.PutID(TestParseInputAppEventRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *TestParseInputAppEventRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode test.parseInputAppEvent#bb0d87f1 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *TestParseInputAppEventRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode test.parseInputAppEvent#bb0d87f1 to nil")
	}
	if err := b.ConsumeID(TestParseInputAppEventRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode test.parseInputAppEvent#bb0d87f1: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *TestParseInputAppEventRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode test.parseInputAppEvent#bb0d87f1 to nil")
	}
	return nil
}

// TestParseInputAppEvent invokes method test.parseInputAppEvent#bb0d87f1 returning error if any.
func (c *Client) TestParseInputAppEvent(ctx context.Context) (*InputAppEvent, error) {
	var result InputAppEvent

	request := &TestParseInputAppEventRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
