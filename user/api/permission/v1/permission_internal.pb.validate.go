// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/api/permission/v1/permission_internal.proto

package v1

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

// Validate checks the field values on CheckSubjectsPermissionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckSubjectsPermissionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckSubjectsPermissionRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CheckSubjectsPermissionRequestMultiError, or nil if none found.
func (m *CheckSubjectsPermissionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckSubjectsPermissionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TenantId

	for idx, item := range m.GetRequirements() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CheckSubjectsPermissionRequestValidationError{
						field:  fmt.Sprintf("Requirements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckSubjectsPermissionRequestValidationError{
						field:  fmt.Sprintf("Requirements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckSubjectsPermissionRequestValidationError{
					field:  fmt.Sprintf("Requirements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CheckSubjectsPermissionRequestMultiError(errors)
	}

	return nil
}

// CheckSubjectsPermissionRequestMultiError is an error wrapping multiple
// validation errors returned by CheckSubjectsPermissionRequest.ValidateAll()
// if the designated constraints aren't met.
type CheckSubjectsPermissionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckSubjectsPermissionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckSubjectsPermissionRequestMultiError) AllErrors() []error { return m }

// CheckSubjectsPermissionRequestValidationError is the validation error
// returned by CheckSubjectsPermissionRequest.Validate if the designated
// constraints aren't met.
type CheckSubjectsPermissionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckSubjectsPermissionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckSubjectsPermissionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckSubjectsPermissionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckSubjectsPermissionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckSubjectsPermissionRequestValidationError) ErrorName() string {
	return "CheckSubjectsPermissionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CheckSubjectsPermissionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckSubjectsPermissionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckSubjectsPermissionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckSubjectsPermissionRequestValidationError{}

// Validate checks the field values on CheckSubjectsPermissionReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckSubjectsPermissionReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckSubjectsPermissionReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckSubjectsPermissionReplyMultiError, or nil if none found.
func (m *CheckSubjectsPermissionReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckSubjectsPermissionReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CheckSubjectsPermissionReplyMultiError(errors)
	}

	return nil
}

// CheckSubjectsPermissionReplyMultiError is an error wrapping multiple
// validation errors returned by CheckSubjectsPermissionReply.ValidateAll() if
// the designated constraints aren't met.
type CheckSubjectsPermissionReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckSubjectsPermissionReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckSubjectsPermissionReplyMultiError) AllErrors() []error { return m }

// CheckSubjectsPermissionReplyValidationError is the validation error returned
// by CheckSubjectsPermissionReply.Validate if the designated constraints
// aren't met.
type CheckSubjectsPermissionReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckSubjectsPermissionReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckSubjectsPermissionReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckSubjectsPermissionReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckSubjectsPermissionReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckSubjectsPermissionReplyValidationError) ErrorName() string {
	return "CheckSubjectsPermissionReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CheckSubjectsPermissionReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckSubjectsPermissionReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckSubjectsPermissionReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckSubjectsPermissionReplyValidationError{}
