package service

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"runtime"
)

var RodBrowser *rod.Browser
var RodPool rod.PagePool
var RodCreate func() *rod.Page

func RodStep() {

	u := launcher.New().
		Headless(true).
		NoSandbox(true).
		MustLaunch()

	RodBrowser = rod.New().ControlURL(u).MustConnect()

	//defer RodBrowser.MustClose()

	//页面池可以使用页面池来辅助同时控制和复用多个页面。
	numCPU := runtime.NumCPU()

	if numCPU == 1 {
		numCPU = 2
	} else {
		numCPU = 4
	}

	RodPool = rod.NewPagePool(numCPU) //同时最多可以打开
	// Create a page if needed
	RodCreate = func() *rod.Page {
		// 使用MustIncognito来隔离页面
		return RodBrowser.MustIncognito().MustPage()
	}

	fmt.Println("rod 初始化完成")
}
