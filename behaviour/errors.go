package behaviour

import (
	"github.com/pkg/errors"
)

type notFound struct {
	error
}

// NewNotFound returns an error which satisfies IsProductEmpty
func NewNotFound(format string, args ...interface{}) error {
	return &notFound{errors.Errorf(format, args...)}
}

// WrapNotFound returns an error which wraps err that satisfies IsNotFound
func WrapNotFound(err error, format string, args ...interface{}) error {
	return &notFound{errors.Wrapf(err, format, args...)}
}

// IsNotFound reports wheteher err was created with NewNotFound or WrapNotFound
func IsNotFound(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*notFound)
	return ok
}



// example

func SearchSomething() error {
	// code here...
	return NewNotFound("not found data for ID: %s", "anyID")
}

func SearchWithWraping() error {
	//err from repository
	err := errors.New("the repository is failing")
	return WrapNotFound(err, "not found data for ID: %s", "anyID")
}


