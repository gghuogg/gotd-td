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

// StoriesSendStoryRequest represents TL type `stories.sendStory#bcb73644`.
type StoriesSendStoryRequest struct {
	// Flags field of StoriesSendStoryRequest.
	Flags bin.Fields
	// Pinned field of StoriesSendStoryRequest.
	Pinned bool
	// Noforwards field of StoriesSendStoryRequest.
	Noforwards bool
	// Peer field of StoriesSendStoryRequest.
	Peer InputPeerClass
	// Media field of StoriesSendStoryRequest.
	Media InputMediaClass
	// MediaAreas field of StoriesSendStoryRequest.
	//
	// Use SetMediaAreas and GetMediaAreas helpers.
	MediaAreas []MediaAreaClass
	// Caption field of StoriesSendStoryRequest.
	//
	// Use SetCaption and GetCaption helpers.
	Caption string
	// Entities field of StoriesSendStoryRequest.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// PrivacyRules field of StoriesSendStoryRequest.
	PrivacyRules []InputPrivacyRuleClass
	// RandomID field of StoriesSendStoryRequest.
	RandomID int64
	// Period field of StoriesSendStoryRequest.
	//
	// Use SetPeriod and GetPeriod helpers.
	Period int
}

// StoriesSendStoryRequestTypeID is TL type id of StoriesSendStoryRequest.
const StoriesSendStoryRequestTypeID = 0xbcb73644

// Ensuring interfaces in compile-time for StoriesSendStoryRequest.
var (
	_ bin.Encoder     = &StoriesSendStoryRequest{}
	_ bin.Decoder     = &StoriesSendStoryRequest{}
	_ bin.BareEncoder = &StoriesSendStoryRequest{}
	_ bin.BareDecoder = &StoriesSendStoryRequest{}
)

func (s *StoriesSendStoryRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Pinned == false) {
		return false
	}
	if !(s.Noforwards == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.Media == nil) {
		return false
	}
	if !(s.MediaAreas == nil) {
		return false
	}
	if !(s.Caption == "") {
		return false
	}
	if !(s.Entities == nil) {
		return false
	}
	if !(s.PrivacyRules == nil) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.Period == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoriesSendStoryRequest) String() string {
	if s == nil {
		return "StoriesSendStoryRequest(nil)"
	}
	type Alias StoriesSendStoryRequest
	return fmt.Sprintf("StoriesSendStoryRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesSendStoryRequest) TypeID() uint32 {
	return StoriesSendStoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesSendStoryRequest) TypeName() string {
	return "stories.sendStory"
}

// TypeInfo returns info about TL type.
func (s *StoriesSendStoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.sendStory",
		ID:   StoriesSendStoryRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pinned",
			SchemaName: "pinned",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Noforwards",
			SchemaName: "noforwards",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Media",
			SchemaName: "media",
		},
		{
			Name:       "MediaAreas",
			SchemaName: "media_areas",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "Caption",
			SchemaName: "caption",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "PrivacyRules",
			SchemaName: "privacy_rules",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "Period",
			SchemaName: "period",
			Null:       !s.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoriesSendStoryRequest) SetFlags() {
	if !(s.Pinned == false) {
		s.Flags.Set(2)
	}
	if !(s.Noforwards == false) {
		s.Flags.Set(4)
	}
	if !(s.MediaAreas == nil) {
		s.Flags.Set(5)
	}
	if !(s.Caption == "") {
		s.Flags.Set(0)
	}
	if !(s.Entities == nil) {
		s.Flags.Set(1)
	}
	if !(s.Period == 0) {
		s.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (s *StoriesSendStoryRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stories.sendStory#bcb73644 as nil")
	}
	b.PutID(StoriesSendStoryRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoriesSendStoryRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stories.sendStory#bcb73644 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field peer: %w", err)
	}
	if s.Media == nil {
		return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field media is nil")
	}
	if err := s.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field media: %w", err)
	}
	if s.Flags.Has(5) {
		b.PutVectorHeader(len(s.MediaAreas))
		for idx, v := range s.MediaAreas {
			if v == nil {
				return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field media_areas element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field media_areas element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(0) {
		b.PutString(s.Caption)
	}
	if s.Flags.Has(1) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field entities element with index %d: %w", idx, err)
			}
		}
	}
	b.PutVectorHeader(len(s.PrivacyRules))
	for idx, v := range s.PrivacyRules {
		if v == nil {
			return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field privacy_rules element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.sendStory#bcb73644: field privacy_rules element with index %d: %w", idx, err)
		}
	}
	b.PutLong(s.RandomID)
	if s.Flags.Has(3) {
		b.PutInt(s.Period)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoriesSendStoryRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stories.sendStory#bcb73644 to nil")
	}
	if err := b.ConsumeID(StoriesSendStoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.sendStory#bcb73644: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoriesSendStoryRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stories.sendStory#bcb73644 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field flags: %w", err)
		}
	}
	s.Pinned = s.Flags.Has(2)
	s.Noforwards = s.Flags.Has(4)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field media: %w", err)
		}
		s.Media = value
	}
	if s.Flags.Has(5) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field media_areas: %w", err)
		}

		if headerLen > 0 {
			s.MediaAreas = make([]MediaAreaClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMediaArea(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field media_areas: %w", err)
			}
			s.MediaAreas = append(s.MediaAreas, value)
		}
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field caption: %w", err)
		}
		s.Caption = value
	}
	if s.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field entities: %w", err)
		}

		if headerLen > 0 {
			s.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field privacy_rules: %w", err)
		}

		if headerLen > 0 {
			s.PrivacyRules = make([]InputPrivacyRuleClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPrivacyRule(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field privacy_rules: %w", err)
			}
			s.PrivacyRules = append(s.PrivacyRules, value)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field random_id: %w", err)
		}
		s.RandomID = value
	}
	if s.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendStory#bcb73644: field period: %w", err)
		}
		s.Period = value
	}
	return nil
}

