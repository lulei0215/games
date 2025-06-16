package utils

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PageInfoTest struct {
	PageInfo request.PageInfo
	Name     string
}

func TestVerify(t *testing.T) {
	PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 0
	testInfo.PageInfo.PageSize = 0
	err := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("，0")
	}
	testInfo.Name = ""
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("，name")
	}
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("，")
	}
}
