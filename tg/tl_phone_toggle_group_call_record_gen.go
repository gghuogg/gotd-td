// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// PhoneToggleGroupCallRecordRequest represents TL type `phone.toggleGroupCallRecord#f128c708`.
// Start or stop recording a group call: the recorded audio and video streams will be
// automatically sent to Saved messages (the chat with ourselves).
//
// See https://core.telegram.org/method/phone.toggleGroupCallRecord for reference.
type PhoneToggleGroupCallRecordRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to start or stop recording
	Start bool
	// Whether to also record video streams
	Video bool
	// The group call or livestream
	Call InputGroupCall
	// Recording title
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// If video stream recording is enabled, whether to record in portrait or landscape mode
	//
	// Use SetVideoPortrait and GetVideoPortrait helpers.
	VideoPortrait bool
}

// PhoneToggleGroupCallRecordRequestTypeID is TL type id of PhoneToggleGroupCallRecordRequest.
const PhoneToggleGroupCallRecordRequestTypeID = 0xf128c708

// Ensuring interfaces in compile-time for PhoneToggleGroupCallRecordRequest.
var (
	_ bin.Encoder     = &PhoneToggleGroupCallRecordRequest{}
	_ bin.Decoder     = &PhoneToggleGroupCallRecordRequest{}
	_ bin.BareEncoder = &PhoneToggleGroupCallRecordRequest{}
	_ bin.BareDecoder = &PhoneToggleGroupCallRecordRequest{}
)

func (t *PhoneToggleGroupCallRecordRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Start == false) {
		return false
	}
	if !(t.Video == false) {
		return false
	}
	if !(t.Call.Zero()) {
		return false
	}
	if !(t.Title == "") {
		return false
	}
	if !(t.VideoPortrait == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *PhoneToggleGroupCallRecordRequest) String() string {
	if t == nil {
		return "PhoneToggleGroupCallRecordRequest(nil)"
	}
	type Alias PhoneToggleGroupCallRecordRequest
	return fmt.Sprintf("PhoneToggleGroupCallRecordRequest%+v", Alias(*t))
}

// FillFrom fills PhoneToggleGroupCallRecordRequest from given interface.
func (t *PhoneToggleGroupCallRecordRequest) FillFrom(from interface {
	GetStart() (value bool)
	GetVideo() (value bool)
	GetCall() (value InputGroupCall)
	GetTitle() (value string, ok bool)
	GetVideoPortrait() (value bool, ok bool)
}) {
	t.Start = from.GetStart()
	t.Video = from.GetVideo()
	t.Call = from.GetCall()
	if val, ok := from.GetTitle(); ok {
		t.Title = val
	}

	if val, ok := from.GetVideoPortrait(); ok {
		t.VideoPortrait = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneToggleGroupCallRecordRequest) TypeID() uint32 {
	return PhoneToggleGroupCallRecordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneToggleGroupCallRecordRequest) TypeName() string {
	return "phone.toggleGroupCallRecord"
}

// TypeInfo returns info about TL type.
func (t *PhoneToggleGroupCallRecordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.toggleGroupCallRecord",
		ID:   PhoneToggleGroupCallRecordRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Start",
			SchemaName: "start",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Video",
			SchemaName: "video",
			Null:       !t.Flags.Has(2),
		},
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !t.Flags.Has(1),
		},
		{
			Name:       "VideoPortrait",
			SchemaName: "video_portrait",
			Null:       !t.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *PhoneToggleGroupCallRecordRequest) SetFlags() {
	if !(t.Start == false) {
		t.Flags.Set(0)
	}
	if !(t.Video == false) {
		t.Flags.Set(2)
	}
	if !(t.Title == "") {
		t.Flags.Set(1)
	}
	if !(t.VideoPortrait == false) {
		t.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (t *PhoneToggleGroupCallRecordRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode phone.toggleGroupCallRecord#f128c708 as nil")
	}
	b.PutID(PhoneToggleGroupCallRecordRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *PhoneToggleGroupCallRecordRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode phone.toggleGroupCallRecord#f128c708 as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.toggleGroupCallRecord#f128c708: field flags: %w", err)
	}
	if err := t.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.toggleGroupCallRecord#f128c708: field call: %w", err)
	}
	if t.Flags.Has(1) {
		b.PutString(t.Title)
	}
	if t.Flags.Has(2) {
		b.PutBool(t.VideoPortrait)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *PhoneToggleGroupCallRecordRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode phone.toggleGroupCallRecord#f128c708 to nil")
	}
	if err := b.ConsumeID(PhoneToggleGroupCallRecordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.toggleGroupCallRecord#f128c708: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *PhoneToggleGroupCallRecordRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode phone.toggleGroupCallRecord#f128c708 to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallRecord#f128c708: field flags: %w", err)
		}
	}
	t.Start = t.Flags.Has(0)
	t.Video = t.Flags.Has(2)
	{
		if err := t.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallRecord#f128c708: field call: %w", err)
		}
	}
	if t.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallRecord#f128c708: field title: %w", err)
		}
		t.Title = value
	}
	if t.Flags.Has(2) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallRecord#f128c708: field video_portrait: %w", err)
		}
		t.VideoPortrait = value
	}
	return nil
}

// SetStart sets value of Start conditional field.
func (t *PhoneToggleGroupCallRecordRequest) SetStart(value bool) {
	if value {
		t.Flags.Set(0)
		t.Start = true
	} else {
		t.Flags.Unset(0)
		t.Start = false
	}
}

// GetStart returns value of Start conditional field.
func (t *PhoneToggleGroupCallRecordRequest) GetStart() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(0)
}

// SetVideo sets value of Video conditional field.
func (t *PhoneToggleGroupCallRecordRequest) SetVideo(value bool) {
	if value {
		t.Flags.Set(2)
		t.Video = true
	} else {
		t.Flags.Unset(2)
		t.Video = false
	}
}

// GetVideo returns value of Video conditional field.
func (t *PhoneToggleGroupCallRecordRequest) GetVideo() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(2)
}

// GetCall returns value of Call field.
func (t *PhoneToggleGroupCallRecordRequest) GetCall() (value InputGroupCall) {
	if t == nil {
		return
	}
	return t.Call
}

// SetTitle sets value of Title conditional field.
func (t *PhoneToggleGroupCallRecordRequest) SetTitle(value string) {
	t.Flags.Set(1)
	t.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (t *PhoneToggleGroupCallRecordRequest) GetTitle() (value string, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(1) {
		return value, false
	}
	return t.Title, true
}

// SetVideoPortrait sets value of VideoPortrait conditional field.
func (t *PhoneToggleGroupCallRecordRequest) SetVideoPortrait(value bool) {
	t.Flags.Set(2)
	t.VideoPortrait = value
}

// GetVideoPortrait returns value of VideoPortrait conditional field and
// boolean which is true if field was set.
func (t *PhoneToggleGroupCallRecordRequest) GetVideoPortrait() (value bool, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(2) {
		return value, false
	}
	return t.VideoPortrait, true
}

// PhoneToggleGroupCallRecord invokes method phone.toggleGroupCallRecord#f128c708 returning error if any.
// Start or stop recording a group call: the recorded audio and video streams will be
// automatically sent to Saved messages (the chat with ourselves).
//
// Possible errors:
//
//	403 GROUPCALL_FORBIDDEN: The group call has already ended.
//	400 GROUPCALL_NOT_MODIFIED: Group call settings weren't modified.
//
// See https://core.telegram.org/method/phone.toggleGroupCallRecord for reference.
func (c *Client) PhoneToggleGroupCallRecord(ctx context.Context, request *PhoneToggleGroupCallRecordRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
