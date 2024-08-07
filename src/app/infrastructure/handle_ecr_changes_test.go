package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertECRNameToLambda(t *testing.T) {
	input := "hello-new-image"
	expected := "HelloNewImage"

	result := convertECRImageNameToLambdaName(input)
	assert.Equal(t, expected, result)
}
