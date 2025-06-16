package plugin

import (
	"github.com/gin-gonic/gin"
)

const (
	OnlyFuncName = "Plugin"
)

// Plugin
type Plugin interface {
	// Register
	Register(group *gin.RouterGroup)

	// RouterPath
	RouterPath() string
}
