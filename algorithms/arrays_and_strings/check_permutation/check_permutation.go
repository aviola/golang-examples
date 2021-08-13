package check_permutation

type CheckPermutationProvider interface {
	CheckPermutation(str1, str2 string) bool
}
