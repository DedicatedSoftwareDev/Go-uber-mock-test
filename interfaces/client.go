package interfaces

import "github.com/Nike-Inc/cerberus-go-client/v3/cerberus"

type ClientInterface interface {
	Secret() SecretInterface
}

type SecretInterface interface {
	Read(path string) (*cerberus.SecureData, error)
}

type CerberusClientWrapper struct {
	Client cerberus.Client
}

func (c *CerberusClientWrapper) Secret() SecretInterface {
	return &SecretWrapper{Secret: c.Client.Secret()}
}

type SecretWrapper struct {
	Secret cerberus.Secret
}

func (s *SecretWrapper) Read(path string) (*cerberus.SecureData, error) {
	return s.Secret.Read(path)
}
