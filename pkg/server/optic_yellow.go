package server

import (
	"context"
	"sync"
	"time"

	"github.com/finn-ball/optic-yellow-server/pkg/proto"
	"github.com/finn-ball/optic-yellow-server/pkg/website"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	proto.UnimplementedOpticYellowServiceServer
	logger   *zap.SugaredLogger
	headless bool
	timeout  time.Duration
	bookings bookings
}

type bookings struct {
	sync.Mutex
	// booking key is the user
	users map[string]user
}

type user struct {
	password string
	// Key is the time the booking is for.
	// Value allows cancellation.
	queue  map[time.Time]context.CancelFunc
	status map[time.Time]proto.BookingResponse_Status
}

// NewServer creates the Optic Yellow server
func NewServer(headless bool) *grpc.Server {
	s := grpc.NewServer()
	proto.RegisterOpticYellowServiceServer(
		s,
		&grpcServer{
			logger:   newLogger().Sugar(),
			headless: headless,
			timeout:  60 * time.Second,
			bookings: bookings{
				Mutex: sync.Mutex{},
				users: map[string]user{},
			}})
	return s
}

// RunRequest will attempt to execute the action asked.
func (s *grpcServer) Run(ctx context.Context, req *proto.RunRequest) (*proto.RunResponse, error) {
	switch r := req.Request.(type) {
	case *proto.RunRequest_Login:
		resp, err := s.login(ctx, r)
		if err != nil {
			s.logger.Errorw(
				"booking",
				"user", r.Login.Username,
				"details", err,
			)
		}
		switch err {
		case nil:
			return resp, nil
		case website.LoginFailed:
			return resp, website.LoginFailed.ToGrpcStatus()
		default:
			return resp, Unknown.ToGrpcStatus()
		}
	case *proto.RunRequest_Booking:
		resp, err := s.booking(ctx, r)
		if err != nil {
			s.logger.Errorw(
				"booking",
				"user", r.Booking.Login.Username,
				"time", r.Booking.Datetime,
				"details", err,
			)
		}
		switch err {
		case nil:
			return resp, nil
		case website.LoginFailed:
			return resp, website.LoginFailed.ToGrpcStatus()
		case website.CourtsUnavailable:
			return resp, website.CourtsUnavailable.ToGrpcStatus()
		case UsernameNotFound:
			return resp, UsernameNotFound.ToGrpcStatus()
		case BookingAlreadyExists:
			return resp, BookingAlreadyExists.ToGrpcStatus()
		default:
			return resp, Unknown.ToGrpcStatus()
		}
	case *proto.RunRequest_List:
		resp, err := s.list(ctx, r)
		if err != nil {
			s.logger.Errorw(
				"list",
				"user", r.List.Username,
				"details", err,
			)
		}
		switch err {
		case nil:
			return resp, nil
		case website.LoginFailed:
			return resp, website.LoginFailed.ToGrpcStatus()
		default:
			return resp, Unknown.ToGrpcStatus()
		}
	case *proto.RunRequest_Cancel:
		resp, err := s.cancel(ctx, r)
		if err != nil {
			s.logger.Errorw(
				"cancel",
				"user", r.Cancel.Login.Username,
				"details", err,
			)
		}
		switch err {
		case nil:
			return resp, nil
		case website.LoginFailed:
			return resp, website.LoginFailed.ToGrpcStatus()
		case UsernameNotFound:
			return resp, UsernameNotFound.ToGrpcStatus()
		case BookingNotFound:
			return resp, BookingNotFound.ToGrpcStatus()
		default:
			return resp, Unknown.ToGrpcStatus()
		}
	default:
		// Should never get here
		s.logger.Errorw(
			"critical error",
			"details", r,
		)
		return &proto.RunResponse{}, status.Error(
			codes.Unimplemented,
			"unimplemented",
		)
	}
}

func (s *grpcServer) updatePassword(username, password string) {
	s.bookings.Lock()
	defer s.bookings.Unlock()
	if entry, ok := s.bookings.users[username]; ok {
		entry.password = password
	} else {
		s.bookings.users[username] = user{
			password: password,
			queue:    map[time.Time]context.CancelFunc{},
		}
	}
}

// getPassword returns false if user does not exist
func (s *grpcServer) getPassword(username string) (string, bool) {
	s.bookings.Lock()
	defer s.bookings.Unlock()
	if entry, ok := s.bookings.users[username]; ok {
		return entry.password, true
	}
	return "", false
}

// updateBookingState returns false if user does not exist
func (s *grpcServer) updateBookingState(username string, t time.Time, state proto.BookingResponse_Status) bool {
	s.bookings.Lock()
	defer s.bookings.Unlock()
	if entry, ok := s.bookings.users[username]; ok {
		entry.status[t] = state
		return true
	}
	return false
}
