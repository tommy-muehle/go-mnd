package config

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

func TestZeroIsIgnoredNumber(t *testing.T) {
	assert := assert.New(t)

	assert.True(DefaultConfig().IsIgnoredNumber("0"))
}

func TestCanIgnoreCustomNumbers(t *testing.T) {
	assert := assert.New(t)

	c := WithOptions(
		WithIgnoredNumbers("1,1000"),
	)

	assert.True(c.IsIgnoredNumber("0"))
	assert.True(c.IsIgnoredNumber("1"))
	assert.True(c.IsIgnoredNumber("1000"))

	assert.False(c.IsIgnoredNumber("2"))
	assert.False(c.IsIgnoredNumber("999"))
}
