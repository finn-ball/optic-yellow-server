package website

import (
	"context"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func click(id string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitReady(id, chromedp.ByID),
		chromedp.Click(id, chromedp.ByID),
	}
}

func fillText(id, q string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitReady(id, chromedp.ByID),
		chromedp.SendKeys(id, q, chromedp.ByID),
	}
}

func takeScreenshot(ctx context.Context, filename string) error {
	var img []byte
	err := chromedp.Run(ctx,
		chromedp.WaitVisible(`#footer`, chromedp.ByQuery),
		fullScreenshot(90, &img),
	)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, img, 0644)
}

func fullScreenshot(quality int, result *[]byte) chromedp.Action {
	return chromedp.Tasks{
		chromedp.FullScreenshot(result, quality),
	}
}

// waitLoaded blocks until a target receives a Page.loadEventFired.
func waitLoaded(ctx context.Context, nLoads int) error {
	ch := make(chan struct{})
	cctx, cancel := context.WithCancel(ctx)
	n := 0
	chromedp.ListenTarget(cctx, func(ev interface{}) {
		switch e := ev.(type) {
		case *page.EventLifecycleEvent:
			if e.Name == `networkAlmostIdle` {
				n += 1
				if n == nLoads {
					cancel()
					close(ch)
				}
			}
		}
	})

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
