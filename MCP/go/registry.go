package main

import (
	"github.com/ip2location-io-ip-geolocation-api/mcp-server/config"
	"github.com/ip2location-io-ip-geolocation-api/mcp-server/models"
	tools_general "github.com/ip2location-io-ip-geolocation-api/mcp-server/tools/general"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_general.CreateGetTool(cfg),
	}
}
