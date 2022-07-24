package main

import "testing"

func Test_isStraight(t *testing.T) {
	type args struct {
		cartas []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "{2, 3, 4, 5, 6}",
			args: args{cartas: []int{2, 3, 4, 5, 6}},
			want: true,
		},
		{
			name: "{14, 5, 4, 2, 3}",
			args: args{cartas: []int{14, 5, 4, 2, 3}},
			want: true,
		},
		{
			name: "{7, 7, 12, 11, 3, 4, 14}",
			args: args{cartas: []int{7, 7, 12, 11, 3, 4, 14}},
			want: false,
		},
		{
			name: "{7, 3, 2}",
			args: args{cartas: []int{7, 3, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.cartas); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}
