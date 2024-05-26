package s10s

import (
	"context"
	"errors"
	"fmt"
	"time"

	api "github.com/ruudiRatlos/s10s/openapi"
)

// ListShips is an alias for GetMyShips
func (c *FleetAPI) ListShips(ctx context.Context) ([]*api.Ship, error) {
	return c.GetMyShips(ctx)
}

// GetMyShips returns a paginated list of all ships under your agent's ownership.
func (c *FleetAPI) GetMyShips(ctx context.Context) ([]*api.Ship, error) {
	var (
		out   []*api.Ship = make([]*api.Ship, 0, 10)
		page  int32       = 1
		limit int32       = 20
	)
	for {
		wait := 1 * time.Second
		req := c.c.FleetAPI.GetMyShips(ctx).Page(page).Limit(limit)
	redo:
		shipsR, httpR, err := req.Execute() //nolint:bodyclose
		err = enhanceErr(err, httpR)
		if errors.Is(err, ErrHTTPStatus429) {
			c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.GetMyShips", "wait", wait)
			if err := c.SleepWithJitter(ctx, wait); err != nil {
				return nil, err
			}
			wait = 2 * wait
			goto redo
		}
		if err != nil {
			return nil, fmt.Errorf("could not GetMyShips: %w", err)
		}
		for _, s := range shipsR.Data {
			c.emitShipChange(ctx, &s) //nolint:gosec
			out = append(out, &s)     //nolint:gosec
		}

		if int(shipsR.Meta.Total) == len(out) {
			return out, nil
		}
		page++
	}
}

// Purchase a ship
func (c *FleetAPI) PurchaseShip(ctx context.Context,
	shipType api.ShipType, wpSym WaypointSymbol) (*api.PurchaseShip201ResponseData, error) {
	purchaseShipRequest := api.NewPurchaseShipRequest(shipType, wpSym.String())

	wait := 1 * time.Second
	req := c.c.FleetAPI.PurchaseShip(ctx).PurchaseShipRequest(*purchaseShipRequest)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.RefuelShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not RefuelShip: %w", err)
	}
	c.emitShipChange(ctx, &res.Data.Ship)
	c.emitAgentChange(ctx, &res.Data.Agent)
	c.emitShipyardTransaction(ctx, &res.Data.Transaction)
	return &res.Data, nil
}

// GetMyShipFromSymbol returns data about a ship
func (c *FleetAPI) GetMyShipFromSymbol(ctx context.Context, shipSymbol string) (*api.Ship, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.GetMyShip(ctx, shipSymbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.GetMyShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetMyShipFromSymbol: %w", err)
	}
	c.emitShipChange(ctx, &res.Data)
	return &res.Data, nil
}

// OrbitShip attempts to move your ship into orbit at its current location.
// The request will only succeed if your ship is capable of moving into orbit
// at the time of the request.
func (c *FleetAPI) OrbitShip(ctx context.Context, ship *api.Ship) error {
	wait := 1 * time.Second
	req := c.c.FleetAPI.OrbitShip(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.OrbitShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return fmt.Errorf("could not OrbitShip: %w", err)
	}
	ship.Nav = res.Data.Nav
	c.emitShipChange(ctx, ship)
	return nil
}

// ShipRefine
func (c *FleetAPI) ShipRefine(ctx context.Context, ship *api.Ship, produce api.TradeSymbol) (*api.ShipRefine201ResponseData, error) {
	wait := 1 * time.Second
	srr := api.NewShipRefineRequest(string(produce))
	req := c.c.FleetAPI.ShipRefine(ctx, ship.Symbol).ShipRefineRequest(*srr)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.ShipRefine", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not ShipRefine: %w", err)
	}
	ship.Cooldown = res.Data.Cooldown
	ship.Cargo = res.Data.Cargo
	c.emitShipChange(ctx, ship)
	return &res.Data, nil
}

