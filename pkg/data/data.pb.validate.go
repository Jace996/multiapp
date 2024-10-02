// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: data/data.proto

package data

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

	structpb "google.golang.org/protobuf/types/known/structpb"
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

	_ = structpb.NullValue(0)
)

// Validate checks the field values on DynamicValue with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DynamicValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DynamicValue with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DynamicValueMultiError, or
// nil if none found.
func (m *DynamicValue) ValidateAll() error {
	return m.validate(true)
}

func (m *DynamicValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Value.(type) {
	case *DynamicValue_IntValue:
		if v == nil {
			err := DynamicValueValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for IntValue
	case *DynamicValue_LongValue:
		if v == nil {
			err := DynamicValueValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for LongValue
	case *DynamicValue_FloatValue:
		if v == nil {
			err := DynamicValueValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for FloatValue
	case *DynamicValue_DoubleValue:
		if v == nil {
			err := DynamicValueValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for DoubleValue
	case *DynamicValue_StringValue:
		if v == nil {
			err := DynamicValueValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for StringValue
	case *DynamicValue_BoolValue:
		if v == nil {
			err := DynamicValueValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for BoolValue
	case *DynamicValue_NullValue:
		if v == nil {
			err := DynamicValueValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for NullValue
	case *DynamicValue_JsonValue:
		if v == nil {
			err := DynamicValueValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetJsonValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DynamicValueValidationError{
						field:  "JsonValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DynamicValueValidationError{
						field:  "JsonValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetJsonValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DynamicValueValidationError{
					field:  "JsonValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return DynamicValueMultiError(errors)
	}

	return nil
}

// DynamicValueMultiError is an error wrapping multiple validation errors
// returned by DynamicValue.ValidateAll() if the designated constraints aren't met.
type DynamicValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DynamicValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DynamicValueMultiError) AllErrors() []error { return m }

// DynamicValueValidationError is the validation error returned by
// DynamicValue.Validate if the designated constraints aren't met.
type DynamicValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DynamicValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DynamicValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DynamicValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DynamicValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DynamicValueValidationError) ErrorName() string { return "DynamicValueValidationError" }

// Error satisfies the builtin error interface
func (e DynamicValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamicValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DynamicValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DynamicValueValidationError{}

// Validate checks the field values on DynamicValueFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DynamicValueFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DynamicValueFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DynamicValueFilterMultiError, or nil if none found.
func (m *DynamicValueFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *DynamicValueFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Filter.(type) {
	case *DynamicValueFilter_IntValue:
		if v == nil {
			err := DynamicValueFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetIntValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "IntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "IntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetIntValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DynamicValueFilterValidationError{
					field:  "IntValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DynamicValueFilter_LongValue:
		if v == nil {
			err := DynamicValueFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetLongValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "LongValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "LongValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLongValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DynamicValueFilterValidationError{
					field:  "LongValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DynamicValueFilter_FloatValue:
		if v == nil {
			err := DynamicValueFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetFloatValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "FloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "FloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFloatValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DynamicValueFilterValidationError{
					field:  "FloatValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DynamicValueFilter_DoubleValue:
		if v == nil {
			err := DynamicValueFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetDoubleValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "DoubleValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "DoubleValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDoubleValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DynamicValueFilterValidationError{
					field:  "DoubleValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DynamicValueFilter_StringValue:
		if v == nil {
			err := DynamicValueFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetStringValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "StringValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "StringValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStringValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DynamicValueFilterValidationError{
					field:  "StringValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DynamicValueFilter_BoolValue:
		if v == nil {
			err := DynamicValueFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetBoolValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "BoolValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "BoolValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetBoolValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DynamicValueFilterValidationError{
					field:  "BoolValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DynamicValueFilter_NullValue:
		if v == nil {
			err := DynamicValueFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetNullValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "NullValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DynamicValueFilterValidationError{
						field:  "NullValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetNullValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DynamicValueFilterValidationError{
					field:  "NullValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return DynamicValueFilterMultiError(errors)
	}

	return nil
}

// DynamicValueFilterMultiError is an error wrapping multiple validation errors
// returned by DynamicValueFilter.ValidateAll() if the designated constraints
// aren't met.
type DynamicValueFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DynamicValueFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DynamicValueFilterMultiError) AllErrors() []error { return m }

// DynamicValueFilterValidationError is the validation error returned by
// DynamicValueFilter.Validate if the designated constraints aren't met.
type DynamicValueFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DynamicValueFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DynamicValueFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DynamicValueFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DynamicValueFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DynamicValueFilterValidationError) ErrorName() string {
	return "DynamicValueFilterValidationError"
}

// Error satisfies the builtin error interface
func (e DynamicValueFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamicValueFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DynamicValueFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DynamicValueFilterValidationError{}
