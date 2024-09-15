package casbinpgxcontextadapter

import (
	"context"
	"time"

	"github.com/CanPacis/casbin-pgx-context-adapter/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Adapter struct {
	pool     *pgxpool.Pool
	query    *db.Queries
	timeout  time.Duration
	filtered bool
}

func (a *Adapter) context() context.Context {
	return context.Background()
}

// FILTERED ADAPTER

// IsFiltered returns true if the loaded policy has been filtered.
func (a *Adapter) IsFiltered() bool {
	return a.IsFilteredCtx(a.context())
}

// CONTEXT FILTERED ADAPTER

// IsFilteredCtx returns true if the loaded policy has been filtered.
func (a *Adapter) IsFilteredCtx(ctx context.Context) bool {
	return a.filtered
}

func New(pool *pgxpool.Pool, options ...Option) *Adapter {
	adapter := &Adapter{
		pool:     pool,
		query:    db.New(pool),
		timeout:  time.Second * 10,
		filtered: false,
	}

	for _, option := range options {
		option(adapter)
	}

	return adapter
}

type Option func(*Adapter)

// WithTimeout can be used to pass a different timeout than DefaultTimeout
// for each request to Postgres
func WithTimeout(timeout time.Duration) Option {
	return func(a *Adapter) {
		a.timeout = timeout
	}
}
