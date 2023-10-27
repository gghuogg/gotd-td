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

// SponsoredWebPage represents TL type `sponsoredWebPage#3db8ec63`.
//
// See https://core.telegram.org/constructor/sponsoredWebPage for reference.
type SponsoredWebPage struct {
	// Flags field of SponsoredWebPage.
	Flags bin.Fields
	// URL field of SponsoredWebPage.
	URL string
	// SiteName field of SponsoredWebPage.
	SiteName string
	// Photo field of SponsoredWebPage.
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo PhotoClass
}

// SponsoredWebPageTypeID is TL type id of SponsoredWebPage.
const SponsoredWebPageTypeID = 0x3db8ec63

// Ensuring interfaces in compile-time for SponsoredWebPage.
var (
	_ bin.Encoder     = &SponsoredWebPage{}
	_ bin.Decoder     = &SponsoredWebPage{}
	_ bin.BareEncoder = &SponsoredWebPage{}
	_ bin.BareDecoder = &SponsoredWebPage{}
)

func (s *SponsoredWebPage) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.URL == "") {
		return false
	}
	if !(s.SiteName == "") {
		return false
	}
	if !(s.Photo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SponsoredWebPage) String() string {
	if s == nil {
		return "SponsoredWebPage(nil)"
	}
	type Alias SponsoredWebPage
	return fmt.Sprintf("SponsoredWebPage%+v", Alias(*s))
}

// FillFrom fills SponsoredWebPage from given interface.
func (s *SponsoredWebPage) FillFrom(from interface {
	GetURL() (value string)
	GetSiteName() (value string)
	GetPhoto() (value PhotoClass, ok bool)
}) {
	s.URL = from.GetURL()
	s.SiteName = from.GetSiteName()
	if val, ok := from.GetPhoto(); ok {
		s.Photo = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SponsoredWebPage) TypeID() uint32 {
	return SponsoredWebPageTypeID
}

// TypeName returns name of type in TL schema.
func (*SponsoredWebPage) TypeName() string {
	return "sponsoredWebPage"
}

// TypeInfo returns info about TL type.
func (s *SponsoredWebPage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sponsoredWebPage",
		ID:   SponsoredWebPageTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "SiteName",
			SchemaName: "site_name",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *SponsoredWebPage) SetFlags() {
	if !(s.Photo == nil) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *SponsoredWebPage) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredWebPage#3db8ec63 as nil")
	}
	b.PutID(SponsoredWebPageTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SponsoredWebPage) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredWebPage#3db8ec63 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredWebPage#3db8ec63: field flags: %w", err)
	}
	b.PutString(s.URL)
	b.PutString(s.SiteName)
	if s.Flags.Has(0) {
		if s.Photo == nil {
			return fmt.Errorf("unable to encode sponsoredWebPage#3db8ec63: field photo is nil")
		}
		if err := s.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode sponsoredWebPage#3db8ec63: field photo: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SponsoredWebPage) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredWebPage#3db8ec63 to nil")
	}
	if err := b.ConsumeID(SponsoredWebPageTypeID); err != nil {
		return fmt.Errorf("unable to decode sponsoredWebPage#3db8ec63: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SponsoredWebPage) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredWebPage#3db8ec63 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sponsoredWebPage#3db8ec63: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredWebPage#3db8ec63: field url: %w", err)
		}
		s.URL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredWebPage#3db8ec63: field site_name: %w", err)
		}
		s.SiteName = value
	}
	if s.Flags.Has(0) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredWebPage#3db8ec63: field photo: %w", err)
		}
		s.Photo = value
	}
	return nil
}

// GetURL returns value of URL field.
func (s *SponsoredWebPage) GetURL() (value string) {
	if s == nil {
		return
	}
	return s.URL
}

// GetSiteName returns value of SiteName field.
func (s *SponsoredWebPage) GetSiteName() (value string) {
	if s == nil {
		return
	}
	return s.SiteName
}

// SetPhoto sets value of Photo conditional field.
func (s *SponsoredWebPage) SetPhoto(value PhotoClass) {
	s.Flags.Set(0)
	s.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (s *SponsoredWebPage) GetPhoto() (value PhotoClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Photo, true
}

// GetPhotoAsNotEmpty returns mapped value of Photo conditional field and
// boolean which is true if field was set.
func (s *SponsoredWebPage) GetPhotoAsNotEmpty() (*Photo, bool) {
	if value, ok := s.GetPhoto(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}
