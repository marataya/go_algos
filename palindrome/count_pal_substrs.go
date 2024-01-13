/*
647. Palindromic Substrings
Medium
Given a string s, return the number of palindromic substrings in it. A string is a palindrome when it reads the same backward as forward. A substring is a contiguous sequence of characters within the string.

Example 1:

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
Example 2:

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

Constraints:
1 <= s.length <= 1000
s consists of lowercase English letters.
*/

package palindrome

var count = 0

func f(s string, i, j int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		count++
		i--
		j++
	}
}

func CountSubstrings(s string) int {
	count = 0
	if len(s) == 0 || len(s) == 1 {
		return 1
	}

	for i := 0; i < len(s); i++ {
		f(s, i, i)
		f(s, i, i+1)
	}
	return count
}
