package website

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

type Session struct {
	ctx      context.Context
	username string
	password string
}

func NewSession(username, password string, headless bool, timeout time.Duration) (*Session, func(), error) {
	var ctx context.Context
	var cancel context.CancelFunc
	var cancels []func()
	if headless {
		ctx = context.Background()
	} else {
		options := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag(`headless`, false),
		)
		ctx, cancel = chromedp.NewExecAllocator(context.Background(), options...)
		cancels = append(cancels, cancel)
	}
	ctx, cancel = chromedp.NewContext(ctx)
	cancels = append(cancels, cancel)
	ctx, cancel = context.WithTimeout(ctx, timeout)
	cancels = append(cancels, cancel)
	return &Session{
			ctx:      ctx,
			username: username,
			password: password,
		}, func() {
			for _, c := range cancels {
				c()
			}
		}, nil
}
