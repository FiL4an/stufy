package simplesql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

// <- http
// <-body
func InsertRow(ctx context.Context, conn *pgx.Conn, tittle string, description string, completed bool, createdAt time.Time,
) error {
	sqlQuery := `INSERT INTO tasks (title, description, completed, created_at)
	 VALUES($1, $2 , $3, $4)`

	_, err := conn.Exec(ctx, sqlQuery, tittle, description, completed, createdAt)
	return err
}
