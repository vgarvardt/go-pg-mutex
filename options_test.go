package pgmutex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithTableName(t *testing.T) {
	randomName := time.Now().String()

	m, err := New(nil, WithTableName(randomName), WithInitTableDisabled())
	require.NoError(t, err)
	assert.Equal(t, randomName, m.tableName)
}