// CreateChart
func (c *FleetAPI) CreateChart(ctx context.Context, ship *api.Ship) (*api.CreateChart201ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.CreateChart(ctx, ship.Symbol)
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
		return nil, fmt.Errorf("could not CreateChart: %w", err)
	}
	c.emitWaypoint(ctx, &res.Data.Waypoint)
	return &res.Data, nil
}

// UpdateCooldown returns cooldown data
func (c *FleetAPI) UpdateCooldown(ctx context.Context, ship *api.Ship) error {
	wait := 1 * time.Second
	req := c.c.FleetAPI.GetShipCooldown(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if httpR.StatusCode == 204 {
		return nil
	}
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.GetShipCooldown", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return fmt.Errorf("could not GetShipCooldown: %w", err)
	}
	ship.Cooldown = res.Data
	c.emitShipChange(ctx, ship)
	return nil
}

// DockShip to a waypoint
func (c *FleetAPI) DockShip(ctx context.Context, ship *api.Ship) error {
	wait := 1 * time.Second
	req := c.c.FleetAPI.DockShip(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.DockShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return fmt.Errorf("could not DockShip: %w", err)
	}
	ship.Nav = res.Data.Nav
	c.emitShipChange(ctx, ship)
	return nil
}

// CreateSurvey returns new survey data
func (c *FleetAPI) CreateSurvey(ctx context.Context, ship *api.Ship) (*api.CreateSurvey201ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.CreateSurvey(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) || errors.Is(err, ErrHTTPStatus409) {
		c.l.DebugContext(ctx, "hit error", "ops", "FleetAPI.CreateSurvey", "error", err, "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not CreateSurvey: %w", err)
	}
	ship.Cooldown = res.Data.Cooldown
	c.emitShipChange(ctx, ship)
	return &res.Data, nil
}

// ExtractResources
func (c *FleetAPI) ExtractResources(ctx context.Context, ship *api.Ship) (*api.ExtractResources201ResponseData, error) {
	wait := 1 * time.Second
	erReq := api.NewExtractResourcesRequestWithDefaults()
	req := c.c.FleetAPI.ExtractResources(ctx, ship.Symbol).ExtractResourcesRequest(*erReq)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) || errors.Is(err, ErrHTTPStatus409) {
		c.l.DebugContext(ctx, "hit error", "ops", "FleetAPI.ExtractResources", "error", err, "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not ExtractResources: %w", err)
	}
	ship.Cargo = res.Data.Cargo
	ship.Cooldown = res.Data.Cooldown
	c.emitShipChange(ctx, ship)
	c.emitExtraction(ctx, &res.Data.Extraction)
	for _, e := range res.Data.Events {
		sce := e.ShipConditionEvent
		c.emitShipCondition(ctx, sce)
	}
	return &res.Data, nil
}

// SiphonResources
func (c *FleetAPI) SiphonResources(ctx context.Context, ship *api.Ship) (*api.SiphonResources201ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.SiphonResources(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) || errors.Is(err, ErrHTTPStatus409) {
		c.l.DebugContext(ctx, "hit error", "ops", "FleetAPI.SiphonResources", "error", err, "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not SiphonResources: %w", err)
	}
	ship.Cooldown = res.Data.Cooldown
	c.emitShipChange(ctx, ship)
	for _, e := range res.Data.Events {
		sce := e.ShipConditionEvent
		c.emitShipCondition(ctx, sce)
	}
	c.emitSiphon(ctx, &res.Data.Siphon)
	return &res.Data, nil
}

// ExtractResourcesWithSurvey
func (c *FleetAPI) ExtractResourcesWithSurvey(ctx context.Context, ship *api.Ship, survey *api.Survey) (*api.ExtractResources201ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.ExtractResourcesWithSurvey(ctx, ship.Symbol).Survey(*survey)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) || errors.Is(err, ErrHTTPStatus409) {
		c.l.DebugContext(ctx, "hit error", "ops", "FleetAPI.ExtractResourcesWithSurvey", "error", err, "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not ExtractResourcesWithSurvey: %w", err)
	}
	ship.Cooldown = res.Data.Cooldown
	ship.Cargo = res.Data.Cargo
	c.emitExtraction(ctx, &res.Data.Extraction)
	c.emitShipChange(ctx, ship)
	for _, e := range res.Data.Events {
		sce := e.ShipConditionEvent
		c.emitShipCondition(ctx, sce)
	}
	return &res.Data, nil
}

