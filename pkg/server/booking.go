package server

import (
	"context"
	"math"
	"time"

	"github.com/finn-ball/optic-yellow-server/pkg/proto"
	"github.com/finn-ball/optic-yellow-server/pkg/website"
)

func (s *grpcServer) booking(ctx context.Context, req *proto.RunRequest_Booking) (*proto.RunResponse, error) {
	username := req.Booking.Login.Username
	password := req.Booking.Login.Password
	session, cancel, err := website.NewSession(
		username,
		password,
		s.headless,
		s.timeout,
	)
	defer cancel()
	if err != nil {
		return &proto.RunResponse{}, err
	}
	if err := session.Login(); err != nil {
		return &proto.RunResponse{}, err
	}
	s.updatePassword(username, password)
	timeToBook := req.Booking.Datetime.AsTime()
	if isValidNow(time.Now(), timeToBook) {
		if err := session.Booking(timeToBook); err != nil {
			return &proto.RunResponse{}, err
		} else {
			return &proto.RunResponse{
				Booking: []*proto.BookingResponse{
					{
						Status:   proto.BookingResponse_SUCCESSFUL,
						Datetime: req.Booking.Datetime,
					},
				},
			}, nil
		}
	} else {
		s.queueBooking(username, req.Booking.Datetime.AsTime())
		return &proto.RunResponse{
			Booking: []*proto.BookingResponse{
				{
					Status:   proto.BookingResponse_PENDING,
					Datetime: req.Booking.Datetime,
				},
			},
		}, nil
	}
}

func (s *grpcServer) queueBooking(username string, timeToBook time.Time) error {
	s.bookings.Lock()
	defer s.bookings.Unlock()
	// program should have already tried a login so user should exist
	user, ok := s.bookings.users[username]
	if !ok {
		return UsernameNotFound
	}
	if _, ok := user.queue[timeToBook]; ok {
		return BookingAlreadyExists
	} else {
		ctx, cancel := context.WithCancel(context.Background())
		user.queue[timeToBook] = cancel
		go s.wait(ctx, username, timeToBook)
	}
	return nil
}

func (s *grpcServer) wait(ctx context.Context, username string, timeToBook time.Time) {
	triggerBooking := make(chan interface{}, 1)
	launchBooking := time.Date(
		timeToBook.Year(),
		timeToBook.Month(),
		timeToBook.Day(),
		0, 0, 0, 0, time.UTC,
	)
	// Bookings come online 3 days in advance
	launchBooking = launchBooking.Add(-3 * 24 * time.Hour)
	// Can login 10s before midnight to be a bit quicker
	launchLogin := launchBooking.Add(-10 * time.Second)
	// login
	go func() {
		<-time.After(time.Until(launchLogin))
		triggerBooking <- true
	}()
	// book
	go func() {
		<-time.After(time.Until(launchBooking))
		triggerBooking <- true
	}()
	select {
	// cancel case
	case <-ctx.Done():
		return
	case <-triggerBooking:
		// We should have already logged in as the user so entry must exist
		password, _ := s.getPassword(username)
		session, cancel, err := website.NewSession(
			username,
			password,
			s.headless,
			s.timeout,
		)
		defer cancel()
		if err != nil {
			s.updateBookingState(
				username,
				timeToBook,
				proto.BookingResponse_FAILED,
			)
			return
		}
		if err = session.Login(); err != nil {
			s.updateBookingState(
				username,
				timeToBook,
				proto.BookingResponse_FAILED,
			)
			return
		}
		<-triggerBooking
		if err = session.Booking(timeToBook); err != nil {
			s.updateBookingState(
				username,
				timeToBook,
				proto.BookingResponse_FAILED,
			)
			return
		}
	}
}

// isValidNow checks to see if it is bookable time frame now.
// It has to be an hour after the current hour and within three days.
func isValidNow(now, check time.Time) bool {
	if check.Before(now) || check == now {
		return false
	} else if h := check.Sub(now).Hours(); h < 1.0 {
		return false
	}
	days := int(math.Ceil(check.Sub(now).Hours() / 24))
	return days <= 3
}
