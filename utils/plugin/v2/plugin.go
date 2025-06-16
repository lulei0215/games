package plugin

import (
	"github.com/gin-gonic/gin"
)

// Plugin v2
type Plugin interface {
	// Register
	Register(group *gin.Engine)
}
