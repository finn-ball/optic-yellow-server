package server

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestBookingsError_Error(t *testing.T) {
	tests := []struct {
		e    BookingsError
		want string
	}{
		{
			e:    Unknown,
			want: "unknown booking error",
		},
		{
			e:    UsernameNotFound,
			want: "username not found",
		},
		{
			e:    BookingNotFound,
			want: "booking not found",
		},
		{
			e:    BookingAlreadyExists,
			want: "booking already exists",
		},
	}
	for _, tt := range tests {
		t.Run(tt.e.Error(), func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("BookingsError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookingsError_ToGrpcStatus(t *testing.T) {
	tests := []struct {
		e    BookingsError
		want codes.Code
	}{
		{
			e:    Unknown,
			want: codes.Unknown,
		},
		{
			e:    UsernameNotFound,
			want: codes.NotFound,
		},
		{
			e:    BookingNotFound,
			want: codes.NotFound,
		},
		{
			e:    BookingAlreadyExists,
			want: codes.AlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.e.Error(), func(t *testing.T) {
			err := tt.e.ToGrpcStatus()
			got, ok := status.FromError(err)
			if !ok {
				t.Errorf("BookingsError.ToGrpcStatus() error = %v, want %v", err, tt.want)
			}
			if got.Code() != tt.want {
				t.Errorf("BookingsError.ToGrpcStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
