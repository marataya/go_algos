package palindrome

import (
	"algos/sliceutils"
	"fmt"
)

func isPalindromeNumber(x int) bool {
	xS := fmt.Sprintf("%d", x)
	return xS == sliceutils.StringReverse(xS)
}
