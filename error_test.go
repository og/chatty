package cha

import (
	gtest "github.com/og/x/test"
	"testing"
)

func TestNewError(t *testing.T) {
	gtest.AS(t).EqualError(newError("abc"), "go-chatty: abc")
}
