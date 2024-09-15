package casbinpgxcontextadapter

import (
	"context"

	"github.com/CanPacis/casbin-pgx-context-adapter/db"
)

// BASE

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	return a.AddPolicyCtx(a.context(), sec, ptype, rule)
}

// CONTEXT

// AddPolicyCtx adds a policy rule to the storage with context.
// This is part of the Auto-Save feature.
func (a *Adapter) AddPolicyCtx(ctx context.Context, sec string, ptype string, rule []string) error {
	return a.AddPoliciesCtx(ctx, sec, ptype, [][]string{rule})
}

// BATCH

// AddPolicies adds policy rules to the storage.
// This is part of the Auto-Save feature.
func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	return a.AddPoliciesCtx(a.context(), sec, ptype, rules)
}

// CONTEXT BATCH

// AddPoliciesCtx adds policy rules to the storage.
// This is part of the Auto-Save feature.
func (a *Adapter) AddPoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) error {
	params := []db.AccessRule{}

	tx, commit, err := transaction(ctx, a.pool)
	if err != nil {
		return err
	}

	for _, rule := range rules {
		ar := db.AccessRule{}
		ar.Scan(ptype, rule)
		params = append(params, ar)
	}

	batch := a.query.WithTx(tx).InsertPolicy(ctx, params)

	var batchErr error
	batch.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
			batch.Close()
		}
	})
	if batchErr != nil {
		return batchErr
	}

	return commit()
}
