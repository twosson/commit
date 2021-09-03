package main

import (
	_ "commit/boot"
	_ "commit/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
