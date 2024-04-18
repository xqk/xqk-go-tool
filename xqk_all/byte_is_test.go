package xqkAll

import "testing"

func TestXqkByteIsLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				b: 'a',
			},
			want: true,
		},
		{
			name: "测试2",
			args: args{
				b: 'b',
			},
			want: true,
		},
		{
			name: "测试3",
			args: args{
				b: 'z',
			},
			want: true,
		},
		{
			name: "测试4",
			args: args{
				b: 'A',
			},
			want: true,
		},
		{
			name: "测试5",
			args: args{
				b: 'Z',
			},
			want: true,
		},
		{
			name: "测试6",
			args: args{
				b: '1',
			},
			want: false,
		},
		{
			name: "测试7",
			args: args{
				b: '*',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkByteIsLetter(tt.args.b); got != tt.want {
				t.Errorf("XqkByteIsLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXqkByteIsLowerLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				b: 'a',
			},
			want: true,
		},
		{
			name: "测试2",
			args: args{
				b: 'b',
			},
			want: true,
		},
		{
			name: "测试3",
			args: args{
				b: 'z',
			},
			want: true,
		},
		{
			name: "测试4",
			args: args{
				b: 'A',
			},
			want: false,
		},
		{
			name: "测试5",
			args: args{
				b: 'Z',
			},
			want: false,
		},
		{
			name: "测试6",
			args: args{
				b: '1',
			},
			want: false,
		},
		{
			name: "测试7",
			args: args{
				b: '*',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkByteIsLowerLetter(tt.args.b); got != tt.want {
				t.Errorf("XqkByteIsLowerLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXqkByteIsUpperLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				b: 'a',
			},
			want: false,
		},
		{
			name: "测试2",
			args: args{
				b: 'b',
			},
			want: false,
		},
		{
			name: "测试3",
			args: args{
				b: 'z',
			},
			want: false,
		},
		{
			name: "测试4",
			args: args{
				b: 'A',
			},
			want: true,
		},
		{
			name: "测试5",
			args: args{
				b: 'Z',
			},
			want: true,
		},
		{
			name: "测试6",
			args: args{
				b: '1',
			},
			want: false,
		},
		{
			name: "测试7",
			args: args{
				b: '*',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkByteIsUpperLetter(tt.args.b); got != tt.want {
				t.Errorf("XqkByteIsUpperLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXqkByteIsNotLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				b: 'a',
			},
			want: false,
		},
		{
			name: "测试2",
			args: args{
				b: 'b',
			},
			want: false,
		},
		{
			name: "测试3",
			args: args{
				b: 'z',
			},
			want: false,
		},
		{
			name: "测试4",
			args: args{
				b: 'A',
			},
			want: false,
		},
		{
			name: "测试5",
			args: args{
				b: 'Z',
			},
			want: false,
		},
		{
			name: "测试6",
			args: args{
				b: '1',
			},
			want: true,
		},
		{
			name: "测试7",
			args: args{
				b: '*',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkByteIsNotLetter(tt.args.b); got != tt.want {
				t.Errorf("XqkByteIsNotLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
