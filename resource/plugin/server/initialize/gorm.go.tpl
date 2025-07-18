package initialize

import (
	"context"
	"fmt"
	"{{.Module}}/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate()
	if err != nil {
		err = errors.Wrap(err, "!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
