package interceptor

import (
	"errors"
	"testing"

	"github.com/mgjules/minion/pkg/validator"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func TestBuildViolations(t *testing.T) {
	t.Parallel()

	fieldErrors := validator.FieldErrors{
		validator.FieldError{
			Name: "Email",
			Err:  errors.New("invalid email format"),
		},
		validator.FieldError{
			Name: "Address",
			Err: validator.FieldErrors{
				validator.FieldError{
					Name: "State",
					Err:  errors.New("invalid state format"),
				},
				validator.FieldError{
					Name: "Zip",
					Err: validator.FieldErrors{
						validator.FieldError{
							Name: "Code",
							Err:  errors.New("invalid code format"),
						},
					},
				},
			},
		},
	}

	violations := buildViolations(fieldErrors, []*errdetails.BadRequest_FieldViolation{}, "")

	assert.Equal(t, 3, len(violations)) //nolint:revive
	assert.Equal(t, "Email", violations[0].Field)
	assert.Equal(t, "Address.State", violations[1].Field)
	assert.Equal(t, "Address.Zip.Code", violations[2].Field)
}
