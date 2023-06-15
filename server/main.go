package main

import (
	"question.com/global"
	"question.com/initializes"
)

func main() {
	r := initializes.Routers()
	r.Run(global.CONF.System.Addr)
}
