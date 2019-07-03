package pgmutex

import (
	"fmt"

	"github.com/vgarvardt/go-pg-adapter"
)

const defaultTableName = "mutex"

// PgMutex is the mutex lock based on PostgreSQL advisory locks
type PgMutex struct {
	adapter pgadapter.Adapter

	tableName         string
	initTableDisabled bool
}

// New instantiates and prepares PgMutex instance
func New(adapter pgadapter.Adapter, options ...Option) (*PgMutex, error) {
	instance := &PgMutex{
		adapter:   adapter,
		tableName: defaultTableName,
	}

	for _, o := range options {
		o(instance)
	}

	if !instance.initTableDisabled {
		if err := instance.initTable(); err != nil {
			return nil, err
		}
	}

	return instance, nil
}

func (m *PgMutex) initTable() error {
	return m.adapter.Exec(fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
  id   BIGSERIAL NOT NULL PRIMARY KEY,
  name TEXT      NOT NULL UNIQUE
);
`, m.tableName))
}

func (m *PgMutex) ensureSemaphoreExists(name string) error {
	return m.adapter.Exec(fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) ON CONFLICT DO NOTHING", m.tableName), name)
}

// Lock puts a named lock or waits until the the resource becomes available
func (m *PgMutex) Lock(name string) error {
	if err := m.ensureSemaphoreExists(name); err != nil {
		return err
	}
	return m.adapter.Exec(fmt.Sprintf("SELECT pg_advisory_lock(id) FROM %s WHERE name = $1", m.tableName), name)
}

// Unlock releases a previously-acquired exclusive session level advisory lock
func (m *PgMutex) Unlock(name string) error {
	return m.adapter.Exec(fmt.Sprintf("SELECT pg_advisory_unlock(id) FROM %s WHERE name = $1", m.tableName), name)
}

// TryLock puts a named lock immediately and returns true or returns false if the lock can not be acquired immediately
func (m *PgMutex) TryLock(name string) (bool, error) {
	if err := m.ensureSemaphoreExists(name); err != nil {
		return false, err
	}

	result := struct {
		Success bool `db:"success"`
	}{}
	return result.Success, m.adapter.SelectOne(&result, fmt.Sprintf("SELECT pg_try_advisory_lock(id) AS success FROM %s WHERE name = $1", m.tableName), name)
}

// TableName returns currently set mutex table name
func (m *PgMutex) TableName() string {
	return m.tableName
}
