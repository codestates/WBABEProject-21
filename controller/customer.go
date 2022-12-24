package controller

import (
	"fmt"
	"net/http"
	md "pkg/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var userInput ReqForm_CreateOrder
	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	var user md.User
	// 유저 테이블을 탐색하고, 해당 닉네임 유저가 존재하면 유저로 아래 로직 진행
	findUser_result, err := md.Muser.FindUserByNickname(userInput.Nickname)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	userId_toString := findUser_result.Id.String()
	userIdCheck := strings.Contains(userId_toString, "000000000000000000000000")

	// True = 유저 없음
	if userIdCheck == true {
		created_user, err := md.Muser.CreateUser(userInput.Nickname)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
		}
		user = created_user
	}
	// // True = 유저가 존재한다면 찾은 유저 삽입
	user = findUser_result

	// Order Table에서 유저 오더기록을 탐색하고,
	var discount int
	userOrderedCount, err := md.Morder.FindOrderCountByUserId(user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	// Order기록이 3회 이상 있다면 단골고객으로 = Discount : 5%;
	if userOrderedCount >= 3 {
		discount = 5
	} else {
		// 만약 아니라면 Discount : 0%;
		discount = 0
	}

	var orderForm md.Order
	orderForm.Menu = userInput.Menu
	orderForm.Status = 1
	orderForm.User = user.Id
	fmt.Println(discount)
	if err := md.Morder.CreateOrder(orderForm, discount); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

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

func GetOrderByUser(ctx *gin.Context) {
	page := ctx.Param("page")
	page_numb, _ := strconv.Atoi(page)
	limit := ctx.Param("limit")
	limit_numb, _ := strconv.Atoi(limit)
	offset := (page_numb - 1) * limit_numb
	userId := ctx.Param("userid")

	ordersResult, err := md.Morder.GetOrderByUser(userId, limit_numb, offset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": ordersResult,
	})
}