// Jettison Cargo ejects the given units of TradeSymbol.
// If units is -1 all units are ejected
func (c *FleetAPI) JettisonCargo(ctx context.Context, ship *api.Ship, good api.TradeSymbol, units int32) error {
	if units == -1 {
		found := false
		for _, inv := range ship.Cargo.Inventory {
			if inv.Symbol == good {
				units = inv.Units
				found = true
			}
		}
		if !found {
			return ErrShipCargoMissing
		}
	}
	jettisonRequest := *api.NewJettisonRequest(good, units)
	wait := 1 * time.Second
	req := c.c.FleetAPI.Jettison(ctx, ship.Symbol).JettisonRequest(jettisonRequest)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.JettisonCargo", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return fmt.Errorf("could not JettisonCargo: %w", err)
	}
	ship.Cargo = res.Data.Cargo
	c.emitShipChange(ctx, ship)
	return nil
}

// JumpShip
func (c *FleetAPI) JumpShip(ctx context.Context, ship *api.Ship, wp string) (*api.JumpShip200ResponseData, error) {
	wait := 1 * time.Second
	jsr := api.NewJumpShipRequest(wp)
	req := c.c.FleetAPI.JumpShip(ctx, ship.Symbol).JumpShipRequest(*jsr)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.JumpShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not JumpShip: %w", err)
	}
	ship.Cooldown = res.Data.Cooldown
	ship.Nav = res.Data.Nav
	c.emitShipChange(ctx, ship)
	c.emitAgentChange(ctx, &res.Data.Agent)
	c.emitTransaction(ctx, &res.Data.Transaction)
	return &res.Data, nil
}

// NavigateShip to a waypoint
func (c *FleetAPI) NavigateShip(ctx context.Context, ship *api.Ship, waypoint string) (*api.NavigateShip200ResponseData, error) {
	nReq := api.NewNavigateShipRequest(waypoint)
	wait := 1 * time.Second
	req := c.c.FleetAPI.NavigateShip(ctx, ship.Symbol).NavigateShipRequest(*nReq)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.NavigateShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not NavigateShip: %w", err)
	}
	ship.Fuel = res.Data.Fuel
	ship.Nav = res.Data.Nav
	c.emitShipChange(ctx, ship)
	for _, e := range res.Data.Events {
		sce := e.ShipConditionEvent
		c.emitShipCondition(ctx, sce)
	}
	return &res.Data, nil
}

// PatchShipNav
func (c *FleetAPI) PatchShipNav(ctx context.Context, ship *api.Ship, fm api.ShipNavFlightMode) error {
	wait := 1 * time.Second
	psr := api.NewPatchShipNavRequest()
	psr.SetFlightMode(fm)
	req := c.c.FleetAPI.PatchShipNav(ctx, ship.Symbol).PatchShipNavRequest(*psr)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) || errors.Is(err, ErrHTTPStatus409) {
		c.l.DebugContext(ctx, "hit error", "ops", "FleetAPI.PatchShipNav", "error", err, "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return fmt.Errorf("could not PatchShipNav: %w", err)
	}
	ship.Nav.FlightMode = res.Data.FlightMode
	ship.Nav.Status = res.Data.Status
	ship.Nav.Route = res.Data.Route
	ship.Nav.SystemSymbol = res.Data.SystemSymbol
	ship.Nav.WaypointSymbol = res.Data.WaypointSymbol
	c.emitShipChange(ctx, ship)
	return nil
}

