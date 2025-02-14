package service

import (
	"context"

	home "github.com/ZHLOVEYY/gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

//	func (h *HomeService) Run(req *home.Empty) (resp *home.Empty, err error) {
//		//defer func() {
//		// hlog.CtxInfof(h.Context, "req = %+v", req)
//		// hlog.CtxInfof(h.Context, "resp = %+v", resp)
//		//}()
//		// todo edit your code
//		return
//	}

// 这部分会传递给home_service.go（利用run函数传递resp） 然后传给main.go，然后调用渲染到前端（template的使用！）
// 利用range，end 遍历
func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-2", "Price": 110, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-3", "Price": 120, "Picture": "/static/image/t-shirt-2.jpeg"},
		{"Name": "T-shirt-4", "Price": 130, "Picture": "/static/image/notebook.jpeg"},
		{"Name": "T-shirt-5", "Price": 140, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-6", "Price": 150, "Picture": "/static/image/t-shirt.jpeg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
