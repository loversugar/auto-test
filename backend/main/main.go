package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"math"
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
	//ctx, cancel = context.WithTimeout(ctx, 50 * time.Second)
	//defer cancel()
	//
	//datas, err := utils.ReadFileToJson("/home/C5311429/go/src/auto-test/record.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//var actions []chromedp.Action
	//actions = append(actions,
	//	chromedp.Navigate("http://localhost:3000"),
	//	chromedp.Sleep(time.Duration(3) * time.Second))
	//
	//actions = utils.ParserObjectsTotActions(actions, datas)
	//
	//if err := chromedp.Run(ctx, actions...); err!= nil {
	//	fmt.Println(err)
	//}

	if err := chromedp.Run(ctx,
		chromedp.Navigate(`http://www.baidu.com`),
		chromedp.Click(`document.querySelector(".s-top-login-btn.c-btn.c-btn-primary.c-btn-mini.lb")`, chromedp.ByJSPath),
		chromedp.Sleep(time.Duration(3) * time.Second),
		chromedp.Click(`TANGRAM__PSP_11__footerULoginBtn`, chromedp.ByID),
		chromedp.Sleep(time.Duration(3) * time.Second),
		chromedp.SendKeys(`TANGRAM__PSP_11__userName`, "username", chromedp.ByID),
		chromedp.Sleep(time.Duration(3) * time.Second),
		chromedp.SendKeys(`TANGRAM__PSP_11__password`, "password", chromedp.ByID),
		chromedp.Sleep(time.Duration(3) * time.Second),
		chromedp.Click(`TANGRAM__PSP_11__submit`, chromedp.ByID),
		chromedp.ActionFunc(func(ctx context.Context) error {
			time.Sleep(time.Second)
			var buf []byte
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			buf, err = page.CaptureScreenshot().
				WithQuality(90).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			if err := ioutil.WriteFile("4.png", buf, 0644); err != nil {
				return err
			}
			return nil
		}),
	); err != nil {
		fmt.Println(err)
	}

}
