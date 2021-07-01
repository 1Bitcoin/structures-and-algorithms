package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLib_LRUCache(t *testing.T) {
	cache := NewLru(2)
	cache.Set("int23", 23)
	cache.Set("string", "hello")

	get1 := cache.Get("int23")
	get2 := cache.Get("string")
	assert.Equal(t, get1, 23)
	assert.Equal(t, get2, "hello")

	response1 := cache.Set("newer", 5)
	assert.Equal(t, response1, true)

	get3 := cache.Get("int23")
	get4 := cache.Get("string")
	get5 := cache.Get("newer")
	assert.Equal(t, get3, false)
	assert.Equal(t, get4, "hello")
	assert.Equal(t, get5, 5)
}
