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

func TestHeading(t *testing.T) {
	testcases := []struct {
		dir string
		wp  Coorder
		h   int
	}{
		{"N", &api.System{X: 0, Y: 1}, 0},
		{"NE", &api.System{X: 1, Y: 1}, 45},
		{"E", &api.System{X: 1, Y: 0}, 90},
		{"SE", &api.Waypoint{X: 1, Y: -1}, 135},
		{"S", &api.Waypoint{X: 0, Y: -1}, 180},
		{"SW", &api.Waypoint{X: -1, Y: -1}, 225},
		{"W", &api.Waypoint{X: -1, Y: 0}, 270},
		{"NW", &api.Waypoint{X: -1, Y: 1}, 315},

		{"0", &api.Waypoint{X: 0, Y: 0}, 0},
	}
	for _, tc := range testcases {
		t.Run(tc.dir, func(t *testing.T) {
			g := Heading(tc.wp)
			if g != tc.h {
				t.Errorf("x=%d y=%d got: %v, want: %v",
					tc.wp.GetX(), tc.wp.GetY(), g, tc.h)
			}
		})
	}
}
