package server

import (
	"context"

	"github.com/finn-ball/optic-yellow-server/pkg/proto"
	"github.com/finn-ball/optic-yellow-server/pkg/website"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *grpcServer) list(ctx context.Context, req *proto.RunRequest_List) (*proto.RunResponse, error) {
	username := req.List.Username
	password := req.List.Password
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
	if err = session.Login(); err != nil {
		return &proto.RunResponse{}, err
	}
	bookings, ok := s.listBookings(username)
	if !ok {
		return &proto.RunResponse{
			Booking: bookings,
		}, UsernameNotFound
	}
	return &proto.RunResponse{
		Booking: bookings,
	}, nil
}

// listBookings returns false if the user is not found.
func (s *grpcServer) listBookings(username string) ([]*proto.BookingResponse, bool) {
	s.bookings.Lock()
	defer s.bookings.Unlock()
	if entry, ok := s.bookings.users[username]; ok {
		bookings := []*proto.BookingResponse{}
		for key, val := range entry.status {
			bookings = append(bookings, &proto.BookingResponse{
				Status:   val,
				Datetime: timestamppb.New(key),
			})
		}
		return bookings, true
	}
	return nil, false
}
