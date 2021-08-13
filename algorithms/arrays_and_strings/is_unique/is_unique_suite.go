package is_unique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	IUP IsUniqueProvider
}

func Test(t *testing.T, config Config) {
	tests := []struct {
		title string
		run   func(t *testing.T, c Config)
	}{
		{"Return true on empty string", testEmptyString},
		{"Return true on string with unique characters", testStringWithUniqueCharacters},
		{"Return false on string with repeated characters", testStringWithRepeatedCharacters},
		{"Should return false on string with length exceeding the number of unique ASCII characters", testStringWithMoreCharactersThanASCIICharSet},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			test.run(t, config)
		})
	}
}

func testEmptyString(t *testing.T, c Config) {
	assert.True(t, c.IUP.IsUnique(""))
}

func testStringWithUniqueCharacters(t *testing.T, c Config) {
	assert.True(t, c.IUP.IsUnique("Hola"))
}

func testStringWithRepeatedCharacters(t *testing.T, c Config) {
	assert.False(t, c.IUP.IsUnique("Hello"))
}

func testStringWithMoreCharactersThanASCIICharSet(t *testing.T, c Config) {
	var asciiSet string = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f" +
		"\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f" +
		"\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf" +
		"\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf" +
		"\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf" +
		"\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf" +
		"\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef" +
		"\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"

	assert.False(t, c.IUP.IsUnique(asciiSet+"a"))
}
