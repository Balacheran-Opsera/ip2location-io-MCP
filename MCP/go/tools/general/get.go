package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ip2location-io-ip-geolocation-api/mcp-server/config"
	"github.com/ip2location-io-ip-geolocation-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("key=%v", val))
		}
		if val, ok := args["ip"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ip=%v", val))
		}
		if val, ok := args["format"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("format=%v", val))
		}
		if val, ok := args["lang"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_",
		mcp.WithDescription("Geolocate user's location information via IP address"),
		mcp.WithString("key", mcp.Required(), mcp.Description("API key.")),
		mcp.WithString("ip", mcp.Required(), mcp.Description("IP address (IPv4 or IPv6) for reverse IP location lookup purposes. If not present, the server IP address will be used for the location lookup.")),
		mcp.WithString("format", mcp.Description("Format of the response message.")),
		mcp.WithString("lang", mcp.Description("Translation information. The translation only applicable for continent, country, region and city name. This parameter is only available for Plus and Security plan only.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetHandler(cfg),
	}
}
