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

// WallPaper represents TL type `wallPaper#a437c3ed`.
type WallPaper struct {
	// ID field of WallPaper.
	ID int64
	// Flags field of WallPaper.
	Flags bin.Fields
	// Creator field of WallPaper.
	Creator bool
	// Default field of WallPaper.
	Default bool
	// Pattern field of WallPaper.
	Pattern bool
	// Dark field of WallPaper.
	Dark bool
	// AccessHash field of WallPaper.
	AccessHash int64
	// Slug field of WallPaper.
	Slug string
	// Document field of WallPaper.
	Document DocumentClass
	// Settings field of WallPaper.
	//
	// Use SetSettings and GetSettings helpers.
	Settings WallPaperSettings
}

// WallPaperTypeID is TL type id of WallPaper.
const WallPaperTypeID = 0xa437c3ed

// construct implements constructor of WallPaperClass.
func (w WallPaper) construct() WallPaperClass { return &w }

// Ensuring interfaces in compile-time for WallPaper.
var (
	_ bin.Encoder     = &WallPaper{}
	_ bin.Decoder     = &WallPaper{}
	_ bin.BareEncoder = &WallPaper{}
	_ bin.BareDecoder = &WallPaper{}

	_ WallPaperClass = &WallPaper{}
)

func (w *WallPaper) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.ID == 0) {
		return false
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Creator == false) {
		return false
	}
	if !(w.Default == false) {
		return false
	}
	if !(w.Pattern == false) {
		return false
	}
	if !(w.Dark == false) {
		return false
	}
	if !(w.AccessHash == 0) {
		return false
	}
	if !(w.Slug == "") {
		return false
	}
	if !(w.Document == nil) {
		return false
	}
	if !(w.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WallPaper) String() string {
	if w == nil {
		return "WallPaper(nil)"
	}
	type Alias WallPaper
	return fmt.Sprintf("WallPaper%+v", Alias(*w))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WallPaper) TypeID() uint32 {
	return WallPaperTypeID
}

// TypeName returns name of type in TL schema.
func (*WallPaper) TypeName() string {
	return "wallPaper"
}

// TypeInfo returns info about TL type.
func (w *WallPaper) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "wallPaper",
		ID:   WallPaperTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Creator",
			SchemaName: "creator",
			Null:       !w.Flags.Has(0),
		},
		{
			Name:       "Default",
			SchemaName: "default",
			Null:       !w.Flags.Has(1),
		},
		{
			Name:       "Pattern",
			SchemaName: "pattern",
			Null:       !w.Flags.Has(3),
		},
		{
			Name:       "Dark",
			SchemaName: "dark",
			Null:       !w.Flags.Has(4),
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
		{
			Name:       "Document",
			SchemaName: "document",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
			Null:       !w.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (w *WallPaper) SetFlags() {
	if !(w.Creator == false) {
		w.Flags.Set(0)
	}
	if !(w.Default == false) {
		w.Flags.Set(1)
	}
	if !(w.Pattern == false) {
		w.Flags.Set(3)
	}
	if !(w.Dark == false) {
		w.Flags.Set(4)
	}
	if !(w.Settings.Zero()) {
		w.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (w *WallPaper) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaper#a437c3ed as nil")
	}
	b.PutID(WallPaperTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WallPaper) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaper#a437c3ed as nil")
	}
	w.SetFlags()
	b.PutLong(w.ID)
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaper#a437c3ed: field flags: %w", err)
	}
	b.PutLong(w.AccessHash)
	b.PutString(w.Slug)
	if w.Document == nil {
		return fmt.Errorf("unable to encode wallPaper#a437c3ed: field document is nil")
	}
	if err := w.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaper#a437c3ed: field document: %w", err)
	}
	if w.Flags.Has(2) {
		if err := w.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode wallPaper#a437c3ed: field settings: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WallPaper) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaper#a437c3ed to nil")
	}
	if err := b.ConsumeID(WallPaperTypeID); err != nil {
		return fmt.Errorf("unable to decode wallPaper#a437c3ed: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WallPaper) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaper#a437c3ed to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field id: %w", err)
		}
		w.ID = value
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field flags: %w", err)
		}
	}
	w.Creator = w.Flags.Has(0)
	w.Default = w.Flags.Has(1)
	w.Pattern = w.Flags.Has(3)
	w.Dark = w.Flags.Has(4)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field access_hash: %w", err)
		}
		w.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field slug: %w", err)
		}
		w.Slug = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field document: %w", err)
		}
		w.Document = value
	}
	if w.Flags.Has(2) {
		if err := w.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field settings: %w", err)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (w *WallPaper) GetID() (value int64) {
	if w == nil {
		return
	}
	return w.ID
}

