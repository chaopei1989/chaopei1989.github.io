package main

func a2i(c byte) int {
	return int(c - '0')
}

func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	for ki := 0; ki < k; ki++ {
		lastMin := -1
		i := 0
		for ; i < len(num); i++ {
			if lastMin < 0 || lastMin <= a2i(num[i]) {
				lastMin = a2i(num[i])
			} else {
				// 首个降序
				break
			}
		}
		// 取 [0:i-1] + [i+1:]
		if i >= len(num) {
			num = num[0:i-1]
		} else {
			num = num[0:i-1] + num[i:]
		}
	}
	ret := num

	// 去掉首0
	i := 0
	for ; i < len(ret); i++ {
		if ret[i] != '0' {
			break
		}
	}
	// 全0
	if i == len(ret) {
		return "0"
	}
	return ret[i:]
}

func main() {
	//log.Print(removeKdigits("1432219", 3))
	//log.Print(removeKdigits("10200", 1))
	//log.Print(removeKdigits("10", 1))
	//log.Print(removeKdigits("9", 1))
	//log.Print(removeKdigits("112", 1))
	//log.Print(removeKdigits("1234567890", 9))
	//log.Print(removeKdigits("1173", 2))
	a:=wrapperResponser{}
	a.Responser.OnData([]byte{1,2,3})
}

type Responser interface {
	OnData(data []byte)
}

type wrapperResponser struct {
	Responser
}

//func (t *wrapperResponser) OnData(data []byte)  {
//
//}