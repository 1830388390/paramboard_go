package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	api "paramboard_go/api/v1"
	"paramboard_go/utilss/middleware"
)

// 路由配置
func Router() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	r.StaticFS("./s", http.Dir("./static"))
	v1 := r.Group("")
	version := "v1"
	{

		v1.GET("ping/", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		//用户操作
		v1.GET("author/addAuthor/"+version, api.AddAuthor)
		//param
		v1.GET("aP/", api.AddParam)
		v1.GET("gP/", api.GetParam)
		v1.GET("dP/", api.DelParam)
	}
	return r
}
