package cha_test

import (
	cha "github.com/og/go-chatty"
	glist "github.com/og/x/list"
	gis "github.com/og/x/test"
	"testing"
)

func TestPickString(t *testing.T) {
	list := []string{"abc", "efg"}
	oneCount := 0
	twoCount := 0
	glist.Run(100, func(i int) (_break bool) {
		s := cha.PickString(list)
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

func TestRun(t *testing.T) {
	{
		data := []int{}
		cha.Run(10, func(i int) (_break bool) {
			data = append(data, i)
			if i==5 {
				return true
			}
			return
		})
		gis.New(t).Eql(data, []int{0,1,2,3,4,5})
	}
	{
		data := []int{}
		cha.Run(10, func(i int) (_break bool) {
			data = append(data, i)
			return
		})
		gis.New(t).Eql(data, []int{0,1,2,3,4,5,6,7,8,9})
	}
}