// GetShipNav
func (c *FleetAPI) GetShipNav(ctx context.Context, ship *api.Ship) error {
	wait := 1 * time.Second
	req := c.c.FleetAPI.GetShipNav(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) || errors.Is(err, ErrHTTPStatus409) {
		c.l.DebugContext(ctx, "hit error", "ops", "FleetAPI.PatchShipNav", "error", err, "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return fmt.Errorf("could not PatchShipNav: %w", err)
	}
	ship.Nav = res.Data
	c.emitShipChange(ctx, ship)
	return nil
}

// WarpShip
func (c *FleetAPI) WarpShip(ctx context.Context, ship *api.Ship, wpSym WaypointSymbol) (*api.WarpShip200ResponseData, error) {
	wait := 1 * time.Second
	nsr := api.NewNavigateShipRequest(wpSym.String())
	req := c.c.FleetAPI.WarpShip(ctx, ship.Symbol).NavigateShipRequest(*nsr)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.JumpShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not JumpShip: %w", err)
	}
	ship.Nav = res.Data.Nav
	ship.Fuel = res.Data.Fuel
	c.emitShipChange(ctx, ship)
	return &res.Data, nil
}

// SellCargo sells the given units of TradeSymbol at the local market.
// If units is -1 all units are sold
func (c *FleetAPI) SellCargo(ctx context.Context, ship *api.Ship, good api.TradeSymbol, units int32) (*api.SellCargo201ResponseData, error) {
	if units == -1 {
		found := false
		for _, inv := range ship.Cargo.Inventory {
			if inv.Symbol == good {
				units = inv.Units
				found = true
			}
		}
		if !found {
			return nil, ErrShipCargoMissing
		}
	}
	sellCargoRequest := *api.NewSellCargoRequest(good, units)
	wait := 1 * time.Second

	req := c.c.FleetAPI.SellCargo(ctx, ship.Symbol).SellCargoRequest(sellCargoRequest)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.SellCargo", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not SellCargo: %w", err)
	}
	ship.Cargo = res.Data.Cargo
	c.emitShipChange(ctx, ship)
	c.emitAgentChange(ctx, &res.Data.Agent)
	c.emitTransaction(ctx, &res.Data.Transaction)
	return &res.Data, nil
}

// Scan Systems
func (c *FleetAPI) ScanSystems(ctx context.Context, ship *api.Ship) (*api.CreateShipSystemScan201ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.CreateShipSystemScan(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.RefuelShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not RefuelShip: %w", err)
	}
	ship.Cooldown = res.Data.Cooldown
	c.emitShipChange(ctx, ship)
	return &res.Data, nil
}

// Scan Waypoints
func (c *FleetAPI) ScanWaypoints(ctx context.Context, ship *api.Ship) (*api.CreateShipWaypointScan201ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.CreateShipWaypointScan(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.RefuelShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not RefuelShip: %w", err)
	}
	ship.Cooldown = res.Data.Cooldown
	c.emitShipChange(ctx, ship)
	return &res.Data, nil
}

// Scan Ships
func (c *FleetAPI) ScanShips(ctx context.Context, ship *api.Ship) (*api.CreateShipShipScan201ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.CreateShipShipScan(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.RefuelShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not RefuelShip: %w", err)
	}
	ship.Cooldown = res.Data.Cooldown
	c.emitShipChange(ctx, ship)
	return &res.Data, nil
}

// Refuels a ship.
// If fromCargo is false, fuel is bought from the local market.
// If units<0, ship is refueled to the maximum.
func (c *FleetAPI) RefuelShip(ctx context.Context, ship *api.Ship, fromCargo bool, units int32) (*api.RefuelShip200ResponseData, error) {
	rsr := api.NewRefuelShipRequest()
	rsr.FromCargo = &fromCargo
	if units > 0 {
		rsr.Units = &units
	}
	wait := 1 * time.Second
	req := c.c.FleetAPI.RefuelShip(ctx, ship.Symbol).RefuelShipRequest(*rsr)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.RefuelShip", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not RefuelShip: %w", err)
	}
	c.emitAgentChange(ctx, &res.Data.Agent)
	ship.Fuel = res.Data.Fuel
	c.emitTransaction(ctx, &res.Data.Transaction)
	return &res.Data, nil
}

