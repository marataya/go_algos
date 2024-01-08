/*
125. Valid Palindrome

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise
*/
package palindrome

import (
	"go_algos/sliceutils"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	clean_s := regexp.MustCompile("[^A-Za-z0-9]").ReplaceAllString(s, "")
	//    if len(clean_s) == 0 {return false}
	return strings.ToLower(clean_s) == strings.ToLower(sliceutils.StringReverse(clean_s))
}
