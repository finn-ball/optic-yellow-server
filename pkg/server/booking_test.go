package server

import (
	"testing"
	"time"
)

func Test_isValidNow(t *testing.T) {
	now := time.Date(2022, 1, 1, 1, 0, 0, 0, time.UTC)
	type args struct {
		now   time.Time
		check time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal time",
			args: args{now, now},
			want: false,
		},
		{
			name: "add hour",
			args: args{now, now.Add(time.Hour)},
			want: true,
		},
		{
			name: "add two hours",
			args: args{now, now.Add(2 * time.Hour)},
			want: true,
		},
		{
			name: "add three days",
			args: args{now, now.Add(3 * 24 * time.Hour)},
			want: true,
		},
		{
			name: "add four days",
			args: args{now, now.Add(4 * 24 * time.Hour)},
			want: false,
		},
		{
			name: "add second",
			args: args{now, now.Add(time.Second)},
			want: false,
		},
		{
			name: "add minute",
			args: args{now, now.Add(time.Minute)},
			want: false,
		},
		{
			name: "subtract hour",
			args: args{now, now.Add(-time.Hour)},
			want: false,
		},
		{
			name: "subtract day",
			args: args{now, now.Add(-24 * time.Hour)},
			want: false,
		},
		{
			name: "add 4 days less than 96h",
			args: args{
				time.Date(2022, 1, 1, 23, 59, 59, 0, time.UTC),
				time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "add 3 days less than 72h, 1",
			args: args{
				time.Date(2022, 1, 1, 23, 59, 59, 0, time.UTC),
				time.Date(2022, 1, 4, 23, 59, 59, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "add 3 days less than 72h, 2",
			args: args{
				time.Date(2022, 1, 1, 23, 59, 59, 0, time.UTC),
				time.Date(2022, 1, 4, 23, 58, 59, 0, time.UTC),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidNow(tt.args.now, tt.args.check); got != tt.want {
				t.Errorf("isValidNow() = %v, want %v", got, tt.want)
			}
		})
	}
}
