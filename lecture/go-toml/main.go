package main

import (
	"flag"
	"fmt"
	// conf "go-example/lecture/go-add/config"
)

func main() {
	// conf.GetConfig("config/config.toml")

	var port string
	flag.StringVar(&port, "port", "7070", "port to listen on")

	var configFlag = flag.String("config", "./config/config.toml", "toml file to use for configuration")

	// var conf string
	// flag.StringVar(&conf, "config", "./conf.toml", "config file to use")

	pMod := flag.String("mode", "debug", "service mode")
	flag.Parse()
	fmt.Println(port)
	// fmt.Println(conf)
	fmt.Println(*pMod)
	fmt.Println(*configFlag)
	fmt.Println(configFlag)

}