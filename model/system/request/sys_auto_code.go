package request

import (
	"encoding/json"
	"fmt"
	"go/token"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
)

type AutoCode struct {
	Package             string                 `json:"package"`
	PackageT            string                 `json:"-"`
	TableName           string                 `json:"tableName" example:""`                //
	BusinessDB          string                 `json:"businessDB" example:""`               //
	StructName          string                 `json:"structName" example:"Struct"`         // Struct
	PackageName         string                 `json:"packageName" example:""`              //
	Description         string                 `json:"description" example:"Struct"`        // Struct
	Abbreviation        string                 `json:"abbreviation" example:"Struct"`       // Struct
	HumpPackageName     string                 `json:"humpPackageName" example:"go"`        // go
	GvaModel            bool                   `json:"gvaModel" example:"false"`            // gvaModel
	AutoMigrate         bool                   `json:"autoMigrate" example:"false"`         //
	AutoCreateResource  bool                   `json:"autoCreateResource" example:"false"`  //
	AutoCreateApiToSql  bool                   `json:"autoCreateApiToSql" example:"false"`  // api
	AutoCreateMenuToSql bool                   `json:"autoCreateMenuToSql" example:"false"` // menu
	AutoCreateBtnAuth   bool                   `json:"autoCreateBtnAuth" example:"false"`   //
	OnlyTemplate        bool                   `json:"onlyTemplate" example:"false"`        //
	IsTree              bool                   `json:"isTree" example:"false"`              //
	TreeJson            string                 `json:"treeJson" example:"json"`             // json
	IsAdd               bool                   `json:"isAdd" example:"false"`               //
	Fields              []*AutoCodeField       `json:"fields"`
	GenerateWeb         bool                   `json:"generateWeb" example:"true"`    // web
	GenerateServer      bool                   `json:"generateServer" example:"true"` // server
	Module              string                 `json:"-"`
	DictTypes           []string               `json:"-"`
	PrimaryField        *AutoCodeField         `json:"primaryField"`
	DataSourceMap       map[string]*DataSource `json:"-"`
	HasPic              bool                   `json:"-"`
	HasFile             bool                   `json:"-"`
	HasTimer            bool                   `json:"-"`
	NeedSort            bool                   `json:"-"`
	NeedJSON            bool                   `json:"-"`
	HasRichText         bool                   `json:"-"`
	HasDataSource       bool                   `json:"-"`
	HasSearchTimer      bool                   `json:"-"`
	HasArray            bool                   `json:"-"`
	HasExcel            bool                   `json:"-"`
}

type DataSource struct {
	DBName       string `json:"dbName"`
	Table        string `json:"table"`
	Label        string `json:"label"`
	Value        string `json:"value"`
	Association  int    `json:"association"` //  1  2
	HasDeletedAt bool   `json:"hasDeletedAt"`
}

func (r *AutoCode) Apis() []model.SysApi {
	return []model.SysApi{
		{
			Path:        "/" + r.Abbreviation + "/" + "create" + r.StructName,
			Description: "" + r.Description,
			ApiGroup:    r.Description,
			Method:      "POST",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "delete" + r.StructName,
			Description: "" + r.Description,
			ApiGroup:    r.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "delete" + r.StructName + "ByIds",
			Description: "" + r.Description,
			ApiGroup:    r.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "update" + r.StructName,
			Description: "" + r.Description,
			ApiGroup:    r.Description,
			Method:      "PUT",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "find" + r.StructName,
			Description: "ID" + r.Description,
			ApiGroup:    r.Description,
			Method:      "GET",
		},
		{
			Path:        "/" + r.Abbreviation + "/" + "get" + r.StructName + "List",
			Description: "" + r.Description + "",
			ApiGroup:    r.Description,
			Method:      "GET",
		},
	}
}

func (r *AutoCode) Menu(template string) model.SysBaseMenu {
	component := fmt.Sprintf("view/%s/%s/%s.vue", r.Package, r.PackageName, r.PackageName)
	if template != "package" {
		component = fmt.Sprintf("plugin/%s/view/%s.vue", r.Package, r.PackageName)
	}
	return model.SysBaseMenu{
		ParentId:  0,
		Path:      r.Abbreviation,
		Name:      r.Abbreviation,
		Component: component,
		Meta: model.Meta{
			Title: r.Description,
		},
	}
}

