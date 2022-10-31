package server

import (
	"context"

	"github.com/finn-ball/optic-yellow-server/pkg/proto"
	"github.com/finn-ball/optic-yellow-server/pkg/website"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcServer) login(ctx context.Context, req *proto.RunRequest_Login) (*proto.RunResponse, error) {
	username := req.Login.Username
	password := req.Login.Password
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
	err = session.Login()
	if err == nil {
		s.updatePassword(username, password)
	}
	switch err {
	case website.Unknown:
		return &proto.RunResponse{}, status.Error(codes.Unknown, err.Error())
	case website.LoginFailed:
		return &proto.RunResponse{}, status.Error(codes.Unauthenticated, err.Error())
	}
	s.logger.Infow(
		"user", req.Login.Username,
		"login", "successful",
	)
	return &proto.RunResponse{}, err
}
