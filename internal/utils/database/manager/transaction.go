package manager

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type key string

const dbTransKey key = "db_trans"

type GORMx interface {
	//method from gorm
	WithContext(ctx context.Context) *gorm.DB
	DB() (*sql.DB, error)
}

// DBTransactionGormFromContext
func DBTransactionFromContext(ctx context.Context) (GORMx, bool) {
	conn, ok := ctx.Value(dbTransKey).(GORMx)
	return conn, ok
}
