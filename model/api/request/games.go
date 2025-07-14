package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GamesSearch struct {
	request.PageInfo
	// 多语言搜索字段
	Name      string `json:"name" form:"name"`           // 名称搜索
	NameEn    string `json:"nameEn" form:"nameEn"`       // 英文名称搜索
	NamePt    string `json:"namePt" form:"namePt"`       // 葡萄牙语名称搜索
	Title     string `json:"title" form:"title"`         // 标题搜索
	TitleEn   string `json:"titleEn" form:"titleEn"`     // 英文标题搜索
	TitlePt   string `json:"titlePt" form:"titlePt"`     // 葡萄牙语标题搜索
	Content   string `json:"content" form:"content"`     // 内容搜索
	ContentEn string `json:"contentEn" form:"contentEn"` // 英文内容搜索
	ContentPt string `json:"contentPt" form:"contentPt"` // 葡萄牙语内容搜索
	Status    int    `json:"status" form:"status"`       // 状态搜索
	IsHot     int    `json:"isHot" form:"isHot"`         // 热门搜索
}
