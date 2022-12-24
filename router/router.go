package router

import (
	ctl "pkg/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	router_vendor := r.Group("/api/vendor")
	{
		router_vendor.POST("/menu", ctl.CreateMenu)               // 메뉴 생성
		router_vendor.DELETE("/menu/:id", ctl.DeleteMenu)         // 메뉴 삭제
		router_vendor.PUT("/update/menu/:id", ctl.UpdateMenu)     // 메뉴 업데이트
		router_vendor.PUT("/update/status/:id", ctl.UpdateStatus) // 오더 상태 업데이트
	}
	router_customer := r.Group("/api/customer")
	{
		router_customer.GET("/menu/:page/:limit", ctl.GetMenus)                // 메뉴조회
		router_customer.POST("/order", ctl.CreateOrder)                        // 오더생성
		router_customer.GET("/order/:page/:limit/:userid", ctl.GetOrderByUser) // Order 기록 찾기 By User
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
