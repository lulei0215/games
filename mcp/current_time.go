package mcpTool

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&CurrentTime{})
}

type CurrentTime struct {
}

func (t *CurrentTime) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	//

	timezone, ok := request.Params.Arguments["timezone"].(string)

	if !ok {
		return nil, errors.New("：timezone ")
	}
	// timezone
	loc, err := loadTimeZone(timezone)
	if err != nil {
		return nil, err
	}

	//
	currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05")
	//
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("%s ：%s", timezone, currentTime),
			},
		},
	}, nil
}

func (t *CurrentTime) New() mcp.Tool {
	return mcp.NewTool("currentTime",
		mcp.WithDescription(""),
		mcp.WithString("timezone",
			mcp.Required(),
			mcp.Description(""),
			mcp.Enum("UTC", "CST", "PST", "EST", "GMT", "CET", "JST", "MST", "IST", "AST", "HST"),
		))
}

// IANA
func loadTimeZone(timezone string) (*time.Location, error) {
	//
	timezoneMap := map[string]string{
		"UTC": "UTC",
		"CST": "Asia/Shanghai", //
		"PST": "America/Los_Angeles",
		"EST": "America/New_York",
		"GMT": "GMT",
		"CET": "Europe/Paris",
		"JST": "Asia/Tokyo",
		"MST": "America/Denver",
		"IST": "Asia/Kolkata",
		"AST": "Asia/Riyadh", //
		"HST": "Pacific/Honolulu",
	}

	//
	tzName, exists := timezoneMap[timezone]
	if !exists {
		return nil, errors.New(": " + timezone)
	}

	//
	return time.LoadLocation(tzName)
}
