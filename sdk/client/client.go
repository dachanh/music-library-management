package client

import (
	"music-library-management/sdk/common"
	"time"

	"music-library-management/sdk"
)

// APIClient
type APIClient interface {
	MakeRequest(sdk.APIRequest) *common.APIResponse
	SetDebug(val bool)
}

// APIClientConfiguration
type APIClientConfiguration struct {
	Address       string
	Protocol      string
	Timeout       time.Duration
	MaxRetry      int
	WaitToRetry   time.Duration
	LoggingCol    string
	LogExpiration *time.Duration
	MaxConnection int
	ErrorLogOnly  bool
}

func NewAPIClient(config *APIClientConfiguration) APIClient {
	switch config.Protocol {
	case "HTTP":
		return NewHTTPClient(config)
	}
	return nil
}
