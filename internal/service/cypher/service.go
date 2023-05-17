package cypher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"google.golang.org/protobuf/proto"
)

type Service struct {
	servicePrivateKey *rsa.PrivateKey
}

func NewCypherService() (*Service, error) {
	key, err := rsa.GenerateKey(rand.Reader, 1<<10-1)
	if err != nil {
		return nil, fmt.Errorf("rsa.GenerateKey: %w", err)
	}
	return &Service{
		servicePrivateKey: key,
	}, nil
}

func (c *Service) DecryptProto(cypherText []byte, dest proto.Message) error {
	hash := sha256.New()
	rnd := rand.Reader
	plain, err := rsa.DecryptOAEP(hash, rnd, c.servicePrivateKey, cypherText, []byte("vote"))
	if err != nil {
		return fmt.Errorf("rsa.Decrypt: %w", err)
	}
	err = proto.Unmarshal(plain, dest)
	if err != nil {
		return fmt.Errorf("proto.Unmarshal: %w", err)
	}
	return nil
}

func (c *Service) GetPublicKey() *rsa.PublicKey {
	return &c.servicePrivateKey.PublicKey
}
