package main

func sortString(s string) string {
	dict := make([]int, 26)
	for _, b := range s {
		dict[int(b-'a')]++
	}

	res := ""
	ap := func(b int) bool {
		if dict[b] > 0 {
			res += string('a' + byte(b))
			dict[b]--
			return true
		}
		return false
	}

	for {
		next := false
		for i := 0; i < len(dict); i++ {
			if ap(i) {
				next = true
			}
		}
		for j := len(dict)-1; j >= 0; j-- {
			if ap(j) {
				next = true
			}
		}
		if !next {
			break
		}
	}
	return res
}

func main() {
	sortString("eaimykhkgxr")
}