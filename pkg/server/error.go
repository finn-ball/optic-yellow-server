package server

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookingsError uint64

const (
	Unknown BookingsError = iota
	UsernameNotFound
	BookingNotFound
	BookingAlreadyExists
)

func (e BookingsError) Error() string {
	return [...]string{
		"unknown booking error",
		"username not found",
		"booking not found",
		"booking already exists",
	}[e]
}

func (e BookingsError) ToGrpcStatus() error {
	return [...]error{
		status.Error(
			codes.Unknown,
			e.Error(),
		),
		status.Error(
			codes.NotFound,
			e.Error(),
		),
		status.Error(
			codes.NotFound,
			e.Error(),
		),
		status.Error(
			codes.AlreadyExists,
			e.Error(),
		),
	}[e]
}
