package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ztrue/tracerr"
)

type CustomError struct {
	Err     error
	Message string
	Code    int
}

type CustomErrorJSON struct {
	Message         string   `json:"message"`
	ErrorMessage    string   `json:"error_message"`
	Stack           []string `json:"stack"`
}

func (e *CustomError) WithMessage(message string) *CustomError {
	e.Message = message
	return e
}

func (e *CustomError) WithCode(statusCode int) *CustomError {
	e.Code = statusCode
	return e
}

func (e *CustomError) WithError(err error) *CustomError {
	if customError, ok := err.(*CustomError); ok {
		// Reuse the custom error instead of recreating it
		// in order to keep the stack trace.
		err = customError.Err
	} else {
		err = tracerr.Wrap(err)
	}

	e.Err = err
	return e
}

func (e *CustomError) Sanitize() *CustomError {
	if e.Code == 0 {
		e.WithCode(http.StatusInternalServerError)
	}
	if e.Message == "" {
		e.WithMessage("Unhandled server error")
	}

	return e
}

func (e CustomError) Error() string {
	return e.Message
}

func (e *CustomError) GetStackTrace() []string {
	rawStackTrace := tracerr.StackTrace(e.Err)
	stackTrace := parseStackTrace(rawStackTrace)

	return stackTrace
}

func (e *CustomError) ToJSON() CustomErrorJSON {
	return CustomErrorJSON{
		Message:        e.Message,
		ErrorMessage:   e.Err.Error(),
		Stack:          e.GetStackTrace(),
	}
}

func (e *CustomError) SendErrorResponse(c echo.Context) error {
	e.Sanitize()

	rawStackTrace := tracerr.StackTrace(e.Err)
	stackTrace := parseStackTrace(rawStackTrace)

	var errorMessage string
	if e.Err != nil {
		errorMessage = e.Err.Error()
	}

	response := ErrorResponse{
		Message:      e.Message,
		ErrorMessage: errorMessage,
		Stack:        stackTrace,
	}

	return c.JSON(e.Code, response)
}