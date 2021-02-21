package hasher

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/big"
)

// Hasher enables to calculate an "aggregated" hash of multiple io.Reader inputs
type Hasher struct {
	hash *big.Int
}

// New creates a new Hasher
func New() *Hasher {
	return &Hasher{
		hash: &big.Int{},
	}
}

// HashFile returns md5 hash byte array for an io.Reader
func HashFile(handler io.Reader) ([]byte, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, handler); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

// AddHash constructs a single hash representing multiple files
// by combining (xor) each individual hashes
func (h *Hasher) AddHash(handler io.Reader) error {
	bytes, err := HashFile(handler)
	if err != nil {
		return err
	}
	h.hash = h.hash.Xor(h.hash, (&big.Int{}).SetBytes(bytes))
	return nil
}

// MD5 returns calculated md5 hash
func (h *Hasher) MD5() string {
	return fmt.Sprintf("%x", h.hash)
}
