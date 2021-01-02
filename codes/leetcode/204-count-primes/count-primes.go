package main

func countPrimes(n int) int {
	count := 0
	// 是否为合数
	signs := make([]bool, n)
	for i := 2; i < n; i++ {
		if !signs[i] {
			count++;
			for j := i + i; j < n; j += i {
				signs[j] = true;
			}
		}
	}
	return count
}
