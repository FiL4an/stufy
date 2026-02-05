package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Delete(ctx context.Context, conn pgx.Conn) error {
	sqlQuery := `
 DELETE FROM tasks 
WHERE id=5
`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
