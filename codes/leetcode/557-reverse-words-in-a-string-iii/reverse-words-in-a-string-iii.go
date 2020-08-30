package main


func reverseWords(s string) string {
	ret := make([]byte, len(s))
	for last, i := 0, 0; i <= len(s); i++ {
		if i == len(s) {
			for j := last; j < i; j++ {
				ret[j] = s[i-(j-last)-1]
			}
		} else if s[i] == ' ' {
			for j := last; j < i; j++ {
				ret[j] = s[i-(j-last)-1]
			}
			ret[i] = ' '
			last = i + 1
		}
	}
	return string(ret)
}

func main() {

}
