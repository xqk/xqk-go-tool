package xqkAll

import "testing"

func TestXqkBoolToInt(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正常测试true",
			args: args{
				b: true,
			},
			want: 1,
		},
		{
			name: "正常测试false",
			args: args{
				b: false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkBoolToInt(tt.args.b); got != tt.want {
				t.Errorf("XqkBoolToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXqkBoolToStr(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "正常测试true",
			args: args{
				b: true,
			},
			want: "true",
		},
		{
			name: "正常测试false",
			args: args{
				b: false,
			},
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkBoolToStr(tt.args.b); got != tt.want {
				t.Errorf("XqkBoolToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
