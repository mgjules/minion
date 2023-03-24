package interceptor

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mgjules/minion/pkg/logger"
	"github.com/mgjules/minion/pkg/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validator is a server interceptor which validates incoming RPC requests.
type Validator struct {
	logger *logger.Logger
}

// NewValidator returns a new validation interceptor.
func NewValidator(logger *logger.Logger) *Validator {
	return &Validator{logger}
}

// Unary returns a server interceptor function to validate incoming Unary RPC requests.
func (v *Validator) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if vreq, ok := req.(validation.Validatable); ok {
			if err := vreq.Validate(); err != nil {
				if fieldErrors, ok := err.(validator.FieldErrors); ok { //nolint:errorlint
					if err = invalidArgumentError(fieldErrors); err != nil {
						v.logger.Ctx(ctx).Errorf("invalid argument error: %v", err)

						return nil, err
					}
				}

				// Internal validation error
				v.logger.Ctx(ctx).Errorf("validate request struct '%T': %v", req, err)

				return nil, internalError("validate request parameters")
			}
		}

		return handler(ctx, req)
	}
}

func internalError(reason string) error {
	if reason == "" {
		reason = "unknown error"
	}

	statusInternal := status.New(codes.Internal, "internal server error")
	statusDetails, err := statusInternal.WithDetails(&errdetails.ErrorInfo{
		Reason: reason,
	})
	if err != nil {
		return statusInternal.Err()
	}

	return statusDetails.Err()
}

func invalidArgumentError(fieldErrors validator.FieldErrors) error {
	if len(fieldErrors) == 0 {
		return nil
	}

	var violations []*errdetails.BadRequest_FieldViolation
	violations = buildViolations(fieldErrors, violations, "")

	statusInvalidArg := status.New(codes.InvalidArgument, "invalid request parameters")
	statusDetails, err := statusInvalidArg.WithDetails(&errdetails.BadRequest{
		FieldViolations: violations,
	})
	if err != nil {
		return statusInvalidArg.Err()
	}

	return statusDetails.Err()
}

func buildViolations(
	fieldErrors validator.FieldErrors,
	violations []*errdetails.BadRequest_FieldViolation,
	fieldPrefix string,
) []*errdetails.BadRequest_FieldViolation {
	for i := range fieldErrors {
		fieldError := fieldErrors[i]

		fieldName := fieldError.Name
		if fieldPrefix != "" {
			fieldName = fieldPrefix + "." + fieldName
		}

		if errs, ok := fieldError.Err.(validator.FieldErrors); ok { //nolint:errorlint
			return buildViolations(errs, violations, fieldName)
		}

		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       fieldName,
			Description: fieldError.Err.Error(),
		})
	}

	return violations
}
