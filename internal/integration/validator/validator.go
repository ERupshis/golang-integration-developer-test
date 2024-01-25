package validator

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/erupshis/golang-integration-developer-test/internal/common/consts"
)

var (
	ErrInvalidToken          = fmt.Errorf("invalid token")
	ErrInvalidPlatform       = fmt.Errorf("invalid platform")
	ErrInvalidCurrencyCode   = fmt.Errorf("invalid currency code")
	ErrInvalidCurrencyName   = fmt.Errorf("invalid currency name")
	ErrInvalidPlayerNickName = fmt.Errorf("invalid player  nickname")
)

var (
	regexToken          = regexp.MustCompile(`^[a-zA-Z0-9]{36}$`)
	regexCurrencyName   = regexp.MustCompile(`^[a-zA-Z\s]+$`)
	regexPlayerNickName = regexp.MustCompile(`^[a-zA-Z_]+$`)
)

func CheckID(rawID string) (int64, []error) {
	var errs []error
	ID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		errs = append(errs, fmt.Errorf("parse game ID: %w", err))
	} else if ID < 1 {
		errs = append(errs, fmt.Errorf("invalid game ID"))
	}

	return ID, errs
}

func CheckToken(token string) (bool, []error) {
	var errs []error
	if !regexToken.MatchString(token) {
		errs = append(errs, ErrInvalidToken)
	}

	return len(errs) == 0, errs
}

func CheckPlatform(platform string) (bool, []error) {
	switch platform {
	case consts.PlatformPC:
		fallthrough
	case consts.PlatformBrowser:
		return true, nil
	default:
		return false, []error{ErrInvalidPlatform}
	}
}

func CheckCurrency(code, name string) (bool, []error) {
	var errs []error

	if !regexCurrencyName.MatchString(name) {
		errs = append(errs, ErrInvalidCurrencyName)
	}

	switch code {
	case consts.CurrencyUSD:
	// NOTHING TO DO.
	default:
		errs = append(errs, ErrInvalidCurrencyCode)
	}

	return len(errs) == 0, errs
}

func CheckPlayer(rawID, nickname string) (bool, []error) {
	var errs []error

	ID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		errs = append(errs, fmt.Errorf("parse game ID: %w", err))
	} else if ID < 1 {
		errs = append(errs, fmt.Errorf("invalid game ID: %w", err))
	}

	if !regexPlayerNickName.MatchString(nickname) {
		errs = append(errs, ErrInvalidPlayerNickName)
	}

	return len(errs) == 0, errs
}
