package is_unique_optimize_for_cpu

type IsUniqueOptimizeForCpu struct{}

func (IsUniqueOptimizeForCpu) IsUnique(str string) bool {
	if len(str) > 128 {
		return false
	}

	var charSet [128]bool
	for i := 0; i < len(str); i++ {
		var val int = int(str[i])
		if charSet[val] {
			return false
		}
		charSet[val] = true
	}

	return true
}
