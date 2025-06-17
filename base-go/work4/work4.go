package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	shortest := len(strs[0])
	var result string = ""
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < shortest {
			shortest = len(strs[i])
		}
	}
	for i := 0; i < shortest; i++ {
		temp := ""
		for j := 0; j < len(strs); j++ {
			temp += string(strs[j][i])
		}
		if allCharsEqual(temp) {
			result += string(temp[0])
		} else {
			break
		}
	}
	return result
}
func allCharsEqual(s string) bool {
	if len(s) == 0 {
		return true // 空字符串可视为所有字符“相等”
	}

	first := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != first {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
