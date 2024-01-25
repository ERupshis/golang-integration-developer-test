package client

import (
	"fmt"
	"net/url"
)

// URI configurator.
type URIParam struct {
	key string
	val string
}

func enrichURI(URI string, URIParams []URIParam) (string, error) {
	u, err := url.Parse(URI)
	if err != nil {
		return "", fmt.Errorf("enrich URI: %w", err)
	}

	params := url.Values{}
	for _, kv := range URIParams {
		params.Add(kv.key, kv.val)
	}

	u.RawQuery = params.Encode()
	return u.String(), nil
}
