package test

import (
	"fmt"
	"github.com/Dzeqkon/ethclient/tools"
	"testing"
)

func TestToHex16(t *testing.T) {
	num := "0"
	rs := tools.ToHex16(num)
	fmt.Println(rs)
}

func TestPriKeyToAddress(t *testing.T) {
	priKey := "87ff9ec48300e8df51cdf5cf92197629f02b7b4c3f2c19b8a7882b41d9791a4e"
	accoount, fromKey, _ := tools.PriKeyToAddress(priKey)
	fmt.Println(accoount)
	fmt.Println(fromKey)
}
