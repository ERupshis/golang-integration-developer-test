package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/erupshis/golang-integration-developer-test/internal/common/consts"
	"github.com/erupshis/golang-integration-developer-test/internal/common/logger"
	"github.com/erupshis/golang-integration-developer-test/internal/common/utils/deferutils"
	"github.com/erupshis/golang-integration-developer-test/internal/players/models"
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

	return r
}

// playerHandler func.
// @Description provides player data by playerID
// @Tags player
// @Summary select player by ID
// @ID player-select
// @Produce json
// @Param id query string true  "player search by id"
// @Success 200 {object} models.UserDataP
// @Success 404 {string} string "user not found"
// @Failure 400 {string} string "incorrect query in request"
// @Failure 500 {string} string "something wrong with storage"
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

	playerData, err := c.userStorage.GetUserByID(playerID)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			http.Error(w, storage.ErrUserNotFound.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "something wrong with storage", http.StatusInternalServerError)
		return
	}

	writePlayerDataInResponse(w, playerData)
}

// withdrawHandler func.
// @Description withdraw currency from player balance
// @Tags player
// @Summary withdraw currency from player account by ID
// @ID player-withdraw
// @Accept json
// @Produce json
// @Param input body models.UserWithdrawP true "user withdraw amount"
// @Success 200 {object} models.UserDataP
// @Failure 400 {string} string "incorrect request"
// @Failure 404 {string} string "user not found"
// @Failure 405 {string} string "insufficient funds"
// @Failure 500 {string} string "something wrong with storage"
// @Failure 500 {string} string "unexpected marshalling error"
// @Router /withdraw [patch]
func (c *Controller) withdrawHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("unexpected during reading req body: %v", err), http.StatusInternalServerError)
		return
	}
	defer deferutils.ExecSilent(r.Body.Close)

	playerWithdrawData := &models.UserWithdrawP{}
	if err = json.Unmarshal(buf.Bytes(), playerWithdrawData); err != nil {
		http.Error(w, "incorrect request", http.StatusBadRequest)
		return
	}

	balance, err := c.userStorage.WithdrawBalance(playerWithdrawData.ID, playerWithdrawData.Amount)
	if err != nil {
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

	writePlayerDataInResponse(w, &models.UserDataP{ID: playerWithdrawData.ID, Balance: balance})
}

func writePlayerDataInResponse(w http.ResponseWriter, playerData *models.UserDataP) {
	respBody, err := json.Marshal(*playerData)
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
