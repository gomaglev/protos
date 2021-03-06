// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/proto/category/types.proto

package category

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _types_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Category with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Category) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Status

	// no validation rules for ParentId

	// no validation rules for Description

	// no validation rules for UserId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CategoryValidationError is the validation error returned by
// Category.Validate if the designated constraints aren't met.
type CategoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryValidationError) ErrorName() string { return "CategoryValidationError" }

// Error satisfies the builtin error interface
func (e CategoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryValidationError{}

// Validate checks the field values on Categories with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Categories) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CategoriesValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoriesValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CategoriesValidationError is the validation error returned by
// Categories.Validate if the designated constraints aren't met.
type CategoriesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoriesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoriesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoriesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoriesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoriesValidationError) ErrorName() string { return "CategoriesValidationError" }

// Error satisfies the builtin error interface
func (e CategoriesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategories.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoriesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoriesValidationError{}
