package casbinpgxcontextadapter

import (
	"context"

	"github.com/canpacis/casbin-pgx-context-adapter/db"
)

// UPDATE

// UpdatePolicy updates a policy rule from storage.
// This is part of the Auto-Save feature.
func (a *Adapter) UpdatePolicy(sec string, ptype string, oldRule, newRule []string) error {
	return a.UpdatePolicyCtx(a.context(), sec, ptype, oldRule, newRule)
}

// UpdatePolicies updates some policy rules to storage, like db, redis.
func (a *Adapter) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
	return a.UpdatePoliciesCtx(a.context(), sec, ptype, oldRules, newRules)
}

// UpdateFilteredPolicies deletes old rules and adds new rules.
func (a *Adapter) UpdateFilteredPolicies(sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	return a.UpdateFilteredPoliciesCtx(a.context(), sec, ptype, newRules, fieldIndex, fieldValues...)
}

// CONTEXT UPDATE

// UpdatePolicyCtx updates a policy rule from storage.
// This is part of the Auto-Save feature.
func (a *Adapter) UpdatePolicyCtx(ctx context.Context, sec string, ptype string, oldRule, newRule []string) error {
	return a.UpdatePoliciesCtx(ctx, sec, ptype, [][]string{oldRule}, [][]string{newRule})
}

// UpdatePoliciesCtx updates some policy rules to storage, like db, redis.
func (a *Adapter) UpdatePoliciesCtx(ctx context.Context, sec string, ptype string, oldRules, newRules [][]string) error {
	tx, commit, err := transaction(ctx, a.pool)
	if err != nil {
		return err
	}

	params := []db.AccessRule{}

	for i, oldRule := range oldRules {
		old := db.AccessRule{}
		old.Scan(ptype, oldRule)

		newRule := db.AccessRule{}
		newRule.Scan(ptype, newRules[i])
		newRule.ID = old.ID

		params = append(params, newRule)
	}

	batch := a.query.WithTx(tx).UpdatePolicy(ctx, params)

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

// UpdateFilteredPoliciesCtx deletes old rules and adds new rules.
func (a *Adapter) UpdateFilteredPoliciesCtx(
	ctx context.Context,
	sec string,
	ptype string,
	newRules [][]string,
	fieldIndex int,
	fieldValues ...string,
) ([][]string, error) {
	tx, commit, err := transaction(ctx, a.pool)
	if err != nil {
		return nil, err
	}

	filter := make([]string, 6)

	for i, value := range fieldValues {
		index := i + fieldIndex
		if index >= len(filter) {
			continue
		}
		filter[i+fieldIndex] = value
	}

	filterRule := db.AccessRule{}
	filterRule.Scan(ptype, filter)

	params := []db.AccessRule{}
	for _, newRule := range newRules {
		param := db.AccessRule{}
		param.Scan(ptype, newRule)
		params = append(params, param)
	}

	batch := a.query.WithTx(tx).UpdateFilteredPolicy(ctx, filterRule, params)

	var batchErr error
	batch.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
			batch.Close()
		}
	})
	if batchErr != nil {
		return nil, batchErr
	}

	return [][]string{}, commit()
}
