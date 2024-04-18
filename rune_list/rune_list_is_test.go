package xqkRuneList

import (
	"testing"

	xqkAll "github.com/xqk/xqk-go-tool/xqk_all"
)

func TestIsLeByUnicodeAndLowerBeforeUpper(t *testing.T) {
	type args struct {
		listA []rune
		listB []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				listA: xqkAll.XqkStrToRuneList("xqkVar"),
				listB: xqkAll.XqkStrToRuneList("xqkAll"),
			},
			want: false,
		},
		{
			name: "t2",
			args: args{
				listA: xqkAll.XqkStrToRuneList("xqkvar"),
				listB: xqkAll.XqkStrToRuneList("xqkAll"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeByUnicodeAndLowerBeforeUpper(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("IsLeByUnicodeAndLowerBeforeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGtByUnicodeAndLowerBeforeUpper(t *testing.T) {
	type args struct {
		listA []rune
		listB []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				listA: xqkAll.XqkStrToRuneList("xqkVar"),
				listB: xqkAll.XqkStrToRuneList("xqkAll"),
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				listA: xqkAll.XqkStrToRuneList("xqkvar"),
				listB: xqkAll.XqkStrToRuneList("xqkAll"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGtByUnicodeAndLowerBeforeUpper(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("IsGtByUnicodeAndLowerBeforeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLeByUnicode(t *testing.T) {
	type args struct {
		listA []rune
		listB []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				listA: xqkAll.XqkStrToRuneList("IsTypeAndroid"),
				listB: xqkAll.XqkStrToRuneList("IsTypeLinux"),
			},
			want: true,
		},
		{
			name: "t1",
			args: args{
				listA: xqkAll.XqkStrToRuneList("IsGoarchMips64p32"),
				listB: xqkAll.XqkStrToRuneList("IsGoarchMips64p32le"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeByUnicode(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("IsLeByUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}
