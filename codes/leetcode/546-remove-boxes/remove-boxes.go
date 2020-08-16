/*546. 移除盒子
题目链接: https://leetcode-cn.com/problems/remove-boxes/

给出一些不同颜色的盒子，盒子的颜色由数字表示，即不同的数字表示不同的颜色。
你将经过若干轮操作去去掉盒子，直到所有的盒子都去掉为止。每一轮你可以移除具有相同颜色的连续 k 个盒子（k >= 1），这样一轮之后你将得到 k*k 个积分。
当你将所有盒子都去掉之后，求你能获得的最大积分和。



示例：

输入：boxes = [1,3,2,2,2,3,4,3,1]
输出：23
解释：
[1, 3, 2, 2, 2, 3, 4, 3, 1]
----> [1, 3, 3, 4, 3, 1] (3*3=9 分)
----> [1, 3, 3, 3, 1] (1*1=1 分)
----> [1, 1] (3*3=9 分)
----> [] (2*2=4 分)


提示：

1 <= boxes.length <= 100
1 <= boxes[i] <= 100
*/
package main

func removeBoxes(boxes []int) int {
	dp := [100][100][100]int{}
	var calculatePoints func(boxes []int, l, r, k int) int
	calculatePoints = func(boxes []int, l, r, k int) int {
		if l > r {
			return 0
		}
		if dp[l][r][k] != 0 {
			return dp[l][r][k]
		}
		for r > l && boxes[r] == boxes[r-1] {
			r--
			k++
		}
		dp[l][r][k] = calculatePoints(boxes, l, r-1, 0) + (k+1)*(k+1)
		for i := l; i < r; i++ {
			if boxes[i] == boxes[r] {
				dp[l][r][k] = max(dp[l][r][k], calculatePoints(boxes, l, i, k+1)+calculatePoints(boxes, i+1, r-1, 0))
			}
		}
		return dp[l][r][k]
	}
	return calculatePoints(boxes, 0, len(boxes)-1, 0)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
