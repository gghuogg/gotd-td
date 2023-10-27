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

// PhoneCreateGroupCallRequest represents TL type `phone.createGroupCall#48cdc6d8`.
type PhoneCreateGroupCallRequest struct {
	// Flags field of PhoneCreateGroupCallRequest.
	Flags bin.Fields
	// RtmpStream field of PhoneCreateGroupCallRequest.
	RtmpStream bool
	// Peer field of PhoneCreateGroupCallRequest.
	Peer InputPeerClass
	// RandomID field of PhoneCreateGroupCallRequest.
	RandomID int
	// Title field of PhoneCreateGroupCallRequest.
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// ScheduleDate field of PhoneCreateGroupCallRequest.
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
}

// PhoneCreateGroupCallRequestTypeID is TL type id of PhoneCreateGroupCallRequest.
const PhoneCreateGroupCallRequestTypeID = 0x48cdc6d8

// Ensuring interfaces in compile-time for PhoneCreateGroupCallRequest.
var (
	_ bin.Encoder     = &PhoneCreateGroupCallRequest{}
	_ bin.Decoder     = &PhoneCreateGroupCallRequest{}
	_ bin.BareEncoder = &PhoneCreateGroupCallRequest{}
	_ bin.BareDecoder = &PhoneCreateGroupCallRequest{}
)

func (c *PhoneCreateGroupCallRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.RtmpStream == false) {
		return false
	}
	if !(c.Peer == nil) {
		return false
	}
	if !(c.RandomID == 0) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.ScheduleDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *PhoneCreateGroupCallRequest) String() string {
	if c == nil {
		return "PhoneCreateGroupCallRequest(nil)"
	}
	type Alias PhoneCreateGroupCallRequest
	return fmt.Sprintf("PhoneCreateGroupCallRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneCreateGroupCallRequest) TypeID() uint32 {
	return PhoneCreateGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneCreateGroupCallRequest) TypeName() string {
	return "phone.createGroupCall"
}

// TypeInfo returns info about TL type.
func (c *PhoneCreateGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.createGroupCall",
		ID:   PhoneCreateGroupCallRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RtmpStream",
			SchemaName: "rtmp_stream",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "ScheduleDate",
			SchemaName: "schedule_date",
			Null:       !c.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *PhoneCreateGroupCallRequest) SetFlags() {
	if !(c.RtmpStream == false) {
		c.Flags.Set(2)
	}
	if !(c.Title == "") {
		c.Flags.Set(0)
	}
	if !(c.ScheduleDate == 0) {
		c.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (c *PhoneCreateGroupCallRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode phone.createGroupCall#48cdc6d8 as nil")
	}
	b.PutID(PhoneCreateGroupCallRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *PhoneCreateGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode phone.createGroupCall#48cdc6d8 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.createGroupCall#48cdc6d8: field flags: %w", err)
	}
	if c.Peer == nil {
		return fmt.Errorf("unable to encode phone.createGroupCall#48cdc6d8: field peer is nil")
	}
	if err := c.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.createGroupCall#48cdc6d8: field peer: %w", err)
	}
	b.PutInt(c.RandomID)
	if c.Flags.Has(0) {
		b.PutString(c.Title)
	}
	if c.Flags.Has(1) {
		b.PutInt(c.ScheduleDate)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *PhoneCreateGroupCallRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode phone.createGroupCall#48cdc6d8 to nil")
	}
	if err := b.ConsumeID(PhoneCreateGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.createGroupCall#48cdc6d8: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *PhoneCreateGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode phone.createGroupCall#48cdc6d8 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.createGroupCall#48cdc6d8: field flags: %w", err)
		}
	}
	c.RtmpStream = c.Flags.Has(2)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.createGroupCall#48cdc6d8: field peer: %w", err)
		}
		c.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.createGroupCall#48cdc6d8: field random_id: %w", err)
		}
		c.RandomID = value
	}
	if c.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.createGroupCall#48cdc6d8: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.createGroupCall#48cdc6d8: field schedule_date: %w", err)
		}
		c.ScheduleDate = value
	}
	return nil
}

// SetRtmpStream sets value of RtmpStream conditional field.
func (c *PhoneCreateGroupCallRequest) SetRtmpStream(value bool) {
	if value {
		c.Flags.Set(2)
		c.RtmpStream = true
	} else {
		c.Flags.Unset(2)
		c.RtmpStream = false
	}
}

// GetRtmpStream returns value of RtmpStream conditional field.
func (c *PhoneCreateGroupCallRequest) GetRtmpStream() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(2)
}

// GetPeer returns value of Peer field.
func (c *PhoneCreateGroupCallRequest) GetPeer() (value InputPeerClass) {
	if c == nil {
		return
	}
	return c.Peer
}

// GetRandomID returns value of RandomID field.
func (c *PhoneCreateGroupCallRequest) GetRandomID() (value int) {
	if c == nil {
		return
	}
	return c.RandomID
}

// SetTitle sets value of Title conditional field.
func (c *PhoneCreateGroupCallRequest) SetTitle(value string) {
	c.Flags.Set(0)
	c.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (c *PhoneCreateGroupCallRequest) GetTitle() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.Title, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (c *PhoneCreateGroupCallRequest) SetScheduleDate(value int) {
	c.Flags.Set(1)
	c.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (c *PhoneCreateGroupCallRequest) GetScheduleDate() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.ScheduleDate, true
}

// PhoneCreateGroupCall invokes method phone.createGroupCall#48cdc6d8 returning error if any.
func (c *Client) PhoneCreateGroupCall(ctx context.Context, request *PhoneCreateGroupCallRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
