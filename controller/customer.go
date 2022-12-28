package controller

import (
	"fmt"
	"log"
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
	/*
	고유한 id가 아니라 닉네임으로 검색을 하는 이유는 무엇인가요?
	유저를 생성할 때 이미 존재하는 닉네임인지 체크를 해서 저장한다고 가정을 해봅시다. 
	하지만, 다른 경로를 통해서 닉네임이 변경이 된다면 unique함을 보장받을 수 있을까요? 두개 이상의 결과가 나오면 어떻게 되나요?
	(여기서 다른 경로란 운영상의 이슈로 직접 Database를 수정하는 것과 같은 상황등을 말합니다.)
	*/
	// 유저 테이블을 탐색하고, 해당 닉네임 유저가 존재하면 유저로 아래 로직 진행
	findUser_result, err := md.Muser.FindUserByNickname(userInput.Nickname)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	userId_toString := findUser_result.Id.String()
	/*
	000000의 값을 통해서 비교한 후 True, False를 설정하는 이유는 무엇인가요?
	*/
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

	/* 
	할인율의 경우엔 기본적으로 매우 변동성이 높은 값입니다. 
	5%의 할인율을 유지하다가 10%로 변경하고 싶은 경우, 또는 5회 이상의 주문으로 기준을 바꾸고 싶은 경우가 생긴다면 
	그 상황마다 코드를 수정하고 다시 서비스 배포를 해야하는 번거로움이 존재합니다.
	어떻게 해결할 수 있을까요? 생각해보시면 좋을 것 같습니다. 
	*/
	// Order기록이 3회 이상 있다면 단골고객으로 = Discount : 5%;
	if userOrderedCount >= 3 {
		discount = 5
	} else {
		// 만약 아니라면 Discount : 0%;
		discount = 0
	}

	var orderForm md.Order
	orderForm.Menu = userInput.Menu
	/* 
	다른 곳에서도 코멘트를 드렸지만, Order의 상태값이 int인 경우 사용되는 부분이 많아지면 많아질수록 가독성이 저하됩니다.
	Enum 사용을 고려해보시면 좋겠습니다.
	*/
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

func ChangeMenu(ctx *gin.Context) {
	// 1. Order를 찾고, Status가 1보다 큰지 확인
	var userInput ReqForm_ChangeMenu

	err := ctx.ShouldBind(&userInput)
	if err != nil {
		log.Fatal(err)
	}

	result_recent_order, err := md.Morder.GetOrderByUser(userInput.UserId, 1, 1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	fmt.Println("Get Order : ", result_recent_order)
	if len(result_recent_order) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "No Order Exist",
		})
		return
	}
	/*
	대소 비교를 하는 것보다 변경할 수 없는 상태 값들을 모아둔 리스트를 만들고, result_recent_order[0].Status 값이
	해당 리스트 안에 들어있는지 여부를 판단하는 것은 어떨까요?
	*/
	// Status가 1보다 크다면 이미 진행중인 상태이므로 변경 불가.
	if result_recent_order[0].Status > 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "Food Being Cooked",
		})
		return
	}

	/*
	아니면 이라는 주석을 달기보다는 else로 처리하고 주석을 제거하는 것은 어떨까요?
	*/
	// 아니면, 변경 성공
	err = md.Morder.UpdateOrder(result_recent_order[0], userInput.NewMenu)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
