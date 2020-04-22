package cha_test

import (
	cha "github.com/og/go-chatty"
	ge "github.com/og/x/error"
	gis "github.com/og/x/test"
	"regexp"
	"testing"
)

func TestLetter(t *testing.T) {
	is := gis.New(t)
	is.Eql(len(cha.Letter(10)), 10)
	is.False(ge.Bool(regexp.MatchString(`[^a-z]`, cha.Letter(10000))))
}
type MockLetter struct {
	Name string `cha:"Letter(10)"`
}
func TestUnsafeMockLetter(t *testing.T) {
	is := gis.New(t)
	v := MockLetter{}
	cha.UnsafeMock(&v)
	is.Eql(len(v.Name), 10)
	is.False(ge.Bool(regexp.MatchString(`[^a-z]`, v.Name)))
}