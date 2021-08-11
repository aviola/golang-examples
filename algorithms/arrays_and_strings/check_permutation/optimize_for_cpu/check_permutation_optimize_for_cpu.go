package check_permutation_optimize_for_cpu

type CheckPermutationOptimizeForCpu struct{}

// Assumption: ASCII
func (CheckPermutationOptimizeForCpu) CheckPermutation(str1, str2 string) bool {
	strLen := len(str1)

	if strLen != len(str2) {
		return false
	}

	if strLen == 0 {
		return true
	}

	var letters [128]int
	for i := 0; i < strLen; i++ {
		asciiCode := int(str1[i])
		letters[asciiCode]++
	}

	for i := 0; i < strLen; i++ {
		asciiCode := int(str2[i])
		if letters[asciiCode] == 0 {
			return false
		} else {
			letters[asciiCode]--
		}
	}

	return true
}
