package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{
			Path:        "/info/createInfo",
			Description: "",
			ApiGroup:    "",
			Method:      "POST",
		},
		{
			Path:        "/info/deleteInfo",
			Description: "",
			ApiGroup:    "",
			Method:      "DELETE",
		},
		{
			Path:        "/info/deleteInfoByIds",
			Description: "",
			ApiGroup:    "",
			Method:      "DELETE",
		},
		{
			Path:        "/info/updateInfo",
			Description: "",
			ApiGroup:    "",
			Method:      "PUT",
		},
		{
			Path:        "/info/findInfo",
			Description: "ID",
			ApiGroup:    "",
			Method:      "GET",
		},
		{
			Path:        "/info/getInfoList",
			Description: "",
			ApiGroup:    "",
			Method:      "GET",
		},
	}
	utils.RegisterApis(entities...)
}
