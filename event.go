package s10s

import (
	"context"

	api "github.com/ruudiRatlos/s10s/openapi"
)

type FireEventer interface {
	FireEvent(name string, data map[string]any) error
}

const (
	EvAgentChange                    string = "s10s.agent.change"
	EvConstructionChange             string = "s10s.construction.change"
	EvTransactionNew                 string = "s10s.transaction.new"
	EvExtractionNew                  string = "s10s.extraction.new"
	EvSiphonNew                      string = "s10s.siphon.new"
	EvShipChange                     string = "s10s.ship.change"
	EvShipCondition                  string = "s10s.ship.condition.event.new"
	EvWaypointChange                 string = "s10s.system.waypoint.change"
	EvContractChange                 string = "s10s.contract.change"
	EvSystem                         string = "s10s.system.change"
	EvJumpGate                       string = "s10s.jumpgate.change"
	EvMarket                         string = "s10s.market.success"
	EvShipyard                       string = "s10s.shipyard.change"
	EvShipyardTransactionNew         string = "s10s.shipyard.transaction.new"
	EvShipModificationTransactionNew string = "s10s.ship.transaction.new"
	EvScrapTransactionNew            string = "s10s.ship.scrap.new"
	EvRepairTransactionNew           string = "s10s.ship.repair.new"
)

func (bc *baseClient) fireEvent(name string, data map[string]any) error {
	if bc.em == nil {
		return nil
	}
	return bc.em.FireEvent(name, data)
}

func (bc *baseClient) emitShipyard(ctx context.Context, sy *api.Shipyard) {
	err := bc.fireEvent(EvShipyard, map[string]any{
		"Context":  ctx,
		"Shipyard": sy,
	})
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", "error", err)
	}
}

func (bc *baseClient) emitMarket(ctx context.Context, m *api.Market) {
	d := map[string]any{
		"Context": ctx,
		"Market":  m,
	}
	err := bc.fireEvent(EvMarket, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", "error", err)
	}
}

func (bc *baseClient) emitWaypoint(ctx context.Context, wp *api.Waypoint) {
	d := map[string]any{
		"Context":  ctx,
		"Waypoint": wp,
	}
	err := bc.fireEvent(EvWaypointChange, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", "error", err)
	}
}

func (bc *baseClient) emitShipChange(ctx context.Context, ship *api.Ship) {
	d := map[string]any{
		"Context": ctx,
		"Ship":    ship,
	}
	err := bc.fireEvent(EvShipChange, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", "error", err)
	}
}

func (bc *baseClient) emitShipCondition(ctx context.Context, sce *api.ShipConditionEvent) {
	d := map[string]any{
		"Context": ctx,
		"Event":   sce,
	}
	err := bc.fireEvent(EvShipCondition, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitTransaction(ctx context.Context, tx *api.MarketTransaction) {
	d := map[string]any{
		"Context": ctx,
		"Tx":      tx,
	}
	err := bc.fireEvent(EvTransactionNew, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitShipyardTransaction(ctx context.Context, tx *api.ShipyardTransaction) {
	d := map[string]any{
		"Context": ctx,
		"Tx":      tx,
	}
	err := bc.fireEvent(EvShipyardTransactionNew, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitShipModificationTransaction(ctx context.Context, tx *api.ShipModificationTransaction) {
	d := map[string]any{
		"Context": ctx,
		"Tx":      tx,
	}
	err := bc.fireEvent(EvShipModificationTransactionNew, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitScrapTransaction(ctx context.Context, tx *api.ScrapTransaction) {
	d := map[string]any{
		"Context": ctx,
		"Tx":      tx,
	}
	err := bc.fireEvent(EvScrapTransactionNew, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitSystem(ctx context.Context, sys *api.System) {
	d := map[string]any{
		"Context": ctx,
		"System":  sys,
	}
	err := bc.fireEvent(EvSystem, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitConstruction(ctx context.Context, ctr *api.Construction) {
	d := map[string]any{
		"Context":      ctx,
		"Construction": ctr,
	}
	err := bc.fireEvent(EvConstructionChange, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitContract(ctx context.Context, contract *api.Contract) {
	d := map[string]any{
		"Context":  ctx,
		"Contract": contract,
	}
	err := bc.fireEvent(EvContractChange, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitJumpGate(ctx context.Context, jg *api.JumpGate) {
	d := map[string]any{
		"Context":  ctx,
		"JumpGate": jg,
	}
	err := bc.fireEvent(EvJumpGate, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitExtraction(ctx context.Context, ex *api.Extraction) {
	d := map[string]any{
		"Context": ctx,
		"Ex":      ex,
	}
	err := bc.fireEvent(EvExtractionNew, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitSiphon(ctx context.Context, siphon *api.Siphon) {
	d := map[string]any{
		"Context": ctx,
		"Siphon":  siphon,
	}
	err := bc.fireEvent(EvSiphonNew, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitAgentChange(ctx context.Context, agent *api.Agent) {
	d := map[string]any{
		"Context": ctx,
		"Agent":   agent,
	}
	err := bc.fireEvent(EvAgentChange, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}

func (bc *baseClient) emitRepairTransaction(ctx context.Context, tx *api.RepairTransaction) {
	d := map[string]any{
		"Context": ctx,
		"Tx":      tx,
	}
	err := bc.fireEvent(EvRepairTransactionNew, d)
	if err != nil {
		bc.l.WarnContext(ctx, "event emit", "e", d, "error", err)
	}
}
