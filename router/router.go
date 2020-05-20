package router

import (
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/nrand/controller"
)

func SetRouter(app *iris.Application) {
	ctr := &controller.IndexController{}

	app.Post("/api/name", controller.H(ctr.GenName))

}
