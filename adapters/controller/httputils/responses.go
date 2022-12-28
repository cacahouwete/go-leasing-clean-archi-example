package httputils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/alexandrevinet/leasing/business/errors"
)

// ResponseCollection struct used to return a collection of any entities.
type ResponseCollection[T interface{}] struct {
	Member     []T  `json:"member"`
	TotalItems uint `json:"totalItems"`
}

// ResponseError struct used to return any kind of error.
type ResponseError struct {
	Error string `json:"error" example:"message"`
}

// ResponseViolations struct used to return collection of violations
// A violation something that bloc the process like missing field, incorrect value, business logic constraint.
type ResponseViolations struct {
	ResponseError
	Violations errors.ViolationErrors `json:"violations"`
}

// ErrorResponse will convert the error in corresponding http code and body.
// By default it respond internal server error with error message.
func ErrorResponse(c *gin.Context, err error) {
	switch err.(type) { //nolint:errorlint // false positive
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError{err.Error()})
	case validator.ValidationErrors:
		ConverValidationErrorsToViolationResponse(c, err)
	case errors.ViolationError:
		ConverViolationErrorToViolationResponse(c, err)
	case errors.ViolationErrors:
		ConverViolationErrorsToViolationResponse(c, err)
	}
}

// NotFoundResponse function will abort with code 404 and return a response error with message not found.
func NotFoundResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, ResponseError{"Not Found."})
}

// BadRequestResponse function will abort with code 400 and return a response error with message request not processable.
func BadRequestResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, ResponseError{"Request not processable."})
}

// ViolationResponse function will abort with code 422 and return a response violations with given violations.
func ViolationResponse(c *gin.Context, violations errors.ViolationErrors) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ResponseViolations{
		ResponseError: ResponseError{
			Error: "Payload is unprocessable du to some violations",
		},
		Violations: violations,
	})
}

func ConverValidationErrorsToViolationResponse(c *gin.Context, err error) {
	errVal, ok := err.(validator.ValidationErrors) //nolint:errorlint // error check with ok
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError{"validatation error not transformable"})

		return
	}

	result := make([]errors.ViolationError, 0, len(errVal))

	for _, fieldError := range errVal {
		result = append(result, errors.ViolationError{
			PropertyPath: fieldError.Field(),
			Message:      fieldError.Error(),
			Code:         fieldError.Tag(),
		})
	}

	ViolationResponse(c, result)
}

func ConverViolationErrorToViolationResponse(c *gin.Context, err error) {
	errVal, ok := err.(errors.ViolationError) //nolint:errorlint // error check with ok
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError{"validatation error not transformable"})

		return
	}

	ViolationResponse(c, errors.ViolationErrors{errVal})
}

func ConverViolationErrorsToViolationResponse(c *gin.Context, err error) {
	errVal, ok := err.(errors.ViolationErrors) //nolint:errorlint // error check with ok
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError{"validatation error not transformable"})

		return
	}

	ViolationResponse(c, errVal)
}
