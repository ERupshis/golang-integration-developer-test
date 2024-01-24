package consts

const (
	GamesHost = "https://www.freetogame.com/api/games"
)

const (
	Platform       = "platform"
	PlatformPC     = "pc"
	PlatformMobile = "mobile"
)

func IsPlatformValid(platform string) bool {
	switch platform {
	case PlatformPC:
		fallthrough
	case PlatformMobile:
		return true
	default:
		return false
	}
}
