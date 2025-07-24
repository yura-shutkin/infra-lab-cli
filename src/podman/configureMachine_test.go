package podman

import (
	"testing"
)

func Test_checkIfParamChanged(t *testing.T) {
	type args struct {
		param        ConfigParam
		currentValue int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Param not changed",
			args: args{
				param:        ConfigParam{ValueFlag: "2"},
				currentValue: 2,
			},
			want: false,
		},
		{
			name: "Param changed",
			args: args{
				param:        ConfigParam{ValueFlag: "4"},
				currentValue: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			param := tt.args.param
			_ = checkIfParamChanged(&param, tt.args.currentValue)
			if param.IsChanged != tt.want {
				t.Errorf("IsChanged = %v, want %v", param.IsChanged, tt.want)
			}
		})
	}
}

func Test_checkIfMemoryChanged(t *testing.T) {
	type args struct {
		param        ConfigParam
		currentValue int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Memory MiB not changed without postfix",
			args: args{
				param:        ConfigParam{ValueFlag: "2048"},
				currentValue: 2048,
			},
			want: false,
		},
		{
			name: "Memory MiB not changed with postfix",
			args: args{
				param:        ConfigParam{ValueFlag: "2048M"},
				currentValue: 2048,
			},
			want: false,
		},
		{
			name: "Memory GiB not changed",
			args: args{
				param:        ConfigParam{ValueFlag: "2G"},
				currentValue: 2048,
			},
			want: false,
		},
		{
			name: "Memory GiB not changed with fraction",
			args: args{
				param:        ConfigParam{ValueFlag: "2.5G"},
				currentValue: 2560,
			},
			want: false,
		},
		{
			name: "Memory MiB changed without postfix",
			args: args{
				param:        ConfigParam{ValueFlag: "2049"},
				currentValue: 2048,
			},
			want: true,
		},
		{
			name: "Memory MiB changed with postfix",
			args: args{
				param:        ConfigParam{ValueFlag: "2049M"},
				currentValue: 2048,
			},
			want: true,
		},
		{
			name: "Memory GiB changed",
			args: args{
				param:        ConfigParam{ValueFlag: "3G"},
				currentValue: 2048,
			},
			want: true,
		},
		{
			name: "Memory GiB changed with fraction",
			args: args{
				param:        ConfigParam{ValueFlag: "2.5G"},
				currentValue: 2048,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			param := tt.args.param
			_ = checkIfMemoryChanged(&param, tt.args.currentValue)
			if param.IsChanged != tt.want {
				t.Errorf("IsChanged = %v, want %v", param.IsChanged, tt.want)
			}
		})
	}
}
