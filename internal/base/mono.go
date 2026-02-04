package base

func Mono(nums []int) bool {
	asc, desc := true, true
	for i := 0; i < len(nums)-1; i++ {
		if asc && nums[i] > nums[i+1] {
			asc = false
		}
		if desc && nums[i] < nums[i+1] {
			desc = false
		}
	}
	return asc || desc
}
