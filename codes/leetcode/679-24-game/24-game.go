package main
const (
	PLUS   = 0
	MINUS  = 1
	MULTI  = 2
	DIVIDE = 3
)

func isFloatEqual(n1, n2 float64) bool {
	if n1-n2 < 0.000001 && n1-n2 > -0.000001 {
		return true
	}
    return false
}

func judgePoint24Float64(nums []float64) bool {
	currLen := len(nums)
	for i := 0; i < currLen; i++ {
		for j := 0; j < currLen; j++ {
			if i == j {
				continue
			}
			for cal := 0; cal < 4; cal++ {
				next := make([]float64, 0)
				for m := 0; m < currLen; m++ {
					if m == i || m == j {
						continue
					}
					next = append(next, nums[m])
				}
				var res float64
				switch cal {
				case PLUS:
					res = nums[i] + nums[j]
					break
				case MINUS:
					res = nums[i] - nums[j]
					break
				case MULTI:
					res = nums[i] * nums[j]
					break
				case DIVIDE:
					res = nums[i] / nums[j]
					break
				}
				next = append(next, res)
				if len(next) == 1 && isFloatEqual(next[0], 24) {
					return true
				} else if judgePoint24Float64(next) {
					return true
				}
			}
		}
	}
	return false
}

func judgePoint24(nums []int) bool {
	var numsFloat []float64 = make([]float64, len(nums))
	for i := 0; i < len(nums); i++ {
		numsFloat[i] = float64(nums[i])
	}
	return judgePoint24Float64(numsFloat)
}

func main() {

}
