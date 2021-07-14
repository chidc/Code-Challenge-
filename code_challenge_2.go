// Maximum Length of Repeated Subarray
//https://github.com/hoangst27/code-challenge/blob/main/code_challenge_2.md

func findLength(nums1 []int, nums2 []int) int {
	maxLength := 0
	for i := 0; i < len(nums1); i++ {
		indexA := i
		indexB := 0
		current := 0
		for indexA < len(nums1) && indexB < len(nums2) {
			if nums1[indexA] == nums2[indexB] {
				current++
				if current > maxLength {
					maxLength = current
				}
			} else {
				current = 0
			}
			indexA++
			indexB++
		}
	}

	for i := 0; i < len(nums2); i++ {
		indexA := 0
		indexB := i
		current := 0
		for indexA < len(nums1) && indexB < len(nums2) {
			if nums1[indexA] == nums2[indexB] {
				current++
				if current > maxLength {
					maxLength = current

				}
			} else {
				current = 0
			}
			indexA++
			indexB++
		}
	}

	return maxLength
}