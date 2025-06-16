package mcpTool

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// McpTool MCP
type McpTool interface {
	// Handle
	Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
	// New
	New() mcp.Tool
}

var toolRegister = make(map[string]McpTool)

// RegisterTool init，
func RegisterTool(tool McpTool) {
	mcpTool := tool.New()
	toolRegister[mcpTool.Name] = tool
}

// RegisterAllTools MCP
func RegisterAllTools(mcpServer *server.MCPServer) {
	for _, tool := range toolRegister {
		mcpServer.AddTool(tool.New(), tool.Handle)
	}
}
