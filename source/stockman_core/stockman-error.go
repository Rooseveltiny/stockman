package core

import "fmt"

/*
system error allows to declare system's error which occur during system performing
SystemError implements error interface
*/
type SystemError struct {
	message   string
	errorText string
}

func (se *SystemError) Error() string {
	if se.errorText != "" {
		return fmt.Sprintf("%s\n%s", se.message, se.errorText)
	} else {
		return se.message
	}
}

func (se *SystemError) AppendErrorText(t string) *SystemError {
	return &SystemError{message: se.message, errorText: t}
}

func NewSystemError(message string) *SystemError {
	return &SystemError{message: message, errorText: ""}
}
