package main

/*****************************************************************
 *
 * 		回文中心的解法
 *
 ****************************************************************/

func countSubstrings(s string) int {
	length := len(s)
	count := length
	// 回文中心有 2*length-1 个
	for i := 0; i < 2*length-1; i++ {
		for left, right := i/2-(1-i%2), i/2+1; left >= 0 && right < length; {
			if s[left] == s[right] {
				count++
				left--
				right++
			} else {
				break
			}
		}
	}
	return count
}

func main() {

}
