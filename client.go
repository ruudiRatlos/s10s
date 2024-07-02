package s10s

import (
	"context"
	"errors"
	"time"

	api "github.com/ruudiRatlos/s10s/openapi"
)

// Client represents a retrying REST API Client for spacetraders
type Client struct {
	*baseClient

	AgentsAPI    *AgentsAPI
	SystemsAPI   *SystemsAPI
	FleetAPI     *FleetAPI
	ContractsAPI *ContractsAPI
	FactionsAPI  *FactionsAPI
}

// NewClient creates a new Spacetraders client
func NewClient(opts ...ClientOption) (*Client, error) {
	bc := &baseClient{
		cfg: api.NewConfiguration(),
	}
	bc.cfg.UserAgent = "s10s/v1/go"

	c := &Client{
		baseClient:   bc,
		AgentsAPI:    &AgentsAPI{bc},
		SystemsAPI:   &SystemsAPI{bc},
		FleetAPI:     &FleetAPI{bc},
		FactionsAPI:  &FactionsAPI{bc},
		ContractsAPI: &ContractsAPI{bc},
	}

	if len(opts) == 0 {
		opts = DefaultOptions
	}

	for _, o := range opts {
		err := o(c)
		if err != nil {
			return nil, err
		}
	}

	bc.c = api.NewAPIClient(bc.cfg)

	ctx := contextWithToken(context.Background(), c.token)

	wait := 1 * time.Second
redo:
	_, httpR, err := c.c.DefaultAPI.GetStatus(ctx).Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		time.Sleep(wait)
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}
