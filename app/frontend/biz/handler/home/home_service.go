package home

import (
	"context"

	"github.com/ZHLOVEYY/gomall/app/frontend/biz/service"
	"github.com/ZHLOVEYY/gomall/app/frontend/biz/utils"
	home "github.com/ZHLOVEYY/gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//resp := &home.Empty{}
	//导致赋值赋不上
	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
	c.HTML(consts.StatusOK, "home.tmpl", resp)
	//Hertz渲染html的方法
}
