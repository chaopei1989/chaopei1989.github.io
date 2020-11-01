package main

func swap(nums []int, l, r int)  {
	if l == r {
		return
	}
	nums[l] ^= nums[r]
	nums[r] ^= nums[l]
	nums[l] ^= nums[r]
}

func sortColors(nums []int) {
	// 分割线, 也是交换的位置
	split01 := 0
	split12 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			swap(nums, i, split01)
			split01++
			split12++
		} else if nums[i] == 1 {
			swap(nums, i, split12)
			split12++
		} else { // 2
			// do nothing
		}
	}
}