// SetCreator sets value of Creator conditional field.
func (w *WallPaper) SetCreator(value bool) {
	if value {
		w.Flags.Set(0)
		w.Creator = true
	} else {
		w.Flags.Unset(0)
		w.Creator = false
	}
}

// GetCreator returns value of Creator conditional field.
func (w *WallPaper) GetCreator() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(0)
}

// SetDefault sets value of Default conditional field.
func (w *WallPaper) SetDefault(value bool) {
	if value {
		w.Flags.Set(1)
		w.Default = true
	} else {
		w.Flags.Unset(1)
		w.Default = false
	}
}

// GetDefault returns value of Default conditional field.
func (w *WallPaper) GetDefault() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(1)
}

// SetPattern sets value of Pattern conditional field.
func (w *WallPaper) SetPattern(value bool) {
	if value {
		w.Flags.Set(3)
		w.Pattern = true
	} else {
		w.Flags.Unset(3)
		w.Pattern = false
	}
}

// GetPattern returns value of Pattern conditional field.
func (w *WallPaper) GetPattern() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(3)
}

// SetDark sets value of Dark conditional field.
func (w *WallPaper) SetDark(value bool) {
	if value {
		w.Flags.Set(4)
		w.Dark = true
	} else {
		w.Flags.Unset(4)
		w.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (w *WallPaper) GetDark() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(4)
}

// GetAccessHash returns value of AccessHash field.
func (w *WallPaper) GetAccessHash() (value int64) {
	if w == nil {
		return
	}
	return w.AccessHash
}

// GetSlug returns value of Slug field.
func (w *WallPaper) GetSlug() (value string) {
	if w == nil {
		return
	}
	return w.Slug
}

// GetDocument returns value of Document field.
func (w *WallPaper) GetDocument() (value DocumentClass) {
	if w == nil {
		return
	}
	return w.Document
}

// SetSettings sets value of Settings conditional field.
func (w *WallPaper) SetSettings(value WallPaperSettings) {
	w.Flags.Set(2)
	w.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (w *WallPaper) GetSettings() (value WallPaperSettings, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(2) {
		return value, false
	}
	return w.Settings, true
}

// WallPaperNoFile represents TL type `wallPaperNoFile#e0804116`.
type WallPaperNoFile struct {
	// ID field of WallPaperNoFile.
	ID int64
	// Flags field of WallPaperNoFile.
	Flags bin.Fields
	// Default field of WallPaperNoFile.
	Default bool
	// Dark field of WallPaperNoFile.
	Dark bool
	// Settings field of WallPaperNoFile.
	//
	// Use SetSettings and GetSettings helpers.
	Settings WallPaperSettings
}

// WallPaperNoFileTypeID is TL type id of WallPaperNoFile.
const WallPaperNoFileTypeID = 0xe0804116

// construct implements constructor of WallPaperClass.
func (w WallPaperNoFile) construct() WallPaperClass { return &w }

// Ensuring interfaces in compile-time for WallPaperNoFile.
var (
	_ bin.Encoder     = &WallPaperNoFile{}
	_ bin.Decoder     = &WallPaperNoFile{}
	_ bin.BareEncoder = &WallPaperNoFile{}
	_ bin.BareDecoder = &WallPaperNoFile{}

	_ WallPaperClass = &WallPaperNoFile{}
)

func (w *WallPaperNoFile) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.ID == 0) {
		return false
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Default == false) {
		return false
	}
	if !(w.Dark == false) {
		return false
	}
	if !(w.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WallPaperNoFile) String() string {
	if w == nil {
		return "WallPaperNoFile(nil)"
	}
	type Alias WallPaperNoFile
	return fmt.Sprintf("WallPaperNoFile%+v", Alias(*w))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WallPaperNoFile) TypeID() uint32 {
	return WallPaperNoFileTypeID
}

// TypeName returns name of type in TL schema.
func (*WallPaperNoFile) TypeName() string {
	return "wallPaperNoFile"
}

// TypeInfo returns info about TL type.
func (w *WallPaperNoFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "wallPaperNoFile",
		ID:   WallPaperNoFileTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Default",
			SchemaName: "default",
			Null:       !w.Flags.Has(1),
		},
		{
			Name:       "Dark",
			SchemaName: "dark",
			Null:       !w.Flags.Has(4),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
			Null:       !w.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (w *WallPaperNoFile) SetFlags() {
	if !(w.Default == false) {
		w.Flags.Set(1)
	}
	if !(w.Dark == false) {
		w.Flags.Set(4)
	}
	if !(w.Settings.Zero()) {
		w.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (w *WallPaperNoFile) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaperNoFile#e0804116 as nil")
	}
	b.PutID(WallPaperNoFileTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WallPaperNoFile) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaperNoFile#e0804116 as nil")
	}
	w.SetFlags()
	b.PutLong(w.ID)
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaperNoFile#e0804116: field flags: %w", err)
	}
	if w.Flags.Has(2) {
		if err := w.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode wallPaperNoFile#e0804116: field settings: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WallPaperNoFile) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaperNoFile#e0804116 to nil")
	}
	if err := b.ConsumeID(WallPaperNoFileTypeID); err != nil {
		return fmt.Errorf("unable to decode wallPaperNoFile#e0804116: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WallPaperNoFile) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaperNoFile#e0804116 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaperNoFile#e0804116: field id: %w", err)
		}
		w.ID = value
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaperNoFile#e0804116: field flags: %w", err)
		}
	}
	w.Default = w.Flags.Has(1)
	w.Dark = w.Flags.Has(4)
	if w.Flags.Has(2) {
		if err := w.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaperNoFile#e0804116: field settings: %w", err)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (w *WallPaperNoFile) GetID() (value int64) {
	if w == nil {
		return
	}
	return w.ID
}

