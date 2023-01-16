package server

import (
	"context"

	"github.com/finn-ball/optic-yellow-server/pkg/proto"
	"github.com/finn-ball/optic-yellow-server/pkg/website"
)

func (s *grpcServer) cancel(ctx context.Context, req *proto.RunRequest_Cancel) (*proto.RunResponse, error) {
	username := req.Cancel.Login.Username
	password := req.Cancel.Login.Password
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
	s.bookings.Lock()
	defer s.bookings.Unlock()
	// At the moment there is no guarantee that the cancellation was successful.
	if entry, ok := s.bookings.users[username]; ok {
		if c, ok := entry.queue[req.Cancel.Datetime.AsTime()]; ok {
			c()
			return &proto.RunResponse{}, nil
		} else {
			return &proto.RunResponse{}, BookingNotFound
		}
	} else {
		return &proto.RunResponse{}, UsernameNotFound
	}
}
