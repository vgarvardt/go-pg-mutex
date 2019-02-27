package pgmutex

// Option is the configuration options type for PgMutex
type Option func(mutex *PgMutex)

// WithTableName returns option that sets PgMutex table name
func WithTableName(tableName string) Option {
	return func(mutex *PgMutex) {
		mutex.tableName = tableName
	}
}

// WithInitTableDisabled returns option that disables table creation on PgMutex instantiation
func WithInitTableDisabled() Option {
	return func(mutex *PgMutex) {
		mutex.initTableDisabled = true
	}
}
