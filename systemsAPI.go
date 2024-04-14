package s10s

import (
	"context"
	"errors"
	"fmt"
	"time"

	api "github.com/ruudiRatlos/s10s/api"
)

// AllSystems returns a channel of all systems and the number of systems to be returned
func (c *systemsAPI) AllSystems(ctx context.Context) (<-chan *api.System, int, error) {
	out := make(chan *api.System)
	var (
		page  int32 = 1
		limit int32 = 20
	)
	doRequest := func(page, limit int32) (*api.GetSystems200Response, error) {
		wait := 1 * time.Second
		req := c.c.SystemsAPI.GetSystems(ctx).Page(page).Limit(limit)
	redo:
		sR, httpR, err := req.Execute() //nolint:bodyclose
		err = enhanceErr(err, httpR)
		if errors.Is(err, ErrHTTPStatus429) {
			c.l.DebugContext(ctx, "hit ratelimit", "ops", "SystemsAPI.GetSystems", "wait", wait)
			if err := c.SleepWithJitter(ctx, wait); err != nil {
				return nil, err
			}
			wait = 2 * wait
			goto redo
		}
		return sR, err
	}
	// fetch first page to get the total
	sR, err := doRequest(page, limit)
	if err != nil {
		return nil, 0, err
	}
	systems := sR.Data
	total := sR.Meta.Total
	page++

	go func() {
		defer close(out)
		for {
			sR, err = doRequest(page, limit)
			if errors.Is(err, context.Canceled) {
				return
			}
			if err != nil {
				c.l.WarnContext(ctx, "could not GetSystems", "error", err)
				return
			}
			systems = append(systems, sR.Data...)
			for _, system := range systems {
				// do not emit systems as waypoint info is not complete
				select {
				case out <- &system: //nolint:gosec
				case <-ctx.Done():
					return
				}
			}
			systems = make([]api.System, 0, limit)

			if limit*page >= sR.Meta.Total {
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

// GetSystem returns the details of a system
func (c *systemsAPI) GetSystem(ctx context.Context, systemSymbol string) (*api.System, error) {
	wait := 1 * time.Second
	req := c.c.SystemsAPI.GetSystem(ctx, systemSymbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "SystemsAPI.GetSystem", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetSystem: %w", err)
	}
	c.emitSystem(ctx, &res.Data)
	return &res.Data, nil
}

func (c *systemsAPI) GetSystemWaypoints(ctx context.Context, sys SystemSymbol) ([]*api.Waypoint, error) {
	var (
		out   []*api.Waypoint = make([]*api.Waypoint, 0, 10)
		page  int32           = 1
		limit int32           = 20
	)
	for {
		wait := 1 * time.Second
		req := c.c.SystemsAPI.GetSystemWaypoints(ctx, string(sys)).Page(page).Limit(limit)
	redo:
		sR, httpR, err := req.Execute() //nolint:bodyclose
		err = enhanceErr(err, httpR)
		if errors.Is(err, ErrHTTPStatus429) {
			c.l.DebugContext(ctx, "hit ratelimit", "ops", "SystemsAPI.GetSystemWaypoints", "wait", wait)
			if err := c.SleepWithJitter(ctx, wait); err != nil {
				return nil, err
			}
			wait = 2 * wait
			goto redo
		}
		if err != nil {
			return nil, fmt.Errorf("could not GetSystemWaypoints: %w", err)
		}
		for _, s := range sR.Data {
			c.emitWaypoint(ctx, &s) //nolint:gosec
			out = append(out, &s)   //nolint:gosec
		}

		if limit*page >= sR.Meta.Total {
			return out, nil
		}
		page++
	}
}

// GetWaypoint
func (c *systemsAPI) GetWaypoint(ctx context.Context, wpSym WaypointSymbol) (*api.Waypoint, error) {
	wait := 1 * time.Second
	req := c.c.SystemsAPI.GetWaypoint(ctx, string(wpSym.SystemSymbol()), string(wpSym))
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.CreateChart", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetWaypoint: %w", err)
	}
	c.emitWaypoint(ctx, &res.Data)
	return &res.Data, nil
}

// GetMarket retrieve imports, exports and exchange data from a marketplace.
// Requires a waypoint that has the Marketplace trait to use.
//
// Send a ship to the waypoint to access trade good prices and recent transactions.
func (c *systemsAPI) GetMarket(ctx context.Context, wpSym WaypointSymbol) (*api.Market, error) {
	wait := 1 * time.Second
	system := wpSym.SystemSymbol()
redo:
	mRes, httpR, err := c.c.SystemsAPI.GetMarket(ctx, string(system), string(wpSym)).Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "SystemsAPI.GetMarket", "wait", wait)
		err := c.SleepWithJitter(ctx, wait)
		if err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetMarket: %w", err)
	}
	c.emitMarket(ctx, &mRes.Data)
	return &mRes.Data, nil
}

// GetShipyard gets the shipyard for a waypoint.
// Requires a waypoint that has the Shipyard trait to use.
//
// Send a ship to the waypoint to access data on ships that are currently available for purchase and recent transactions.
func (c *systemsAPI) GetShipyard(ctx context.Context, wpSym WaypointSymbol) (*api.Shipyard, error) {
	wait := 1 * time.Second
	req := c.c.SystemsAPI.GetShipyard(ctx, wpSym.SystemSymbol().String(), wpSym.String())
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "SystemsAPI.GetShipyard", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetShipyard (%q): %w", string(wpSym), err)
	}
	c.emitShipyard(ctx, &res.Data)
	return &res.Data, nil
}

// GetJumpGate gets jump gate details for a waypoint.
// Requires a waypoint of type JUMP_GATE to use.
func (c *systemsAPI) GetJumpGate(ctx context.Context, wpSym WaypointSymbol) (*api.JumpGate, error) {
	wait := 1 * time.Second
	req := c.c.SystemsAPI.GetJumpGate(ctx, wpSym.SystemSymbol().String(), wpSym.String())
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.GetJumpGate", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetJumpGate: %w", err)
	}
	c.emitJumpGate(ctx, &res.Data)
	return &res.Data, nil
}

// GetConstruction gets construction details for a waypoint.
// Requires a waypoint with a property of isUnderConstruction to be true.
func (c *systemsAPI) GetConstruction(ctx context.Context, wp *api.Waypoint) (*api.Construction, error) {
	wait := 1 * time.Second
	req := c.c.SystemsAPI.GetConstruction(ctx, wp.SystemSymbol, wp.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "SystemsAPI.GetConstruction", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetConstruction: %w", err)
	}
	c.emitConstruction(ctx, &res.Data)
	return &res.Data, nil
}

// SupplyConstruction supply a construction site with the specified good.
// Requires a waypoint with a property of isUnderConstruction to be true.
//
// The good must be in your ship's cargo. The good will be removed from
// your ship's cargo and added to the construction site's materials.
func (c *systemsAPI) SupplyConstruction(ctx context.Context, ship *api.Ship, good *api.TradeSymbol, units int32) (*api.SupplyConstruction201ResponseData, error) {
	wait := 1 * time.Second
	scRequest := api.NewSupplyConstructionRequest(ship.Symbol, string(*good), units)
	req := c.c.SystemsAPI.SupplyConstruction(ctx, ship.Nav.SystemSymbol, ship.Nav.WaypointSymbol).
		SupplyConstructionRequest(*scRequest)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "SystemsAPI.SupplyConstruction", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not SupplyConstruction: %w", err)
	}
	ship.Cargo = res.Data.Cargo
	c.emitShipChange(ctx, ship)
	c.emitConstruction(ctx, &res.Data.Construction)
	return &res.Data, nil
}