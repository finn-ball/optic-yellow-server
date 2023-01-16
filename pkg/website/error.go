package website

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WebsiteError uint64

const (
	Unknown WebsiteError = iota
	LoginFailed
	CourtsUnavailable
	BookingPageFailed
	TraverseSlotsFailed
	TraverseMainFailed
	DateUnavailable
)

func (e WebsiteError) Error() string {
	return [...]string{
		"unknown website error",
		"login failed",
		"courts unavailable",
		"could not get to booking page",
		"unable to traverse slots",
		"unable to traverse main bookings page",
		"date unavailable",
	}[e]
}

func (e WebsiteError) ToGrpcStatus() error {
	return [...]error{
		status.Error(
			codes.Unknown,
			e.Error(),
		),
		status.Error(
			codes.Unauthenticated,
			e.Error(),
		),
		status.Error(
			codes.NotFound,
			e.Error(),
		),
		status.Error(
			codes.Unknown,
			e.Error(),
		),
		status.Error(
			codes.Unknown,
			e.Error(),
		),
		status.Error(
			codes.Unknown,
			e.Error(),
		),
		status.Error(
			codes.OutOfRange,
			e.Error(),
		),
	}[e]
}

func (e WebsiteError) wrapError(err error) error {
	return fmt.Errorf("%s: %w", e, err)
}
