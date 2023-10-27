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

// InvokeAfterMsgRequest represents TL type `invokeAfterMsg#cb9f372d`.
type InvokeAfterMsgRequest struct {
	// MsgID field of InvokeAfterMsgRequest.
	MsgID int64
	// Query field of InvokeAfterMsgRequest.
	Query bin.Object
}

// InvokeAfterMsgRequestTypeID is TL type id of InvokeAfterMsgRequest.
const InvokeAfterMsgRequestTypeID = 0xcb9f372d

// Ensuring interfaces in compile-time for InvokeAfterMsgRequest.
var (
	_ bin.Encoder     = &InvokeAfterMsgRequest{}
	_ bin.Decoder     = &InvokeAfterMsgRequest{}
	_ bin.BareEncoder = &InvokeAfterMsgRequest{}
	_ bin.BareDecoder = &InvokeAfterMsgRequest{}
)

func (i *InvokeAfterMsgRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.MsgID == 0) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeAfterMsgRequest) String() string {
	if i == nil {
		return "InvokeAfterMsgRequest(nil)"
	}
	type Alias InvokeAfterMsgRequest
	return fmt.Sprintf("InvokeAfterMsgRequest%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InvokeAfterMsgRequest) TypeID() uint32 {
	return InvokeAfterMsgRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*InvokeAfterMsgRequest) TypeName() string {
	return "invokeAfterMsg"
}

// TypeInfo returns info about TL type.
func (i *InvokeAfterMsgRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "invokeAfterMsg",
		ID:   InvokeAfterMsgRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InvokeAfterMsgRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeAfterMsg#cb9f372d as nil")
	}
	b.PutID(InvokeAfterMsgRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InvokeAfterMsgRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeAfterMsg#cb9f372d as nil")
	}
	b.PutLong(i.MsgID)
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeAfterMsg#cb9f372d: field query: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InvokeAfterMsgRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeAfterMsg#cb9f372d to nil")
	}
	if err := b.ConsumeID(InvokeAfterMsgRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeAfterMsg#cb9f372d: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InvokeAfterMsgRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeAfterMsg#cb9f372d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode invokeAfterMsg#cb9f372d: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeAfterMsg#cb9f372d: field query: %w", err)
		}
	}
	return nil
}

// GetMsgID returns value of MsgID field.
func (i *InvokeAfterMsgRequest) GetMsgID() (value int64) {
	if i == nil {
		return
	}
	return i.MsgID
}

// GetQuery returns value of Query field.
func (i *InvokeAfterMsgRequest) GetQuery() (value bin.Object) {
	if i == nil {
		return
	}
	return i.Query
}
