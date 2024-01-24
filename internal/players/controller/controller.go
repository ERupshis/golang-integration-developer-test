package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/erupshis/golang-integration-developer-test/internal/common/consts"
	"github.com/erupshis/golang-integration-developer-test/internal/common/logger"
	"github.com/erupshis/golang-integration-developer-test/internal/common/utils/deferutils"
	"github.com/erupshis/golang-integration-developer-test/internal/players/storage"
	"github.com/erupshis/golang-integration-developer-test/internal/players/storage/inmem"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	logs        logger.BaseLogger
	userStorage *inmem.UserStorage
}

func NewController(userStorage *inmem.UserStorage, logger logger.BaseLogger) *Controller {
	return &Controller{
		userStorage: userStorage,
		logs:        logger,
	}
}

// Route returns a new chi.Mux router configured with middleware and handlers for HTTPController.
func (c *Controller) Route() *chi.Mux {
	r := chi.NewRouter()

	r.Use(c.logs.LogHandler)

	r.Get("/player", c.playerHandler)
	r.Patch("/withdraw", c.withdrawHandler)
	r.NotFound(c.badRequestHandler)

	return r
}

func (c *Controller) badRequestHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

// playerHandler func.
// @Description provides player data by playerID
// @Tags player
// @Summary select player by ID
// @ID player-select
// @Produce json
// @Param id query string true  "player search by id"
// @Success 200 {object} models.UserData
// @Success 204
// @Failure 400 {string} string incorrect query in request
// @Failure 500 {string} string something wrong with storage
// @Failure 500 {string} string unexpected marshalling error
// @Router /player [get]
func (c *Controller) playerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer deferutils.ExecSilent(r.Body.Close)
	}

	queryParams := r.URL.Query()
	rawID := queryParams.Get(consts.ID)

	playerID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		http.Error(w, "incorrect query in request", http.StatusBadRequest)
		return
	}

	userData, err := c.userStorage.GetUserByID(playerID)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, "something wrong with storage", http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(*userData)
	if err != nil {
		http.Error(w, "unexpected marshalling error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respBody)
	if err != nil {
		http.Error(w, "unexpected write response error", http.StatusInternalServerError)
		return
	}
}

// withdrawHandler func.
// @Description withdraw currency from player balance
// @Tags player
// @Summary withdraw currency from player account by ID
// @ID player-withdraw
// @Produce plain
// @Param id query string true  "player search by id"
// @Param amount query string true  "currency amount"
// @Success 200 {object} string "OK"
// @Failure 404 {string} string "user not found"
// @Failure 400 {string} string "incorrect query in request"
// @Failure 405 {string} string "insufficient funds"
// @Failure 500 {string} string "something wrong with storage"
// @Failure 500 {string} string "unexpected marshalling error"
// @Router /withdraw [patch]
func (c *Controller) withdrawHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer deferutils.ExecSilent(r.Body.Close)
	}

	queryParams := r.URL.Query()
	rawID := queryParams.Get(consts.ID)
	rawAmount := queryParams.Get(consts.Amount)

	playerID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		http.Error(w, "incorrect query in request", http.StatusBadRequest)
		return
	}

	currencyAmount, err := strconv.ParseInt(rawAmount, 10, 64)
	if err != nil {
		http.Error(w, "incorrect query in request", http.StatusBadRequest)
		return
	}

	if err = c.userStorage.WithdrawBalance(playerID, currencyAmount); err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		} else if errors.Is(err, storage.ErrUserInSufficientFunds) {
			http.Error(w, "insufficient funds", http.StatusMethodNotAllowed)
			return
		}

		http.Error(w, "something wrong with storage", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		http.Error(w, "unexpected write response error", http.StatusInternalServerError)
		return
	}
}
