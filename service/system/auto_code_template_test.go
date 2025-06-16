package system

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

func Test_autoCodeTemplate_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		info request.AutoCode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &autoCodeTemplate{}
			if err := s.Create(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_autoCodeTemplate_Preview(t *testing.T) {
	type args struct {
		ctx  context.Context
		info request.AutoCode
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: " package",
			args: args{
				ctx:  context.Background(),
				info: request.AutoCode{},
			},
			wantErr: false,
		},
		{
			name: " plugin",
			args: args{
				ctx:  context.Background(),
				info: request.AutoCode{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testJson := `{"structName":"SysUser","tableName":"sys_users","packageName":"sysUsers","package":"gva","abbreviation":"sysUsers","description":"sysUsers","businessDB":"","autoCreateApiToSql":true,"autoCreateMenuToSql":true,"autoMigrate":true,"gvaModel":true,"autoCreateResource":false,"fields":[{"fieldName":"Uuid","fieldDesc":"UUID","fieldType":"string","dataType":"varchar","fieldJson":"uuid","primaryKey":false,"dataTypeLong":"191","columnName":"uuid","comment":"UUID","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Username","fieldDesc":"","fieldType":"string","dataType":"varchar","fieldJson":"username","primaryKey":false,"dataTypeLong":"191","columnName":"username","comment":"","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Password","fieldDesc":"","fieldType":"string","dataType":"varchar","fieldJson":"password","primaryKey":false,"dataTypeLong":"191","columnName":"password","comment":"","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"NickName","fieldDesc":"","fieldType":"string","dataType":"varchar","fieldJson":"nickName","primaryKey":false,"dataTypeLong":"191","columnName":"nick_name","comment":"","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"SideMode","fieldDesc":"","fieldType":"string","dataType":"varchar","fieldJson":"sideMode","primaryKey":false,"dataTypeLong":"191","columnName":"side_mode","comment":"","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"HeaderImg","fieldDesc":"","fieldType":"string","dataType":"varchar","fieldJson":"headerImg","primaryKey":false,"dataTypeLong":"191","columnName":"header_img","comment":"","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"BaseColor","fieldDesc":"","fieldType":"string","dataType":"varchar","fieldJson":"baseColor","primaryKey":false,"dataTypeLong":"191","columnName":"base_color","comment":"","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"AuthorityId","fieldDesc":"ID","fieldType":"int","dataType":"bigint","fieldJson":"authorityId","primaryKey":false,"dataTypeLong":"20","columnName":"authority_id","comment":"ID","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Phone","fieldDesc":"","fieldType":"string","dataType":"varchar","fieldJson":"phone","primaryKey":false,"dataTypeLong":"191","columnName":"phone","comment":"","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Email","fieldDesc":"","fieldType":"string","dataType":"varchar","fieldJson":"email","primaryKey":false,"dataTypeLong":"191","columnName":"email","comment":"","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Enable","fieldDesc":" 1 2","fieldType":"int","dataType":"bigint","fieldJson":"enable","primaryKey":false,"dataTypeLong":"19","columnName":"enable","comment":" 1 2","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}}],"humpPackageName":"sys_users"}`
			err := json.Unmarshal([]byte(testJson), &tt.args.info)
			if err != nil {
				t.Error(err)
				return
			}
			err = tt.args.info.Pretreatment()
			if err != nil {
				t.Error(err)
				return
			}
			got, err := AutoCodeTemplate.Preview(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("Preview() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Preview() got = %v, want %v", got, tt.want)
			}
		})
	}
}
