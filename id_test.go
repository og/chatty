package cha_test

import (
	"github.com/og/go-chatty"
	gconv "github.com/og/x/conv"
	ge "github.com/og/x/error"
	glist "github.com/og/x/list"
	gis "github.com/og/x/test"
	"regexp"
	"testing"
)
func TestUUID(t *testing.T) {
	is := gis.New(t)
	glist.Run(100, func(i int) (_break bool) {
		is.Eql(len(cha.UUID()), 36)
		is.True(ge.Bool(regexp.MatchString("[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}", cha.UUID())))
		return
	})
}
func TestIncrID(t *testing.T) {
	is := gis.New(t)
	userIncrID := cha.IncrID()
	userStringID := cha.IncrID()
	glist.Run(100, func(i int) (_break bool) {
		id := i+1
		is.Eql(id, userIncrID.Int())
		return
	})
	glist.Run(100, func(i int) (_break bool) {
		id := i+1
		is.Eql(string(id), userStringID.String())
		return
	})
}
func TestNameIncrID(t *testing.T) {
	is := gis.New(t)
	glist.Run(100, func(i int) (_break bool) {
		id := gconv.IntString(i+1)
		is.Eql(id, cha.NameIncrID("user"))
		return
	})
}