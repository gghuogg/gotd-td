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

// StatsAbsValueAndPrev represents TL type `statsAbsValueAndPrev#cb43acde`.
type StatsAbsValueAndPrev struct {
	// Current field of StatsAbsValueAndPrev.
	Current float64
	// Previous field of StatsAbsValueAndPrev.
	Previous float64
}

// StatsAbsValueAndPrevTypeID is TL type id of StatsAbsValueAndPrev.
const StatsAbsValueAndPrevTypeID = 0xcb43acde

// Ensuring interfaces in compile-time for StatsAbsValueAndPrev.
var (
	_ bin.Encoder     = &StatsAbsValueAndPrev{}
	_ bin.Decoder     = &StatsAbsValueAndPrev{}
	_ bin.BareEncoder = &StatsAbsValueAndPrev{}
	_ bin.BareDecoder = &StatsAbsValueAndPrev{}
)

func (s *StatsAbsValueAndPrev) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Current == 0) {
		return false
	}
	if !(s.Previous == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsAbsValueAndPrev) String() string {
	if s == nil {
		return "StatsAbsValueAndPrev(nil)"
	}
	type Alias StatsAbsValueAndPrev
	return fmt.Sprintf("StatsAbsValueAndPrev%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsAbsValueAndPrev) TypeID() uint32 {
	return StatsAbsValueAndPrevTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsAbsValueAndPrev) TypeName() string {
	return "statsAbsValueAndPrev"
}

// TypeInfo returns info about TL type.
func (s *StatsAbsValueAndPrev) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsAbsValueAndPrev",
		ID:   StatsAbsValueAndPrevTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Current",
			SchemaName: "current",
		},
		{
			Name:       "Previous",
			SchemaName: "previous",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsAbsValueAndPrev) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsAbsValueAndPrev#cb43acde as nil")
	}
	b.PutID(StatsAbsValueAndPrevTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsAbsValueAndPrev) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsAbsValueAndPrev#cb43acde as nil")
	}
	b.PutDouble(s.Current)
	b.PutDouble(s.Previous)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsAbsValueAndPrev) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsAbsValueAndPrev#cb43acde to nil")
	}
	if err := b.ConsumeID(StatsAbsValueAndPrevTypeID); err != nil {
		return fmt.Errorf("unable to decode statsAbsValueAndPrev#cb43acde: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsAbsValueAndPrev) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsAbsValueAndPrev#cb43acde to nil")
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode statsAbsValueAndPrev#cb43acde: field current: %w", err)
		}
		s.Current = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode statsAbsValueAndPrev#cb43acde: field previous: %w", err)
		}
		s.Previous = value
	}
	return nil
}

// GetCurrent returns value of Current field.
func (s *StatsAbsValueAndPrev) GetCurrent() (value float64) {
	if s == nil {
		return
	}
	return s.Current
}

// GetPrevious returns value of Previous field.
func (s *StatsAbsValueAndPrev) GetPrevious() (value float64) {
	if s == nil {
		return
	}
	return s.Previous
}
