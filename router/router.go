package router

import (
	ctl "pkg/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	/*
	1. vender, customer로 네이밍을 분리하신 부분 좋습니다. 직관적이네요!

	2. 메뉴 하나에 오더가 여러개인 1:N 구조의 경우라면, 다음과 같이 nested된 URI로 구성이 가능합니다.
		/menu/1/order/{order_id}
		/menu/1/order/{order_id}
		/menu/2/order/{order_id}
	*/
	router_vendor := r.Group("/api/vendor")
	{
		router_vendor.POST("/menu", ctl.CreateMenu)               // 메뉴 생성
		router_vendor.DELETE("/menu/:id", ctl.DeleteMenu)         // 메뉴 삭제
		/*
		update라는 단어는 들어가지 않아도 충분해 보입니다.
		*/
		router_vendor.PUT("/update/menu/:id", ctl.UpdateMenu)     // 메뉴 업데이트
		router_vendor.PUT("/update/status/:id", ctl.UpdateStatus) // 오더 상태 업데이트
	}
	router_customer := r.Group("/api/customer")
	{
		router_customer.GET("/menu/:page/:limit", ctl.GetMenus)                // 메뉴조회
		router_customer.POST("/order", ctl.CreateOrder)                        // 오더생성
		router_customer.GET("/order/:page/:limit/:userid", ctl.GetOrderByUser) // Order 기록 찾기 By User
		/*
		order의 id 값을 이용해 업데이트를 할 수 있어야 REST API를 충족시킬 수 있습니다.
		메뉴의 id 값으로 업데이트를 하는 부분과 동일하게 구현하시면 될 것 같습니다.
		*/
		router_customer.PUT("/order", ctl.ChangeMenu)                          // Order 업데이트
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
