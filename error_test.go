package cha

import (
	gtest "github.com/og/x/test"
	"testing"
)

func TestNewError(t *testing.T) {
	gtest.NewAS(t).ErrorString(newError("abc"), "go-chatty: abc")
}
