package s10s

import (
	"context"
	"errors"
	"fmt"
	"time"

	api "github.com/ruudiRatlos/s10s/api"
)

// ListContracts returns a list of all your contracts.
func (c *contractsAPI) ListContracts(ctx context.Context) ([]*api.Contract, error) {
	var (
		out   []*api.Contract = make([]*api.Contract, 0, 20)
		page  int32           = 1
		limit int32           = 20
	)
	for {
		wait := 1 * time.Second
		req := c.c.ContractsAPI.GetContracts(ctx).Page(page).Limit(limit)
	redo:
		sR, httpR, err := req.Execute() //nolint:bodyclose
		err = enhanceErr(err, httpR)
		if errors.Is(err, ErrHTTPStatus429) {
			c.l.DebugContext(ctx, "hit ratelimit", "ops", "SystemsAPI.ListContracts", "wait", wait)
			if err := c.SleepWithJitter(ctx, wait); err != nil {
				return nil, err
			}
			wait = 2 * wait
			goto redo
		}
		if err != nil {
			return nil, fmt.Errorf("could not ListContracts: %w", err)
		}
		for _, contract := range sR.Data {
			c.emitContract(ctx, &contract) //nolint:gosec
			out = append(out, &contract)   //nolint:gosec
		}

		if limit*page >= sR.Meta.Total {
			return out, nil
		}
		page++
	}
}

// GetContract returns data about a contract by ID
func (c *contractsAPI) GetContract(ctx context.Context, contractID string) (*api.Contract, error) {
	wait := 1 * time.Second
	req := c.c.ContractsAPI.GetContract(ctx, contractID)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "ContractsAPI.GetContract", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetContract: %w", err)
	}
	c.emitContract(ctx, &res.Data)
	return &res.Data, nil
}

// AcceptContract accepts a contract by ID
// You can only accept contracts that were offered to you, were not accepted yet,
// and whose deadlines has not passed yet.
func (c *contractsAPI) AcceptContract(ctx context.Context, contract *api.Contract) (*api.AcceptContract200ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.ContractsAPI.AcceptContract(ctx, contract.Id)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "ContractsAPI.AcceptContract", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not AcceptContract: %w", err)
	}
	c.emitContract(ctx, &res.Data.Contract)
	c.emitAgentChange(ctx, &res.Data.Agent)
	return &res.Data, nil
}

// DeliverCargo returns data about a contract
func (c *contractsAPI) DeliverContract(ctx context.Context, contract *api.Contract, ship *api.Ship, good api.TradeSymbol, units int32) (*api.DeliverContract200ResponseData, error) {
	wait := 1 * time.Second
	dcr := api.NewDeliverContractRequest(ship.Symbol, string(good), units)
	req := c.c.ContractsAPI.DeliverContract(ctx, contract.Id).DeliverContractRequest(*dcr)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "ContractsAPI.DeliverContract", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not DeliverContract: %w", err)
	}
	ship.Cargo = res.Data.Cargo
	c.emitShipChange(ctx, ship)
	c.emitContract(ctx, &res.Data.Contract)
	return &res.Data, nil
}

// FulfillContract fulfills a contract.
// Can only be used on contracts that have all of their delivery terms fulfilled.
func (c *contractsAPI) FulfillContract(ctx context.Context, contract *api.Contract) (*api.AcceptContract200ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.ContractsAPI.FulfillContract(ctx, contract.Id)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "ContractsAPI.FulfillContract", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not FulfillContract: %w", err)
	}
	c.emitContract(ctx, &res.Data.Contract)
	c.emitAgentChange(ctx, &res.Data.Agent)
	return &res.Data, nil
}
