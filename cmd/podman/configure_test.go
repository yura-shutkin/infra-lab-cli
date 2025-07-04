package podman

import (
	"testing"
)

func Test_convertToMiB(t *testing.T) {
	type args struct {
		size string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "plain MiB value",
			args:    args{size: "2048"},
			want:    "2048",
			wantErr: false,
		},
		{
			name:    "MiB with M suffix",
			args:    args{size: "2048M"},
			want:    "2048",
			wantErr: false,
		},
		{
			name:    "MiB with m suffix",
			args:    args{size: "2048m"},
			want:    "2048",
			wantErr: false,
		},
		{
			name:    "GiB with G suffix",
			args:    args{size: "2G"},
			want:    "2048",
			wantErr: false,
		},
		{
			name:    "GiB with g suffix",
			args:    args{size: "2g"},
			want:    "2048",
			wantErr: false,
		},
		{
			name:    "Fractional GiB",
			args:    args{size: "2.5G"},
			want:    "2560",
			wantErr: false,
		},
		{
			name:    "Invalid input",
			args:    args{size: "abc"},
			want:    "0",
			wantErr: true,
		},
		{
			name:    "Invalid input GB",
			args:    args{size: "1GB"},
			want:    "0",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToMiB(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToMiB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertToMiB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isParamChanged(t *testing.T) {
	type args struct {
		param        string
		currentValue string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Param changed",
			args:    args{param: "3", currentValue: "2"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Param not provided",
			args:    args{param: "0", currentValue: "2"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Param not changed",
			args:    args{param: "2", currentValue: "2"},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isParamChanged(tt.args.param, tt.args.currentValue)
			if got != tt.want {
				t.Errorf("convertToMiB() got = %v, want %v", got, tt.want)
			}
		})
	}
}
