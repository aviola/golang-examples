package check_permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPermutation(t *testing.T) {

	t.Run("Return false if the two strings have different lenght", func(t *testing.T) {
		assert.False(t, CheckPermutation("hello", "hello!"))
	})

	t.Run("Return false if the two strings have the same lenght and one is not a permutation of the other", func(t *testing.T) {
		assert.False(t, CheckPermutation("hello", "h3ll0"))
	})

	t.Run("Return true if the two strings have the same lenght and one is a permutation of the other", func(t *testing.T) {
		assert.True(t, CheckPermutation("hola", "halo"))
	})

	t.Run("Return true if the two strings are empty", func(t *testing.T) {
		assert.True(t, CheckPermutation("", ""))
	})
}
