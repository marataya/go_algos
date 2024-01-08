/*
Given an integer x, return true if x is a palindrome, and false otherwise.
*/
package palindrome

import (
	"fmt"
	"go_algos/sliceutils"
)

func isPalindromeNumber(x int) bool {
	xS := fmt.Sprintf("%d", x)
	return xS == sliceutils.StringReverse(xS)
}
