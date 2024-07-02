package s10s

import (
	"context"
	"errors"
	"time"

	api "github.com/ruudiRatlos/s10s/openapi"
)

// Register registers a new agent
func Register(faction, name string, email string) (*api.Register201ResponseData, error) {
	registerRequest := *api.NewRegisterRequest(api.FactionSymbol(faction), name)
	registerRequest.Email = &email
	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)
	resp, httpR, err := apiClient.DefaultAPI.Register(context.Background()).RegisterRequest(registerRequest).Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// GetStatus returns the current status
func GetStatus() (*api.GetStatus200Response, error) {
	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)
	wait := 1 * time.Second
redo:
	resp, httpR, err := apiClient.DefaultAPI.GetStatus(context.Background()).Execute() //nolint:bodyclose
	err = enhanceErr(err, httpR)
	if errors.Is(err, ErrHTTPStatus429) {
		time.Sleep(wait)
		wait = 2 * wait
		goto redo
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
