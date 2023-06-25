package main

import (
	"question.com/global"
	"question.com/initializes"
)

func main() {
	r := initializes.Routers()
	initializes.Viper()
	initializes.K8s()
	panic(r.Run(global.CONF.System.Addr))
}
