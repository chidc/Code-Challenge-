// Find Peak Element
// https://github.com/hoangst27/code-challenge/blob/main/code_challenge_3.md

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}