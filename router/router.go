package router

import (
	ctl "pkg/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	router_vendor := r.Group("/api/vendor")
	{
		router_vendor.POST("/menu", ctl.CreateMenu)
		router_vendor.DELETE("/menu/:id", ctl.DeleteMenu)
		router_vendor.PUT("/update/menu/:id", ctl.UpdateMenu)
		router_vendor.PUT("/update/status/:id", ctl.UpdateStatus) // 오더 상태 업데이트
	}
	router_customer := r.Group("/api/customer")
	{
		router_customer.GET("/menu/:page/:limit", ctl.GetMenus)
		router_customer.POST("/order", ctl.CreateOrder) // 오더생성
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
