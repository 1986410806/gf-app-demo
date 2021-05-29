package main

import (
	_ "master_home_api/boot"
	_ "master_home_api/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}