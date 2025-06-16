package {{ .Package }}

import (
	"context"
	"{{.Module}}/plugin/{{ .Package }}/initialize"
	interfaces "{{.Module}}/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

// ，config.Config，config.yamlkey
// initialize.Viper()
// api.Api
// initialize.Api(ctx)
// api.Menu
// initialize.Menu(ctx)
func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background() 
	initialize.Gorm(ctx)
	initialize.Router(group)
}
