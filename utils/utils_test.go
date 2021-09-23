// +build all

package utils

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	assert := assert.New(t)

	assert.True(Exists(RootDir() + "/fixtures/www/foo.com/config.json"))
	assert.False(Exists(RootDir() + "/fixtures/www/foo.com/_config.json"))
	assert.False(Exists(RootDir() + "/fixtures/www/baz.com/config.json"))
}

// BenchmarkExists-12    	  243614	      4351 ns/op	     552 B/op	       5 allocs/op
func BenchmarkExists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Exists(RootDir() + "/fixtures/www/foo.com/config.json")
	}
}
