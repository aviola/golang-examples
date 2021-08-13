package check_permutation_optimize_for_memory

import (
	"sort"
	"strings"
)

type CheckPermutationOptimizeForMemory struct{}

func (CheckPermutationOptimizeForMemory) CheckPermutation(str1, str2 string) bool {
	strLen := len(str1)

	if strLen != len(str2) {
		return false
	}

	if strLen == 0 || str1 == str2 {
		return true
	}

	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str2)

	return sortedStr1 == sortedStr2
}

func sortString(s string) string {
	strSlice := strings.Split(s, "")
	sort.Strings(strSlice)
	return strings.Join(strSlice, "")
}
