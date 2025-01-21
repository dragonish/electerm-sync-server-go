package logger

import (
	"errors"
	"os"
)

// Debug outputs debug level log.
func Debug(msg string, args ...any) {
	instance.Debug(msg, args...)
}

// Info outputs information level log.
func Info(msg string, args ...any) {
	instance.Info(msg, args...)
}

// Warn outputs warning level log.
func Warn(msg string, args ...any) {
	instance.Warn(msg, args...)
}

// Warn outputs warning level log with a err field.
func WarnWithErr(msg string, err error, args ...any) {
	slice := make([]any, 0)
	slice = append(slice, "err", ErrMsg(err))
	slice = append(slice, args...)
	instance.Warn(msg, slice...)
}

// Err outputs error level log.
//
// Returns new error with msg.
func Err(msg string, err error, args ...any) error {
	slice := make([]any, 0)
	slice = append(slice, "err", ErrMsg(err))
	slice = append(slice, args...)
	instance.Error(msg, slice...)

	return errors.New(msg)
}

// Panic outputs error level log and causes panic.
func Panic(msg string, err error, args ...any) {
	slice := make([]any, 0)
	slice = append(slice, "severity", "Panic")
	slice = append(slice, args...)
	aErr := Err(msg, err, slice...)

	panic(aErr)
}

// Fatal outputs error level log and exit the program.
func Fatal(msg string, err error, args ...any) {
	slice := make([]any, 0)
	slice = append(slice, "severity", "Fatal")
	slice = append(slice, args...)
	Err(msg, err, slice...)

	os.Exit(1)
}

// NewErr outputs error level log and returns new error with err.
func NewErr(errMsg string, args ...any) error {
	slice := make([]any, 0)
	slice = append(slice, args...)
	instance.Error(errMsg, slice...)

	return errors.New(errMsg)
}
