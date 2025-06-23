from typing import Any
import httpx
from mcp.server.fastmcp import FastMCP
import os

# Initialize FastMCP server
mcp = FastMCP("ip2locationio")

# Constants
IPLIO_API_BASE = "https://api.ip2location.io"
USER_AGENT = "ip2locationio-app/1.0"

def get_api_key() -> str | None:
    """Retrieve the API key from MCP server config."""
    return os.getenv("IP2LOCATION_API_KEY")

async def make_request(url: str, params: dict[str, str]) -> dict[str, Any] | None:
    """Make a request to the IP2Location.io API with proper error handling."""
    headers = {
        "User-Agent": USER_AGENT,
        "Accept": "application/json"
    }
    async with httpx.AsyncClient() as client:
        try:
            response = await client.get(url, headers=headers, params=params, timeout=30.0)
            response.raise_for_status()
            return response.json()
        except Exception:
            return None

@mcp.tool()
async def get_geolocation(ip: str) -> str:
    """Fetch geolocation for the given IP address."""
    params = {"ip": ip}
    api_key = get_api_key()
    if api_key:
        params["key"] = api_key  # IP2Location.io API key parameter

    geolocation_result = await make_request(IPLIO_API_BASE, params)

    if not geolocation_result:
        return f"Unable to fetch geolocation for IP {ip}."

    return geolocation_result

if __name__ == "__main__":
    mcp.run(transport='stdio')
