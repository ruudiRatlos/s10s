package s10s

import (
	"context"

	api "github.com/ruudiRatlos/s10s/api"
)

// Client represents a retrying REST API Client for spacetraders
type Client struct {
	*baseClient

	AgentsAPI    *agentsAPI
	SystemsAPI   *systemsAPI
	FleetAPI     *fleetAPI
	ContractsAPI *contractsAPI
	FactionsAPI  *factionsAPI
}

// NewClient creates a new Spacetraders client
func NewClient(opts ...ClientOption) (*Client, error) {
	bc := &baseClient{
		cfg: api.NewConfiguration(),
	}
	bc.cfg.UserAgent = "s10s/v1/go"

	c := &Client{
		baseClient:   bc,
		AgentsAPI:    &agentsAPI{bc},
		SystemsAPI:   &systemsAPI{bc},
		FleetAPI:     &fleetAPI{bc},
		FactionsAPI:  &factionsAPI{bc},
		ContractsAPI: &contractsAPI{bc},
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

	_, httpR, err := c.c.DefaultAPI.GetStatus(ctx).Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if err != nil {
		return nil, err
	}

	return c, nil
}