// Pretreatment
// Author [SliverHorn](https://github.com/SliverHorn)
func (r *AutoCode) Pretreatment() error {
	r.Module = global.GVA_CONFIG.AutoCode.Module
	if token.IsKeyword(r.Abbreviation) {
		r.Abbreviation = r.Abbreviation + "_"
	} // go
	if strings.HasSuffix(r.HumpPackageName, "test") {
		r.HumpPackageName = r.HumpPackageName + "_"
	} // test
	length := len(r.Fields)
	dict := make(map[string]string, length)
	r.DataSourceMap = make(map[string]*DataSource, length)
	for i := 0; i < length; i++ {
		if r.Fields[i].Excel {
			r.HasExcel = true
		}
		if r.Fields[i].DictType != "" {
			dict[r.Fields[i].DictType] = ""
		}
		if r.Fields[i].Sort {
			r.NeedSort = true
		}
		switch r.Fields[i].FieldType {
		case "file":
			r.HasFile = true
			r.NeedJSON = true
		case "json":
			r.NeedJSON = true
		case "array":
			r.NeedJSON = true
			r.HasArray = true
		case "video":
			r.HasPic = true
		case "richtext":
			r.HasRichText = true
		case "picture":
			r.HasPic = true
		case "pictures":
			r.HasPic = true
			r.NeedJSON = true
		case "time.Time":
			r.HasTimer = true
			if r.Fields[i].FieldSearchType != "" && r.Fields[i].FieldSearchType != "BETWEEN" && r.Fields[i].FieldSearchType != "NOT BETWEEN" {
				r.HasSearchTimer = true
			}
		}
		if r.Fields[i].DataSource != nil {
			if r.Fields[i].DataSource.Table != "" && r.Fields[i].DataSource.Label != "" && r.Fields[i].DataSource.Value != "" {
				r.HasDataSource = true
				r.Fields[i].CheckDataSource = true
				r.DataSourceMap[r.Fields[i].FieldJson] = r.Fields[i].DataSource
			}
		}
		if !r.GvaModel && r.PrimaryField == nil && r.Fields[i].PrimaryKey {
			r.PrimaryField = r.Fields[i]
		} //
	}
	{
		for key := range dict {
			r.DictTypes = append(r.DictTypes, key)
		}
	} // DictTypes =>
	{
		if r.GvaModel {
			r.PrimaryField = &AutoCodeField{
				FieldName:    "ID",
				FieldType:    "uint",
				FieldDesc:    "ID",
				FieldJson:    "ID",
				DataTypeLong: "20",
				Comment:      "ID",
				ColumnName:   "id",
			}
		}
	} // GvaModel
	{
		if r.IsAdd && r.PrimaryField == nil {
			r.PrimaryField = new(AutoCodeField)
		}
	} //
	if r.Package == "" {
		return errors.New("Package!")
	} // ：Package
	packages := []rune(r.Package)
	if len(packages) > 0 {
		if packages[0] >= 97 && packages[0] <= 122 {
			packages[0] = packages[0] - 32
		}
		r.PackageT = string(packages)
	} // PackageT  Package
	return nil
}

func (r *AutoCode) History() SysAutoHistoryCreate {
	bytes, _ := json.Marshal(r)
	return SysAutoHistoryCreate{
		Table:       r.TableName,
		Package:     r.Package,
		Request:     string(bytes),
		StructName:  r.StructName,
		BusinessDB:  r.BusinessDB,
		Description: r.Description,
	}
}

type AutoCodeField struct {
	FieldName       string `json:"fieldName"`       // Field
	FieldDesc       string `json:"fieldDesc"`       //
	FieldType       string `json:"fieldType"`       // Field
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataTypeLong    string `json:"dataTypeLong"`    //
	Comment         string `json:"comment"`         //
	ColumnName      string `json:"columnName"`      //
	FieldSearchType string `json:"fieldSearchType"` //
	FieldSearchHide bool   `json:"fieldSearchHide"` //
	DictType        string `json:"dictType"`        //
	//Front           bool        `json:"front"`           //
	Form            bool        `json:"form"`            // /
	Table           bool        `json:"table"`           //
	Desc            bool        `json:"desc"`            //
	Excel           bool        `json:"excel"`           // /
	Require         bool        `json:"require"`         //
	DefaultValue    string      `json:"defaultValue"`    //
	ErrorText       string      `json:"errorText"`       //
	Clearable       bool        `json:"clearable"`       //
	Sort            bool        `json:"sort"`            //
	PrimaryKey      bool        `json:"primaryKey"`      //
	DataSource      *DataSource `json:"dataSource"`      //
	CheckDataSource bool        `json:"checkDataSource"` //
	FieldIndexType  string      `json:"fieldIndexType"`  //
}

type AutoFunc struct {
	Package         string `json:"package"`
	FuncName        string `json:"funcName"`        //
	Router          string `json:"router"`          //
	FuncDesc        string `json:"funcDesc"`        //
	BusinessDB      string `json:"businessDB"`      //
	StructName      string `json:"structName"`      // Struct
	PackageName     string `json:"packageName"`     //
	Description     string `json:"description"`     // Struct
	Abbreviation    string `json:"abbreviation"`    // Struct
	HumpPackageName string `json:"humpPackageName"` // go
	Method          string `json:"method"`          //
	IsPlugin        bool   `json:"isPlugin"`        //
	IsAuth          bool   `json:"isAuth"`          //
	IsPreview       bool   `json:"isPreview"`       //
	IsAi            bool   `json:"isAi"`            // AI
	ApiFunc         string `json:"apiFunc"`         // API
	ServerFunc      string `json:"serverFunc"`      //
	JsFunc          string `json:"jsFunc"`          // JS
}

type InitMenu struct {
	PlugName   string `json:"plugName"`
	ParentMenu string `json:"parentMenu"`
	Menus      []uint `json:"menus"`
}

type InitApi struct {
	PlugName string `json:"plugName"`
	APIs     []uint `json:"apis"`
}

type LLMAutoCode struct {
	Prompt string `json:"prompt" form:"prompt" gorm:"column:prompt;comment:;type:text;"` //
	Mode   string `json:"mode" form:"mode" gorm:"column:mode;comment:;type:text;"`       //
}
