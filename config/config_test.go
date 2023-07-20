package config_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tommy-muehle/go-mnd/v2/config"
)

func TestWithCustomChecks(t *testing.T) {
	assert := assert.New(t)

	c := config.WithOptions(
		config.WithCustomChecks("return,operation"),
	)

	assert.True(c.IsCheckEnabled("return"))
	assert.True(c.IsCheckEnabled("operation"))

	assert.False(c.IsCheckEnabled("assign"))
	assert.False(c.IsCheckEnabled("case"))
	assert.False(c.IsCheckEnabled("argument"))
	assert.False(c.IsCheckEnabled("condition"))
}

func TestIgnoredNumbers(t *testing.T) {
	assert := assert.New(t)

	// Sync with DefaultConfig()
	// Ignored numbers are: 0, 1, and their two floating expressions
	ignoredNumbers := []string{ "0", "0.0", "0.", "1", "1.0", "1." }

	for _, n := range ignoredNumbers {
		assert.True(config.DefaultConfig().IsIgnoredNumber(n))
	}
}

func TestCanIgnoreCustomNumbers(t *testing.T) {
	assert := assert.New(t)

	c := config.WithOptions(
		config.WithIgnoredNumbers("1,1000"),
	)

	assert.True(c.IsIgnoredNumber("0"))
	assert.True(c.IsIgnoredNumber("1"))
	assert.True(c.IsIgnoredNumber("1000"))

	assert.False(c.IsIgnoredNumber("2"))
	assert.False(c.IsIgnoredNumber("999"))
}

func TestWithIgnoredFiles(t *testing.T) {
	assert := assert.New(t)

	c := config.WithOptions(
		config.WithIgnoredFiles(".*"),
	)

	assert.Contains(c.IgnoredFiles, regexp.MustCompile(".*"))
}
