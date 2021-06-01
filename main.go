package main

import (
	_ "gf-app-demo/boot"
	_ "gf-app-demo/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
