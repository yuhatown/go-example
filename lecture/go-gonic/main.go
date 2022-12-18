package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

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

func index() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORS())

	route := r.Group("/route/v01", liteAuth())
	{
		route.POST("/post", post)
		route.PUT("/put", put)
		route.GET("/get", get)
	}
	return r
}

func post(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
	return
}

func put(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
	return
}


func get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
		"info": "data",
	})
}

func main() {
	mapi := &http.Server {
		Addr: ":8080",
		Handler: index(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,			
		MaxHeaderBytes: 1 << 20,
	}

	g.Go(func() error {
		return mapi.ListenAndServe()
	})

	stopSig := make(chan os.Signal)
	signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM)
	<- stopSig
	log.Println("Shudown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mapi.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	select {
	case <- ctx.Done():
		fmt.Println("timeout 5 seconds.")
	}
	fmt.Println("Server stop")

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}