package sentinels

import (
	"github.com/pkg/errors"
)

var ErrEmptyData = errors.New("not found")

func SearchSomething() error {
	// code here...

	return ErrEmptyData
}

func SearchWithWraping() error {

	//err from repository
	err := errors.New("the repository is failing")
	return errors.Wrap(err, "something happens with the repository")
}
