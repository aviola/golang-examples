package check_permutation_optimize_for_cpu

import (
	"testing"

	"github.com/aviola/golang-examples/algorithms/arrays_and_strings/check_permutation"
)

func TestCheckPermutation(t *testing.T) {
	optimizeForCpu := CheckPermutationOptimizeForCpu{}
	check_permutation.Test(t, check_permutation.Config{CPP: optimizeForCpu})
}
