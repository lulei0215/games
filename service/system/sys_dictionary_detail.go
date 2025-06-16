package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSysDictionaryDetail
//@description:
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

type DictionaryDetailService struct{}

var DictionaryDetailServiceApp = new(DictionaryDetailService)

func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Create(&sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionaryDetail
//@description:
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Delete(&sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysDictionaryDetail
//@description:
//@param: sysDictionaryDetail *model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(sysDictionaryDetail *system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Save(sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionaryDetail
//@description: id
//@param: id uint
//@return: sysDictionaryDetail system.SysDictionaryDetail, err error

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (sysDictionaryDetail system.SysDictionaryDetail, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionaryDetailInfoList
//@description:
//@param: info request.SysDictionaryDetailSearch
//@return: list interface{}, total int64, err error

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	//
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != "" {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&sysDictionaryDetails).Error
	return sysDictionaryDetails, total, err
}

// id
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryList(dictionaryID uint) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	err = global.GVA_DB.Find(&sysDictionaryDetails, "sys_dictionary_id = ?", dictionaryID).Error
	return sysDictionaryDetails, err
}

// type
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryListByType(t string) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id")
	err = db.Debug().Find(&sysDictionaryDetails, "type = ?", t).Error
	return sysDictionaryDetails, err
}

// id+value
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryInfoByValue(dictionaryID uint, value string) (detail system.SysDictionaryDetail, err error) {
	var sysDictionaryDetail system.SysDictionaryDetail
	err = global.GVA_DB.First(&sysDictionaryDetail, "sys_dictionary_id = ? and value = ?", dictionaryID, value).Error
	return sysDictionaryDetail, err
}

// type+value
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryInfoByTypeValue(t string, value string) (detail system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails system.SysDictionaryDetail
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id")
	err = db.First(&sysDictionaryDetails, "sys_dictionaries.type = ? and sys_dictionary_details.value = ?", t, value).Error
	return sysDictionaryDetails, err
}
