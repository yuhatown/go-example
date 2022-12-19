package main

import (
	"fmt"
	ctl "go-example/lecture/go-swag/controller"
	"go-example/lecture/go-swag/model"
	rt "go-example/lecture/go-swag/router"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	if mod, err := model.NewModel(); err != nil {
		panic(err)
	} else if controller, err := ctl.NewCTL(mod); err != nil {
		panic(err)
	} else if rt, err := rt.NewRouter(controller); err != nil {
		panic(err)
	} else {
		mapi := &http.Server {
			Addr: ":8080",
			Handler: rt.Idx(),
			ReadTimeout: 5 * time.Second,
			WriteTimeout: 10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}