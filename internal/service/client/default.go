package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/erupshis/golang-integration-developer-test/internal/common/consts"
	"github.com/erupshis/golang-integration-developer-test/internal/common/retrier"
	"github.com/erupshis/golang-integration-developer-test/internal/common/utils/deferutils"
	"github.com/erupshis/golang-integration-developer-test/internal/models"
)

var repeatableErrors []error

var (
	_ BaseClient = (*Default)(nil)
)

type Default struct {
	client *http.Client
	host   string
}

func NewDefault(host string) BaseClient {
	return &Default{
		client: http.DefaultClient,
		host:   host,
	}
}

func (d *Default) GetGames(ctx context.Context, platform string) (models.Games, error) {
	errMsg := "get games list: %w"

	if !consts.IsPlatformValid(platform) {
		return nil, fmt.Errorf(errMsg, ErrInvalidPlatform)
	}

	fullURI, err := enrichURIWithPlatform(consts.GamesHost, platform)
	if err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}

	request := func(context context.Context) ([]byte, error) {
		return d.makeRequest(context, http.MethodPost, fullURI, nil)
	}

	rawGames, err := retrier.RetryCallWithTimeout[[]byte](ctx, nil, repeatableErrors, request)
	if err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}

	games := make(models.Games, 0)
	if err = json.Unmarshal(rawGames, &games); err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}

	return games, nil
}

func (d *Default) makeRequest(ctx context.Context, method string, URI string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, URI, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("new request generation: %w", err)
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}

	defer deferutils.ExecSilent(resp.Body.Close)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	return respBody, nil
}

func (d *Default) GetBalance(ctx context.Context, playerID string) (int64, error) {
	return 0, nil
}

func enrichURIWithPlatform(URI, platform string) (string, error) {
	u, err := url.Parse(URI)
	if err != nil {
		return "", fmt.Errorf("enrich URI with game platform: %w", err)
	}

	params := url.Values{}
	params.Add(consts.Platform, platform)

	u.RawQuery = params.Encode()
	return u.String(), nil
}
