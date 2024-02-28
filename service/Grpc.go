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
	timeout := 10 * time.Second

	page := RodPool.Get(RodCreate)
	defer RodPool.Put(page)

	//page.MustNavigate(htmlInfo.Url).MustWaitLoad()
	err := page.Timeout(timeout).MustNavigate(url).MustWaitLoad().WaitStable(time.Second)
	if err != nil {
		fmt.Println("采集页面超时:"+url, "err:", err.Error())
		page.Timeout(timeout).MustNavigate(url).MustWaitLoad().WaitStable(time.Second)
	}
	//加载完成且页面稳定

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
