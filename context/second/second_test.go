package second

import (
	"testing"

	"gotest.tools/assert"

	first "something.org/first"
)

func TestSecond(t *testing.T) {

	assert.Equal(t, "second", Second())
}

func TestFirst(t *testing.T) {

	assert.Equal(t, "first", first.First())
}
