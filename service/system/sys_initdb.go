package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gorm.io/gorm"
)

const (
	Mysql           = "mysql"
	Pgsql           = "pgsql"
	Sqlite          = "sqlite"
	Mssql           = "mssql"
	InitSuccess     = "\n[%v] --> !\n"
	InitDataExist   = "\n[%v] --> %v !\n"
	InitDataFailed  = "\n[%v] --> %v ! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v !\n"
)

const (
	InitOrderSystem   = 10
	InitOrderInternal = 1000
	InitOrderExternal = 100000
)

var (
	ErrMissingDBContext        = errors.New("missing db in context")
	ErrMissingDependentContext = errors.New("missing dependent value in context")
	ErrDBTypeMismatch          = errors.New("db type mismatch")
)

// SubInitializer  source/*/init() ， initializer
type SubInitializer interface {
	InitializerName() string // ，
	MigrateTable(ctx context.Context) (next context.Context, err error)
	InitializeData(ctx context.Context) (next context.Context, err error)
	TableCreated(ctx context.Context) bool
	DataInserted(ctx context.Context) bool
}

// TypedDBInitHandler  initializer
type TypedDBInitHandler interface {
	EnsureDB(ctx context.Context, conf *request.InitDB) (context.Context, error) // ， fatal error， panic
	WriteConfig(ctx context.Context) error                                       //
	InitTables(ctx context.Context, inits initSlice) error                       //  handler
	InitData(ctx context.Context, inits initSlice) error                         //  handler
}

// orderedInitializer ，
type orderedInitializer struct {
	order int
	SubInitializer
}

// initSlice  initializer
type initSlice []*orderedInitializer

var (
	initializers initSlice
	cache        map[string]*orderedInitializer
)

// RegisterInit ， InitDB()
func RegisterInit(order int, i SubInitializer) {
	if initializers == nil {
		initializers = initSlice{}
	}
	if cache == nil {
		cache = map[string]*orderedInitializer{}
	}
	name := i.InitializerName()
	if _, existed := cache[name]; existed {
		panic(fmt.Sprintf("Name conflict on %s", name))
	}
	ni := orderedInitializer{order, i}
	initializers = append(initializers, &ni)
	cache[name] = &ni
}

/* ---- * service * ---- */

type InitDBService struct{}

// InitDB
func (initDBService *InitDBService) InitDB(conf request.InitDB) (err error) {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "adminPassword", conf.AdminPassword)
	if len(initializers) == 0 {
		return errors.New("，")
	}
	sort.Sort(&initializers) //  initializer
	// Note:  initializer ， B=A+1, C=A+1;  BC ，
	// ， C=A+B, D=A+B+C, E=A+1;
	// C>A|B，AB，D>A|B|C，ABC，EA，CD，ECD
	var initHandler TypedDBInitHandler
	switch conf.DBType {
	case "mysql":
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	case "pgsql":
		initHandler = NewPgsqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "pgsql")
	case "sqlite":
		initHandler = NewSqliteInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "sqlite")
	case "mssql":
		initHandler = NewMssqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mssql")
	default:
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	}
	ctx, err = initHandler.EnsureDB(ctx, &conf)
	if err != nil {
		return err
	}

	db := ctx.Value("db").(*gorm.DB)
	global.GVA_DB = db

	if err = initHandler.InitTables(ctx, initializers); err != nil {
		return err
	}
	if err = initHandler.InitData(ctx, initializers); err != nil {
		return err
	}

	if err = initHandler.WriteConfig(ctx); err != nil {
		return err
	}
	initializers = initSlice{}
	cache = map[string]*orderedInitializer{}
	return nil
}

// createDatabase （ EnsureDB()  ）
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

// createTables （ dbInitHandler.initTables ）
func createTables(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.TableCreated(next) {
			continue
		}
		if n, err := init.MigrateTable(next); err != nil {
			return err
		} else {
			next = n
		}
	}
	return nil
}

/* -- sortable interface -- */

func (a initSlice) Len() int {
	return len(a)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}

func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
