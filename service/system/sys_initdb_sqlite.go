package system

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"github.com/gookit/color"
	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SqliteInitHandler struct{}

func NewSqliteInitHandler() *SqliteInitHandler {
	return &SqliteInitHandler{}
}

// WriteConfig mysql
func (h SqliteInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Sqlite)
	if !ok {
		return errors.New("sqlite config invalid")
	}
	global.GVA_CONFIG.System.DbType = "sqlite"
	global.GVA_CONFIG.Sqlite = c
	global.GVA_CONFIG.JWT.SigningKey = uuid.New().String()
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	global.GVA_ACTIVE_DBNAME = &c.Dbname
	return global.GVA_VP.WriteConfig()
}

// EnsureDB  sqlite
func (h SqliteInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "sqlite" {
		return ctx, ErrDBTypeMismatch
	}

	c := conf.ToSqliteConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	} // ,

	dsn := conf.SqliteEmptyDsn()

	var db *gorm.DB
	if db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return ctx, err
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h SqliteInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h SqliteInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf(InitDataExist, Sqlite, init.InitializerName())
			continue
		}
		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Sqlite, init.InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Sqlite, init.InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Sqlite)
	return nil
}
