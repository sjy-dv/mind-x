package embeddings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextEmbeddings(t *testing.T) {
	err := InitEmbeddings()
	assert.NoError(t, err)

	input := "test code is cool"
	vec, err := TextEmbedding(input)
	assert.NoError(t, err)
	assert.NotEmpty(t, vec)
	t.Log(vec)
}
