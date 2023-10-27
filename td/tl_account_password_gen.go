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

// AccountPassword represents TL type `account.password#957b50fb`.
type AccountPassword struct {
	// Flags field of AccountPassword.
	Flags bin.Fields
	// HasRecovery field of AccountPassword.
	HasRecovery bool
	// HasSecureValues field of AccountPassword.
	HasSecureValues bool
	// HasPassword field of AccountPassword.
	HasPassword bool
	// CurrentAlgo field of AccountPassword.
	//
	// Use SetCurrentAlgo and GetCurrentAlgo helpers.
	CurrentAlgo PasswordKdfAlgoClass
	// SRPB field of AccountPassword.
	//
	// Use SetSRPB and GetSRPB helpers.
	SRPB []byte
	// SRPID field of AccountPassword.
	//
	// Use SetSRPID and GetSRPID helpers.
	SRPID int64
	// Hint field of AccountPassword.
	//
	// Use SetHint and GetHint helpers.
	Hint string
	// EmailUnconfirmedPattern field of AccountPassword.
	//
	// Use SetEmailUnconfirmedPattern and GetEmailUnconfirmedPattern helpers.
	EmailUnconfirmedPattern string
	// NewAlgo field of AccountPassword.
	NewAlgo PasswordKdfAlgoClass
	// NewSecureAlgo field of AccountPassword.
	NewSecureAlgo SecurePasswordKdfAlgoClass
	// SecureRandom field of AccountPassword.
	SecureRandom []byte
	// PendingResetDate field of AccountPassword.
	//
	// Use SetPendingResetDate and GetPendingResetDate helpers.
	PendingResetDate int
	// LoginEmailPattern field of AccountPassword.
	//
	// Use SetLoginEmailPattern and GetLoginEmailPattern helpers.
	LoginEmailPattern string
}

// AccountPasswordTypeID is TL type id of AccountPassword.
const AccountPasswordTypeID = 0x957b50fb

// Ensuring interfaces in compile-time for AccountPassword.
var (
	_ bin.Encoder     = &AccountPassword{}
	_ bin.Decoder     = &AccountPassword{}
	_ bin.BareEncoder = &AccountPassword{}
	_ bin.BareDecoder = &AccountPassword{}
)

