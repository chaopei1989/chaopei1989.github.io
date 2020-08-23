package main

func repeatedSubstringPattern(s string) bool {
	strLen := len(s)
	// i 代表每个字符串的字数
	for i := 1; i <= strLen/2; i++ {
		if judge(s, i) {
			return true
		}
	}
	return false
}

func judge(s string, i int) bool {
	strLen := len(s)
	if strLen%i == 0 {
		// m 代表遍历字符串内的每个字符
		for m := 0; m < i; m++ {
			tmp := s[m]
			// j 代表重复的次数
			for j := 0; j < strLen/i; j++ {
				if s[j*i+m] != tmp {
					return false
				}
			}
		}
		return true
	}
	return false
}

func main() {

}
