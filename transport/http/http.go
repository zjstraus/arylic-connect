package http

import (
	"arylic-connect/transport"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Transport struct {
	client *http.Client
	target *url.URL
}

func New() (*Transport, error) {
	return &Transport{
		client: &http.Client{},
	}, nil
}

func (t *Transport) Connect(target string) error {
	closeErr := t.Close()
	if closeErr != nil {
		return closeErr
	}

	targetUrl, urlParseErr := url.Parse(target)
	if urlParseErr != nil {
		return urlParseErr
	}

	t.target = targetUrl
	return nil
}

func (t *Transport) MakeRequest(ctx context.Context, command string, params ...string) ([]byte, error) {
	if t.client == nil {
		return nil, errors.New("no http target")
	}

	allParams := []string{command}
	allParams = append(allParams, params...)
	joinedParams := strings.Join(allParams, ":")

	targetUrl := *t.target
	targetQuery := targetUrl.Query()
	targetQuery.Add("command", joinedParams)
	targetUrl.RawQuery = targetQuery.Encode()

	req, reqErr := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl.String(), nil)
	if reqErr != nil {
		return nil, reqErr
	}

	resp, reqErr := t.client.Do(req)
	if reqErr != nil {
		return nil, reqErr
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (t *Transport) Close() error {
	if t.client != nil {
		t.client.CloseIdleConnections()
	}

	return nil
}

func (t *Transport) Flavor() transport.InterfaceFlavor {
	return transport.Flavor_HTTP
}

func (t *Transport) Target() string {
	if t.client == nil {
		return ""
	}
	return t.target.String()
}
