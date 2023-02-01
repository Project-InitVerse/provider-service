package ubicutil

import (
	"context"
	"io"
	"net/http"
)

// ServiceDiscoveryAgent is interface
type ServiceDiscoveryAgent interface {
	Stop()
	GetClient(ctx context.Context, isHTTPS, secure bool) (ServiceClient, error)
	DiscoverNow()
}

// ServiceClient is interface
type ServiceClient interface {
	CreateRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error)
	DoRequest(req *http.Request) (*http.Response, error)
}

type serviceDiscoveryRequest struct {
	errCh    chan<- error
	resultCh chan<- clientFactory
}

type clientFactory func(isHttps, secure bool) ServiceClient

type httpWrapperServiceClient struct {
	httpClient *http.Client
	url        string
}
