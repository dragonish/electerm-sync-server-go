package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskWithStars(t *testing.T) {
	assert.Equal(t, "*", MaskWithStars("a"))
	assert.Equal(t, "**", MaskWithStars("ab"))
	assert.Equal(t, "***", MaskWithStars("abc"))
	assert.Equal(t, "****", MaskWithStars("abcd"))
	assert.Equal(t, "*****", MaskWithStars("abcde"))
	assert.Equal(t, "******", MaskWithStars("abcdef"))
	assert.Equal(t, "a******", MaskWithStars("abcdefg"))
	assert.Equal(t, "a******h", MaskWithStars("abcdefgh"))
}
