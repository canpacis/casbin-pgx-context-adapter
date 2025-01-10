package casbinpgxcontextadapter

import (
	"context"

	"github.com/canpacis/casbin-pgx-context-adapter/db"
)

// BASE

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return a.RemovePolicyCtx(a.context(), sec, ptype, rule)
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return a.RemoveFilteredPolicyCtx(a.context(), sec, ptype, fieldIndex, fieldValues...)
}

// CONTEXT

// RemovePolicyCtx removes a policy rule from the storage with context.
// This is part of the Auto-Save feature.
func (a *Adapter) RemovePolicyCtx(ctx context.Context, sec string, ptype string, rule []string) error {
	return a.RemovePoliciesCtx(ctx, sec, ptype, [][]string{rule})
}

// RemoveFilteredPolicyCtx removes policy rules that match the filter from the storage with context.
// This is part of the Auto-Save feature.
func (a *Adapter) RemoveFilteredPolicyCtx(ctx context.Context, sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	// this will be populated with given values
	filter := make([]string, 6)

	for i, value := range fieldValues {
		index := i + fieldIndex
		if index >= len(filter) {
			continue
		}
		filter[i+fieldIndex] = value
	}

	params := db.AccessRule{}
	params.Scan(ptype, filter)

	return a.query.FilteredRemovePolicy(ctx, params)
}

// BATCH

// RemovePolicies removes policy rules from the storage.
// This is part of the Auto-Save feature.
func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	return a.RemovePoliciesCtx(a.context(), sec, ptype, rules)
}

// CONTEXT BATCH

// RemovePoliciesCtx removes policy rules from the storage.
// This is part of the Auto-Save feature.
func (a *Adapter) RemovePoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) error {
	tx, commit, err := transaction(ctx, a.pool)
	if err != nil {
		return err
	}

	ids := []string{}

	// turn rules into access rules and calculate their ids
	for _, rule := range rules {
		ar := db.AccessRule{}
		ar.Scan(ptype, rule)
		ids = append(ids, ar.GetID())
	}

	batch := a.query.WithTx(tx).RemovePolicy(ctx, ids)
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
