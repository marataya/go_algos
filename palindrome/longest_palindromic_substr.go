package palindrome

import (
	"bytes"
	"go_algos/utils"
)

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
			if utils.StringReverse(subStr) == subStr {
				if len(longestPalSubstr) < len(subStr) {
					longestPalSubstr = subStr
				}
			}
		}
	}
	return longestPalSubstr
}

/* O(n^2) solution */
func LongestPalindrome1(s string) string {
	p := []int{}
	if len(s) == 0 {
		return s
	}
	lLen := 2*len(s) + 1
	for i := 0; i < lLen; i++ {
		L := i / 2
		R := L + i%2
		for L > 0 && R < len(s) && s[L-1] == s[R] {
			L -= 1
			R += 1
		}
		p = append(p, R-L)
	}

	max_c := 0
	max_len := 0
	for i := 0; i < len(p); i++ {
		if max_len < p[i] {
			max_len = p[i]
			max_c = i
		}
	}
	return s[(max_c-max_len)/2 : (max_c+max_len)/2]
}

/* Manacher solution */
func ManacherAlgo(s string) string {
	buf := bytes.Buffer{}
	buf.WriteRune('#')
	for _, rune := range s {
		buf.WriteRune(rune)
		buf.WriteRune('#')
	}
	modified_s := buf.String()
	n := len(modified_s)
	p := make([]int, n)
	c, r := 0, 0

	max_len := 0
	max_c := 0

	for i := 0; i < n; i++ {
		if i < r {
			mirror := 2*c - i
			if r-i <= p[mirror] {
				p[i] = r - i
			} else {
				p[i] = p[mirror]
			}
		}
		a, b := i+(1+p[i]), i-(1+p[i])

		for a < n && b >= 0 && modified_s[a] == modified_s[b] {
			p[i]++
			a++
			b--
		}
		if i+p[i] > r {
			c, r = i, i+p[i]
		}
		if p[i] > max_len {
			max_len = p[i]
			max_c = i
		}
	}
	start := (max_c - max_len) / 2
	end := start + max_len
	return s[start:end]
}
