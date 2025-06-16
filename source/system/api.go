package system

import (
	"context"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initApi struct{}

const initOrderApi = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderApi, &initApi{})
}

func (i *initApi) InitializerName() string {
	return sysModel.SysApi{}.TableName()
}

func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysApi{})
}

func (i *initApi) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysApi{})
}

func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysApi{
		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "jwt(，)"},

		{ApiGroup: "", Method: "DELETE", Path: "/user/deleteUser", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/user/admin_register", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/user/getUserList", Description: ""},
		{ApiGroup: "", Method: "PUT", Path: "/user/setUserInfo", Description: ""},
		{ApiGroup: "", Method: "PUT", Path: "/user/setSelfInfo", Description: "()"},
		{ApiGroup: "", Method: "GET", Path: "/user/getUserInfo", Description: "()"},
		{ApiGroup: "", Method: "POST", Path: "/user/setUserAuthorities", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/user/changePassword", Description: "（)"},
		{ApiGroup: "", Method: "POST", Path: "/user/setUserAuthority", Description: "()"},
		{ApiGroup: "", Method: "POST", Path: "/user/resetPassword", Description: ""},
		{ApiGroup: "", Method: "PUT", Path: "/user/setSelfSetting", Description: ""},

		{ApiGroup: "api", Method: "POST", Path: "/api/createApi", Description: "api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/deleteApi", Description: "Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/updateApi", Description: "Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiList", Description: "api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getAllApis", Description: "api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiById", Description: "api"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: "api"},
		{ApiGroup: "api", Method: "GET", Path: "/api/syncApi", Description: "API"},
		{ApiGroup: "api", Method: "GET", Path: "/api/getApiGroups", Description: ""},
		{ApiGroup: "api", Method: "POST", Path: "/api/enterSyncApi", Description: "API"},
		{ApiGroup: "api", Method: "POST", Path: "/api/ignoreApi", Description: "API"},

		{ApiGroup: "", Method: "POST", Path: "/authority/copyAuthority", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/authority/createAuthority", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/authority/deleteAuthority", Description: ""},
		{ApiGroup: "", Method: "PUT", Path: "/authority/updateAuthority", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/authority/getAuthorityList", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/authority/setDataAuthority", Description: ""},

		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: "api"},
		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: ""},

		{ApiGroup: "", Method: "POST", Path: "/menu/addBaseMenu", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/menu/getMenu", Description: "()"},
		{ApiGroup: "", Method: "POST", Path: "/menu/deleteBaseMenu", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/menu/updateBaseMenu", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/menu/getBaseMenuById", Description: "id"},
		{ApiGroup: "", Method: "POST", Path: "/menu/getMenuList", Description: "menu"},
		{ApiGroup: "", Method: "POST", Path: "/menu/getBaseMenuTree", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/menu/getMenuAuthority", Description: "menu"},
		{ApiGroup: "", Method: "POST", Path: "/menu/addMenuAuthority", Description: "menu"},

		{ApiGroup: "", Method: "GET", Path: "/fileUploadAndDownload/findFile", Description: "（）"},
		{ApiGroup: "", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: ""},

		{ApiGroup: "", Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "（）"},
		{ApiGroup: "", Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/fileUploadAndDownload/editFileName", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/fileUploadAndDownload/importURL", Description: "URL"},

		{ApiGroup: "", Method: "POST", Path: "/system/getServerInfo", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/system/getSystemConfig", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/system/setSystemConfig", Description: ""},

		{ApiGroup: "", Method: "PUT", Path: "/customer/customer", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/customer/customer", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/customer/customer", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/customer/customer", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/customer/customerList", Description: ""},

		{ApiGroup: "", Method: "GET", Path: "/autoCode/getDB", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/autoCode/getTables", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/createTemp", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/preview", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/autoCode/getColumn", Description: "table"},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/installPlugin", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/pubPlug", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/mcp", Description: " MCP Tool "},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/mcpTest", Description: "MCP Tool "},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/mcpList", Description: " MCP ToolList"},

		{ApiGroup: "", Method: "POST", Path: "/autoCode/createPackage", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/autoCode/getTemplates", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/getPackage", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/delPackage", Description: ""},

		{ApiGroup: "", Method: "POST", Path: "/autoCode/getMeta", Description: "meta"},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/rollback", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/getSysHistory", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/delSysHistory", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/autoCode/addFunc", Description: ""},

		{ApiGroup: "", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "ID"},
		{ApiGroup: "", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: ""},

		{ApiGroup: "", Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: ""},
		{ApiGroup: "", Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "ID（）"},
		{ApiGroup: "", Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: ""},

		{ApiGroup: "", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "ID"},
		{ApiGroup: "", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: ""},

		{ApiGroup: "()", Method: "POST", Path: "/simpleUploader/upload", Description: ""},
		{ApiGroup: "()", Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: ""},
		{ApiGroup: "()", Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: ""},

		{ApiGroup: "email", Method: "POST", Path: "/email/emailTest", Description: ""},
		{ApiGroup: "email", Method: "POST", Path: "/email/sendEmail", Description: ""},

		{ApiGroup: "", Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: ""},

		{ApiGroup: "", Method: "POST", Path: "/sysExportTemplate/createSysExportTemplate", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplate", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplateByIds", Description: ""},
		{ApiGroup: "", Method: "PUT", Path: "/sysExportTemplate/updateSysExportTemplate", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/sysExportTemplate/findSysExportTemplate", Description: "ID"},
		{ApiGroup: "", Method: "GET", Path: "/sysExportTemplate/getSysExportTemplateList", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/sysExportTemplate/exportExcel", Description: "Excel"},
		{ApiGroup: "", Method: "GET", Path: "/sysExportTemplate/exportTemplate", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/sysExportTemplate/importExcel", Description: "Excel"},

		{ApiGroup: "", Method: "POST", Path: "/info/createInfo", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/info/deleteInfo", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/info/deleteInfoByIds", Description: ""},
		{ApiGroup: "", Method: "PUT", Path: "/info/updateInfo", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/info/findInfo", Description: "ID"},
		{ApiGroup: "", Method: "GET", Path: "/info/getInfoList", Description: ""},

		{ApiGroup: "", Method: "POST", Path: "/sysParams/createSysParams", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/sysParams/deleteSysParams", Description: ""},
		{ApiGroup: "", Method: "DELETE", Path: "/sysParams/deleteSysParamsByIds", Description: ""},
		{ApiGroup: "", Method: "PUT", Path: "/sysParams/updateSysParams", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/sysParams/findSysParams", Description: "ID"},
		{ApiGroup: "", Method: "GET", Path: "/sysParams/getSysParamsList", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/sysParams/getSysParam", Description: ""},
		{ApiGroup: "", Method: "GET", Path: "/attachmentCategory/getCategoryList", Description: ""},
		{ApiGroup: "", Method: "POST", Path: "/attachmentCategory/addCategory", Description: "/"},
		{ApiGroup: "", Method: "POST", Path: "/attachmentCategory/deleteCategory", Description: ""},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysApi{}.TableName()+"!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ? AND method = ?", "/authorityBtn/canRemoveAuthorityBtn", "POST").
		First(&sysModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
