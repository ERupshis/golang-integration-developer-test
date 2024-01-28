package integr

import (
	"context"
	"errors"
	"strconv"

	"github.com/erupshis/golang-integration-developer-test/internal/integration/models"
	"github.com/erupshis/golang-integration-developer-test/internal/integration/validator"
	"github.com/erupshis/golang-integration-developer-test/internal/service/client"
	serviceModels "github.com/erupshis/golang-integration-developer-test/internal/service/models"
	pb_integr "github.com/erupshis/golang-integration-developer-test/pb/integration"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Controller struct {
	pb_integr.UnimplementedServiceServer

	client client.BaseClient
}

func NewController(client client.BaseClient) *Controller {
	return &Controller{
		client: client,
	}
}

func (c *Controller) GetBalance(ctx context.Context, in *pb_integr.GetBalanceRequest) (*pb_integr.GetBalanceResponse, error) {
	errs := validateGeneral(in.GetGeneral())
	if len(errs) != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", errors.Join(errs...))
	}

	balance, err := c.client.GetBalance(ctx, in.GetGeneral().GetPlayer().GetId())
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "%v", err)
	}

	games, err := c.client.GetGames(ctx, in.GetGeneral().GetPlatform())
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "%v", err)
	}

	game := findGameByID(&games, in.GetGeneral().GetGameId())
	if game == nil {
		return nil, status.Errorf(codes.NotFound, "game not found by id")
	}

	return &pb_integr.GetBalanceResponse{
		Balance: int32(balance),
		Game:    models.ConvertGameToGRPC(game),
	}, nil
}

func (c *Controller) SendBet(ctx context.Context, in *pb_integr.SendBetRequest) (*pb_integr.SendBetResponse, error) {
	errs := validateGeneral(in.GetGeneral())
	if len(errs) != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", errors.Join(errs...))
	}

	balance, err := c.client.WithdrawBalance(ctx, in.GetGeneral().GetPlayer().GetId(), int64(in.GetAmount()))
	if err != nil {
		if errors.Is(err, client.ErrUserNotFound) {
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}

		return nil, status.Errorf(codes.Unavailable, "%v", err)
	}

	return &pb_integr.SendBetResponse{Balance: int32(balance)}, nil
}

func validateGeneral(in *pb_integr.General) []error {
	var errs []error
	_, IDErrs := validator.CheckID(in.GetGameId())
	errs = append(errs, IDErrs...)

	_, tokenErrs := validator.CheckToken(in.GetToken())
	errs = append(errs, tokenErrs...)

	_, platformErrs := validator.CheckPlatform(in.GetPlatform())
	errs = append(errs, platformErrs...)

	_, currencyErrs := validator.CheckCurrency(in.GetCurrency().GetCode(), in.GetCurrency().GetName())
	errs = append(errs, currencyErrs...)

	_, playerErrs := validator.CheckPlayer(in.GetPlayer().GetId(), in.GetPlayer().GetNickname())
	errs = append(errs, playerErrs...)

	return errs
}

func findGameByID(games *serviceModels.Games, rawGameID string) *models.Game {
	gameID, _ := strconv.ParseInt(rawGameID, 10, 64)
	game := games.FindGameByID(gameID)
	if game == nil {
		return nil
	}

	return &models.Game{
		ID:               rawGameID,
		Title:            game.Title,
		ShortDescription: game.ShortDescription,
		GameURL:          game.GameURL,
	}
}