func (p *AccountPassword) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.HasRecovery == false) {
		return false
	}
	if !(p.HasSecureValues == false) {
		return false
	}
	if !(p.HasPassword == false) {
		return false
	}
	if !(p.CurrentAlgo == nil) {
		return false
	}
	if !(p.SRPB == nil) {
		return false
	}
	if !(p.SRPID == 0) {
		return false
	}
	if !(p.Hint == "") {
		return false
	}
	if !(p.EmailUnconfirmedPattern == "") {
		return false
	}
	if !(p.NewAlgo == nil) {
		return false
	}
	if !(p.NewSecureAlgo == nil) {
		return false
	}
	if !(p.SecureRandom == nil) {
		return false
	}
	if !(p.PendingResetDate == 0) {
		return false
	}
	if !(p.LoginEmailPattern == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *AccountPassword) String() string {
	if p == nil {
		return "AccountPassword(nil)"
	}
	type Alias AccountPassword
	return fmt.Sprintf("AccountPassword%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountPassword) TypeID() uint32 {
	return AccountPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountPassword) TypeName() string {
	return "account.password"
}

// TypeInfo returns info about TL type.
func (p *AccountPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.password",
		ID:   AccountPasswordTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasRecovery",
			SchemaName: "has_recovery",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "HasSecureValues",
			SchemaName: "has_secure_values",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "HasPassword",
			SchemaName: "has_password",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "CurrentAlgo",
			SchemaName: "current_algo",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "SRPB",
			SchemaName: "srp_B",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "SRPID",
			SchemaName: "srp_id",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "Hint",
			SchemaName: "hint",
			Null:       !p.Flags.Has(3),
		},
		{
			Name:       "EmailUnconfirmedPattern",
			SchemaName: "email_unconfirmed_pattern",
			Null:       !p.Flags.Has(4),
		},
		{
			Name:       "NewAlgo",
			SchemaName: "new_algo",
		},
		{
			Name:       "NewSecureAlgo",
			SchemaName: "new_secure_algo",
		},
		{
			Name:       "SecureRandom",
			SchemaName: "secure_random",
		},
		{
			Name:       "PendingResetDate",
			SchemaName: "pending_reset_date",
			Null:       !p.Flags.Has(5),
		},
		{
			Name:       "LoginEmailPattern",
			SchemaName: "login_email_pattern",
			Null:       !p.Flags.Has(6),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *AccountPassword) SetFlags() {
	if !(p.HasRecovery == false) {
		p.Flags.Set(0)
	}
	if !(p.HasSecureValues == false) {
		p.Flags.Set(1)
	}
	if !(p.HasPassword == false) {
		p.Flags.Set(2)
	}
	if !(p.CurrentAlgo == nil) {
		p.Flags.Set(2)
	}
	if !(p.SRPB == nil) {
		p.Flags.Set(2)
	}
	if !(p.SRPID == 0) {
		p.Flags.Set(2)
	}
	if !(p.Hint == "") {
		p.Flags.Set(3)
	}
	if !(p.EmailUnconfirmedPattern == "") {
		p.Flags.Set(4)
	}
	if !(p.PendingResetDate == 0) {
		p.Flags.Set(5)
	}
	if !(p.LoginEmailPattern == "") {
		p.Flags.Set(6)
	}
}

// Encode implements bin.Encoder.
func (p *AccountPassword) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode account.password#957b50fb as nil")
	}
	b.PutID(AccountPasswordTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *AccountPassword) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode account.password#957b50fb as nil")
	}
	p.SetFlags()
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.password#957b50fb: field flags: %w", err)
	}
	if p.Flags.Has(2) {
		if p.CurrentAlgo == nil {
			return fmt.Errorf("unable to encode account.password#957b50fb: field current_algo is nil")
		}
		if err := p.CurrentAlgo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.password#957b50fb: field current_algo: %w", err)
		}
	}
	if p.Flags.Has(2) {
		b.PutBytes(p.SRPB)
	}
	if p.Flags.Has(2) {
		b.PutLong(p.SRPID)
	}
	if p.Flags.Has(3) {
		b.PutString(p.Hint)
	}
	if p.Flags.Has(4) {
		b.PutString(p.EmailUnconfirmedPattern)
	}
	if p.NewAlgo == nil {
		return fmt.Errorf("unable to encode account.password#957b50fb: field new_algo is nil")
	}
	if err := p.NewAlgo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.password#957b50fb: field new_algo: %w", err)
	}
	if p.NewSecureAlgo == nil {
		return fmt.Errorf("unable to encode account.password#957b50fb: field new_secure_algo is nil")
	}
	if err := p.NewSecureAlgo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.password#957b50fb: field new_secure_algo: %w", err)
	}
	b.PutBytes(p.SecureRandom)
	if p.Flags.Has(5) {
		b.PutInt(p.PendingResetDate)
	}
	if p.Flags.Has(6) {
		b.PutString(p.LoginEmailPattern)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *AccountPassword) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode account.password#957b50fb to nil")
	}
	if err := b.ConsumeID(AccountPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode account.password#957b50fb: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *AccountPassword) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode account.password#957b50fb to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field flags: %w", err)
		}
	}
	p.HasRecovery = p.Flags.Has(0)
	p.HasSecureValues = p.Flags.Has(1)
	p.HasPassword = p.Flags.Has(2)
	if p.Flags.Has(2) {
		value, err := DecodePasswordKdfAlgo(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field current_algo: %w", err)
		}
		p.CurrentAlgo = value
	}
	if p.Flags.Has(2) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field srp_B: %w", err)
		}
		p.SRPB = value
	}
	if p.Flags.Has(2) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field srp_id: %w", err)
		}
		p.SRPID = value
	}
	if p.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field hint: %w", err)
		}
		p.Hint = value
	}
	if p.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field email_unconfirmed_pattern: %w", err)
		}
		p.EmailUnconfirmedPattern = value
	}
	{
		value, err := DecodePasswordKdfAlgo(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field new_algo: %w", err)
		}
		p.NewAlgo = value
	}
	{
		value, err := DecodeSecurePasswordKdfAlgo(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field new_secure_algo: %w", err)
		}
		p.NewSecureAlgo = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field secure_random: %w", err)
		}
		p.SecureRandom = value
	}
	if p.Flags.Has(5) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field pending_reset_date: %w", err)
		}
		p.PendingResetDate = value
	}
	if p.Flags.Has(6) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#957b50fb: field login_email_pattern: %w", err)
		}
		p.LoginEmailPattern = value
	}
	return nil
}

// SetHasRecovery sets value of HasRecovery conditional field.
func (p *AccountPassword) SetHasRecovery(value bool) {
	if value {
		p.Flags.Set(0)
		p.HasRecovery = true
	} else {
		p.Flags.Unset(0)
		p.HasRecovery = false
	}
}

