package palindrome

import "go_algos/sliceutils"

/*5. Longest Palindromic Substring
Given a string s, return the longest palindromic substring in s.
*/

// Brute force algorithm complexity O(n^3)
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	longestPalSubstr := s[0:1]
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s)+1; j++ {
			subStr := s[i:j]
			if sliceutils.StringReverse(subStr) == subStr {
				if len(longestPalSubstr) < len(subStr) {
					longestPalSubstr = subStr
				}
			}
		}
	}
	return longestPalSubstr
}

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	longestPalSubstr := s[0:1]
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s)+1; j++ {
			subStr := s[i:j]
			if sliceutils.StringReverse(subStr) == subStr {
				if len(longestPalSubstr) < len(subStr) {
					longestPalSubstr = subStr
				}
			}
		}
	}
	return longestPalSubstr
}
