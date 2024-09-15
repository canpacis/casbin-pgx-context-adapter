package casbinpgxcontextadapter

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func transaction(ctx context.Context, pgxdb *pgxpool.Pool) (pgx.Tx, func() error, error) {
	tx, err := pgxdb.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() error {
		defer tx.Rollback(context.Background())

		if err := tx.Commit(ctx); err != nil {
			return err
		}

		return nil
	}

	return tx, cleanup, nil
}
