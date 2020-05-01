package cha

import (
	glist "github.com/og/x/list"
	"github.com/og/x/test"
	"strings"
	"testing"
)

func TestCFirstName(t *testing.T) {
	nameMap := map[string]int{}
	Run(100, func(i int) (_break bool) {
		name := CFirstName()
		nameMap[name]++
		return
	})
	testPickString(t, seed.ChineseFirstName, nameMap)
}
func TestFirstName(t *testing.T) {
	nameMap := map[string]int{}
	Run(100, func(i int) (_break bool) {
		name := FirstName()
		nameMap[name]++
		return
	})
	testPickString(t, seed.FirstName, nameMap)
}

func TestCLastName(t *testing.T) {
	nameMap := map[string]int{}
	Run(100, func(i int) (_break bool) {
		name := CLastName()
		nameMap[name]++
		return
	})
	testPickString(t, seed.ChineseLastName, nameMap)
}

func TestLastName(t *testing.T) {
	nameMap := map[string]int{}
	Run(100, func(i int) (_break bool) {
		name := LastName()
		nameMap[name]++
		return
	})
	testPickString(t, seed.LastName, nameMap)
}

func TestName(t *testing.T) {
	as := gtest.NewAS(t)
	nameMap := map[string]int{}
	Run(100, func(i int) (_break bool) {
		name := Name()
		nameMap[name]++
		return
	})
	as.True(len(nameMap)!=0)
	for name, _:= range nameMap {
		names := strings.Split(name, " ")
		firstName := names[0]
		lastName := names[1]
		as.True(glist.StringList{seed.FirstName}.In(firstName))
		as.True(glist.StringList{seed.LastName}.In(lastName))
	}
}

func TestCName(t *testing.T) {
	as := gtest.NewAS(t)
	nameMap := map[string]int{}
	Run(100, func(i int) (_break bool) {
		name := CName()
		nameMap[name]++
		return
	})
	as.True(len(nameMap)!=0)
	as.True(len(nameMap) > 10)
}

func TestFullName(t *testing.T) {
	as := gtest.NewAS(t)
	nameMap := map[string]int{}
	Run(100, func(i int) (_break bool) {
		name := FullName()
		nameMap[name]++
		return
	})
	as.True(len(nameMap)!=0)
	for name, _:= range nameMap {
		names := strings.Split(name, " ")
		firstName := names[0]
		middleName := names[1]
		lastName := names[2]
		as.True(glist.StringList{seed.FirstName}.In(firstName))
		as.True(glist.StringList{seed.MiddleName}.In(middleName))
		as.True(glist.StringList{seed.LastName}.In(lastName))
	}
}

type mockNames struct {
	FirstName string `cha:"FirstName()"`
	LastName string `cha:"LastName()"`
	Name string `cha:"Name()"`
	FullName string `cha:"FullName()"`
	CFirstName string `cha:"CFirstName()"`
	CLastName string `cha:"CLastName()"`
	CName string `cha:"CName()"`
}
func TestUnsafeMockNames(t *testing.T) {
	as := gtest.NewAS(t)
	Run(100, func(i int) (_break bool) {
		v := mockNames{}
		UnsafeMock(&v)
		as.True(glist.StringList{seed.FirstName}.In(v.FirstName))
		as.True(glist.StringList{seed.LastName}.In(v.LastName))
		as.True(glist.StringList{seed.ChineseFirstName}.In(v.CFirstName))
		as.True(glist.StringList{seed.ChineseLastName}.In(v.CLastName))
		{
			names := strings.Split(v.Name, " ")
			as.True(glist.StringList{seed.FirstName}.In(names[0]))
			as.True(glist.StringList{seed.LastName}.In(names[1]))
		}
		{
			names := strings.Split(v.FullName, " ")
			as.True(glist.StringList{seed.FirstName}.In(names[0]))
			as.True(glist.StringList{seed.MiddleName}.In(names[1]))
			as.True(glist.StringList{seed.LastName}.In(names[2]))
		}
		{
			as.True(len(v.CName) > 2)
		}
		return
	})
	// as.Equal()

}