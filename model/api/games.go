// Games
package api

import (
	"time"
)

// games表   Games
type Games struct {
	Id        uint       `json:"id" form:"id" gorm:"primarykey;comment:id;column:id;size:19;"`             //id
	Img       string     `json:"img" form:"img" gorm:"comment:img;column:img;size:255;"`                   //img
	ImgEn     string     `json:"imgEn" form:"imgEn" gorm:"comment:img_en;column:img_en;size:255;"`         //img_en
	ImgPt     string     `json:"imgPt" form:"imgPt" gorm:"comment:img_pt;column:img_pt;size:255;"`         //img_pt
	Name      string     `json:"name" form:"name" gorm:"comment:name;column:name;size:255;"`               //name
	NameEn    string     `json:"nameEn" form:"nameEn" gorm:"comment:name_en;column:name_en;size:255;"`     //name_en
	NamePt    string     `json:"namePt" form:"namePt" gorm:"comment:name_pt;column:name_pt;size:255;"`     //name_pt
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                     //createdAt字段
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                     //updatedAt字段
	SessionId string     `json:"sessionId" form:"sessionId" gorm:"column:session_id;size:255;"`            //sessionId字段
	Status    int        `json:"status" form:"status" gorm:"comment:status;column:status;"`                //status
	Title     string     `json:"title" form:"title" gorm:"comment:title;column:title;size:255;"`           //title
	TitleEn   string     `json:"titleEn" form:"titleEn" gorm:"comment:title_en;column:title_en;size:255;"` //title_en
	TitlePt   string     `json:"titlePt" form:"titlePt" gorm:"comment:title_pt;column:title_pt;size:255;"` //title_pt
	Sort      int        `json:"sort" form:"sort" gorm:"comment:sort;column:sort;"`                        //sort
	IsHot     int        `json:"isHot" form:"isHot" gorm:"column:is_hot;"`                                 //isHot字段
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`                     //deletedAt字段
	Content   string     `json:"content" form:"content" gorm:"column:content;size:500;"`                   //content字段
	ContentEn string     `json:"contentEn" form:"contentEn" gorm:"column:content_en;size:500;"`            //content_en字段
	ContentPt string     `json:"contentPt" form:"contentPt" gorm:"column:content_pt;size:500;"`            //content_pt字段
	Url       string     `json:"url" form:"url" gorm:"column:url;size:500;"`                               //content字段
}

// TableName games表 Games games
func (Games) TableName() string {
	return "games"
}
