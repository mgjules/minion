package validator_test

import (
	"regexp"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mgjules/minion/pkg/validator"
	"github.com/stretchr/testify/assert"
)

type Customer struct {
	Name    string
	Gender  string
	Email   string
	Address *Address
}

func (c *Customer) Validate() error {
	return validator.ValidateStruct(c,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&c.Name, validation.Required, validation.Length(5, 20)), //nolint:revive
		// Gender is optional, and should be either "Female" or "Male".
		validation.Field(&c.Gender, validation.In("Female", "Male")),
		// Email cannot be empty and should be in a valid email format.
		validation.Field(&c.Email, validation.Required, is.Email),
		// Validate Address using its own validation rules
		validation.Field(&c.Address),
	)
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a *Address) Validate() error {
	return validator.ValidateStruct(a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)), //nolint:revive
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.City, validation.Required, validation.Length(5, 50)), //nolint:revive
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// State cannot be empty, and must be a string consisting of five digits
		validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
}

func TestValidateStruct(t *testing.T) {
	t.Parallel()

	c := Customer{
		Name:  "Random Guy",
		Email: "q",
		Address: &Address{
			Street: "123 Main Street",
			City:   "Somewhere",
			State:  "Unknown",
			Zip:    "12345",
		},
	}

	err := c.Validate()
	assert.NotNil(t, err)

	fieldErrors, ok := err.(validator.FieldErrors) //nolint:errorlint
	assert.True(t, ok)

	assert.Equalf(t, 2, len(fieldErrors), "Number of field errors do not match expected value")
}
