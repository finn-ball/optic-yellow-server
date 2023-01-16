package website

import (
	"errors"
	"sync"

	"github.com/chromedp/chromedp"
)

// Login will navigate to the website, fill in the user data and then click.
func (session *Session) Login() error {
	var errorLoad error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err := waitLoaded(session.ctx, 2); err != nil {
			errorLoad = err
		}
		wg.Done()
	}()

	url := "https://members.glasgowclub.org/Connect/mrmLogin.aspx"
	if err := chromedp.Run(session.ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#footer`, chromedp.ByQuery),
	); err != nil {
		return err
	}
	err := chromedp.Run(session.ctx,
		fillText(
			`ctl00_MainContent_InputLogin`,
			session.username,
		),
		fillText(
			`ctl00_MainContent_InputPassword`,
			session.password,
		),
		click(`ctl00_MainContent_btnLogin`),
	)
	if err != nil {
		return err
	}
	wg.Wait()
	if errorLoad != nil {
		return LoginFailed
		// return status.Error(codes.Unauthenticated, "failed login")
	}
	// Check to see if the login failed
	var location string
	chromedp.Run(
		session.ctx,
		chromedp.Location(&location),
	)
	// If the login failed, try to find out why
	if location != "https://members.glasgowclub.org/Connect/memberHomePage.aspx" {
		chromedp.Run(
			session.ctx,
			chromedp.Text(`ctl00_MainContent_errorbox`, &location, chromedp.ByID),
		)
		if len(location) > 0 {
			return errors.New(location)
		} else {
			return Unknown
		}
	}
	return nil
}
