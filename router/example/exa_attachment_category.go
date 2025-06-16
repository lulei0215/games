package example

import (
	"github.com/gin-gonic/gin"
)

type AttachmentCategoryRouter struct{}

func (r *AttachmentCategoryRouter) InitAttachmentCategoryRouterRouter(Router *gin.RouterGroup) {
	router := Router.Group("attachmentCategory")
	{
		router.GET("getCategoryList", attachmentCategoryApi.GetCategoryList) //
		router.POST("addCategory", attachmentCategoryApi.AddCategory)        // /
		router.POST("deleteCategory", attachmentCategoryApi.DeleteCategory)  //
	}
}
