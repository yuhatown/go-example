package route

import (
	"fmt"

	ctl "go-example/lecture/go-loger/controller"
	"go-example/lecture/go-loger/logger"

	"github.com/gin-gonic/gin"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl}

	return r, nil
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

func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort()
			return
		}
		auth := c.GetHeader("Authorization")
		fmt.Println("Authorization-word ", auth)

		c.Next()
	}
}

func (p *Router) Idx() *gin.Engine {
	// 컨피그나 상황에 맞게 gin 모드 설정
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)

	e := gin.Default()
	// e.Use(gin.Logger())
	// e.Use(gin.Recovery())
  // 기존의 logger, recovery 대신 logger에서 선언한 미들웨어 사용
	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	e.Use(CORS())

	logger.Info("start server")
	e.GET("/health")

	account := e.Group("acc/v01", liteAuth())
	{
		fmt.Println(account)
	}

	return e
}