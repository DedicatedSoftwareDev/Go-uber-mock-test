package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Nike-Inc/cerberus-go-client/v3/auth"
	"github.com/Nike-Inc/cerberus-go-client/v3/cerberus"
	"nike.com/aiml/jwt-validator/interfaces"
)

type CacheValue struct {
	value    string
	creation time.Time
}

var tokenMap = make(map[string]CacheValue)
var tokenTTL = 60 * 60 * 20.0 // 20 hours in seconds

var cerberusClientFactory = func(authMethod auth.Auth) (interfaces.ClientInterface, error) {
	client, err := cerberus.NewClient(authMethod, nil)
	if err != nil {
		return nil, err
	}
	return &interfaces.CerberusClientWrapper{Client: client}, nil
}

func GetPlatformToken(endpointEnv string) (string, error) {
	key := BuildKey(endpointEnv)

	// Check cache first
	cacheValue, ok := tokenMap[key]
	if ok {
		// Cache hit.  Check if expired
		duration := time.Since(cacheValue.creation)
		if duration.Seconds() < tokenTTL {
			// Not expired.  Return it
			return cacheValue.value, nil
		}
	}
	// Not found or expired.  Get new token, update cache, and return it
	log.Printf("Getting and caching new platform token for key: %s", key)
	token, err := GetTokenFromCerberus(key)
	if err != nil {
		return "", err
	}
	tokenMap[key] = CacheValue{value: token, creation: time.Now()}
	return token, nil
}

func BuildKey(endpointEnv string) string {
	workspace := GetSoleHostName()
	databricksRole := GetDatabricksRole(endpointEnv)
	return fmt.Sprintf("ServicePrincipal_%s_App.NikeSole.aimltf_whq.%s", workspace, databricksRole)
}

func GetTokenFromCerberus(key string) (string, error) {
	sdbPath := "app/aiml-sole-serviceprincipals/tokensaimltf_whq"

	authMethod, err := auth.NewSTSAuth("https://prod.cerberus.nikecloud.com", "us-west-2")
	if err != nil {
		return "", fmt.Errorf("error getting Cerberus STS Auth: %w", err)
	}

	client, err := cerberusClientFactory(authMethod)
	if err != nil {
		return "", fmt.Errorf("error getting Cerberus client: %w", err)
	}
	sec, err := client.Secret().Read(sdbPath)
	if err != nil {
		return "", fmt.Errorf("error reading secrets with SDB path: %s: %w", sdbPath, err)
	}
	if sec.Data != nil {
		value, ok := sec.Data[key]
		if ok {
			return value.(string), nil
		} else {
			return "", fmt.Errorf("no value found in path %s with key %s", sdbPath, key)
		}
	} else {
		return "", fmt.Errorf("no value found at path %s", sdbPath)
	}

}

type DataFetcher interface {
	FetchData() string
}

type RealFetcher struct{}

func (r *RealFetcher) FetchData() string {
	return "Hello, World!"
}

func main() {
	fetcher := &RealFetcher{}
	fmt.Println(ProcessData(fetcher))
}

func ProcessData(df DataFetcher) string {
	return "Processed: " + df.FetchData()
}
