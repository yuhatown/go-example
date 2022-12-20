package main

import (
	conf "go-example/lecture/go-add/config"
)

func main() {
	conf.GetConfig("config/config.toml")
}