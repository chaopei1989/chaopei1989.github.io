package main

func lr(size int) (int, int) {
	if size%2 == 0 {
		return size/2 - 1, size / 2
	} else {
		return size/2 - 1, size/2 + 1
	}
}

func shortestPalindrome(str string) string {
	// 1. 以找中心为出发点, 找回前缀回文串
	for s := len(str); s >= 0; s-- {
		l, r := lr(s)
		for ; l >= 0 && str[l] == str[r]; {
			l--
			r++
		}
		if l < 0 {
			// ok
			if s < len(str) {
				restRev := make([]byte, 0)
				for i := len(str) - 1; i >= s; i-- {
					restRev = append(restRev, str[i])
				}
				return string(restRev) + str
			} else {
				return str
			}
		}
	}
	return ""
}

func main() {

}
