package main

func rangeBitwiseAnd(m int, n int) int {
	mTmp := m
	nTmp := n
	shift := 0
	for {
		if mTmp != nTmp {
			shift++
			mTmp >>= 1
			nTmp >>= 1
			continue
		} else {
			break
		}
	}
	for i := 0; i < shift; i++ {
		mTmp <<= 1
	}
	return mTmp
}

func main() {

}
