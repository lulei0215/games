package mcpTool

import (
	"context"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/mark3labs/mcp-go/mcp"
	"gorm.io/gorm"
)

func init() {
	RegisterTool(&GetNickname{})
}

type GetNickname struct{}

// usernamenickname
func (t *GetNickname) New() mcp.Tool {
	return mcp.NewTool("getNickname",
		mcp.WithDescription("usernamenickname"),
		mcp.WithString("username",
			mcp.Required(),
			mcp.Description("username"),
		))
}

// Handle
func (t *GetNickname) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 1.
	username, ok := request.Params.Arguments["username"].(string)
	if !ok {
		return nil, errors.New("：username ")
	}

	if username == "" {
		return nil, errors.New("：username ")
	}

	// 2.
	global.GVA_LOG.Info("getNickname ")

	// 3. ，
	var user struct {
		NickName string
	}

	err := global.GVA_DB.Model(&system.SysUser{}).
		Select("nick_name").
		Where("username = ?", username).
		First(&user).Error

	// 4.
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf(" %s ", username),
					},
				},
			}, nil
		}
		global.GVA_LOG.Error("")
		return nil, errors.New("，")
	}

	//
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf(" %s  %s", username, user.NickName),
			},
		},
	}, nil
}
