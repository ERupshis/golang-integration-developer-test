package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/erupshis/golang-integration-developer-test/internal/common/consts"
	"github.com/erupshis/golang-integration-developer-test/internal/common/retrier"
	"github.com/erupshis/golang-integration-developer-test/internal/common/utils/deferutils"
	"github.com/erupshis/golang-integration-developer-test/internal/service/models"
)

const (
	getBalanceURI      = "http://%s/api/v1/player"
	withdrawBalanceURI = "http://%s/api/v1/withdraw"
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

	fullURI, err := enrichURI(
		consts.GamesHost,
		[]URIParam{{consts.Platform, platform}},
	)
	if err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}

	request := func(context context.Context) ([]byte, error) {
		return d.makeRequest(context, http.MethodGet, fullURI, nil)
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

func (d *Default) GetBalance(ctx context.Context, playerID string) (int64, error) {
	errMsg := "get player balance: %w"

	fullURI, err := enrichURI(
		fmt.Sprintf(getBalanceURI, d.host),
		[]URIParam{{consts.ID, playerID}},
	)
	if err != nil {
		return -1, fmt.Errorf(errMsg, err)
	}

	request := func(context context.Context) ([]byte, error) {
		return d.makeRequest(context, http.MethodGet, fullURI, nil)
	}

	rawPlayerData, err := retrier.RetryCallWithTimeout[[]byte](ctx, nil, repeatableErrors, request)
	if err != nil {
		return -1, fmt.Errorf(errMsg, err)
	}

	playerData := models.UserData{}
	if err = json.Unmarshal(rawPlayerData, &playerData); err != nil {
		return -1, fmt.Errorf(errMsg, err)
	}

	return playerData.Balance, nil
}

func (d *Default) WithdrawBalance(ctx context.Context, playerID string, amount int64) (int64, error) {
	errMsg := "withdraw player balance: %w"

	ID, err := strconv.ParseInt(playerID, 10, 64)
	if err != nil {
		return -1, fmt.Errorf(errMsg, err)
	}

	requestPlayerData := models.UserWithdraw{
		ID:     ID,
		Amount: amount,
	}

	reqBody, err := json.Marshal(requestPlayerData)
	if err != nil {
		return -1, fmt.Errorf(errMsg, err)
	}

	request := func(context context.Context) ([]byte, error) {
		return d.makeRequest(context, http.MethodPatch, fmt.Sprintf(withdrawBalanceURI, d.host), reqBody)
	}

	rawPlayerData, err := retrier.RetryCallWithTimeout[[]byte](ctx, nil, repeatableErrors, request)
	if err != nil {
		return -1, fmt.Errorf(errMsg, err)
	}

	if rawPlayerData != nil && rawPlayerData[0] != '{' {
		switch string(rawPlayerData) {
		case ErrUserNotFound.Error():
			return -1, ErrUserNotFound
		case ErrInsufficientFunds.Error():
			return -1, ErrInsufficientFunds
		default:
			return -1, fmt.Errorf(string(rawPlayerData))
		}
	}

	fmt.Println(string(rawPlayerData))
	playerData := models.UserData{}
	if err = json.Unmarshal(rawPlayerData, &playerData); err != nil {
		return -1, fmt.Errorf(errMsg, err)
	}

	return playerData.Balance, nil
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
