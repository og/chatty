package cha

import (
	"crypto/rand"
	ge "github.com/og/x/error"
	grand "github.com/og/x/rand"
	"math/big"
)

// equal TrueLikelihood(50)
func Bool() bool{
	return TrueLikelihood(50)
}
// It panics if likelihood < 0 or likelihood > 100 .
func TrueLikelihood(likelihood int) bool {
	if likelihood < 0 {
		panic(newError("BoolLikelihood(likelihood int) likelihood can not less than 0%"))
	}
	if likelihood > 100 {
		panic(newError("BoolLikelihood(likelihood int) likelihood can not greater than 100%"))
	}
	if likelihood == 0 {
		return false
	}
	return bool(Int(1, 100) <= likelihood )
}

func String(min int, max int) string {
	size := Int(min, max)
	return grand.StringBySeed("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]",size)
}
func Letter(size int) string {
	return grand.StringBySeed("abcdefghijklmnopqrstuvwxyz",size)
}
func CapitalLetter(size int) string {
	return grand.StringBySeed("ABCDEFGHIJKLMNOPQRSTUVWXYZ",size)
}

func randomBig(max int) *big.Int {
	random, err := rand.Int(rand.Reader, big.NewInt(int64(max))) ; ge.Check(err)
	return random
}
func Int(min int, max int) int {
	if min > max {
		min, max = max, min
		WarningLog("Int(min int, max int) min can not greater than max")
	}
	if min == max { return max }
	rangeValue := max - min + 1
	random := randomBig(rangeValue)
	return int(random.Int64()) + min
	// min 6 max 6
	// return 6

	// min 0 max 6
	// random: 0 ~ 6(6-0)
	// return 6 + 0 = 6

	// min 4 max 10
	// random: 0 ~ 6(10-4)
	// return 4 + 4 = 8

	// min -2 max 4
	// random: 0 ~ 6(4-(-2))
	// return 1 + -2 = -1


}