// PurchaseCargo buys the given units of TradeSymbol at the local market.
func (c *FleetAPI) PurchaseCargo(ctx context.Context, ship *api.Ship, good api.TradeSymbol, units int32) (*api.SellCargo201ResponseData, error) {
	wait := 1 * time.Second
	purchaseCargoRequest := *api.NewPurchaseCargoRequest(good, units)
	req := c.c.FleetAPI.PurchaseCargo(ctx, ship.Symbol).PurchaseCargoRequest(purchaseCargoRequest)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.PurchaseCargo", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not PurchaseCargo: %w", err)
	}
	ship.Cargo = res.Data.Cargo
	c.emitShipChange(ctx, ship)
	c.emitAgentChange(ctx, &res.Data.Agent)
	c.emitTransaction(ctx, &res.Data.Transaction)
	return &res.Data, nil
}

// Transfer Cargo transfer the given units of TradeSymbol from source to to target.
// If units is -1 all units are transferred.
func (c *FleetAPI) TransferCargo(ctx context.Context, target, source *api.Ship, good api.TradeSymbol, units int32) error {
	if units == -1 {
		found := false
		for _, inv := range source.Cargo.Inventory {
			if inv.Symbol == good {
				units = inv.Units
				found = true
			}
		}
		if !found {
			return ErrShipCargoMissing
		}
	}
	transferCargoRequest := *api.NewTransferCargoRequest(good, units, target.Symbol)
	wait := 1 * time.Second
	req := c.c.FleetAPI.TransferCargo(ctx, source.Symbol).TransferCargoRequest(transferCargoRequest)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit", "ops", "FleetAPI.TransferCargo", "wait", wait)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return fmt.Errorf("could not TransferCargo: %w", err)
	}
	source.Cargo = res.Data.Cargo
	c.emitShipChange(ctx, source)
	found := false
	for i, inv := range target.Cargo.Inventory {
		if inv.Symbol != good {
			continue
		}
		target.Cargo.Inventory[i].Units += units
		found = true
		break
	}
	if !found {
		inv := api.NewShipCargoItem(good, "", "", units)
		target.Cargo.Inventory = append(target.Cargo.Inventory, *inv)
	}
	target.Cargo.Units += units
	c.emitShipChange(ctx, target)
	return nil
}

// NegotiateContract
func (c *FleetAPI) NegotiateContract(ctx context.Context, ship *api.Ship) (*api.Contract, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.NegotiateContract(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit",
			"ops", "FleetAPI.NegotiateContract",
			"wait", wait,
			"ship", ship.Symbol)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not NegotiateContract: %w", err)
	}
	c.emitContract(ctx, &res.Data.Contract)
	return &res.Data.Contract, nil
}

// GetMounts gets the mounts installed on a ship.
func (c *FleetAPI) GetMounts(ctx context.Context, ship *api.Ship) ([]api.ShipMount, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.GetMounts(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit",
			"ops", "FleetAPI.GetMounts",
			"wait", wait,
			"ship", ship.Symbol)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetMounts: %w", err)
	}
	return res.Data, nil
}

// InstallMount installs a mount on a ship.
func (c *FleetAPI) InstallMount(ctx context.Context, ship *api.Ship, mountSym string) (*api.InstallMount201ResponseData, error) {
	wait := 1 * time.Second
	imr := api.NewInstallMountRequest(mountSym)
	req := c.c.FleetAPI.InstallMount(ctx, ship.Symbol).InstallMountRequest(*imr)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit",
			"ops", "FleetAPI.InstallMount",
			"wait", wait,
			"ship", ship.Symbol)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not InstallMount: %w", err)
	}
	c.emitAgentChange(ctx, &res.Data.Agent)
	ship.Mounts = res.Data.Mounts
	ship.Cargo = res.Data.Cargo
	c.emitShipChange(ctx, ship)
	c.emitShipModificationTransaction(ctx, &res.Data.Transaction)
	return &res.Data, nil
}

