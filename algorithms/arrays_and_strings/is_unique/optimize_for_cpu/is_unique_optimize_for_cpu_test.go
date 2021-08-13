package is_unique_optimize_for_cpu

import (
	"testing"

	"github.com/aviola/golang-examples/algorithms/arrays_and_strings/is_unique"
)

func TestIsUnique(t *testing.T) {
	optimizeForCpu := IsUniqueOptimizeForCpu{}
	is_unique.Test(t, is_unique.Config{IUP: optimizeForCpu})
}
