package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/google/uuid"
)

type Login interface {
	GetUsername() string
	GetNickname() string
	GetUUID() uuid.UUID
	GetUserId() uint
	GetAuthorityId() uint
	GetUserInfo() any
}

var _ Login = new(SysUser)

type SysUser struct {
	global.GVA_MODEL
	UUID             uuid.UUID      `json:"uuid" gorm:"index;comment:UUID"`                                                                   // UUID
	Username         string         `json:"userName" gorm:"index;comment:"`                                                                   //
	Password         string         `json:"-"  gorm:"comment:"`                                                                               //
	NickName         string         `json:"nickName" gorm:"default:;comment:"`                                                                //
	HeaderImg        string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:"`                 //
	AuthorityId      uint           `json:"authorityId" gorm:"default:888;comment:ID"`                                                        // ID
	Authority        SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:"`                          //
	Authorities      []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`                                                 //
	Phone            string         `json:"phone"  gorm:"comment:"`                                                                           //
	Email            string         `json:"email"  gorm:"comment:"`                                                                           //
	Enable           int            `json:"enable" gorm:"default:1;comment: 1 2"`                                                             // 1 2
	OriginSetting    common.JSONMap `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:;"` //
	WithdrawPassword string         `json:"withdrawPassword" gorm:"comment:"`                                                                 //
	Balance          float64        `json:"balance" gorm:"type:decimal(20,2);default:0.00;comment:"`                                          //
	Birthday         *time.Time     `json:"birthday" gorm:"type:date;comment:"`                                                               //
	Facebook         string         `json:"facebook" gorm:"comment:Facebook"`                                                                 // Facebook
	Whatsapp         string         `json:"whatsapp" gorm:"comment:WhatsApp"`                                                                 // WhatsApp
	Telegram         string         `json:"telegram" gorm:"comment:Telegram"`                                                                 // Telegram
	Twitter          string         `json:"twitter" gorm:"comment:Twitter"`                                                                   // Twitter
	VipLevel         uint8          `json:"vipLevel" gorm:"type:tinyint(4);default:0;comment:VIP 0- 1-5VIP"`                                  // VIP
	VipExpireTime    *time.Time     `json:"vipExpireTime" gorm:"comment:VIP"`
	UserType         int            `json:"userType" gorm:"default:2;comment:2putong"`
	Level            int            `json:"level" gorm:"default:3;comment:2putong"`
}
type ApiSysUser struct {
	UUID             uuid.UUID  `json:"uuid" gorm:"index;comment:UUID"`                                                   // UUID
	Username         string     `json:"userName" gorm:"index;comment:"`                                                   //
	Password         string     `json:"-"  gorm:"comment:"`                                                               //
	NickName         string     `json:"nickName" gorm:"default:;comment:"`                                                //
	HeaderImg        string     `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:"` //
	Phone            string     `json:"phone"  gorm:"comment:"`                                                           //
	Email            string     `json:"email"  gorm:"comment:"`                                                           //
	Enable           int        `json:"enable" gorm:"default:1;comment: 1 2"`                                             // 1 2
	WithdrawPassword string     `json:"withdrawPassword" gorm:"comment:"`                                                 //
	Balance          float64    `json:"balance" gorm:"type:decimal(20,2);default:0.00;comment:"`                          //
	Birthday         *time.Time `json:"birthday" gorm:"type:date;comment:"`                                               //
	Facebook         string     `json:"facebook" gorm:"comment:Facebook"`                                                 // Facebook
	Whatsapp         string     `json:"whatsapp" gorm:"comment:WhatsApp"`                                                 // WhatsApp
	Telegram         string     `json:"telegram" gorm:"comment:Telegram"`                                                 // Telegram
	Twitter          string     `json:"twitter" gorm:"comment:Twitter"`                                                   // Twitter
	VipLevel         uint8      `json:"vipLevel" gorm:"type:tinyint(4);default:0;comment:VIP 0- 1-5VIP"`                  // VIP
	VipExpireTime    *time.Time `json:"vipExpireTime" gorm:"comment:VIP"`
	UserType         int        `json:"userType" gorm:"default:2;comment:2putong"`
	Level            int        `json:"level" gorm:"default:3;comment:2putong"`
}

func (SysUser) TableName() string {
	return "sys_users"
}

func (s *SysUser) GetUsername() string {
	return s.Username
}

func (s *SysUser) GetNickname() string {
	return s.NickName
}

func (s *SysUser) GetUUID() uuid.UUID {
	return s.UUID
}

func (s *SysUser) GetUserId() uint {
	return s.ID
}

func (s *SysUser) GetAuthorityId() uint {
	return s.AuthorityId
}

func (s *SysUser) GetUserInfo() any {
	return *s
}
