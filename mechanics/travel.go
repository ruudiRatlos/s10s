package mechanics

import (
	"time"

	"github.com/ruudiRatlos/s10s/api"
)

// CalcTravelTime returns the duration of travel for the given ship and distance
func CalcTravelTime(ship *api.Ship, dist int) time.Duration {
	return CalcTravelTimeRaw(ship.Engine.Speed, ship.Nav.FlightMode, dist)
}

// Source: https://github.com/SpaceTradersAPI/api-docs/wiki/Travel-Fuel-and-Time#travel-time
var timeMult = map[api.ShipNavFlightMode]float64{
	api.SHIPNAVFLIGHTMODE_CRUISE:  25,
	api.SHIPNAVFLIGHTMODE_DRIFT:   250,
	api.SHIPNAVFLIGHTMODE_BURN:    12.5,
	api.SHIPNAVFLIGHTMODE_STEALTH: 30,
}

// CalcTravelTimeRaw returns the duration of travel for the given speed, flightmode and distance
func CalcTravelTimeRaw(speed int32, fm api.ShipNavFlightMode, dist int) time.Duration {
	secs := int(float64(max(1, dist))*(timeMult[fm]/float64(speed)) + 15)
	return time.Second * time.Duration(secs)
}

// Source: https://github.com/SpaceTradersAPI/api-docs/wiki/Travel-Fuel-and-Time#fuel-costs
var fuelMult = map[api.ShipNavFlightMode]func(int) int{
	api.SHIPNAVFLIGHTMODE_CRUISE:  func(c int) int { return c },
	api.SHIPNAVFLIGHTMODE_DRIFT:   func(_ int) int { return 1 },
	api.SHIPNAVFLIGHTMODE_BURN:    func(c int) int { return 2 * c },
	api.SHIPNAVFLIGHTMODE_STEALTH: func(c int) int { return c },
}

// CalcTravelFuelCost returns the fuel consumption for a given distance and flightmode
func CalcTravelFuelCost(dist int, fm api.ShipNavFlightMode) int {
	return fuelMult[fm](dist)
}
