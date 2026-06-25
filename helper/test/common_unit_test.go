package test

import (
	"ewallet-transaction/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateReference_WhenCalled_ReturnsValidReference(t *testing.T) {
	// Act
	result := helper.GenerateReference()

	// Assert
	assert.NotEmpty(t, result)
	assert.Regexp(t, `^\d+$`, result)
	assert.GreaterOrEqual(t, len(result), 15)
	assert.LessOrEqual(t, len(result), 16)
}
