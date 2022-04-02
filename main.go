package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

var username, password, date string
var hour uint
var timeout = 20 * time.Second
var timeToBook time.Time

func main() {
	cmd()

	if err := run(username, password, timeToBook); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("Timeout.")
			log.Println("Timeout currently set to", timeout)
			log.Fatal(err)
		}
		log.Fatal(err)
	}
}

// cmd will parse the user input.
func cmd() {
	flag.StringVar(&username, "u", "", `Username`)
	flag.StringVar(&password, "p", "", `Password`)
	flag.StringVar(&date, "d", "", `Date in the format DD/MM/YY`)
	flag.UintVar(&hour, "h", 0, `Hour in the format HH (e.g "13" for 1pm)`)
	flag.Parse()
	if !isFlagPassed("u") {
		log.Fatal("Username not defined.")
	}
	if !isFlagPassed("p") {
		log.Fatal("Password not defined.")
	}
	if !isFlagPassed("h") {
		log.Fatal("Hour not defined.")
	}
	if hour < 9 || hour > 20 {
		log.Fatal("Hour needs to be in the range of 9 - 20")
	}
	if !isFlagPassed("d") {
		log.Fatal("Date not defined.")
	}
	// Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)
	_, err := time.Parse("02/01/06", date)
	if err != nil {
		log.Println("Could not parse date.")
		log.Println("Format: DD/MM/YY")
		log.Println("Today is: ", time.Now().Format("02/01/06"))
		log.Fatal(err)
	}
	timeToBook, err = time.Parse("02/01/06 15:04:05", fmt.Sprintf("%s %d:00:00", date, hour))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Username: ", username)
	log.Println("To book: ", timeToBook.Format("02/01/06 15:04:05"))
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	if !found {
		flag.PrintDefaults()
	}
	return found
}

// run the program.
// This will cause the bot to login, navigate to the booking page, book a court
// and finally screenshot the result.
func run(user, password string, date time.Time) error {
	// Set up the various contexts.
	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag(`headless`, false),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()
	// Login to te website.
	if err := login(ctx, user, password); err != nil {
		return err
	}
	// Get to the bookings page.
	if err := bookingPage(ctx); err != nil {
		return err
	}
	// Traverse the first grid.
	btnID, err := traverseSlotsGrid(ctx, date)
	if err != nil {
		return err
	}
	// With the button ID, traverse the main court booking grid.
	if err := traverseMainGrid(ctx, btnID); err != nil {
		return err
	}
	return nil
}

// login will navigate to the website, fill in the user data and then click.
func login(ctx context.Context, user string, password string) error {
	url := "https://members.glasgowclub.org/Connect/mrmLogin.aspx"
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#footer`, chromedp.ByQuery),
	); err != nil {
		return err
	}
	return chromedp.Run(ctx,
		fillText(
			url,
			`ctl00_MainContent_InputLogin`,
			user,
		),
		fillText(
			url,
			`ctl00_MainContent_InputPassword`,
			password,
		),
		click(`ctl00_MainContent_btnLogin`),
	)
}

// bookingPage will navigate to the specified time and date.
// We go through the "Make a booking" option first as it gives a better overview
// of the dates available.
// It will also attempt to take a screenshot of the availability.
func bookingPage(ctx context.Context) error {
	defer takeScreenshot(ctx, "availability.png")
	return chromedp.Run(ctx,
		// Bookings drop down box.
		click("ctl00_ctl12_Li3"),
		// Make a booking.
		click("ctl00_ctl12_MakeBookingli"),
		// Glasgow club in the parks.
		click("ctl00_MainContent_sitesGrid_ctrl8_lnkListCommand"),
		// Tennis court bookings
		click("ctl00_MainContent_activityGroupsGrid_ctrl0_lnkListCommand"),
		// Queen's park tennis
		click("ctl00_MainContent_activitiesGrid_ctrl2_lnkListCommand"),
	)
}

// traverseSlotsGrid will iterate through the table to find an available date.
func traverseSlotsGrid(ctx context.Context, date time.Time) (string, error) {
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		// Ensure whole table is visible.
		chromedp.WaitVisible("slotsGrid", chromedp.ByID),
		// Get the nodes of the table.
		chromedp.Nodes("slotsGrid", &nodes, chromedp.ByID),
		// Ask chromedp to populate the subtree of a node,
		chromedp.ActionFunc(func(c context.Context) error {
			// Depth -1 for the entire subtree
			return dom.RequestChildNodes(nodes[0].NodeID).WithDepth(-1).Do(c)
		}),
		// Wait a little while for dom.EventSetChildNodes to be fired and handled.
		chromedp.Sleep(time.Second),
	)
	if err != nil {
		return "", err
	}
	btnID, available, err := findSlots(nodes, date)
	if err != nil {
		return "", err
	}
	if btnID == "" {
		return "", errors.New("could not find date")
	}
	if !available {
		return "", errors.New("date not available")
	}
	return btnID, nil
}

