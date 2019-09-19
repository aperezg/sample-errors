package behaviour113

import (
	"errors"
	"fmt"
)

type notFound struct {
	error
}

// NewNotFound returns an error which satisfies IsProductEmpty
func NewNotFound(format string, args ...interface{}) error {
	return &notFound{fmt.Errorf(format, args...)}
}

// WrapNotFound returns an error which wraps err that satisfies IsNotFound
func WrapNotFound(err error, format string, args ...interface{}) error {
	args = append(args, err)

	return &notFound{fmt.Errorf(format+": %w", args...)}
}

// IsNotFound reports err was created with NewNotFound or WrapNotFound
func IsNotFound(err error) bool {
	var target *notFound
	return errors.As(err, &target)
}

// example

func SearchSomething() error {
	// code here...
	return NewNotFound("not found data for ID: %s", "anyID 113")
}

func SearchWithWraping() error {
	//err from repository
	err := errors.New("the repository is failing")
	return WrapNotFound(err, "not found data for ID: %s", "anyID 113")
}
