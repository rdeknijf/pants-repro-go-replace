package first

import (
	"testing"

	"gotest.tools/assert"
)

func TestFirst(t *testing.T) {

	assert.Equal(t, "first", First())
}
