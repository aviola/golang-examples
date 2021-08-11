package check_permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	CPP CheckPermutationProvider
}

func Test(t *testing.T, config Config) {
	tests := []struct {
		title string
		run   func(t *testing.T, c Config)
	}{
		{"Return false if the two strings have different lenght", testStringsWithDifferentLength},
		{"Return false if the two strings have the same lenght and one is not a permutation of the other", testStringsWithSameLengthNotBeingPermutation},
		{"Return false if the two strings have the same lenght and one is a permutation of the other", testStringsWithSameLengthBeingPermutation},
		{"Return true if the two strings are empty", testEmptyStrings},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			test.run(t, config)
		})
	}
}

func testStringsWithDifferentLength(t *testing.T, c Config) {
	assert.False(t, c.CPP.CheckPermutation("hello", "hello!"))
}

func testStringsWithSameLengthNotBeingPermutation(t *testing.T, c Config) {
	assert.False(t, c.CPP.CheckPermutation("hello", "h3ll0"))
}

func testStringsWithSameLengthBeingPermutation(t *testing.T, c Config) {
	assert.True(t, c.CPP.CheckPermutation("hola", "halo"))
}

func testEmptyStrings(t *testing.T, c Config) {
	assert.True(t, c.CPP.CheckPermutation("", ""))
}
