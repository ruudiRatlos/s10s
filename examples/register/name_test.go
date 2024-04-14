package main

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	for _ = range 100_000 {
		g := randomName()
		if g == "" {
			t.Fail()
		}
		if !strings.HasSuffix(g, "-TESTER") {
			t.Errorf("got: %q, expected \"-TESTER\"", g)
		}
	}
}
