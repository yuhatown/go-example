package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	conf "go-example/lecture/go-loger/config"
	ctl "go-example/lecture/go-loger/controller"
	"go-example/lecture/go-loger/logger"
	"go-example/lecture/go-loger/model"
	rt "go-example/lecture/go-loger/route"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	var configFlag = flag.String("config", "./config/config.toml", "toml file to use for configuration")
	flag.Parse()
	cf := conf.NewConfig(*configFlag)

	// 로그 초기화
	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	mod, _ := model.NewModel()	
	controller, _ := ctl.NewCTL(mod)

	logger.Debug("ready server....")

	if rt, err := rt.NewRouter(controller); err != nil {
		panic(fmt.Errorf("router.NewRouter > %v", err))
	} else {
		mapi := &http.Server{
			Addr:           cf.Server.Port,
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Warn("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			logger.Error("Server Shutdown:", err)
		}

		select {
		case <-ctx.Done():
			logger.Info("timeout of 5 seconds.")
		}

		logger.Info("Server exiting")
	}

	if err := g.Wait(); err != nil {
		logger.Error(err)
	}

}