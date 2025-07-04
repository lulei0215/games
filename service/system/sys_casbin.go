package system

import (
	"errors"
	"strconv"

	"gorm.io/gorm"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	_ "github.com/go-sql-driver/mysql"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCasbin
//@description: casbin
//@param: authorityId string, casbinInfos []request.CasbinInfo
//@return: error

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (casbinService *CasbinService) UpdateCasbin(adminAuthorityID, AuthorityID uint, casbinInfos []request.CasbinInfo) error {

	err := AuthorityServiceApp.CheckAuthorityIDAuth(adminAuthorityID, AuthorityID)
	if err != nil {
		return err
	}

	if global.GVA_CONFIG.System.UseStrictAuth {
		apis, e := ApiServiceApp.GetAllApis(adminAuthorityID)
		if e != nil {
			return e
		}

		for i := range casbinInfos {
			hasApi := false
			for j := range apis {
				if apis[j].Path == casbinInfos[i].Path && apis[j].Method == casbinInfos[i].Method {
					hasApi = true
					break
				}
			}
			if !hasApi {
				return errors.New("api")
			}
		}
	}

	authorityId := strconv.Itoa(int(AuthorityID))
	casbinService.ClearCasbin(0, authorityId)
	rules := [][]string{}
	//
	deduplicateMap := make(map[string]bool)
	for _, v := range casbinInfos {
		key := authorityId + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
	}
	if len(rules) == 0 {
		return nil
	} //  AddPolicies
	e := utils.GetCasbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("api,,")
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCasbinApi
//@description: API
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GVA_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	if err != nil {
		return err
	}

	e := utils.GetCasbin()
	return e.LoadPolicy()
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPolicyPathByAuthorityId
//@description:
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo

func (casbinService *CasbinService) GetPolicyPathByAuthorityId(AuthorityID uint) (pathMaps []request.CasbinInfo) {
	e := utils.GetCasbin()
	authorityId := strconv.Itoa(int(AuthorityID))
	list, _ := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ClearCasbin
//@description:
//@param: v int, p ...string
//@return: bool

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := utils.GetCasbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RemoveFilteredPolicy
//@description: politicy FreshCasbin
//@param: db *gorm.DB, authorityId string
//@return: error

func (casbinService *CasbinService) RemoveFilteredPolicy(db *gorm.DB, authorityId string) error {
	return db.Delete(&gormadapter.CasbinRule{}, "v0 = ?", authorityId).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SyncPolicy
//@description: policy FreshCasbin
//@param: db *gorm.DB, authorityId string, rules [][]string
//@return: error

func (casbinService *CasbinService) SyncPolicy(db *gorm.DB, authorityId string, rules [][]string) error {
	err := casbinService.RemoveFilteredPolicy(db, authorityId)
	if err != nil {
		return err
	}
	return casbinService.AddPolicies(db, rules)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddPolicies
//@description:
//@param: v int, p ...string
//@return: bool

func (casbinService *CasbinService) AddPolicies(db *gorm.DB, rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return db.Create(&casbinRules).Error
}

func (casbinService *CasbinService) FreshCasbin() (err error) {
	e := utils.GetCasbin()
	err = e.LoadPolicy()
	return err
}
