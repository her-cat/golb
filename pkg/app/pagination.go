package app

import (
	"github.com/gin-gonic/gin"
	"github.com/her-cat/golb/global"
	"github.com/her-cat/golb/pkg/convert"
)

func GetPage(ctx *gin.Context) int {
	page := convert.StrTo(ctx.Query("page")).MustInt()
	if page < 0 {
		return 1
	}
	return page
}

func GetPageSize(ctx *gin.Context) int {
	pageSize := convert.StrTo(ctx.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.DefaultPageSize {
		return global.AppSetting.DefaultPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
