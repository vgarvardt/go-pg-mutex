package pgmutex

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vgarvardt/go-pg-adapter/pgxadapter"
)

var uri string

func TestMain(m *testing.M) {
	uri = os.Getenv("PG_URI")
	if uri == "" {
		fmt.Println("Env variable PG_URI is required to run the tests")
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	pgxConnConfig, err := pgx.ParseURI(uri)
	require.NoError(t, err)

	pgxConn1, err := pgx.Connect(pgxConnConfig)
	pgxConn2, err := pgx.Connect(pgxConnConfig)
	require.NoError(t, err)

	adapter1 := pgxadapter.NewConn(pgxConn1)
	adapter2 := pgxadapter.NewConn(pgxConn2)

	tableName := fmt.Sprintf("mutex_%d", time.Now().UnixNano())
	lockName := fmt.Sprintf("lock_%d", time.Now().UnixNano())

	m1, err := New(adapter1, WithTableName(tableName))
	require.NoError(t, err)

	m2, err := New(adapter2, WithTableName(tableName))
	require.NoError(t, err)

	err = m1.Lock(lockName)
	require.NoError(t, err)

	success, err := m2.TryLock(lockName)
	require.NoError(t, err)
	assert.False(t, success)

	err = m1.Unlock(lockName)
	require.NoError(t, err)

	success, err = m2.TryLock(lockName)
	require.NoError(t, err)
	assert.True(t, success)

	success, err = m1.TryLock(lockName)
	require.NoError(t, err)
	assert.False(t, success)

	var m1AcquiredLock bool
	go func() {
		err := m1.Lock(lockName)
		require.NoError(t, err)
		m1AcquiredLock = true
	}()

	time.Sleep(time.Second)

	assert.False(t, m1AcquiredLock)

	err = m2.Unlock(lockName)
	require.NoError(t, err)

	time.Sleep(time.Second)

	assert.True(t, m1AcquiredLock)

	err = m1.Unlock(lockName)
	require.NoError(t, err)
}
