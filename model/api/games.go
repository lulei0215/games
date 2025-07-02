// Games
package api

import (
	"time"
)

// games表   Games
type Games struct {
	Id        *int       `json:"id" form:"id" gorm:"primarykey;comment:id;column:id;size:19;"`   //id
	Img       string     `json:"img" form:"img" gorm:"comment:img;column:img;size:255;"`         //img
	Name      *string    `json:"name" form:"name" gorm:"comment:name;column:name;size:255;"`     //name
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`           //createdAt字段
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`           //updatedAt字段
	SessionId *string    `json:"sessionId" form:"sessionId" gorm:"column:session_id;size:255;"`  //sessionId字段
	Status    *bool      `json:"status" form:"status" gorm:"comment:status;column:status;"`      //status
	Title     *string    `json:"title" form:"title" gorm:"comment:title;column:title;size:255;"` //title
	Sort      *bool      `json:"sort" form:"sort" gorm:"comment:sort;column:sort;"`              //sort
	IsHot     *bool      `json:"isHot" form:"isHot" gorm:"column:is_hot;"`                       //isHot字段
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`           //deletedAt字段
	Content   *string    `json:"content" form:"content" gorm:"column:content;size:500;"`         //content字段
	Url       *string    `json:"url" form:"url" gorm:"column:url;size:500;"`                     //content字段
}

// TableName games表 Games games
func (Games) TableName() string {
	return "games"
}
