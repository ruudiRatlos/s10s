package s10s

import (
	"context"
	"errors"
	"fmt"
	"time"

	api "github.com/ruudiRatlos/s10s/openapi"
)

// GetMyAgent
func (aa *AgentsAPI) GetMyAgent(ctx context.Context) (*api.Agent, error) {
	wait := 1 * time.Second
	req := aa.c.AgentsAPI.GetMyAgent(ctx)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		aa.l.DebugContext(ctx, "hit ratelimit", "ops", "AgentsAPI.GetMyAgent", "wait", wait)
		if err := aa.Sleep(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetMyAgent: %w", err)
	}

	aa.emitAgentChange(ctx, &res.Data)
	return &res.Data, nil
}

// AllAgents returns a channel of all agents and the number of agents to be returned
func (c *AgentsAPI) AllAgents(ctx context.Context) (<-chan *api.Agent, int, error) {
	out := make(chan *api.Agent)
	var (
		page  int32 = 1
		limit int32 = 20
	)
	doRequest := func(page, limit int32) (*api.GetAgents200Response, error) {
		wait := 1 * time.Second
		req := c.c.AgentsAPI.GetAgents(ctx).Page(page).Limit(limit)
	redo:
		sR, httpR, err := req.Execute() //nolint:bodyclose
		err = enhanceErr(err, httpR)
		if errors.Is(err, ErrHTTPStatus429) {
			c.l.DebugContext(ctx, "hit ratelimit", "ops", "AgentsAPI.GetAgents", "wait", wait)
			if err := c.SleepWithJitter(ctx, wait); err != nil {
				return nil, err
			}
			wait = 2 * wait
			goto redo
		}
		return sR, err
	}
	// fetch first page to get the total
	aR, err := doRequest(page, limit)
	if err != nil {
		return nil, 0, err
	}
	agents := aR.Data
	total := aR.Meta.Total
	page++

	go func() {
		defer close(out)
		for {
			aR, err = doRequest(page, limit)
			if errors.Is(err, context.Canceled) {
				return
			}
			if err != nil {
				c.l.WarnContext(ctx, "could not GetAgents", "error", err)
				return
			}
			agents = append(agents, aR.Data...)
			for _, agent := range agents {
				c.emitAgentChange(ctx, &agent) //nolint:gosec
				select {
				case out <- &agent: //nolint:gosec
				case <-ctx.Done():
					return
				}
			}
			agents = make([]api.Agent, 0, limit)

			if limit*page >= aR.Meta.Total {
				return
			}
			page++

			err = c.SleepWithJitter(ctx, 1*time.Second)
			if err != nil {
				return
			}
		}
	}()
	return out, int(total), nil
}

// GetPublicAgent
func (aa *AgentsAPI) GetPublicAgent(ctx context.Context, agentSym string) (*api.Agent, error) {
	wait := 1 * time.Second
	req := aa.c.AgentsAPI.GetAgent(ctx, agentSym)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		aa.l.DebugContext(ctx, "hit ratelimit", "ops", "AgentsAPI.GetPublicAgent", "wait", wait)
		if err := aa.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetPublicAgent: %w", err)
	}

	aa.emitAgentChange(ctx, &res.Data)
	return &res.Data, nil
}
