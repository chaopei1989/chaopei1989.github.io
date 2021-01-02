package main

func searchRange(nums []int, target int) []int {
	return searchRangeFT(nums, 0, len(nums)-1, target)
}

func searchRangeFT(nums []int, from, to, target int) []int {
	if from > to {
		return []int{-1,-1}
	}
	half := (to - from + 1) / 2 + from
	if nums[half] < target {
		return searchRangeFT(nums, half+1, to, target)
	} else if nums[half] > target {
		return searchRangeFT(nums, from, half-1, target)
	} else {
		var retF, retT, tmp int
		for tmp = half-1; tmp >= 0; tmp-- {
			if nums[tmp] != target {
				retF = tmp + 1
				break
			}
		}
		if tmp < 0 {
			retF = 0
		}
		for tmp = half+1; tmp < len(nums); tmp++ {
			if nums[tmp] != target {
				retT = tmp - 1
				break
			}
		}
		if tmp >= len(nums) {
			retT = len(nums)-1
		}
		return []int{retF, retT}
	}
}

func main() {
	searchRange([]int{1,3}, 1)
}
