package xqkAll

import "testing"

func TestXqkRuneListIsLeByUnicodeAndLowerBeforeUpper(t *testing.T) {
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
				listA: XqkStrToRuneList("xqkVar"),
				listB: XqkStrToRuneList("xqkAll"),
			},
			want: false,
		},
		{
			name: "t2",
			args: args{
				listA: XqkStrToRuneList("xqkvar"),
				listB: XqkStrToRuneList("xqkAll"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkRuneListIsLeByUnicodeAndLowerBeforeUpper(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("XqkRuneListIsLeByUnicodeAndLowerBeforeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXqkRuneListIsGtByUnicodeAndLowerBeforeUpper(t *testing.T) {
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
				listA: XqkStrToRuneList("xqkVar"),
				listB: XqkStrToRuneList("xqkAll"),
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				listA: XqkStrToRuneList("xqkvar"),
				listB: XqkStrToRuneList("xqkAll"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkRuneListIsGtByUnicodeAndLowerBeforeUpper(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("XqkRuneListIsGtByUnicodeAndLowerBeforeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXqkRuneListIsLeByUnicode(t *testing.T) {
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
				listA: XqkStrToRuneList("IsTypeAndroid"),
				listB: XqkStrToRuneList("IsTypeLinux"),
			},
			want: true,
		},
		{
			name: "t1",
			args: args{
				listA: XqkStrToRuneList("IsGoarchMips64p32"),
				listB: XqkStrToRuneList("IsGoarchMips64p32le"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkRuneListIsLeByUnicode(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("XqkRuneListIsLeByUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}
