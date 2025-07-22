package podman

import (
	"testing"
)

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
