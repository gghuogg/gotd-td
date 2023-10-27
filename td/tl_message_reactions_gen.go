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

// MessageReactions represents TL type `messageReactions#4f2b9479`.
type MessageReactions struct {
	// Flags field of MessageReactions.
	Flags bin.Fields
	// Min field of MessageReactions.
	Min bool
	// CanSeeList field of MessageReactions.
	CanSeeList bool
	// Results field of MessageReactions.
	Results []ReactionCount
	// RecentReactions field of MessageReactions.
	//
	// Use SetRecentReactions and GetRecentReactions helpers.
	RecentReactions []MessagePeerReaction
}

// MessageReactionsTypeID is TL type id of MessageReactions.
const MessageReactionsTypeID = 0x4f2b9479

// Ensuring interfaces in compile-time for MessageReactions.
var (
	_ bin.Encoder     = &MessageReactions{}
	_ bin.Decoder     = &MessageReactions{}
	_ bin.BareEncoder = &MessageReactions{}
	_ bin.BareDecoder = &MessageReactions{}
)

func (m *MessageReactions) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Min == false) {
		return false
	}
	if !(m.CanSeeList == false) {
		return false
	}
	if !(m.Results == nil) {
		return false
	}
	if !(m.RecentReactions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReactions) String() string {
	if m == nil {
		return "MessageReactions(nil)"
	}
	type Alias MessageReactions
	return fmt.Sprintf("MessageReactions%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageReactions) TypeID() uint32 {
	return MessageReactionsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageReactions) TypeName() string {
	return "messageReactions"
}

// TypeInfo returns info about TL type.
func (m *MessageReactions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageReactions",
		ID:   MessageReactionsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Min",
			SchemaName: "min",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "CanSeeList",
			SchemaName: "can_see_list",
			Null:       !m.Flags.Has(2),
		},
		{
			Name:       "Results",
			SchemaName: "results",
		},
		{
			Name:       "RecentReactions",
			SchemaName: "recent_reactions",
			Null:       !m.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (m *MessageReactions) SetFlags() {
	if !(m.Min == false) {
		m.Flags.Set(0)
	}
	if !(m.CanSeeList == false) {
		m.Flags.Set(2)
	}
	if !(m.RecentReactions == nil) {
		m.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (m *MessageReactions) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReactions#4f2b9479 as nil")
	}
	b.PutID(MessageReactionsTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageReactions) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReactions#4f2b9479 as nil")
	}
	m.SetFlags()
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageReactions#4f2b9479: field flags: %w", err)
	}
	b.PutVectorHeader(len(m.Results))
	for idx, v := range m.Results {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageReactions#4f2b9479: field results element with index %d: %w", idx, err)
		}
	}
	if m.Flags.Has(1) {
		b.PutVectorHeader(len(m.RecentReactions))
		for idx, v := range m.RecentReactions {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messageReactions#4f2b9479: field recent_reactions element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageReactions) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReactions#4f2b9479 to nil")
	}
	if err := b.ConsumeID(MessageReactionsTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReactions#4f2b9479: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageReactions) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReactions#4f2b9479 to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageReactions#4f2b9479: field flags: %w", err)
		}
	}
	m.Min = m.Flags.Has(0)
	m.CanSeeList = m.Flags.Has(2)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messageReactions#4f2b9479: field results: %w", err)
		}

		if headerLen > 0 {
			m.Results = make([]ReactionCount, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ReactionCount
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messageReactions#4f2b9479: field results: %w", err)
			}
			m.Results = append(m.Results, value)
		}
	}
	if m.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messageReactions#4f2b9479: field recent_reactions: %w", err)
		}

		if headerLen > 0 {
			m.RecentReactions = make([]MessagePeerReaction, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessagePeerReaction
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messageReactions#4f2b9479: field recent_reactions: %w", err)
			}
			m.RecentReactions = append(m.RecentReactions, value)
		}
	}
	return nil
}

// SetMin sets value of Min conditional field.
func (m *MessageReactions) SetMin(value bool) {
	if value {
		m.Flags.Set(0)
		m.Min = true
	} else {
		m.Flags.Unset(0)
		m.Min = false
	}
}

// GetMin returns value of Min conditional field.
func (m *MessageReactions) GetMin() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(0)
}

// SetCanSeeList sets value of CanSeeList conditional field.
func (m *MessageReactions) SetCanSeeList(value bool) {
	if value {
		m.Flags.Set(2)
		m.CanSeeList = true
	} else {
		m.Flags.Unset(2)
		m.CanSeeList = false
	}
}

// GetCanSeeList returns value of CanSeeList conditional field.
func (m *MessageReactions) GetCanSeeList() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(2)
}

// GetResults returns value of Results field.
func (m *MessageReactions) GetResults() (value []ReactionCount) {
	if m == nil {
		return
	}
	return m.Results
}

// SetRecentReactions sets value of RecentReactions conditional field.
func (m *MessageReactions) SetRecentReactions(value []MessagePeerReaction) {
	m.Flags.Set(1)
	m.RecentReactions = value
}

// GetRecentReactions returns value of RecentReactions conditional field and
// boolean which is true if field was set.
func (m *MessageReactions) GetRecentReactions() (value []MessagePeerReaction, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(1) {
		return value, false
	}
	return m.RecentReactions, true
}
