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

// StatsDateRangeDays represents TL type `statsDateRangeDays#b637edaf`.
type StatsDateRangeDays struct {
	// MinDate field of StatsDateRangeDays.
	MinDate int
	// MaxDate field of StatsDateRangeDays.
	MaxDate int
}

// StatsDateRangeDaysTypeID is TL type id of StatsDateRangeDays.
const StatsDateRangeDaysTypeID = 0xb637edaf

// Ensuring interfaces in compile-time for StatsDateRangeDays.
var (
	_ bin.Encoder     = &StatsDateRangeDays{}
	_ bin.Decoder     = &StatsDateRangeDays{}
	_ bin.BareEncoder = &StatsDateRangeDays{}
	_ bin.BareDecoder = &StatsDateRangeDays{}
)

func (s *StatsDateRangeDays) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.MinDate == 0) {
		return false
	}
	if !(s.MaxDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsDateRangeDays) String() string {
	if s == nil {
		return "StatsDateRangeDays(nil)"
	}
	type Alias StatsDateRangeDays
	return fmt.Sprintf("StatsDateRangeDays%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsDateRangeDays) TypeID() uint32 {
	return StatsDateRangeDaysTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsDateRangeDays) TypeName() string {
	return "statsDateRangeDays"
}

// TypeInfo returns info about TL type.
func (s *StatsDateRangeDays) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsDateRangeDays",
		ID:   StatsDateRangeDaysTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MinDate",
			SchemaName: "min_date",
		},
		{
			Name:       "MaxDate",
			SchemaName: "max_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsDateRangeDays) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsDateRangeDays#b637edaf as nil")
	}
	b.PutID(StatsDateRangeDaysTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsDateRangeDays) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsDateRangeDays#b637edaf as nil")
	}
	b.PutInt(s.MinDate)
	b.PutInt(s.MaxDate)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsDateRangeDays) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsDateRangeDays#b637edaf to nil")
	}
	if err := b.ConsumeID(StatsDateRangeDaysTypeID); err != nil {
		return fmt.Errorf("unable to decode statsDateRangeDays#b637edaf: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsDateRangeDays) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsDateRangeDays#b637edaf to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsDateRangeDays#b637edaf: field min_date: %w", err)
		}
		s.MinDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsDateRangeDays#b637edaf: field max_date: %w", err)
		}
		s.MaxDate = value
	}
	return nil
}

// GetMinDate returns value of MinDate field.
func (s *StatsDateRangeDays) GetMinDate() (value int) {
	if s == nil {
		return
	}
	return s.MinDate
}

// GetMaxDate returns value of MaxDate field.
func (s *StatsDateRangeDays) GetMaxDate() (value int) {
	if s == nil {
		return
	}
	return s.MaxDate
}
