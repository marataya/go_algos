"""
125. Valid Palindrome
Solved
Easy
Topics
Companies
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

"""

import re
class Solution:
    def isPalindrome(self, s: str) -> bool:
        clean_s = re.sub('[^A-Za-z0-9]', '', s)
        return clean_s.lower() == clean_s[::-1].lower()
