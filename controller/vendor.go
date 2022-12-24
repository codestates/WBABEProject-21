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
	ctx.JSON(http.StatusOK, nil)
}
