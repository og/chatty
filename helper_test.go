package cha

import (
	glist "github.com/og/x/list"
	"testing"
)

func TestPickString(t *testing.T) {
	list := []string{"abc", "efg"}
	oneCount := 0
	twoCount := 0
	glist.Run(100, func(i int) (_break bool) {
		s := PickString(list)
		if s == "abc" {
			oneCount++
		}
		if s == "efg" {
			twoCount++
		}
		if s != "abc" && s !="efg" {
			t.Log("pickString error")
			t.Fail()
		}
		return
	})
	if twoCount == 0 {
		t.Log("pickString one error")
		t.Fail()
	}
	if oneCount == 0 {
		t.Log("pickString two error")
		t.Fail()
	}
}