// findSlots will try to return the button ID and if there is availabilty.
func findSlots(nodes []*cdp.Node, dateToFind time.Time) (string, bool, error) {
	for _, node := range nodes {
		if node.NodeName == "INPUT" {
			var key string
			attrs := make(map[string]string)
			for i, attr := range node.Attributes {
				// Data is in the form:
				// [key, val, key, val]
				if i%2 == 0 {
					key = attr
				} else {
					attrs[key] = attr
				}
			}
			data, ok := attrs["data-qa-id"]
			if !ok {
				return "", false, errors.New("no data field")
			}
			x := strings.Split(data, "Date=")
			if len(x) > 1 {
				y := strings.Split(x[1], " Availability=")
				// Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)
				dt, err := time.Parse("02/01/2006 15:04:05", y[0])
				if err != nil {
					return "", false, err
				}
				if dt == timeToBook {
					availability := attrs["value"] == "Available"
					return attrs["id"], availability, nil
				}
			}
		}
		if node.ChildNodeCount > 0 {
			btnID, available, err := findSlots(node.Children, dateToFind)
			if err != nil {
				return "", false, err
			}
			if btnID != "" {
				return btnID, available, nil
			}
		}
	}
	return "", false, nil
}

func traverseMainGrid(ctx context.Context, btnID string) error {
	defer takeScreenshot(ctx, "bookingConfirmation.png")
	var nodes []*cdp.Node
	// Click on the button to book.
	err := chromedp.Run(ctx,
		// Click on the booking date.
		click(btnID),
		// Ensure whole table is visible.
		chromedp.WaitVisible("ctl00_MainContent_grdResourceView", chromedp.ByID),
		// Get the nodes of the table.
		chromedp.Nodes("ctl00_MainContent_grdResourceView", &nodes, chromedp.ByID),
		// Ask chromedp to populate the subtree of a node,
		chromedp.ActionFunc(func(c context.Context) error {
			// Depth -1 for the entire subtree
			return dom.RequestChildNodes(nodes[0].NodeID).WithDepth(-1).Do(c)
		}),
		// Wait a little while for dom.EventSetChildNodes to be fired and handled.
		chromedp.Sleep(time.Second),
	)
	if err != nil {
		return err
	}
	courtAvailability := [5]string{"", "", "", "", ""}
	err = findCourt(nodes, timeToBook, &courtAvailability)
	if err != nil {
		return err
	}
	// Order in which to search for courts.
	// This is just my personal preference.
	courtPreferenceOrder := [5]int{5, 4, 3, 1, 2}
	for _, p := range courtPreferenceOrder {
		// Arrays start at 0, not 1, hence the offset.
		btn := courtAvailability[p-1]
		if btn != "" {
			return chromedp.Run(ctx,
				// No ID is given so we search for the CSS button instead.
				chromedp.Click(btn, chromedp.BySearch),
				// Finally... Book!
				click("ctl00_MainContent_btnBasket"),
			)
		}
	}
	return nil
}

func findCourt(nodes []*cdp.Node, dateToFind time.Time, availibility *[5]string) error {
	for _, node := range nodes {
		if node.NodeName == "INPUT" {
			var key string
			attrs := make(map[string]string)
			for i, attr := range node.Attributes {
				// Data is in the form:
				// [key, val, key, val]
				if i%2 == 0 {
					key = attr
				} else {
					attrs[key] = attr
				}
			}
			// Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)
			t := fmt.Sprintf("%s %s", date, attrs["value"])
			dt, err := time.Parse("02/01/06 15:04", t)
			if err != nil {
				return err
			}
			if dt == timeToBook {
				x := strings.Split(attrs["data-qa-id"], "Court=Queens Park Tennis Ct ")[1]
				i, err := strconv.Atoi(x)
				if err != nil {
					return err
				}
				availibility[i-1] = attrs["name"]
			}
		}
		if node.ChildNodeCount > 0 {
			findCourt(node.Children, dateToFind, availibility)
		}
	}
	return nil
}

func fillText(urlstr, id, q string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitReady(id, chromedp.ByID),
		chromedp.SendKeys(id, q, chromedp.ByID),
	}
}

func click(id string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitReady(id, chromedp.ByID),
		chromedp.Click(id, chromedp.ByID),
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
	return ioutil.WriteFile(filename, img, 0644)
}

func fullScreenshot(quality int, result *[]byte) chromedp.Action {
	return chromedp.Tasks{
		chromedp.FullScreenshot(result, quality),
	}
}