// SetPinned sets value of Pinned conditional field.
func (s *StoriesSendStoryRequest) SetPinned(value bool) {
	if value {
		s.Flags.Set(2)
		s.Pinned = true
	} else {
		s.Flags.Unset(2)
		s.Pinned = false
	}
}

// GetPinned returns value of Pinned conditional field.
func (s *StoriesSendStoryRequest) GetPinned() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(2)
}

// SetNoforwards sets value of Noforwards conditional field.
func (s *StoriesSendStoryRequest) SetNoforwards(value bool) {
	if value {
		s.Flags.Set(4)
		s.Noforwards = true
	} else {
		s.Flags.Unset(4)
		s.Noforwards = false
	}
}

// GetNoforwards returns value of Noforwards conditional field.
func (s *StoriesSendStoryRequest) GetNoforwards() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(4)
}

// GetPeer returns value of Peer field.
func (s *StoriesSendStoryRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetMedia returns value of Media field.
func (s *StoriesSendStoryRequest) GetMedia() (value InputMediaClass) {
	if s == nil {
		return
	}
	return s.Media
}

// SetMediaAreas sets value of MediaAreas conditional field.
func (s *StoriesSendStoryRequest) SetMediaAreas(value []MediaAreaClass) {
	s.Flags.Set(5)
	s.MediaAreas = value
}

// GetMediaAreas returns value of MediaAreas conditional field and
// boolean which is true if field was set.
func (s *StoriesSendStoryRequest) GetMediaAreas() (value []MediaAreaClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(5) {
		return value, false
	}
	return s.MediaAreas, true
}

// SetCaption sets value of Caption conditional field.
func (s *StoriesSendStoryRequest) SetCaption(value string) {
	s.Flags.Set(0)
	s.Caption = value
}

// GetCaption returns value of Caption conditional field and
// boolean which is true if field was set.
func (s *StoriesSendStoryRequest) GetCaption() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Caption, true
}

// SetEntities sets value of Entities conditional field.
func (s *StoriesSendStoryRequest) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(1)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *StoriesSendStoryRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.Entities, true
}

// GetPrivacyRules returns value of PrivacyRules field.
func (s *StoriesSendStoryRequest) GetPrivacyRules() (value []InputPrivacyRuleClass) {
	if s == nil {
		return
	}
	return s.PrivacyRules
}

// GetRandomID returns value of RandomID field.
func (s *StoriesSendStoryRequest) GetRandomID() (value int64) {
	if s == nil {
		return
	}
	return s.RandomID
}

// SetPeriod sets value of Period conditional field.
func (s *StoriesSendStoryRequest) SetPeriod(value int) {
	s.Flags.Set(3)
	s.Period = value
}

// GetPeriod returns value of Period conditional field and
// boolean which is true if field was set.
func (s *StoriesSendStoryRequest) GetPeriod() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Period, true
}

// StoriesSendStory invokes method stories.sendStory#bcb73644 returning error if any.
func (c *Client) StoriesSendStory(ctx context.Context, request *StoriesSendStoryRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
