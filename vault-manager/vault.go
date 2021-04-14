package main

import (
	"fmt"

	"github.com/hashicorp/vault/api"
	"github.com/labstack/echo/v4"
)

// https://pkg.go.dev/github.com/hashicorp/vault/api

type VaultManager struct {
	Client    *api.Client
	SysClient *api.Sys
}

func NewVaultManager() *VaultManager {
	config := api.Config{
		Address: "http://127.0.0.1:8200",
	}
	client, err := api.NewClient(&config)
	if err != nil {
		fmt.Printf("Failed to create Echo client\n")
		return nil
	}
	return &VaultManager{
		Client:    client,
		SysClient: client.Sys(),
	}
}

// Return information about the Vault configuration
func (v *VaultManager) getInfo(c echo.Context) {

}

// Return current status of the tool
func (v *VaultManager) status(c echo.Context) {
	// number of shards added already

}
