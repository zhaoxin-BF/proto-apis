// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: stream/v1/service.proto

package stream

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

// Validate checks the field values on Order with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Order) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Order with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OrderMultiError, or nil if none found.
func (m *Order) ValidateAll() error {
	return m.validate(true)
}

func (m *Order) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for OrderSn

	// no validation rules for Date

	if len(errors) > 0 {
		return OrderMultiError(errors)
	}

	return nil
}

// OrderMultiError is an error wrapping multiple validation errors returned by
// Order.ValidateAll() if the designated constraints aren't met.
type OrderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderMultiError) AllErrors() []error { return m }

// OrderValidationError is the validation error returned by Order.Validate if
// the designated constraints aren't met.
type OrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderValidationError) ErrorName() string { return "OrderValidationError" }

// Error satisfies the builtin error interface
func (e OrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderValidationError{}

// Validate checks the field values on OrderListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OrderListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderListResponseMultiError, or nil if none found.
func (m *OrderListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOrder()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderListResponseValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderListResponseValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderListResponseValidationError{
				field:  "Order",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OrderListResponseMultiError(errors)
	}

	return nil
}

// OrderListResponseMultiError is an error wrapping multiple validation errors
// returned by OrderListResponse.ValidateAll() if the designated constraints
// aren't met.
type OrderListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderListResponseMultiError) AllErrors() []error { return m }

// OrderListResponseValidationError is the validation error returned by
// OrderListResponse.Validate if the designated constraints aren't met.
type OrderListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderListResponseValidationError) ErrorName() string {
	return "OrderListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OrderListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderListResponseValidationError{}

// Validate checks the field values on OrderSearchParams with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OrderSearchParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderSearchParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderSearchParamsMultiError, or nil if none found.
func (m *OrderSearchParams) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderSearchParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OrderSearchParamsMultiError(errors)
	}

	return nil
}

// OrderSearchParamsMultiError is an error wrapping multiple validation errors
// returned by OrderSearchParams.ValidateAll() if the designated constraints
// aren't met.
type OrderSearchParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderSearchParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderSearchParamsMultiError) AllErrors() []error { return m }

// OrderSearchParamsValidationError is the validation error returned by
// OrderSearchParams.Validate if the designated constraints aren't met.
type OrderSearchParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderSearchParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderSearchParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderSearchParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderSearchParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderSearchParamsValidationError) ErrorName() string {
	return "OrderSearchParamsValidationError"
}

// Error satisfies the builtin error interface
func (e OrderSearchParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderSearchParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderSearchParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderSearchParamsValidationError{}

// Validate checks the field values on Image with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Image) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Image with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ImageMultiError, or nil if none found.
func (m *Image) ValidateAll() error {
	return m.validate(true)
}

func (m *Image) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FileName

	// no validation rules for File

	if len(errors) > 0 {
		return ImageMultiError(errors)
	}

	return nil
}

// ImageMultiError is an error wrapping multiple validation errors returned by
// Image.ValidateAll() if the designated constraints aren't met.
type ImageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageMultiError) AllErrors() []error { return m }

// ImageValidationError is the validation error returned by Image.Validate if
// the designated constraints aren't met.
type ImageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageValidationError) ErrorName() string { return "ImageValidationError" }

// Error satisfies the builtin error interface
func (e ImageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageValidationError{}

// Validate checks the field values on ImageList with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ImageList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImageList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ImageListMultiError, or nil
// if none found.
func (m *ImageList) ValidateAll() error {
	return m.validate(true)
}

func (m *ImageList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetImage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ImageListValidationError{
					field:  "Image",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImageListValidationError{
					field:  "Image",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetImage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageListValidationError{
				field:  "Image",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ImageListMultiError(errors)
	}

	return nil
}

// ImageListMultiError is an error wrapping multiple validation errors returned
// by ImageList.ValidateAll() if the designated constraints aren't met.
type ImageListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageListMultiError) AllErrors() []error { return m }

// ImageListValidationError is the validation error returned by
// ImageList.Validate if the designated constraints aren't met.
type ImageListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageListValidationError) ErrorName() string { return "ImageListValidationError" }

// Error satisfies the builtin error interface
func (e ImageListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageListValidationError{}

// Validate checks the field values on UploadResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UploadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UploadResponseMultiError,
// or nil if none found.
func (m *UploadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UploadResponseMultiError(errors)
	}

	return nil
}

// UploadResponseMultiError is an error wrapping multiple validation errors
// returned by UploadResponse.ValidateAll() if the designated constraints
// aren't met.
type UploadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadResponseMultiError) AllErrors() []error { return m }

// UploadResponseValidationError is the validation error returned by
// UploadResponse.Validate if the designated constraints aren't met.
type UploadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadResponseValidationError) ErrorName() string { return "UploadResponseValidationError" }

// Error satisfies the builtin error interface
func (e UploadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadResponseValidationError{}

// Validate checks the field values on SumDataRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SumDataRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SumDataRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SumDataRequestMultiError,
// or nil if none found.
func (m *SumDataRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SumDataRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Number

	if len(errors) > 0 {
		return SumDataRequestMultiError(errors)
	}

	return nil
}

// SumDataRequestMultiError is an error wrapping multiple validation errors
// returned by SumDataRequest.ValidateAll() if the designated constraints
// aren't met.
type SumDataRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SumDataRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SumDataRequestMultiError) AllErrors() []error { return m }

// SumDataRequestValidationError is the validation error returned by
// SumDataRequest.Validate if the designated constraints aren't met.
type SumDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SumDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SumDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SumDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SumDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SumDataRequestValidationError) ErrorName() string { return "SumDataRequestValidationError" }

// Error satisfies the builtin error interface
func (e SumDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSumDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SumDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SumDataRequestValidationError{}

// Validate checks the field values on SumDataResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SumDataResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SumDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SumDataResponseMultiError, or nil if none found.
func (m *SumDataResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SumDataResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return SumDataResponseMultiError(errors)
	}

	return nil
}

// SumDataResponseMultiError is an error wrapping multiple validation errors
// returned by SumDataResponse.ValidateAll() if the designated constraints
// aren't met.
type SumDataResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SumDataResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SumDataResponseMultiError) AllErrors() []error { return m }

// SumDataResponseValidationError is the validation error returned by
// SumDataResponse.Validate if the designated constraints aren't met.
type SumDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SumDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SumDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SumDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SumDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SumDataResponseValidationError) ErrorName() string { return "SumDataResponseValidationError" }

// Error satisfies the builtin error interface
func (e SumDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSumDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SumDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SumDataResponseValidationError{}