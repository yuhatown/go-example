package router

import (
	"fmt"

	ctl "lecture/go-mvc/controller"

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
	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(CORS())

	account := e.Group("person/v01", liteAuth())
	{
		account.GET("/name", p.ct.GetName)
		account.GET("/pnum", p.ct.GetPnum)
		account.POST("/", p.ct.Post)
		account.DELETE("/pnum", p.ct.DeletePnum)
		account.PUT("/pnum", p.ct.DeletePnum)
	}
	return e
}