// SetDefault sets value of Default conditional field.
func (w *WallPaperNoFile) SetDefault(value bool) {
	if value {
		w.Flags.Set(1)
		w.Default = true
	} else {
		w.Flags.Unset(1)
		w.Default = false
	}
}

// GetDefault returns value of Default conditional field.
func (w *WallPaperNoFile) GetDefault() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(1)
}

// SetDark sets value of Dark conditional field.
func (w *WallPaperNoFile) SetDark(value bool) {
	if value {
		w.Flags.Set(4)
		w.Dark = true
	} else {
		w.Flags.Unset(4)
		w.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (w *WallPaperNoFile) GetDark() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(4)
}

// SetSettings sets value of Settings conditional field.
func (w *WallPaperNoFile) SetSettings(value WallPaperSettings) {
	w.Flags.Set(2)
	w.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (w *WallPaperNoFile) GetSettings() (value WallPaperSettings, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(2) {
		return value, false
	}
	return w.Settings, true
}

// WallPaperClassName is schema name of WallPaperClass.
const WallPaperClassName = "WallPaper"

// WallPaperClass represents WallPaper generic type.
//
// Example:
//
//	g, err := td.DecodeWallPaper(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *td.WallPaper: // wallPaper#a437c3ed
//	case *td.WallPaperNoFile: // wallPaperNoFile#e0804116
//	default: panic(v)
//	}
type WallPaperClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() WallPaperClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// ID field of WallPaper.
	GetID() (value int64)
	// Default field of WallPaper.
	GetDefault() (value bool)
	// Dark field of WallPaper.
	GetDark() (value bool)
	// Settings field of WallPaper.
	GetSettings() (value WallPaperSettings, ok bool)
}

// DecodeWallPaper implements binary de-serialization for WallPaperClass.
func DecodeWallPaper(buf *bin.Buffer) (WallPaperClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case WallPaperTypeID:
		// Decoding wallPaper#a437c3ed.
		v := WallPaper{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WallPaperClass: %w", err)
		}
		return &v, nil
	case WallPaperNoFileTypeID:
		// Decoding wallPaperNoFile#e0804116.
		v := WallPaperNoFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WallPaperClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode WallPaperClass: %w", bin.NewUnexpectedID(id))
	}
}

// WallPaper boxes the WallPaperClass providing a helper.
type WallPaperBox struct {
	WallPaper WallPaperClass
}

// Decode implements bin.Decoder for WallPaperBox.
func (b *WallPaperBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode WallPaperBox to nil")
	}
	v, err := DecodeWallPaper(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.WallPaper = v
	return nil
}

// Encode implements bin.Encode for WallPaperBox.
func (b *WallPaperBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.WallPaper == nil {
		return fmt.Errorf("unable to encode WallPaperClass as nil")
	}
	return b.WallPaper.Encode(buf)
}
