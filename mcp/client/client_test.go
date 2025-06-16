package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

// MCP
func TestMcpClientConnection(t *testing.T) {
	c, err := NewClient("http://localhost:8888/sse", "test-client", "1.0.0", "gin-vue-admin MCP")
	defer c.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestTools(t *testing.T) {
	t.Run("currentTime", func(t *testing.T) {
		c, err := NewClient("http://localhost:8888/sse", "test-client", "1.0.0", "gin-vue-admin MCP")
		defer c.Close()
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}
		ctx := context.Background()

		request := mcp.CallToolRequest{}
		request.Params.Name = "currentTime"
		request.Params.Arguments = map[string]interface{}{
			"timezone": "UTC+8",
		}

		result, err := c.CallTool(ctx, request)
		if err != nil {
			t.Fatalf(": %v", err)
		}

		if len(result.Content) != 1 {
			t.Errorf("1， %d", len(result.Content))
		}
		if content, ok := result.Content[0].(mcp.TextContent); ok {
			t.Logf("%s", content.Text)
		} else {
			t.Logf("%+v", content)
		}
	})

	t.Run("getNickname", func(t *testing.T) {

		c, err := NewClient("http://localhost:8888/sse", "test-client", "1.0.0", "gin-vue-admin MCP")
		defer c.Close()
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}
		ctx := context.Background()

		// Initialize
		initRequest := mcp.InitializeRequest{}
		initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
		initRequest.Params.ClientInfo = mcp.Implementation{
			Name:    "test-client",
			Version: "1.0.0",
		}

		_, err = c.Initialize(ctx, initRequest)
		if err != nil {
			t.Fatalf(": %v", err)
		}

		request := mcp.CallToolRequest{}
		request.Params.Name = "getNickname"
		request.Params.Arguments = map[string]interface{}{
			"username": "admin",
		}

		result, err := c.CallTool(ctx, request)
		if err != nil {
			t.Fatalf(": %v", err)
		}

		if len(result.Content) != 1 {
			t.Errorf("1， %d", len(result.Content))
		}
		if content, ok := result.Content[0].(mcp.TextContent); ok {
			t.Logf("%s", content.Text)
		} else {
			t.Logf("%+v", content)
		}
	})
}

func TestGetTools(t *testing.T) {
	c, err := NewClient("http://localhost:8888/sse", "test-client", "1.0.0", "gin-vue-admin MCP")
	defer c.Close()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	ctx := context.Background()

	toolsRequest := mcp.ListToolsRequest{}

	toolListResult, err := c.ListTools(ctx, toolsRequest)
	if err != nil {
		t.Fatalf(": %v", err)
	}
	for i := range toolListResult.Tools {
		tool := toolListResult.Tools[i]
		fmt.Printf(": %s\n", tool.Name)
		fmt.Printf(": %s\n", tool.Description)

		//
		if tool.InputSchema.Properties != nil {
			fmt.Println(":")
			for paramName, prop := range tool.InputSchema.Properties {
				required := ""
				//
				for _, reqField := range tool.InputSchema.Required {
					if reqField == paramName {
						required = ""
						break
					}
				}
				fmt.Printf("  - %s (: %s, : %s, : %s)\n",
					paramName, prop.(map[string]any)["type"], prop.(map[string]any)["description"], required)
			}
		} else {
			fmt.Println("")
		}
		fmt.Println("-------------------")
	}
}