// GetHasRecovery returns value of HasRecovery conditional field.
func (p *AccountPassword) GetHasRecovery() (value bool) {
	if p == nil {
		return
	}
	return p.Flags.Has(0)
}

// SetHasSecureValues sets value of HasSecureValues conditional field.
func (p *AccountPassword) SetHasSecureValues(value bool) {
	if value {
		p.Flags.Set(1)
		p.HasSecureValues = true
	} else {
		p.Flags.Unset(1)
		p.HasSecureValues = false
	}
}

// GetHasSecureValues returns value of HasSecureValues conditional field.
func (p *AccountPassword) GetHasSecureValues() (value bool) {
	if p == nil {
		return
	}
	return p.Flags.Has(1)
}

// SetHasPassword sets value of HasPassword conditional field.
func (p *AccountPassword) SetHasPassword(value bool) {
	if value {
		p.Flags.Set(2)
		p.HasPassword = true
	} else {
		p.Flags.Unset(2)
		p.HasPassword = false
	}
}

// GetHasPassword returns value of HasPassword conditional field.
func (p *AccountPassword) GetHasPassword() (value bool) {
	if p == nil {
		return
	}
	return p.Flags.Has(2)
}

// SetCurrentAlgo sets value of CurrentAlgo conditional field.
func (p *AccountPassword) SetCurrentAlgo(value PasswordKdfAlgoClass) {
	p.Flags.Set(2)
	p.CurrentAlgo = value
}

// GetCurrentAlgo returns value of CurrentAlgo conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetCurrentAlgo() (value PasswordKdfAlgoClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.CurrentAlgo, true
}

// SetSRPB sets value of SRPB conditional field.
func (p *AccountPassword) SetSRPB(value []byte) {
	p.Flags.Set(2)
	p.SRPB = value
}

// GetSRPB returns value of SRPB conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetSRPB() (value []byte, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.SRPB, true
}

// SetSRPID sets value of SRPID conditional field.
func (p *AccountPassword) SetSRPID(value int64) {
	p.Flags.Set(2)
	p.SRPID = value
}

// GetSRPID returns value of SRPID conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetSRPID() (value int64, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.SRPID, true
}

// SetHint sets value of Hint conditional field.
func (p *AccountPassword) SetHint(value string) {
	p.Flags.Set(3)
	p.Hint = value
}

// GetHint returns value of Hint conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetHint() (value string, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(3) {
		return value, false
	}
	return p.Hint, true
}

// SetEmailUnconfirmedPattern sets value of EmailUnconfirmedPattern conditional field.
func (p *AccountPassword) SetEmailUnconfirmedPattern(value string) {
	p.Flags.Set(4)
	p.EmailUnconfirmedPattern = value
}

// GetEmailUnconfirmedPattern returns value of EmailUnconfirmedPattern conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetEmailUnconfirmedPattern() (value string, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.EmailUnconfirmedPattern, true
}

// GetNewAlgo returns value of NewAlgo field.
func (p *AccountPassword) GetNewAlgo() (value PasswordKdfAlgoClass) {
	if p == nil {
		return
	}
	return p.NewAlgo
}

// GetNewSecureAlgo returns value of NewSecureAlgo field.
func (p *AccountPassword) GetNewSecureAlgo() (value SecurePasswordKdfAlgoClass) {
	if p == nil {
		return
	}
	return p.NewSecureAlgo
}

// GetSecureRandom returns value of SecureRandom field.
func (p *AccountPassword) GetSecureRandom() (value []byte) {
	if p == nil {
		return
	}
	return p.SecureRandom
}

// SetPendingResetDate sets value of PendingResetDate conditional field.
func (p *AccountPassword) SetPendingResetDate(value int) {
	p.Flags.Set(5)
	p.PendingResetDate = value
}

// GetPendingResetDate returns value of PendingResetDate conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetPendingResetDate() (value int, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(5) {
		return value, false
	}
	return p.PendingResetDate, true
}

// SetLoginEmailPattern sets value of LoginEmailPattern conditional field.
func (p *AccountPassword) SetLoginEmailPattern(value string) {
	p.Flags.Set(6)
	p.LoginEmailPattern = value
}

// GetLoginEmailPattern returns value of LoginEmailPattern conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetLoginEmailPattern() (value string, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(6) {
		return value, false
	}
	return p.LoginEmailPattern, true
}
