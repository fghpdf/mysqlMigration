package errorEnum

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	UNKNOW = ErrorCode(iota)
	FILE_TYPE_INVALID
)

type ErrorCode uint

type customError struct {
	errorCode   ErrorCode
	originError error
}

func (error customError) Error() string {
	return error.originError.Error()
}

// generate a new custom error
func (code ErrorCode) New(message string) error {
	return customError{errorCode: code, originError: errors.New(message)}
}

// format error message and generate a new custom error
func (code ErrorCode) NewFormat(message string, args ...interface{}) error {
	formatError := fmt.Errorf(message, args)

	return customError{errorCode: code, originError: formatError}
}

// wrap a new error
func (code ErrorCode) Wrap(err error, message string) error {
	return code.WrapFormat(err, message)
}

// format error message and wrap a new error
func (code ErrorCode) WrapFormat(err error, message string, args ...interface{}) error {
	formatError := errors.Wrapf(err, message, args)

	return customError{errorCode: code, originError: formatError}
}
