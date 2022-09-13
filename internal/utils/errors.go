package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"strconv"
)

// NewError returns an error that formats as the given text.
func NewError(status int, message string) error {
	return &ServerError{
		Status:  status,
		Message: message,
	}
}

func NewValidationError(status int, message string, fields []FieldError) error {
	return &ValidationError{
		ServerError: ServerError{
			Status:  status,
			Message: message,
		},
		Fields: fields,
	}
}

type ServerError struct {
	Status  int
	Message string
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ValidationError struct {
	ServerError
	Fields []FieldError
}

func (e *ServerError) Error() string {
	return strconv.Itoa(e.Status) + " - " + e.Message
}

type ErrorResponse struct {
	Status    int          `json:"status"`
	Message   string       `json:"message"`
	RequestID string       `json:"request_id,omitempty"`
	Fields    []FieldError `json:"fields,omitempty"`
} //@name ErrorResponse

func ValidatePayload(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		log.Error().Err(err).Msg("validation failed")

		var fields []FieldError
		for _, field := range validationError {
			fields = append(fields, FieldError{
				Field: field.Field(),
				Error: field.Error(),
			})
		}

		return NewValidationError(fiber.StatusUnprocessableEntity, "validation failed", fields)
	}

	return nil
}
