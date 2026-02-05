package simpleconnection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:5432@localhost:5432/postgres")

}
