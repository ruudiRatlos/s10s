package s10s

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type ClientOption func(*Client) error

var DefaultOptions []ClientOption = []ClientOption{WithTokenFromEnv, WithoutLogging}

func WithLogger(logger *slog.Logger) ClientOption {
	return func(c *Client) error {
		c.l = logger
		return nil
	}
}

func WithEvents(em FireEventer) ClientOption {
	return func(c *Client) error {
		c.em = em
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		return nil
	}
}

const tokenEnvName = "S10S_API_TOKEN" //nolint:gosec

func WithTokenFromEnv(c *Client) error {
	token := os.Getenv(tokenEnvName)
	if token == "" {
		return fmt.Errorf("Cannot initialize, environment variable %q is unset", tokenEnvName)
	}
	c.token = token
	return nil
}

func WithToken(token string) func(*Client) error {
	return func(c *Client) error {
		if token == "" {
			return errors.New("Cannot initialize, token is empty")
		}
		c.token = token
		return nil
	}
}

func WithoutLogging(c *Client) error {
	c.l = slog.New(&discardHandler{})
	return nil
}
