package system

import (
	"context"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/autocode"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	utilsAst "github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var AutoCodeTemplate = new(autoCodeTemplate)

type autoCodeTemplate struct{}

func (s *autoCodeTemplate) checkPackage(Pkg string, template string) (err error) {
	switch template {
	case "package":
		apiEnter := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", Pkg, "enter.go")
		_, err = os.Stat(apiEnter)
		if err != nil {
			return fmt.Errorf("package,api/v1/%s/enter.go", Pkg)
		}
		serviceEnter := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", Pkg, "enter.go")
		_, err = os.Stat(serviceEnter)
		if err != nil {
			return fmt.Errorf("package,service/%s/enter.go", Pkg)
		}
		routerEnter := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", Pkg, "enter.go")
		_, err = os.Stat(routerEnter)
		if err != nil {
			return fmt.Errorf("package,router/%s/enter.go", Pkg)
		}
	case "plugin":
		pluginEnter := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", Pkg, "plugin.go")
		_, err = os.Stat(pluginEnter)
		if err != nil {
			return fmt.Errorf("plugin,plugin/%s/plugin.go", Pkg)
		}
	}
	return nil
}

// Create
func (s *autoCodeTemplate) Create(ctx context.Context, info request.AutoCode) error {
	history := info.History()
	var autoPkg model.SysAutoCodePackage
	err := global.GVA_DB.WithContext(ctx).Where("package_name = ?", info.Package).First(&autoPkg).Error
	if err != nil {
		return errors.Wrap(err, "!")
	}
	err = s.checkPackage(info.Package, autoPkg.Template)
	if err != nil {
		return err
	}
	// : struct
	if AutocodeHistory.Repeat(info.BusinessDB, info.StructName, info.Abbreviation, info.Package) {
		return errors.New(",!")
	}

	generate, templates, injections, err := s.generate(ctx, info, autoPkg)
	if err != nil {
		return err
	}
	for key, builder := range generate {
		err = os.MkdirAll(filepath.Dir(key), os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "[filepath:%s]!", key)
		}
		err = os.WriteFile(key, []byte(builder.String()), 0666)
		if err != nil {
			return errors.Wrapf(err, "[filepath:%s]!", key)
		}
	}

	// api
	if info.AutoCreateApiToSql && !info.OnlyTemplate {
		apis := info.Apis()
		err := global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			for _, v := range apis {
				var api model.SysApi
				var id uint
				err := tx.Where("path = ? AND method = ?", v.Path, v.Method).First(&api).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = tx.Create(&v).Error; err != nil { //
						return err
					}
					id = v.ID
				} else {
					id = api.ID
				}
				history.ApiIDs = append(history.ApiIDs, id)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// menu
	if info.AutoCreateMenuToSql {
		var entity model.SysBaseMenu
		var id uint
		err := global.GVA_DB.WithContext(ctx).First(&entity, "name = ?", info.Abbreviation).Error
		if err == nil {
			id = entity.ID
		} else {
			entity = info.Menu(autoPkg.Template)
			if info.AutoCreateBtnAuth && !info.OnlyTemplate {
				entity.MenuBtn = []model.SysBaseMenuBtn{
					{SysBaseMenuID: entity.ID, Name: "add", Desc: ""},
					{SysBaseMenuID: entity.ID, Name: "batchDelete", Desc: ""},
					{SysBaseMenuID: entity.ID, Name: "delete", Desc: ""},
					{SysBaseMenuID: entity.ID, Name: "edit", Desc: ""},
					{SysBaseMenuID: entity.ID, Name: "info", Desc: ""},
				}
				if info.HasExcel {
					excelBtn := []model.SysBaseMenuBtn{
						{SysBaseMenuID: entity.ID, Name: "exportTemplate", Desc: ""},
						{SysBaseMenuID: entity.ID, Name: "exportExcel", Desc: "Excel"},
						{SysBaseMenuID: entity.ID, Name: "importExcel", Desc: "Excel"},
					}
					entity.MenuBtn = append(entity.MenuBtn, excelBtn...)
				}
			}
			err = global.GVA_DB.WithContext(ctx).Create(&entity).Error
			id = entity.ID
			if err != nil {
				return errors.Wrap(err, "!")
			}
		}
		history.MenuID = id
	}

	if info.HasExcel {
		dbName := info.BusinessDB
		name := info.Package + "_" + info.StructName
		tableName := info.TableName
		fieldsMap := make(map[string]string, len(info.Fields))
		for _, field := range info.Fields {
			if field.Excel {
				fieldsMap[field.ColumnName] = field.FieldDesc
			}
		}
		templateInfo, _ := json.Marshal(fieldsMap)
		sysExportTemplate := model.SysExportTemplate{
			DBName:       dbName,
			Name:         name,
			TableName:    tableName,
			TemplateID:   name,
			TemplateInfo: string(templateInfo),
		}
		err = SysExportTemplateServiceApp.CreateSysExportTemplate(&sysExportTemplate)
		if err != nil {
			return err
		}
		history.ExportTemplateID = sysExportTemplate.ID
	}

	//
	history.Templates = templates
	history.Injections = make(map[string]string, len(injections))
	for key, value := range injections {
		bytes, _ := json.Marshal(value)
		history.Injections[key] = string(bytes)
	}
	err = AutocodeHistory.Create(ctx, history)
	if err != nil {
		return err
	}
	return nil
}

// Preview
func (s *autoCodeTemplate) Preview(ctx context.Context, info request.AutoCode) (map[string]string, error) {
	var entity model.SysAutoCodePackage
	err := global.GVA_DB.WithContext(ctx).Where("package_name = ?", info.Package).First(&entity).Error
	if err != nil {
		return nil, errors.Wrap(err, "!")
	}
	// : struct
	if AutocodeHistory.Repeat(info.BusinessDB, info.StructName, info.Abbreviation, info.Package) && !info.IsAdd {
		return nil, errors.New(",!")
	}

	preview := make(map[string]string)
	codes, _, _, err := s.generate(ctx, info, entity)
	if err != nil {
		return nil, err
	}
	for key, writer := range codes {
		if len(key) > len(global.GVA_CONFIG.AutoCode.Root) {
			key, _ = filepath.Rel(global.GVA_CONFIG.AutoCode.Root, key)
		}
		// key .
		suffix := filepath.Ext(key)[1:]
		var builder strings.Builder
		builder.WriteString("```" + suffix + "\n\n")
		builder.WriteString(writer.String())
		builder.WriteString("\n\n```")
		preview[key] = builder.String()
	}
	return preview, nil
}

func (s *autoCodeTemplate) generate(ctx context.Context, info request.AutoCode, entity model.SysAutoCodePackage) (map[string]strings.Builder, map[string]string, map[string]utilsAst.Ast, error) {
	templates, asts, _, err := AutoCodePackage.templates(ctx, entity, info, false)
	if err != nil {
		return nil, nil, nil, err
	}
	code := make(map[string]strings.Builder)
	for key, create := range templates {
		var files *template.Template
		files, err = template.New(filepath.Base(key)).Funcs(autocode.GetTemplateFuncMap()).ParseFiles(key)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "[filpath:%s]!", key)
		}
		var builder strings.Builder
		err = files.Execute(&builder, info)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "[filpath:%s]!", create)
		}
		code[create] = builder
	} //
	injections := make(map[string]utilsAst.Ast, len(asts))
	for key, value := range asts {
		keys := strings.Split(key, "=>")
		if len(keys) == 2 {
			if keys[1] == utilsAst.TypePluginInitializeV2 {
				continue
			}
			if info.OnlyTemplate {
				if keys[1] == utilsAst.TypePackageInitializeGorm || keys[1] == utilsAst.TypePluginInitializeGorm {
					continue
				}
			}
			if !info.AutoMigrate {
				if keys[1] == utilsAst.TypePackageInitializeGorm || keys[1] == utilsAst.TypePluginInitializeGorm {
					continue
				}
			}
			var builder strings.Builder
			parse, _ := value.Parse("", &builder)
			if parse != nil {
				_ = value.Injection(parse)
				err = value.Format("", &builder, parse)
				if err != nil {
					return nil, nil, nil, err
				}
				code[keys[0]] = builder
				injections[keys[1]] = value
				fmt.Println(keys[0], "!")
			}
		}
	}
	//
	return code, templates, injections, nil
}

