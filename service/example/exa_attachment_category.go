package example

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"gorm.io/gorm"
)

type AttachmentCategoryService struct{}

// AddCategory /
func (a *AttachmentCategoryService) AddCategory(req *example.ExaAttachmentCategory) (err error) {
	//
	if (!errors.Is(global.GVA_DB.Take(&example.ExaAttachmentCategory{}, "name = ? and pid = ?", req.Name, req.Pid).Error, gorm.ErrRecordNotFound)) {
		return errors.New("")
	}
	if req.ID > 0 {
		if err = global.GVA_DB.Model(&example.ExaAttachmentCategory{}).Where("id = ?", req.ID).Updates(&example.ExaAttachmentCategory{
			Name: req.Name,
			Pid:  req.Pid,
		}).Error; err != nil {
			return err
		}
	} else {
		if err = global.GVA_DB.Create(&example.ExaAttachmentCategory{
			Name: req.Name,
			Pid:  req.Pid,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

// DeleteCategory
func (a *AttachmentCategoryService) DeleteCategory(id *int) error {
	var childCount int64
	global.GVA_DB.Model(&example.ExaAttachmentCategory{}).Where("pid = ?", id).Count(&childCount)
	if childCount > 0 {
		return errors.New("")
	}
	return global.GVA_DB.Where("id = ?", id).Unscoped().Delete(&example.ExaAttachmentCategory{}).Error
}

// GetCategoryList
func (a *AttachmentCategoryService) GetCategoryList() (res []*example.ExaAttachmentCategory, err error) {
	var fileLists []example.ExaAttachmentCategory
	err = global.GVA_DB.Model(&example.ExaAttachmentCategory{}).Find(&fileLists).Error
	if err != nil {
		return res, err
	}
	return a.getChildrenList(fileLists, 0), nil
}

// getChildrenList
func (a *AttachmentCategoryService) getChildrenList(categories []example.ExaAttachmentCategory, parentID uint) []*example.ExaAttachmentCategory {
	var tree []*example.ExaAttachmentCategory
	for _, category := range categories {
		if category.Pid == parentID {
			category.Children = a.getChildrenList(categories, category.ID)
			tree = append(tree, &category)
		}
	}
	return tree
}
