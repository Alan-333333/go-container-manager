package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildImage(t *testing.T) {
	name := "go-demo"
	got, err := BuildImage(name)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, got, "")
}
