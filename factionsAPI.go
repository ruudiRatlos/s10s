package s10s

import (
	"context"
	"errors"
	"fmt"
	"time"

	api "github.com/ruudiRatlos/s10s/api"
)

// ListFactions returns a list of all factions in the game.
func (c *factionsAPI) ListFactions(ctx context.Context) ([]*api.Faction, error) {
	var (
		out   []*api.Faction = make([]*api.Faction, 0, 20)
		page  int32          = 1
		limit int32          = 20
	)
	for {
		wait := 1 * time.Second
		req := c.c.FactionsAPI.GetFactions(ctx).Page(page).Limit(limit)
	redo:
		fR, httpR, err := req.Execute() //nolint:bodyclose
		err = enhanceErr(err, httpR)
		if errors.Is(err, ErrHTTPStatus429) {
			c.l.DebugContext(ctx, "hit ratelimit", "ops", "FactionsAPI.ListFactions", "wait", wait)
			if err := c.SleepWithJitter(ctx, wait); err != nil {
				return nil, err
			}
			wait = 2 * wait
			goto redo
		}
		if err != nil {
			return nil, fmt.Errorf("could not ListContracts: %w", err)
		}
		for _, faction := range fR.Data {
			out = append(out, &faction) //nolint:gosec
		}

		if limit*page >= fR.Meta.Total {
			return out, nil
		}
		page++
	}
}

// GetFaction returns the details of a faction
func (c *factionsAPI) GetFaction(ctx context.Context, factionSymbol string) (*api.Faction, error) {
	wait := 1 * time.Second
	req := c.c.FactionsAPI.GetFaction(ctx, factionSymbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FactionsAPI.GetFaction", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetSystem: %w", err)
	}
	return &res.Data, nil
}
