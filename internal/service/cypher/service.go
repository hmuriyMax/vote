package cypher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Service struct {
	servicePrivateKey *rsa.PrivateKey
}

func NewCypherService() (*Service, error) {
	key, err := rsa.GenerateKey(rand.Reader, 32)
	if err != nil {
		return nil, fmt.Errorf("rsa.GenerateKey: %w", err)
	}
	return &Service{
		servicePrivateKey: key,
	}, nil
}

func (c *Service) Decrypt(cypherText []byte) (int64, error) {
	plain, err := rsa.DecryptOAEP(sha256.New(), nil, c.servicePrivateKey, cypherText, []byte(""))
	if err != nil {
		return 0, fmt.Errorf("rsa.Decrypt: %w", err)
	}
	res, err := strconv.ParseInt(string(plain), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing int from text: %w", err)
	}
	return res, nil
}
