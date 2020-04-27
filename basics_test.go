package cha_test

import (
	cha "github.com/og/go-chatty"
	gconv "github.com/og/x/conv"
	ge "github.com/og/x/error"
	"github.com/og/x/test"
	"log"
	"regexp"
	"testing"
)

func coreTestInt(t *testing.T,min int, max int, rangeList []int) {
	as:=gtest.AS(t)
	list := []int{}
	cha.Run(1000, func(i int) (_break bool) {
		list = append(list, cha.Int(min, max))
		return
	})
	for _, item := range list {
		as.True(item >=min && item <= max)
	}
	for i:=0;i< len(rangeList);i++ {
		number := rangeList[i]
		foundIt := false
		for _, item := range list {
			if item == number {
				foundIt = true
			}
		}
		if !foundIt {
			log.Print(" can not found " + gconv.IntString(number))
			log.Print(min, max)
			log.Print(list)
			t.Fail()
		}
	}
	{
		v := struct {
			Int int `cha:"Int(1,4)"`
		}{}
		cha.UnsafeMock(&v)
		as.True(v.Int >=1  && v.Int <= 4)
	}
}
func TestInt(t *testing.T) {
	coreTestInt(t,-3, 6, []int{-3,-2,-1,0,1,2,3,4,5,6})
	coreTestInt(t,-2, 6, []int{-2,-1,0,1,2,3,4,5,6})
	coreTestInt(t,-1, 6, []int{-1,0,1,2,3,4,5,6})
	coreTestInt(t, 0, 6, []int{0,1,2,3,4,5,6})
	coreTestInt(t, 1, 6, []int{1,2,3,4,5,6})
	coreTestInt(t, 2, 6, []int{2,3,4,5,6})
	coreTestInt(t, 3, 6, []int{3,4,5,6})
	coreTestInt(t, 4, 6, []int{4,5,6})
	coreTestInt(t, 5, 6, []int{5,6})
	coreTestInt(t, 6, 6, []int{6})
}
func coreTestBool(t *testing.T, likelihood int, trueCount int,falseCount int) {
	as := gtest.AS(t)
	as.Eql(10000, trueCount + falseCount)
	if trueCount < likelihood*100 && trueCount > likelihood*100 {
		t.Log("trueCount", trueCount, " overflow normal range")
		t.Fail()
	}

}
func TestBool(t *testing.T) {
	as := gtest.AS(t)
	{
		trueCount := 0
		falseCount := 0
		cha.Run(10000, func(i int) (_break bool) {
			if cha.Bool() { trueCount++ } else { falseCount++ }
			return
		})
		coreTestBool(t, 50, trueCount, falseCount)
	}
	{
		hasTrue := false
		cha.Run(100, func(i int) (_break bool) {
			v := struct {
				Bool bool `cha:"Bool()"`
			}{}
			cha.UnsafeMock(&v)
			if v.Bool {
				hasTrue = true
			}
			return
		})
		as.True(hasTrue)
	}
}
func TestTrueLikelihood(t *testing.T) {
	{
		trueCount := 0
		falseCount := 0
		cha.Run(10000, func(i int) (_break bool) {
			if cha.TrueLikelihood(0) { trueCount++ } else { falseCount++ }
			return
		})
		coreTestBool(t, 0, trueCount, falseCount)
	}
	{
		trueCount := 0
		falseCount := 0
		cha.Run(10000, func(i int) (_break bool) {
			if cha.TrueLikelihood(10) { trueCount++ } else { falseCount++ }
			return
		})
		coreTestBool(t, 10, trueCount, falseCount)
	}
	{
		trueCount := 0
		falseCount := 0
		cha.Run(10000, func(i int) (_break bool) {
			if cha.TrueLikelihood(20) { trueCount++ } else { falseCount++ }
			return
		})
		coreTestBool(t, 20, trueCount, falseCount)
	}
	{
		trueCount := 0
		falseCount := 0
		cha.Run(10000, func(i int) (_break bool) {
			if cha.TrueLikelihood(100) { trueCount++ } else { falseCount++ }
			return
		})
		coreTestBool(t, 100, trueCount, falseCount)
	}

	{
		trueCount := 0
		falseCount := 0
		cha.Run(10000, func(i int) (_break bool) {
			v := struct {
				Bool bool `cha:"TrueLikelihood(40)"`
			}{}
			cha.UnsafeMock(&v)
			if v.Bool {
				trueCount++
			} else {
				falseCount++
			}
			return
		})
		coreTestBool(t, 40, trueCount, falseCount)
	}
}

func TestLetter(t *testing.T) {
	as := gtest.AS(t)
	as.Eql(len(cha.Letter(10)), 10)
	as.False(ge.Bool(regexp.MatchString(`[^a-z]`, cha.Letter(10000))))
}
func TestCapitalLetter(t *testing.T) {
	as := gtest.AS(t)
	as.Eql(len(cha.CapitalLetter(10)), 10)
	as.False(ge.Bool(regexp.MatchString(`[^A-Z]`, cha.CapitalLetter(10000))))
}
type MockLetter struct {
	Name string `cha:"Letter(10)"`
	Title string `cha:"CapitalLetter(10)"`
}
func TestUnsafeMockLetter(t *testing.T) {
	as := gtest.AS(t)
	v := MockLetter{}
	cha.UnsafeMock(&v)
	as.Eql(len(v.Name), 10)
	as.False(ge.Bool(regexp.MatchString(`[^a-z]`, v.Name)))
	as.Eql(len(v.Title), 10)
	as.False(ge.Bool(regexp.MatchString(`[^A-Z]`, v.Title)))
}

func TestString(t *testing.T) {
	as := gtest.AS(t)
	as.Run(100, func(i int) (_break bool) {
		s := cha.String(1,5)
		as.Range(len(s), 1, 5)
		for i:=0;i<len(s);i++ {
			v := s[i]
			as.StringContains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]",string(v))
		}
		return
	})
}