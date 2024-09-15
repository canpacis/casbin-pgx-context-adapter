package casbinpgxcontextadapter

import (
	"context"
	"errors"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type Filter map[string][][]string

// BASE

// LoadPolicy loads all policy rules from the storage.
func (a *Adapter) LoadPolicy(model model.Model) error {
	return a.LoadPolicyCtx(a.context(), model)
}

// CONTEXT

// LoadPolicyCtx loads all policy rules from the storage with context.
func (a *Adapter) LoadPolicyCtx(ctx context.Context, model model.Model) error {
	rules, err := a.query.LoadPolicy(ctx)
	if err != nil {
		return err
	}

	for _, rule := range rules {
		if err := persist.LoadPolicyLine(rule.PolicyLine(), model); err != nil {
			return err
		}
	}

	return nil
}

// FILTERED

// LoadFilteredPolicy loads only policy rules that match the filter.
func (a *Adapter) LoadFilteredPolicy(model model.Model, filter interface{}) error {
	return a.LoadFilteredPolicyCtx(a.context(), model, filter)
}

// CONTEXT FILTERED

// LoadFilteredPolicyCtx loads only policy rules that match the filter.
func (a *Adapter) LoadFilteredPolicyCtx(ctx context.Context, model model.Model, f interface{}) error {
	if f == nil {
		return a.LoadPolicyCtx(ctx, model)
	}

	return errors.ErrUnsupported
}
