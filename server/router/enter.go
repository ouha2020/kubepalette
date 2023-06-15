package router

import "question.com/router/system"

type RouterGroup struct {
	SystemRouterGroup system.SystemRouter
}

var RouterGroupApp = new(RouterGroup)
