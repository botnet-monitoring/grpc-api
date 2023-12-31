// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common.proto

package gen

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Timestamp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Timestamp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Timestamp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TimestampMultiError, or nil
// if none found.
func (m *Timestamp) ValidateAll() error {
	return m.validate(true)
}

func (m *Timestamp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UnixSeconds

	// no validation rules for UnixNanoSeconds

	if len(errors) > 0 {
		return TimestampMultiError(errors)
	}

	return nil
}

// TimestampMultiError is an error wrapping multiple validation errors returned
// by Timestamp.ValidateAll() if the designated constraints aren't met.
type TimestampMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TimestampMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TimestampMultiError) AllErrors() []error { return m }

// TimestampValidationError is the validation error returned by
// Timestamp.Validate if the designated constraints aren't met.
type TimestampValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimestampValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimestampValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimestampValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimestampValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimestampValidationError) ErrorName() string { return "TimestampValidationError" }

// Error satisfies the builtin error interface
func (e TimestampValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimestamp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimestampValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimestampValidationError{}

// Validate checks the field values on IPv4Address with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IPv4Address) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IPv4Address with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IPv4AddressMultiError, or
// nil if none found.
func (m *IPv4Address) ValidateAll() error {
	return m.validate(true)
}

func (m *IPv4Address) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	if len(errors) > 0 {
		return IPv4AddressMultiError(errors)
	}

	return nil
}

// IPv4AddressMultiError is an error wrapping multiple validation errors
// returned by IPv4Address.ValidateAll() if the designated constraints aren't met.
type IPv4AddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IPv4AddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IPv4AddressMultiError) AllErrors() []error { return m }

// IPv4AddressValidationError is the validation error returned by
// IPv4Address.Validate if the designated constraints aren't met.
type IPv4AddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IPv4AddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IPv4AddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IPv4AddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IPv4AddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IPv4AddressValidationError) ErrorName() string { return "IPv4AddressValidationError" }

// Error satisfies the builtin error interface
func (e IPv4AddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPv4Address.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IPv4AddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IPv4AddressValidationError{}

// Validate checks the field values on IPv6Address with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IPv6Address) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IPv6Address with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IPv6AddressMultiError, or
// nil if none found.
func (m *IPv6Address) ValidateAll() error {
	return m.validate(true)
}

func (m *IPv6Address) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	if len(errors) > 0 {
		return IPv6AddressMultiError(errors)
	}

	return nil
}

// IPv6AddressMultiError is an error wrapping multiple validation errors
// returned by IPv6Address.ValidateAll() if the designated constraints aren't met.
type IPv6AddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IPv6AddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IPv6AddressMultiError) AllErrors() []error { return m }

// IPv6AddressValidationError is the validation error returned by
// IPv6Address.Validate if the designated constraints aren't met.
type IPv6AddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IPv6AddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IPv6AddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IPv6AddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IPv6AddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IPv6AddressValidationError) ErrorName() string { return "IPv6AddressValidationError" }

// Error satisfies the builtin error interface
func (e IPv6AddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPv6Address.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IPv6AddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IPv6AddressValidationError{}

// Validate checks the field values on IPAddress with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IPAddress) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IPAddress with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IPAddressMultiError, or nil
// if none found.
func (m *IPAddress) ValidateAll() error {
	return m.validate(true)
}

func (m *IPAddress) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Address.(type) {
	case *IPAddress_V4:
		if v == nil {
			err := IPAddressValidationError{
				field:  "Address",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetV4()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IPAddressValidationError{
						field:  "V4",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IPAddressValidationError{
						field:  "V4",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetV4()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IPAddressValidationError{
					field:  "V4",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *IPAddress_V6:
		if v == nil {
			err := IPAddressValidationError{
				field:  "Address",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetV6()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IPAddressValidationError{
						field:  "V6",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IPAddressValidationError{
						field:  "V6",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetV6()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IPAddressValidationError{
					field:  "V6",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return IPAddressMultiError(errors)
	}

	return nil
}

// IPAddressMultiError is an error wrapping multiple validation errors returned
// by IPAddress.ValidateAll() if the designated constraints aren't met.
type IPAddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IPAddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IPAddressMultiError) AllErrors() []error { return m }

// IPAddressValidationError is the validation error returned by
// IPAddress.Validate if the designated constraints aren't met.
type IPAddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IPAddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IPAddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IPAddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IPAddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IPAddressValidationError) ErrorName() string { return "IPAddressValidationError" }

// Error satisfies the builtin error interface
func (e IPAddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IPAddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IPAddressValidationError{}

// Validate checks the field values on Port with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Port) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Port with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PortMultiError, or nil if none found.
func (m *Port) ValidateAll() error {
	return m.validate(true)
}

func (m *Port) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Port

	if len(errors) > 0 {
		return PortMultiError(errors)
	}

	return nil
}

// PortMultiError is an error wrapping multiple validation errors returned by
// Port.ValidateAll() if the designated constraints aren't met.
type PortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PortMultiError) AllErrors() []error { return m }

// PortValidationError is the validation error returned by Port.Validate if the
// designated constraints aren't met.
type PortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PortValidationError) ErrorName() string { return "PortValidationError" }

// Error satisfies the builtin error interface
func (e PortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PortValidationError{}

// Validate checks the field values on JSON with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *JSON) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JSON with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in JSONMultiError, or nil if none found.
func (m *JSON) ValidateAll() error {
	return m.validate(true)
}

func (m *JSON) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for JsonString

	if len(errors) > 0 {
		return JSONMultiError(errors)
	}

	return nil
}

// JSONMultiError is an error wrapping multiple validation errors returned by
// JSON.ValidateAll() if the designated constraints aren't met.
type JSONMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JSONMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JSONMultiError) AllErrors() []error { return m }

// JSONValidationError is the validation error returned by JSON.Validate if the
// designated constraints aren't met.
type JSONValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JSONValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JSONValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JSONValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JSONValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JSONValidationError) ErrorName() string { return "JSONValidationError" }

// Error satisfies the builtin error interface
func (e JSONValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJSON.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JSONValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JSONValidationError{}