// RemoveMount removes a mount from a ship.
func (c *FleetAPI) RemoveMount(ctx context.Context, ship *api.Ship, mountSym string) (*api.RemoveMount201ResponseData, error) {
	wait := 1 * time.Second
	imr := api.NewRemoveMountRequest(mountSym)
	req := c.c.FleetAPI.RemoveMount(ctx, ship.Symbol).RemoveMountRequest(*imr)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit",
			"ops", "FleetAPI.RemoveMount",
			"wait", wait,
			"ship", ship.Symbol)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not InstallMount: %w", err)
	}
	c.emitAgentChange(ctx, &res.Data.Agent)
	ship.Mounts = res.Data.Mounts
	ship.Cargo = res.Data.Cargo
	c.emitShipChange(ctx, ship)
	c.emitShipModificationTransaction(ctx, &res.Data.Transaction)
	return &res.Data, nil
}

// GetScrapShipValue gets the amount of value that will be returned when scrapping a ship.
// The ship must be docked in a waypoint that has the Shipyard trait in order to use
// this function.
func (c *FleetAPI) GetScrapShipValue(ctx context.Context, ship *api.Ship) (*api.GetScrapShip200ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.GetScrapShip(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit",
			"ops", "FleetAPI.GetScrapShip",
			"wait", wait,
			"ship", ship.Symbol)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetScrapShipValue: %w", err)
	}
	return &res.Data, nil
}

// ScrapShip removes a ship from the game
// The ship must be docked in a waypoint that has the Shipyard trait in order to use
// this function.
func (c *FleetAPI) ScrapShip(ctx context.Context, ship *api.Ship) (*api.ScrapShip200ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.ScrapShip(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit",
			"ops", "FleetAPI.ScrapShip",
			"wait", wait,
			"ship", ship.Symbol)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not GetScrapShipValue: %w", err)
	}
	c.emitAgentChange(ctx, &res.Data.Agent)
	c.emitScrapTransaction(ctx, &res.Data.Transaction)
	return &res.Data, nil
}

// GetRepairShipCost
func (c *FleetAPI) GetRepairShipCost(ctx context.Context, ship *api.Ship) (*api.RepairTransaction, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.GetRepairShip(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit",
			"ops", "FleetAPI.GetRepairShip",
			"wait", wait,
			"ship", ship.Symbol)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not RepairShip: %w", err)
	}
	return &res.Data.Transaction, nil
}

// RepairShip
func (c *FleetAPI) RepairShip(ctx context.Context, ship *api.Ship) (*api.RepairShip200ResponseData, error) {
	wait := 1 * time.Second
	req := c.c.FleetAPI.RepairShip(ctx, ship.Symbol)
redo:
	res, httpR, err := req.Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		c.l.DebugContext(ctx, "hit ratelimit",
			"ops", "FleetAPI.RepairShip",
			"wait", wait,
			"ship", ship.Symbol)
		if err := c.SleepWithJitter(ctx, wait); err != nil {
			return nil, err
		}
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, fmt.Errorf("could not RepairShip: %w", err)
	}
	CopyShipState(ship, &res.Data.Ship)
	c.emitShipChange(ctx, ship)
	c.emitAgentChange(ctx, &res.Data.Agent)
	c.emitRepairTransaction(ctx, &res.Data.Transaction)
	return &res.Data, nil
}

// CopyShipState changes all public properties of to to the values of from
func CopyShipState(to, from *api.Ship) {
	to.Symbol = from.Symbol
	to.Cargo = from.Cargo
	to.Cooldown = from.Cooldown
	to.Crew = from.Crew
	to.Engine = from.Engine
	to.Frame = from.Frame
	to.Fuel = from.Fuel
	to.Modules = from.Modules
	to.Mounts = from.Mounts
	to.Nav = from.Nav
	to.Reactor = from.Reactor
	to.Registration = from.Registration
}

// CopyAgent creates a new copy of the given agent
func CopyAgent(from *api.Agent) *api.Agent {
	to := api.NewAgentWithDefaults()
	to.AccountId = from.AccountId
	to.Credits = from.Credits
	to.Headquarters = from.Headquarters
	to.ShipCount = from.ShipCount
	to.StartingFaction = from.StartingFaction
	to.Symbol = from.Symbol
	return to
}
