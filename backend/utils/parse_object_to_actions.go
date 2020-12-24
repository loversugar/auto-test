package utils

import (
	"auto-test/datamodals"
	"github.com/chromedp/chromedp"
	"strings"
	"time"
)

func ParserObjectsTotActions(actions []chromedp.Action, datas []*datamodals.Data) []chromedp.Action {

	for _, data := range datas {
		if strings.EqualFold(data.OperationType, "SendKey") {
			actions = append(actions, chromedp.SendKeys(data.ElementId, data.Data, chromedp.ByID),
				chromedp.Sleep(3 * time.Second))

		} else if strings.EqualFold(data.OperationType, "Click") {
			actions = append(actions, chromedp.Click(data.ElementId, chromedp.ByID),
				chromedp.Sleep(3 * time.Second))
		}
	}

	return actions
}
