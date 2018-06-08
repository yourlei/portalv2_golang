package test

import (
	"testing"
	"fmt"
	"strconv"
)

func TestStr2Int(t *testing.T) {
	// s := "128"
	// n := 0
	// for i := 0; i < len(s); i++ {
	// 	n *= 10	+ s[i] 	// base
	// 	fmt.Println(n)
	// }

	n, err := strconv.ParseInt("fg", 10, 8)

	if err != nil {
		t.Error(err, n)
	}
	fmt.Println(n)
}