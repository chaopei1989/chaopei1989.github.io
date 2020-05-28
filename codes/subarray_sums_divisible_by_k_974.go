/**
 * 给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。
 * 输入：A = [4,5,0,-2,-3,1], K = 5
 * 输出：7
 *
 * 有 7 个子数组满足其元素之和可被 K = 5 整除：
 * [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
**/
package main

import "fmt"

// 核心函数
func subarraysDivByK(A []int, K int) int {
	// 1. 创建前缀和数组 preSums
	preSums := createPreSums(A)
	// 2. 此时i到j的数组之和即: preSums[j]-preSums[i-1], 其中i=0时, 可以指定preSums[-1]=0
	// 判断所有的i和j的组合, 让(preSums[j]-preSums[i-1])%K=0
	// 即让preSums[j]%K=preSums[i-1]%K
	// 遍历preSums数组计算上述表达式,
	kModMap := make(map[int]int)
	// 把preSums[-1]加进去
	kModMap[0]++
	for _, sum := range preSums {
		modResult := sum % K
		// 累加该mod的值
		kModMap[modResult]++
	}
	// 3. kModMap每个元素都对应该mod下的前缀数的个数, 统计组合数即可
	var result int
	for k, v := range kModMap {
		fmt.Printf("%d:%d ", k, v)
		result += (v * (v - 1) / 2)
	}
	return result
}

// 创建前缀和数组, 数组下标若为i, 即0~i之和的数组
func createPreSums(A []int) []int {
	var preSums []int = make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if i == 0 {
			// 首值
			preSums[0] = A[0]
		} else {
			// 后续的值用邻接前值累加
			preSums[i] = preSums[i-1] + A[i]
		}
	}
	return preSums
}

// 测试
func main() {
	result := subarraysDivByK([]int{
		-1, 2, 9,
	}, 2)
	fmt.Println(result)
}
