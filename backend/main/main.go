package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx,
		chromedp.WithLogf(log.Printf),)
	defer cancel()

	// 控制整个程序运行时间
	ctx, cancel = context.WithTimeout(ctx, 30 * time.Second)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate("http://localhost:3000"),
		chromedp.Sleep(time.Duration(3) * time.Second),
		chromedp.SendKeys(`username`, "username", chromedp.ByID),
		chromedp.Sleep(3 * time.Second),
		chromedp.SendKeys(`password`, "password", chromedp.ByID),
		chromedp.Sleep(time.Duration(3) * time.Second),); err!= nil {
		fmt.Println(err)
	}

	//if err := chromedp.Run(ctx,
	//	//打开网页
	//	chromedp.Navigate(`http://www.baidu.com`),
	//	//定位登录按钮
	//	chromedp.Click(`document.querySelector(".s-top-login-btn.c-btn.c-btn-primary.c-btn-mini.lb")`, chromedp.ByJSPath),
	//	// 添加卡顿时间
	//	chromedp.Sleep(time.Duration(3) * time.Second),
	//	//使用用户名密码登录
	//	chromedp.Click(`TANGRAM__PSP_11__footerULoginBtn`, chromedp.ByID),
	//	chromedp.Sleep(time.Duration(3) * time.Second),
	//	//输入用户名
	//	chromedp.SendKeys(`TANGRAM__PSP_11__userName`, "username", chromedp.ByID),
	//	chromedp.Sleep(time.Duration(3) * time.Second),
	//	//输入密码
	//	chromedp.SendKeys(`TANGRAM__PSP_11__password`, "password", chromedp.ByID),
	//	chromedp.Sleep(time.Duration(3) * time.Second),
	//	//点击登录按钮
	//	chromedp.Click(`TANGRAM__PSP_11__submit`, chromedp.ByID),
	//	//截全屏
	//	chromedp.ActionFunc(func(ctx context.Context) error {
	//		time.Sleep(time.Second)
	//		var buf []byte
	//		_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
	//		if err != nil {
	//			return err
	//		}
	//
	//		width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))
	//
	//		// force viewport emulation
	//		err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
	//			WithScreenOrientation(&emulation.ScreenOrientation{
	//				Type:  emulation.OrientationTypePortraitPrimary,
	//				Angle: 0,
	//			}).
	//			Do(ctx)
	//		if err != nil {
	//			return err
	//		}
	//
	//		// capture screenshot
	//		buf, err = page.CaptureScreenshot().
	//			WithQuality(90).
	//			WithClip(&page.Viewport{
	//				X:      contentSize.X,
	//				Y:      contentSize.Y,
	//				Width:  contentSize.Width,
	//				Height: contentSize.Height,
	//				Scale:  1,
	//			}).Do(ctx)
	//		if err != nil {
	//			return err
	//		}
	//		if err := ioutil.WriteFile("4.png", buf, 0644); err != nil {
	//			return err
	//		}
	//		return nil
	//	}),
	//); err != nil {
	//	fmt.Println(err)
	//}

}
