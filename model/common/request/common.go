package request

import (
	"gorm.io/gorm"
)

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         //
	PageSize int    `json:"pageSize" form:"pageSize"` //
	Keyword  string `json:"keyword" form:"keyword"`   //
}

func (r *PageInfo) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page <= 0 {
			r.Page = 1
		}
		switch {
		case r.PageSize > 100:
			r.PageSize = 100
		case r.PageSize <= 0:
			r.PageSize = 10
		}
		offset := (r.Page - 1) * r.PageSize
		return db.Offset(offset).Limit(r.PageSize)
	}
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // ID
}

type Empty struct{}
