package mechanics

import (
	"github.com/ruudiRatlos/s10s/openapi"
)

// IsMineableOre indicates if an TradeSymbol can be obtained from mining asteroids
func IsMineableOre(good openapi.TradeSymbol) bool {
	_, r := TraitFromOre(good)
	return r
}

// TraitFromOre returns the TraitSymbol on Asteroids to mine at.
// the returned bool indicates whether an ore is mineable or not.
func TraitFromOre(good openapi.TradeSymbol) (openapi.WaypointTraitSymbol, bool) {
	switch good {
	case openapi.TRADESYMBOL_SILICON_CRYSTALS,
		openapi.TRADESYMBOL_QUARTZ_SAND:
		return openapi.WAYPOINTTRAITSYMBOL_MINERAL_DEPOSITS, true
	case openapi.TRADESYMBOL_IRON_ORE,
		openapi.TRADESYMBOL_COPPER_ORE,
		openapi.TRADESYMBOL_ALUMINUM_ORE:
		return openapi.WAYPOINTTRAITSYMBOL_COMMON_METAL_DEPOSITS, true
	case openapi.TRADESYMBOL_SILVER_ORE,
		openapi.TRADESYMBOL_GOLD_ORE,
		openapi.TRADESYMBOL_PLATINUM_ORE,
		openapi.TRADESYMBOL_URANITE_ORE,
		openapi.TRADESYMBOL_MERITIUM_ORE:
		return openapi.WAYPOINTTRAITSYMBOL_PRECIOUS_METAL_DEPOSITS, true
	default:
		// there is no zero value for WpTS
		return openapi.WAYPOINTTRAITSYMBOL_STRIPPED, false
	}
}
