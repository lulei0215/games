package announcement

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background()
	// ，config.Config，config.yamlkey
	// initialize.Viper()
	// api.Api
	initialize.Api(ctx)
	// api.Menu
	initialize.Menu(ctx)
	initialize.Gorm(ctx)
	initialize.Router(group)
}
