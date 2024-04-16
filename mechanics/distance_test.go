package mechanics

import (
	"testing"

	api "github.com/ruudiRatlos/s10s/openapi"
)

func TestDistance(t *testing.T) {
	a := &api.Waypoint{X: -1, Y: 29}
	b := &api.Waypoint{X: 17, Y: 43}

	dist := Distance(a, b)
	if int(dist) != 22 {
		t.Errorf("got: %v, want: %v", dist, 22)
	}
}

func TestDistanceFromCenter(t *testing.T) {
	testcases := []struct {
		wp Coorder
		d  int
	}{
		{&api.Waypoint{X: -1, Y: 29}, 29},
		{&api.System{X: 17, Y: 43}, 46},
		{&api.System{X: 0, Y: 0}, 0},
	}
	for _, tc := range testcases {
		g := int(DistanceFromCenter(tc.wp))
		if g != tc.d {
			t.Errorf("got: %v, want: %v", g, tc.d)
		}
	}
}
