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
	countMap := map[string]int{}
	cha.Run(100, func(i int) (_break bool) {
		countMap[cha.UUID()] = countMap[cha.UUID()] +1
		if countMap[cha.UUID()] >1 {
			t.Log("uuid repeat!")
			t.Fail()
		}
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
		is.Eql(gconv.IntString(id), userStringID.String())
		return
	})
}
func TestNameIncrID(t *testing.T) {
	is := gis.New(t)
	glist.Run(100, func(i int) (_break bool) {
		id := gconv.IntString(i+1)
		is.Eql(id, cha.NameIncrID("34gv43g43gv"))
		return
	})
}


type User struct {
	ID string `cha:"UUID()"`
	Son struct{
		ID string `cha:"NameIncrID(\"usersonid\")"`
		ID2 string `cha:"NameIncrID(\"usersonid\")"`
	}
}
type User2 struct {
	ID string
	Son struct{
		ID string
		ID2 string
	}
}
func (user *User2) Chatty () {
	user.ID = cha.UUID()
	user.Son.ID = cha.NameIncrID("user2")
	user.Son.ID2 = cha.NameIncrID("user2")
}

func TestMock(t *testing.T) {
	user := User{}
	cha.UnsafeMock(&user)
	is := gis.New(t)
	is.Eql(len(user.ID), 36)
	is.Eql(user.Son.ID, "1")
	is.Eql(user.Son.ID2, "2")
}
func TestSafeMock(t *testing.T) {
	user := User2{}
	cha.Mock(&user)
	is := gis.New(t)
	is.Eql(len(user.ID), 36)
	is.Eql(user.Son.ID, "1")
	is.Eql(user.Son.ID2, "2")
}