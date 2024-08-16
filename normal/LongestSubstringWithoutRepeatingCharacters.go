package main

/*
Given a string s, find the length of the longest
substring
 without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.

*/

func main() {
	// test case
	println(solution("abcabcbb"))         // 3
	println(solution("bbbbb"))            // 1
	println(solution("pwwkew"))           // 3
	println(solution(" "))                // 1
	println(solution("au"))               // 2
	println(solution("dvdf"))             // 3
	println(solution("dvoikdfgzx"))       // 9
	println(solution("dvoikddfgzxop"))    // 7
	println(solution("dvodioutyjsdoiyu")) // 8
	println(solution("abcdefgbas"))       // 8
}

func solution(s string) int {
	n := len(s)
	ans := 0
	index := make(map[byte]int)
	left, right := 0, 0
	for right < n {
		char := s[right]
		if lastSeen, ok := index[char]; ok {
			left = max(lastSeen, left)
		}
		ans = max(ans, right-left+1)
		index[char] = right + 1
		right++
	}
	return ans
}
