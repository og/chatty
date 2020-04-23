package cha

import "github.com/pkg/errors"

func newError(msg string) error {
	return errors.New("go-chatty: "+ msg)
}
