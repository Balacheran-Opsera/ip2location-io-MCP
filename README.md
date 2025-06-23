# IP2Location.io MCP server

This is a simple Model Context Protocol (MCP) server implementation for IP2Location.io API. It will return geolocation information for an IP address.

# Requirement

This MCP server supports to query without an API key, with a limitation of 1,000 queries per day. You can also [sign up](https://www.ip2location.io/sign-up) for a free API key and enjoy up to 50,000 queries per month.

The setup also use uv, which can be install by following [the guide](https://modelcontextprotocol.io/quickstart/server#set-up-your-environment).

# Setup

To use this MCP server with Claude Desktop, just add the following to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "ip2locationio": {
      "command": "uv",
      "args": [
        "--directory",
        "/path/to/ip2locationio",
        "run",
        "ip2locationio.py"
      ],
      "env": {
        "IP2LOCATION_API_KEY": "<YOUR API key HERE>"
      }
    }
  }
}
```

To get your API key, just [login](https://www.ip2location.io/log-in) to your dashboard and get it from there.

Restart the Claude Desktop after save the changes, and you shall see it pops out in the `Search and tools` menu.

# License

See the LICENSE file.