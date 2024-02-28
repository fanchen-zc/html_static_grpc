package Router

import (
	"github.com/gin-gonic/gin"
	"html_static_grpc/Controllers"
	"html_static_grpc/Middlewares"
	"html_static_grpc/config"
	"html_static_grpc/pkg/e"
	"net/http"
	_ "net/http/pprof"
)

func InitRouter() {
	router := gin.Default()
	//要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	router.GET("/test", Controllers.Test)

	router.NoRoute(func(c *gin.Context) {
		res := e.Gin{C: c}
		//返回404状态码
		res.Res(http.StatusNotFound, 404, 1, "page not exists!", "")
	})

	router.Run(":" + config.Configs.HttpListenPort) //

}
