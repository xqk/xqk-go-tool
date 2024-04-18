package xqkAll

import "testing"

func TestXqkStrGetTrimLeftSStr(t *testing.T) {
	type args struct {
		str          string
		targetStrArr []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "正常测试1",
			args: args{
				str: "	◎ 测试字符串",
				targetStrArr: []string{
					" ",
				},
			},
			want: "\t◎ 测试字符串",
		},
		{
			name: "正常测试2",
			args: args{
				str: "	◎ 测试字符串",
				targetStrArr: []string{
					" ",
					"\t",
				},
			},
			want: "◎ 测试字符串",
		},
		{
			name: "正常测试3",
			args: args{
				str: "	◎ 测试字符串",
				targetStrArr: []string{
					" ",
					"◎",
					"\t",
				},
			},
			want: "测试字符串",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkStrGetTrimLeftSStr(tt.args.str, tt.args.targetStrArr...); got != tt.want {
				t.Errorf("XqkStrGetTrimLeftSStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXqkStrGetReplaceEndStr(t *testing.T) {
	type args struct {
		s   string
		end string
		r   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "正常测试1",
			args: args{
				s:   "Hello Xqk!",
				end: "!",
				r:   ".",
			},
			want: "Hello Xqk.",
		},
		{
			name: "正常测试2",
			args: args{
				s:   "Hello Xqk!",
				end: "Fidel!",
				r:   ".",
			},
			want: "Hello Xqk!",
		},
		{
			name: "空测试",
			args: args{
				s:   "",
				end: "",
				r:   ".",
			},
			want: ".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XqkStrGetReplaceEndStr(tt.args.s, tt.args.end, tt.args.r); got != tt.want {
				t.Errorf("XqkStrGetReplaceEndStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
