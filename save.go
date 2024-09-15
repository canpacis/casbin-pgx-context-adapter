package casbinpgxcontextadapter

import (
	"context"

	"github.com/CanPacis/casbin-pgx-context-adapter/db"
	"github.com/casbin/casbin/model"
)

// BASE

// SavePolicy saves all policy rules to the storage.
func (a *Adapter) SavePolicy(model model.Model) error {
	return a.SavePolicyCtx(a.context(), model)
}

// CONTEXT

// SavePolicyCtx saves all policy rules to the storage with context.
func (a *Adapter) SavePolicyCtx(ctx context.Context, model model.Model) error {
	tx, commit, err := transaction(ctx, a.pool)
	if err != nil {
		return err
	}

	rows := []db.AccessRule{}

	for ptype, assertions := range model {
		for _, assertion := range assertions {
			for _, rule := range assertion.Policy {
				ar := db.AccessRule{}
				ar.Scan(ptype, rule)
				rows = append(rows, ar)
			}
		}
	}

	_, err = a.query.WithTx(tx).Copy(ctx, rows)
	if err != nil {
		return err
	}

	return commit()
}
