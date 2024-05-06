package s10s

import (
	"context"
	"log/slog"
	"math/rand"
	"time"

	api "github.com/ruudiRatlos/s10s/openapi"
)

func (bc *baseClient) ContextWithToken(ctx context.Context) context.Context {
	return contextWithToken(ctx, bc.token)
}

func contextWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, api.ContextAccessToken, token)
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

type AgentsAPI service
type SystemsAPI service
type FleetAPI service
type ContractsAPI service
type FactionsAPI service

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
	jitter := rand.Intn(int(wait.Milliseconds()*10) / 100) //nolint:gosec
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