func (s *autoCodeTemplate) AddFunc(info request.AutoFunc) error {
	autoPkg := model.SysAutoCodePackage{}
	err := global.GVA_DB.First(&autoPkg, "package_name = ?", info.Package).Error
	if err != nil {
		return err
	}
	if autoPkg.Template != "package" {
		info.IsPlugin = true
	}
	err = s.addTemplateToFile("api.go", info)
	if err != nil {
		return err
	}
	err = s.addTemplateToFile("server.go", info)
	if err != nil {
		return err
	}
	err = s.addTemplateToFile("api.js", info)
	if err != nil {
		return err
	}
	return s.addTemplateToAst("router", info)
}

func (s *autoCodeTemplate) GetApiAndServer(info request.AutoFunc) (map[string]string, error) {
	autoPkg := model.SysAutoCodePackage{}
	err := global.GVA_DB.First(&autoPkg, "package_name = ?", info.Package).Error
	if err != nil {
		return nil, err
	}
	if autoPkg.Template != "package" {
		info.IsPlugin = true
	}

	apiStr, err := s.getTemplateStr("api.go", info)
	if err != nil {
		return nil, err
	}
	serverStr, err := s.getTemplateStr("server.go", info)
	if err != nil {
		return nil, err
	}
	jsStr, err := s.getTemplateStr("api.js", info)
	if err != nil {
		return nil, err
	}
	return map[string]string{"api": apiStr, "server": serverStr, "js": jsStr}, nil

}

