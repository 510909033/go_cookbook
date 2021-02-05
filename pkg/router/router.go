package router

import (
	"baotian0506.com/go_cookbook/pkg/api"
	"baotian0506.com/go_cookbook/pkg/api/cook/xiangha"
	"github.com/gin-gonic/gin"
	//. "go_cookbook/pkg/api"
)

/*
InnitRouter 路由初始化
*/
func InnitRouter() *gin.Engine {
	//ctx.Writer.Header().Add("Access-Control-Allow-Origin", host)
	//ctx.Writer.Header().Add("Access-Control-Allow-Credentials", "true")

	router := gin.Default()
	router.GET("/users", api.Users)
	router.GET("/menu/add_menu", api.AddMenu)
	router.GET("/get_user", api.GetUser)
	router.GET("/cook/xiangha/init_table", xiangha.InitTable)
	router.GET("/cook/xiangha/update_sitemap", xiangha.UpdateSitemap)
	return router
}
