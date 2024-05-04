package mechanics

import (
	api "github.com/ruudiRatlos/s10s/openapi"
)

func FilterWaypoints(waypoints []*api.Waypoint, criteria api.WaypointTraitSymbol) []*api.Waypoint {
	out := []*api.Waypoint{}
	for _, wp := range waypoints {
		mpF := false
		for _, t := range wp.Traits {
			if t.Symbol != criteria {
				continue
			}
			mpF = true
		}
		if !mpF {
			continue
		}
		out = append(out, wp)
	}
	return out
}

func IsUnchartedWP(wp *api.Waypoint) bool {
	for _, t := range wp.Traits {
		if t.Symbol == api.WAYPOINTTRAITSYMBOL_UNCHARTED {
			return true
		}
	}
	return false
}
