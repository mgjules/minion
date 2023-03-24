package validator

import (
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/samber/lo"
)

// FieldErrors is a collection of field error.
type FieldErrors []FieldError

// Error implements the error interface.
func (fes FieldErrors) Error() string {
	var errStr strings.Builder
	for i := range fes {
		errStr.WriteString(fes[i].Error() + "; ")
	}

	return strings.TrimRight(errStr.String(), "; ")
}

// FieldError hold the attributes for a single field error.
type FieldError struct {
	Name string
	Err  error
}

func (fe FieldError) Error() string {
	return "'" + fe.Name + "': " + fe.Err.Error()
}

// NewFieldErrors returns an error of type FieldErrors.
func NewFieldErrors(err error) error {
	if err == nil {
		return nil
	}

	switch err := err.(type) { //nolint:errorlint
	case validation.InternalError:
		// Check if we got an internal validation error.
		// Usually happens through misuse of the ozzo-validation package.
		return fmt.Errorf("internal validation error: %w", err)
	case validation.Errors:
		// Transforms validation errors to FieldErrors.
		fieldErrors := lo.MapToSlice(err, func(name string, err error) FieldError {
			return FieldError{
				Name: name,
				Err:  err,
			}
		})

		return FieldErrors(fieldErrors)
	default:
		return fmt.Errorf("unknown validation error: %w", err)
	}
}

// ValidateStruct is a wrapper around validation.ValidateStruct that transforms validation.Errors
// to FieldErrors.
func ValidateStruct(structPtr interface{}, fields ...*validation.FieldRules) error {
	return NewFieldErrors(validation.ValidateStruct(structPtr, fields...))
}
