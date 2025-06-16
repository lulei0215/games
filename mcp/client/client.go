package client

import (
	"context"
	"errors"

	mcpClient "github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func NewClient(baseUrl, name, version, serverName string) (*mcpClient.Client, error) {
	client, err := mcpClient.NewSSEMCPClient(baseUrl)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	// client
	if err := client.Start(ctx); err != nil {
		return nil, err
	}

	//
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    name,
		Version: version,
	}

	result, err := client.Initialize(ctx, initRequest)
	if err != nil {
		return nil, err
	}
	if result.ServerInfo.Name != serverName {
		return nil, errors.New("server name mismatch")
	}
	return client, nil
}
