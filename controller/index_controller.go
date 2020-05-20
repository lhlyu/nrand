package controller

import (
	"github.com/lhlyu/nrand/controller/dto"
	"github.com/lhlyu/nrand/result"
	"github.com/lhlyu/nrand/service"
)

type IndexController struct {
}

func (c *IndexController) GenName(ctx *Context) {
	req := &dto.NameDto{}
	if !ctx.GetParams(req, false) {
		return
	}
	svc := service.NewIndexService(ctx.GetTracker())
	datas := svc.GenName(req)
	if len(datas) == 0 {
		ctx.JSON(result.Failure)
		return
	}
	ctx.JSON(result.Success.WithData(datas))
}
