// Package hasher provides hasher for message hash-sum calculation and verification.
package hasher

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"

	"github.com/erupshis/golang-integration-developer-test/internal/common/logger"
)

type HashType int

const (
	TypeSHA256 = HashType(0) // type of used hash algorithm
)

const (
	headerSHA256 = "HashSHA256"
)

const (
	algoSHA256 = iota
)

type readCloserWrapper struct {
	io.Reader
	io.Closer
}

// Hasher stores hash related config models.
type Hasher struct {
	log      logger.BaseLogger
	hashType HashType // type of algorithm
	key      string   // hash key
}

// CreateHasher create method.
func CreateHasher(hashKey string, hashType HashType, log logger.BaseLogger) *Hasher {
	return &Hasher{key: hashKey, hashType: hashType, log: log}
}

// HashMsg returns hash for message.
func (hr *Hasher) HashMsg(msg []byte) (string, error) {
	algo, err := hr.getAlgo()
	if err != nil {
		return "", fmt.Errorf("hash message: %w", err)
	}

	switch algo {
	case algoSHA256:
		return hashMsg(sha256.New, msg, hr.key)
	default:
		panic("unknown algorithm")
	}
}

// hashMsg returns hash for message.
func hashMsg(hashFunc func() hash.Hash, msg []byte, key string) (string, error) {
	var h hash.Hash
	if key != "" {
		h = hmac.New(hashFunc, []byte(key))
	} else {
		// hasher sum w/o authentication.
		h = hashFunc()
	}

	_, err := h.Write(msg)
	if err != nil {
		return "", err
	}

	hashVal := h.Sum(nil)
	return fmt.Sprintf("%x", hashVal), nil
}

// getAlgo returns used algo.
func (hr *Hasher) getAlgo() (int, error) {
	switch hr.hashType {
	case TypeSHA256:
		return algoSHA256, nil
	default:
		return -1, fmt.Errorf("unknow algorithm")
	}
}

// GetKey returns hash key.
func (hr *Hasher) GetKey() string {
	return hr.key
}
