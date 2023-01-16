package website

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

// Booking will attempt to book the court
func (session Session) Booking(date time.Time) error {
	// Get to the bookings page.
	if err := bookingPage(session.ctx); err != nil {
		return BookingPageFailed.wrapError(err)
	}
	// Traverse the first grid.
	btnID, err := traverseSlotsGrid(session.ctx, date)
	if err != nil {
		return TraverseSlotsFailed.wrapError(err)
	}
	// With the button ID, traverse the main court booking grid.
	if err := traverseMainGrid(session.ctx, date, btnID); err != nil {
		return TraverseMainFailed.wrapError(err)
	}
	return nil
}

// bookingPage will navigate to the specified time and date.
// We go through the "Make a booking" option first as it gives a better overview
// of the dates available.
// It will also attempt to take a screenshot of the availability.
func bookingPage(ctx context.Context) error {
	defer takeScreenshot(ctx, "availability_1.png")
	return chromedp.Run(ctx,
		// Get to the list of sites
		chromedp.Navigate("https://members.glasgowclub.org/Connect/mrmselectsite.aspx"),
		fillText("ctl00_MainContent_txtSiteSearch", "Queens Park Tennis"),
		// Click the little dropdown box
		click(`ui-id-3`),
	)
}

// traverseSlotsGrid will iterate through the table to find an available date.
func traverseSlotsGrid(ctx context.Context, date time.Time) (string, error) {
	defer takeScreenshot(ctx, "availability_2.png")
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
		return "", CourtsUnavailable
	}
	if !available {
		return "", DateUnavailable
	}
	return btnID, nil
}

// findSlots will try to return the button ID and if there is availability.
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
				if dt == dateToFind {
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

func traverseMainGrid(ctx context.Context, dateToFind time.Time, btnID string) error {
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
	err = findCourt(nodes, dateToFind, &courtAvailability)
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
	return CourtsUnavailable
}

func findCourt(nodes []*cdp.Node, dateToFind time.Time, availability *[5]string) error {
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
			dt, err := time.Parse("15:04", attrs["value"])
			if err != nil {
				return err
			}
			if dt.Hour() == dateToFind.Hour() {
				x := strings.Split(attrs["data-qa-id"], "Court=Queens Park Tennis Ct ")[1]
				i, err := strconv.Atoi(x)
				if err != nil {
					return err
				}
				availability[i-1] = attrs["name"]
			}
		}
		if node.ChildNodeCount > 0 {
			findCourt(node.Children, dateToFind, availability)
		}
	}
	return nil
}
