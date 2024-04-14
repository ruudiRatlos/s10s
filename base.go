package s10s

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"strings"
	"time"

	api "github.com/ruudiRatlos/s10s/api"
)

func (bc *baseClient) ContextWithToken(ctx context.Context) context.Context {
	return contextWithToken(ctx, bc.token)
}

func contextWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, api.ContextAccessToken, token)
}

// SystemSymbol identifies a system, containing one dash
type SystemSymbol string

// WaypointSymbol identifies a waypoint with a specific system, prefixed with their respective SystemSymbol
type WaypointSymbol string

func NewWaypointSymbol(wp string) (WaypointSymbol, error) {
	parts := strings.SplitN(string(wp), "-", 3)
	if len(parts) != 3 || strings.Count(wp, "-") != 2 {
		return "", fmt.Errorf("%q does not contain 2 dashes. ", wp)
	}
	return WaypointSymbol(wp), nil
}

func MustNewWaypointSymbol(wp string) WaypointSymbol {
	wayp, err := NewWaypointSymbol(wp)
	if err != nil {
		panic(err)
	}
	return wayp
}

func NewSystemSymbol(sys string) SystemSymbol {
	return SystemSymbol(sys)
}

func (sys SystemSymbol) String() string {
	return string(sys)
}

func (wp WaypointSymbol) SystemSymbol() SystemSymbol {
	parts := strings.SplitN(string(wp), "-", 3)
	return SystemSymbol(fmt.Sprintf("%s-%s", parts[0], parts[1]))
}

func (wp WaypointSymbol) String() string {
	return string(wp)
}

type baseClient struct {
	c  *api.APIClient
	l  *slog.Logger
	em FireEventer

	token string
	cfg   *api.Configuration
}

type service struct {
	*baseClient
}

type agentsAPI service
type systemsAPI service
type fleetAPI service
type contractsAPI service
type factionsAPI service

// Sleep for the given duration unless interrupted by a canceled Context
func (s *baseClient) Sleep(ctx context.Context, wait time.Duration) error {
	select {
	case <-time.After(wait):
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

// SleepWithJitter sleeps for at least the given duration and adds a random increment up to 10% of the given duration unless interrupted by a canceled Context
func (s *baseClient) SleepWithJitter(ctx context.Context, wait time.Duration) error {
	jitter := rand.Intn(int(wait.Milliseconds())) //nolint:gosec
	wait = wait + time.Duration(jitter)*time.Millisecond
	select {
	case <-time.After(wait):
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func (s *baseClient) SleepUntil(ctx context.Context, t time.Time) error {
	select {
	case <-time.After(time.Until(t)):
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}
