package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

func New(message string) error {
	return customError{errorCode: UNKNOW, originError: errors.New(message)}
}

func NewFormat(message string, args ...interface{}) error {
	formatMessage := fmt.Sprintf(message, args)

	return customError{errorCode: UNKNOW, originError: errors.New(formatMessage)}
}

func Wrap(err error, message string) error {
	return WrapFormat(err, message)
}

func Cause(err error) error {
	return errors.Cause(err)
}

func WrapFormat(err error, message string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, message, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorCode:   customErr.errorCode,
			originError: customErr.originError,
		}
	}

	return customError{errorCode: UNKNOW, originError: wrappedError}
}
