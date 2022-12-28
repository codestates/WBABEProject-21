package controller

import (
	"net/http"
	md "pkg/model"

	"github.com/gin-gonic/gin"
)

func CreateMenu(ctx *gin.Context) {
	var Input md.Menu

	if err := ctx.ShouldBind(&Input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err := md.Mmenu.CreateMenu(Input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	/*
	리소스 생성의 경우에는 201 created를 return하는 것이 일반적입니다.
	*/
	ctx.JSON(http.StatusOK, nil)

}

func UpdateMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	var Input md.Menu
	if err := ctx.BindJSON(&Input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err := md.Mmenu.UpdateMenu(Input, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)

}

func DeleteMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	err := md.Mmenu.DeleteMenu(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	/*
	삭제 후에는 일반적으로 204 No Content를 return 합니다.
	*/
	ctx.JSON(http.StatusOK, nil)
}

func UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var Input md.Order
	if err := ctx.BindJSON(&Input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err := md.Morder.UpdateStatus(Input, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
