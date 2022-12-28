package errors

import (
	"bytes"
	"strings"
)

type ViolationErrors []ViolationError

func (v ViolationErrors) Error() string {
	buff := bytes.NewBufferString("")

	for _, violation := range v {
		buff.WriteString(violation.Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// ViolationError represent all information about the violation.
type ViolationError struct {
	PropertyPath string `json:"propertyPath"`
	Message      string `json:"message"`
	Code         string `json:"code"`
}

func (vi ViolationError) Error() string {
	return vi.Message
}
