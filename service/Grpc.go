package service

import (
	"context"
	"fmt"
	"html_static_grpc/service/html_static"
	"time"
)

type HtmlServiceServer struct{}

func (s *HtmlServiceServer) GetHTMLContent(ctx context.Context, req *html_static.HTMLRequest) (*html_static.HTMLResponse, error) {
	url := req.Url

	page := RodPool.Get(RodCreate)
	defer RodPool.Put(page)
	//加载完成且页面稳定
	err := page.MustNavigate(url).MustWaitLoad().WaitStable(time.Second * 2)

	if err != nil {
		fmt.Println("采集页面超时:"+url, "err:", err.Error())
	}

	// 获取页面的 HTML 内容
	htmlContent, err := page.HTML()
	if err != nil {
		return nil, err
	}
	fmt.Printf("采集完成:%s - %d\n", url, len(htmlContent))

	return &html_static.HTMLResponse{
		Content: htmlContent,
	}, nil
}
