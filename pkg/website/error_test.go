package website

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestWebsiteError_Error(t *testing.T) {
	tests := []struct {
		e    WebsiteError
		want string
	}{
		{
			e:    Unknown,
			want: "unknown website error",
		},
		{
			e:    LoginFailed,
			want: "login failed",
		},
		{
			e:    CourtsUnavailable,
			want: "courts unavailable",
		},
		{
			e:    BookingPageFailed,
			want: "could not get to booking page",
		},
		{
			e:    TraverseSlotsFailed,
			want: "unable to traverse slots",
		},
		{
			e:    TraverseMainFailed,
			want: "unable to traverse main bookings page",
		},
		{
			e:    DateUnavailable,
			want: "date unavailable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("WebsiteError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebsiteError_ToGrpcStatus(t *testing.T) {
	tests := []struct {
		e    WebsiteError
		want codes.Code
	}{
		{
			e:    Unknown,
			want: codes.Unknown,
		},
		{
			e:    LoginFailed,
			want: codes.Unauthenticated,
		},
		{
			e:    CourtsUnavailable,
			want: codes.NotFound,
		},
		{
			e:    BookingPageFailed,
			want: codes.Unknown,
		},
		{
			e:    TraverseSlotsFailed,
			want: codes.Unknown,
		},
		{
			e:    TraverseMainFailed,
			want: codes.Unknown,
		},
		{
			e:    DateUnavailable,
			want: codes.OutOfRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.e.Error(), func(t *testing.T) {
			err := tt.e.ToGrpcStatus()
			got, ok := status.FromError(err)
			if !ok {
				t.Errorf("WebsiteError.ToGrpcStatus() error = %v, want %v", err, tt.want)
			}
			if got.Code() != tt.want {
				t.Errorf("WebsiteError.ToGrpcStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
