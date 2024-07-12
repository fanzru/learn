package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	start, maxLength := 0, 0

	for i, char := range s {
		if lastIndex, found := charIndexMap[char]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		charIndexMap[char] = i
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

// func lengthOfLongestSubstring(s string) int {
// 	result := map[rune]bool{}
// 	var start, finish int = 0, 0
// 	maxLength := 0
// 	found := false
// 	for i, v := range s {
// 		finish = i
// 		if result[v] {
// 			if maxLength < finish-start {
// 				maxLength = finish - start
// 			}
// 			start = i
// 			found = true
// 		}

// 		result[v] = true
// 	}
// 	if found {
// 		return maxLength
// 	} else {
// 		return len(result)
// 	}

// }
