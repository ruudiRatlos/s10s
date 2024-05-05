package s10s

import (
	"fmt"
	"strings"

	api "github.com/ruudiRatlos/s10s/openapi"
)

// SystemSymbol identifies a system, containing one dash
type SystemSymbol string

// WaypointSymbol identifies a waypoint with a specific system, prefixed with their respective SystemSymbol
type WaypointSymbol string

func WaypointSymbolFromValue(wp string) (WaypointSymbol, error) {
	parts := strings.SplitN(string(wp), "-", 3)
	if len(parts) != 3 || strings.Count(wp, "-") != 2 {
		return "", fmt.Errorf("%q does not contain 2 dashes. ", wp)
	}
	return WaypointSymbol(wp), nil
}

func MustWaypointSymbolFromValue(wp string) WaypointSymbol {
	wayp, err := WaypointSymbolFromValue(wp)
	if err != nil {
		panic(err)
	}
	return wayp
}

func SystemSymbolFromValue(sys string) SystemSymbol {
	return SystemSymbol(sys)
}

func (sys SystemSymbol) String() string {
	return string(sys)
}

func (sys SystemSymbol) Equals(sys2 SystemSymbol) bool {
	return string(sys) == string(sys2)
}

func (wp WaypointSymbol) SystemSymbol() SystemSymbol {
	parts := strings.SplitN(string(wp), "-", 3)
	return SystemSymbol(fmt.Sprintf("%s-%s", parts[0], parts[1]))
}

func (wp WaypointSymbol) String() string {
	return string(wp)
}

func (wp WaypointSymbol) Equals(wp2 WaypointSymbol) bool {
	return string(wp) == string(wp2)
}

type SysLocationeers interface {
	*api.System | *api.Construction | *api.Waypoint | *api.Shipyard | *api.Agent | *api.JumpGate | *api.Ship
	GetSymbol() string
}

func SystemSymbolFrom[T SysLocationeers](v T) SystemSymbol {
	if a, ok := any(v).(*api.Agent); ok {
		return MustWaypointSymbolFromValue(a.Headquarters).SystemSymbol()
	}
	if s, ok := any(v).(*api.Ship); ok {
		return SystemSymbolFromValue(s.Nav.SystemSymbol)
	}
	if _, ok := any(v).(*api.System); ok {
		return SystemSymbolFromValue(v.GetSymbol())
	}
	return MustWaypointSymbolFromValue(v.GetSymbol()).SystemSymbol()
}

type Locationeers interface {
	*api.Construction | *api.Waypoint | *api.Shipyard | *api.Agent | *api.JumpGate | *api.Ship | *api.Survey | *api.Market
	GetSymbol() string
}

func WaypointSymbolFrom[T Locationeers](v T) WaypointSymbol {
	if a, ok := any(v).(*api.Agent); ok {
		return MustWaypointSymbolFromValue(a.Headquarters)
	}
	if s, ok := any(v).(*api.Ship); ok {
		return MustWaypointSymbolFromValue(s.Nav.WaypointSymbol)
	}
	return MustWaypointSymbolFromValue(v.GetSymbol())
}
