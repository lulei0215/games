package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type ApiLoginResponse struct {
	User      system.ApiSysUser `json:"user"`
	Token     string            `json:"token"`
	ExpiresAt int64             `json:"expiresAt"`
}
