package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

type VaultClient struct {
	client *api.Client
}



// DatabaseCredentials from Vault
type DatabaseCredentials struct {
    Username string
    Password string
    TTL      int
}


// NewVaultclient creates Vault client using .env configuration


func NewVaultClient() (*VaultClient, error) {
	config := api.DefaultConfig()

	config.Address = os.Getenv("VAULT_ADDR")

	//TODO - add approle auth and using tls to connect to Vault
	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create vault client : %w", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	return &VaultClient{client: client}, nil


}


//GetDatabaseCredentials gets fresh creds for your RSS agg role
func (v *VaultClient) GetDatabaseCredentials() (*DatabaseCredentials, error) {
	role := os.Getenv("VAULT_POSTGRES_ROLE")

	secret, err := v.client.Logical().Read(fmt.Sprintf("postgres/creds/%s", role))

	 if err != nil {
        return nil, fmt.Errorf("failed to read credentials: %w", err)
    }

	if secret == nil || secret.Data == nil {
		return nil, fmt.Errorf("no credentials found for role %s", role)
	}

	username, ok := secret.Data["username"].(string)

	if !ok {
		return nil, fmt.Errorf("username not found in response")
	}

	password, ok := secret.Data["password"].(string)

	if !ok {
		return nil, fmt.Errorf("password not found in response")
	}

	return &DatabaseCredentials{
		Username: username,
		Password: password,
		TTL: secret.LeaseDuration,
	}, nil


} 
