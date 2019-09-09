package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {
	if s, err := NewServer("../config.json"); assert.NoError(t, err) {
		assert.NotNil(t, s)
	}
}
