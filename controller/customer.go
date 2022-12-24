package controller

import (
	"net/http"
	md "pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMenus(ctx *gin.Context) {
	page := ctx.Param("page")
	page_numb, _ := strconv.Atoi(page)
	limit := ctx.Param("limit")
	limit_numb, _ := strconv.Atoi(limit)
	offset := (page_numb - 1) * limit_numb
	curResult, err := md.Mmenu.GetMenuPaging(limit_numb, offset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"payload": curResult,
	})
}
