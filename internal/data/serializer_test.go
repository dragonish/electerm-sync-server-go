package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	assert.Empty(t, GetUsers(""))
	assert.Empty(t, GetUsers(" "))
	assert.Equal(t, []string{"abc"}, GetUsers("abc"))
	assert.Equal(t, []string{"abc"}, GetUsers(",abc"))
	assert.Equal(t, []string{"abc"}, GetUsers("abc,"))
	assert.Equal(t, []string{"abc"}, GetUsers(" ,abc, "))
	assert.Equal(t, []string{"abc", "efg"}, GetUsers(" abc, efg "))
}
