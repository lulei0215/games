package system

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"github.com/pkg/errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	request "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"go.uber.org/zap"
)

var AutocodeHistory = new(autoCodeHistory)

type autoCodeHistory struct{}

// Create
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) Create(ctx context.Context, info request.SysAutoHistoryCreate) error {
	create := info.Create()
	err := global.GVA_DB.WithContext(ctx).Create(&create).Error
	if err != nil {
		return errors.Wrap(err, "!")
	}
	return nil
}

// First id
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) First(ctx context.Context, info common.GetById) (string, error) {
	var meta string
	err := global.GVA_DB.WithContext(ctx).Model(model.SysAutoCodeHistory{}).Where("id = ?", info.ID).Pluck("request", &meta).Error
	if err != nil {
		return "", errors.Wrap(err, "!")
	}
	return meta, nil
}

// Repeat
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) Repeat(businessDB, structName, abbreviation, Package string) bool {
	var count int64
	global.GVA_DB.Model(&model.SysAutoCodeHistory{}).Where("business_db = ? and (struct_name = ? OR abbreviation = ?) and package = ? and flag = ?", businessDB, structName, abbreviation, Package, 0).Count(&count).Debug()
	return count > 0
}

// RollBack
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) RollBack(ctx context.Context, info request.SysAutoHistoryRollBack) error {
	var history model.SysAutoCodeHistory
	err := global.GVA_DB.Where("id = ?", info.ID).First(&history).Error
	if err != nil {
		return err
	}
	if history.ExportTemplateID != 0 {
		err = global.GVA_DB.Delete(&model.SysExportTemplate{}, "id = ?", history.ExportTemplateID).Error
		if err != nil {
			return err
		}
	}
	if info.DeleteApi {
		ids := info.ApiIds(history)
		err = ApiServiceApp.DeleteApisByIds(ids)
		if err != nil {
			global.GVA_LOG.Error("ClearTag DeleteApiByIds:", zap.Error(err))
		}
	} // API
	if info.DeleteMenu {
		err = BaseMenuServiceApp.DeleteBaseMenu(int(history.MenuID))
		if err != nil {
			return errors.Wrap(err, "!")
		}
	} //
	if info.DeleteTable {
		err = s.DropTable(history.BusinessDB, history.Table)
		if err != nil {
			return errors.Wrap(err, "!")
		}
	} //
	templates := make(map[string]string, len(history.Templates))
	for key, template := range history.Templates {
		{
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			keys := strings.Split(key, "/")
			key = filepath.Join(keys...)
			key = strings.TrimPrefix(key, server)
		} // key
		{
			web := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot())
			server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
			slices := strings.Split(template, "/")
			template = filepath.Join(slices...)
			ext := path.Ext(template)
			switch ext {
			case ".js", ".vue":
				template = filepath.Join(web, template)
			case ".go":
				template = filepath.Join(server, template)
			}
		} // value
		templates[key] = template
	}
	history.Templates = templates
	for key, value := range history.Injections {
		var injection ast.Ast
		switch key {
		case ast.TypePackageApiEnter, ast.TypePackageRouterEnter, ast.TypePackageServiceEnter:

		case ast.TypePackageApiModuleEnter, ast.TypePackageRouterModuleEnter, ast.TypePackageServiceModuleEnter:
			var entity ast.PackageModuleEnter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePackageInitializeGorm:
			var entity ast.PackageInitializeGorm
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePackageInitializeRouter:
			var entity ast.PackageInitializeRouter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePluginGen:
			var entity ast.PluginGen
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePluginApiEnter, ast.TypePluginRouterEnter, ast.TypePluginServiceEnter:
			var entity ast.PluginEnter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePluginInitializeGorm:
			var entity ast.PluginInitializeGorm
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast.TypePluginInitializeRouter:
			var entity ast.PluginInitializeRouter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		}
		if injection == nil {
			continue
		}
		file, _ := injection.Parse("", nil)
		if file != nil {
			_ = injection.Rollback(file)
			err = injection.Format("", nil, file)
			if err != nil {
				return err
			}
			fmt.Printf("[filepath:%s]!\n", key)
		}
	} //
	removeBasePath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, "rm_file", strconv.FormatInt(int64(time.Now().Nanosecond()), 10))
	for _, value := range history.Templates {
		if !filepath.IsAbs(value) {
			continue
		}
		removePath := filepath.Join(removeBasePath, strings.TrimPrefix(value, global.GVA_CONFIG.AutoCode.Root))
		err = utils.FileMove(value, removePath)
		if err != nil {
			return errors.Wrapf(err, "[src:%s][dst:%s]!", value, removePath)
		}
	} //
	err = global.GVA_DB.WithContext(ctx).Model(&model.SysAutoCodeHistory{}).Where("id = ?", info.ID).Update("flag", 1).Error
	if err != nil {
		return errors.Wrap(err, "!")
	}
	return nil
}

// Delete
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) Delete(ctx context.Context, info common.GetById) error {
	err := global.GVA_DB.WithContext(ctx).Where("id = ?", info.Uint()).Delete(&model.SysAutoCodeHistory{}).Error
	if err != nil {
		return errors.Wrap(err, "!")
	}
	return nil
}

// GetList
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (s *autoCodeHistory) GetList(ctx context.Context, info common.PageInfo) (list []model.SysAutoCodeHistory, total int64, err error) {
	var entities []model.SysAutoCodeHistory
	db := global.GVA_DB.WithContext(ctx).Model(&model.SysAutoCodeHistory{})
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Scopes(info.Paginate()).Order("updated_at desc").Find(&entities).Error
	return entities, total, err
}

// DropTable ,
// @author: [piexlmax](https://github.com/piexlmax)
func (s *autoCodeHistory) DropTable(BusinessDb, tableName string) error {
	if BusinessDb != "" {
		return global.MustGetGlobalDBByDBName(BusinessDb).Exec("DROP TABLE " + tableName).Error
	} else {
		return global.GVA_DB.Exec("DROP TABLE " + tableName).Error
	}
}
