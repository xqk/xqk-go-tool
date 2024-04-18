package xqkStrList

import (
	"reflect"
	"testing"
)

func TestGetAscUnicode(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "t1",
			args: args{
				list: []string{
					"IsTypeAndroid",
					"IsTypeLinux",
					"IsTypeFreebsd",
					"IsGoarchArm",
				},
			},
			want: []string{
				"IsGoarchArm",
				"IsTypeAndroid",
				"IsTypeFreebsd",
				"IsTypeLinux",
			},
		},
		{
			name: "t2",
			args: args{
				list: []string{
					"IsGoarchMips64p32le",
					"IsGoarchMips64p32",
				},
			},
			want: []string{
				"IsGoarchMips64p32",
				"IsGoarchMips64p32le",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAscUnicode(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAscUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}
