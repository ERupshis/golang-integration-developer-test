# Golang test for developer (integrations)

## Description

Implement simple integration grpc service.

Create HTTP client based on ``Client`` interface in `service/client.go`. Method ``GetGames`` should 
go to [freetogame](https://www.freetogame.com/api/games?platform=pc) source 
(in query key `platform` you should pass only `pc` or `mobile`, implement validation for this) and receive list of games
in json format, parse all available fields to ``Games`` struct.

Generate GRPC server based on `proto/integration.proto` and implement methods ``GetBalance`` and 
``SendBet``.

- ``GetBalance`` should validate following fields: `token`,`player`,`platform`,`currency`,`game_id`
then make request to service method `GetGames` and find game according to `game_id` from grpc request
if game is found then fill in field in `GetBalanceResponse.game`(field namings are the same) and
make `GetBalance` request to service by `Player.id` (Here use any in-memory storage with prefabricated data by player_id)


- ``SendBet`` should validate all fields and check in-memory storage 
if player has sufficient balance otherwise thow error, subtract amount from balance and update in-memory
storage and return updated balance.

Fork this repository so we can see it, update `README.md` with full documentation how to deploy and 
run this service. Optionally it will be nice to have unit tests and any usage of middlewares throughout
this project.

## Credits
Thank you for [FreeToGame.com](https://www.freetogame.com) portal for providing free API.

## How to start:
- make docker_up - to run containers.
- grpc server will be available on :18081 port.
