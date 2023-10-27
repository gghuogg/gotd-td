//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// AccountEmailVerifiedClassArray is adapter for slice of AccountEmailVerifiedClass.
type AccountEmailVerifiedClassArray []AccountEmailVerifiedClass

// Sort sorts slice of AccountEmailVerifiedClass.
func (s AccountEmailVerifiedClassArray) Sort(less func(a, b AccountEmailVerifiedClass) bool) AccountEmailVerifiedClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountEmailVerifiedClass.
func (s AccountEmailVerifiedClassArray) SortStable(less func(a, b AccountEmailVerifiedClass) bool) AccountEmailVerifiedClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountEmailVerifiedClass.
func (s AccountEmailVerifiedClassArray) Retain(keep func(x AccountEmailVerifiedClass) bool) AccountEmailVerifiedClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AccountEmailVerifiedClassArray) First() (v AccountEmailVerifiedClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountEmailVerifiedClassArray) Last() (v AccountEmailVerifiedClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountEmailVerifiedClassArray) PopFirst() (v AccountEmailVerifiedClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountEmailVerifiedClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountEmailVerifiedClassArray) Pop() (v AccountEmailVerifiedClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAccountEmailVerified returns copy with only AccountEmailVerified constructors.
func (s AccountEmailVerifiedClassArray) AsAccountEmailVerified() (to AccountEmailVerifiedArray) {
	for _, elem := range s {
		value, ok := elem.(*AccountEmailVerified)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAccountEmailVerifiedLogin returns copy with only AccountEmailVerifiedLogin constructors.
func (s AccountEmailVerifiedClassArray) AsAccountEmailVerifiedLogin() (to AccountEmailVerifiedLoginArray) {
	for _, elem := range s {
		value, ok := elem.(*AccountEmailVerifiedLogin)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AccountEmailVerifiedArray is adapter for slice of AccountEmailVerified.
type AccountEmailVerifiedArray []AccountEmailVerified

// Sort sorts slice of AccountEmailVerified.
func (s AccountEmailVerifiedArray) Sort(less func(a, b AccountEmailVerified) bool) AccountEmailVerifiedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountEmailVerified.
func (s AccountEmailVerifiedArray) SortStable(less func(a, b AccountEmailVerified) bool) AccountEmailVerifiedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountEmailVerified.
func (s AccountEmailVerifiedArray) Retain(keep func(x AccountEmailVerified) bool) AccountEmailVerifiedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AccountEmailVerifiedArray) First() (v AccountEmailVerified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountEmailVerifiedArray) Last() (v AccountEmailVerified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountEmailVerifiedArray) PopFirst() (v AccountEmailVerified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountEmailVerified
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountEmailVerifiedArray) Pop() (v AccountEmailVerified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AccountEmailVerifiedLoginArray is adapter for slice of AccountEmailVerifiedLogin.
type AccountEmailVerifiedLoginArray []AccountEmailVerifiedLogin

// Sort sorts slice of AccountEmailVerifiedLogin.
func (s AccountEmailVerifiedLoginArray) Sort(less func(a, b AccountEmailVerifiedLogin) bool) AccountEmailVerifiedLoginArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountEmailVerifiedLogin.
func (s AccountEmailVerifiedLoginArray) SortStable(less func(a, b AccountEmailVerifiedLogin) bool) AccountEmailVerifiedLoginArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountEmailVerifiedLogin.
func (s AccountEmailVerifiedLoginArray) Retain(keep func(x AccountEmailVerifiedLogin) bool) AccountEmailVerifiedLoginArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AccountEmailVerifiedLoginArray) First() (v AccountEmailVerifiedLogin, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountEmailVerifiedLoginArray) Last() (v AccountEmailVerifiedLogin, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountEmailVerifiedLoginArray) PopFirst() (v AccountEmailVerifiedLogin, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountEmailVerifiedLogin
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountEmailVerifiedLoginArray) Pop() (v AccountEmailVerifiedLogin, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
