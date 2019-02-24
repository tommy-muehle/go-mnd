package magic_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithCustomChecks(t *testing.T) {
	assert := assert.New(t)

	c := WithOptions(
		WithCustomChecks("return,operation"),
	)

	assert.True(c.IsCheckEnabled("return"))
	assert.True(c.IsCheckEnabled("operation"))

	assert.False(c.IsCheckEnabled("assign"))
	assert.False(c.IsCheckEnabled("case"))
	assert.False(c.IsCheckEnabled("argument"))
	assert.False(c.IsCheckEnabled("condition"))
}
