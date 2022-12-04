//go:build fast

package internal

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "-1 error",
			args: args{
				n: -1,
			},
			want: 0,
			wantErr: true,
		},
		{
			name: "0 test",
			args: args{
				n: 0,
			},
			want: 0,
			wantErr: false,
		},
		{
			name: "1 test",
			args: args{
				n: 1,
			},
			want: 1,
			wantErr: false,
		},
		{
			name: "3 test",
			args: args{
				n: 3,
			},
			want: 2,
			wantErr: false,
		},
		{
			name: "5 test",
			args: args{
				n: 5,
			},
			want: 5,
			wantErr: false,
		},
		{
			name: "10 test",
			args: args{
				n: 10,
			},
			want: 55,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fib(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("fib() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