func (s *autoCodeTemplate) getTemplateStr(t string, info request.AutoFunc) (string, error) {
	tempPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", "function", t+".tpl")
	files, err := template.New(filepath.Base(tempPath)).Funcs(autocode.GetTemplateFuncMap()).ParseFiles(tempPath)
	if err != nil {
		return "", errors.Wrapf(err, "[filepath:%s]!", tempPath)
	}
	var builder strings.Builder
	err = files.Execute(&builder, info)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.Wrapf(err, "[filpath:%s]!", tempPath)
	}
	return builder.String(), nil
}

func (s *autoCodeTemplate) addTemplateToAst(t string, info request.AutoFunc) error {
	tPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", info.Package, info.HumpPackageName+".go")
	funcName := fmt.Sprintf("Init%sRouter", info.StructName)

	routerStr := "RouterWithoutAuth"
	if info.IsAuth {
		routerStr = "Router"
	}

	stmtStr := fmt.Sprintf("%s%s.%s(\"%s\", %sApi.%s)", info.Abbreviation, routerStr, info.Method, info.Router, info.Abbreviation, info.FuncName)
	if info.IsPlugin {
		tPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", info.Package, "router", info.HumpPackageName+".go")
		stmtStr = fmt.Sprintf("group.%s(\"%s\", api%s.%s)", info.Method, info.Router, info.StructName, info.FuncName)
		funcName = "Init"
	}

	src, err := os.ReadFile(tPath)
	if err != nil {
		return err
	}

	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	if err != nil {
		return err
	}
	funcDecl := utilsAst.FindFunction(astFile, funcName)
	stmtNode := utilsAst.CreateStmt(stmtStr)

	if info.IsAuth {
		for i := 0; i < len(funcDecl.Body.List); i++ {
			st := funcDecl.Body.List[i]
			// stmt
			if blockStmt, ok := st.(*ast.BlockStmt); ok {
				// ，
				blockStmt.List = append(blockStmt.List, stmtNode)
				break
			}
		}
	} else {
		for i := len(funcDecl.Body.List) - 1; i >= 0; i-- {
			st := funcDecl.Body.List[i]
			// stmt
			if blockStmt, ok := st.(*ast.BlockStmt); ok {
				// ，
				blockStmt.List = append(blockStmt.List, stmtNode)
				break
			}
		}
	}

	//
	f, err := os.Create(tPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := format.Node(f, fileSet, astFile); err != nil {
		return err
	}
	return err
}

func (s *autoCodeTemplate) addTemplateToFile(t string, info request.AutoFunc) error {
	getTemplateStr, err := s.getTemplateStr(t, info)
	if err != nil {
		return err
	}
	var target string

	switch t {
	case "api.go":
		if info.IsAi && info.ApiFunc != "" {
			getTemplateStr = info.ApiFunc
		}
		target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", info.Package, info.HumpPackageName+".go")
	case "server.go":
		if info.IsAi && info.ServerFunc != "" {
			getTemplateStr = info.ServerFunc
		}
		target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", info.Package, info.HumpPackageName+".go")
	case "api.js":
		if info.IsAi && info.JsFunc != "" {
			getTemplateStr = info.JsFunc
		}
		target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Web, "api", info.Package, info.PackageName+".js")
	}
	if info.IsPlugin {
		switch t {
		case "api.go":
			target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", info.Package, "api", info.HumpPackageName+".go")
		case "server.go":
			target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", info.Package, "service", info.HumpPackageName+".go")
		case "api.js":
			target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Web, "plugin", info.Package, "api", info.PackageName+".js")
		}
	}

	// ，
	file, err := os.OpenFile(target, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	//
	_, err = fmt.Fprintln(file, getTemplateStr)
	if err != nil {
		fmt.Printf(": %s\n", err.Error())
		return err
	}

	return nil
}
