package xqkByteList

import (
	"reflect"
	"strings"
	"testing"

	xqkAll "github.com/xqk/xqk-go-tool/xqk_all"
)

func TestGetDeduplicate(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "去重测试1",
			args: args{
				list: xqkAll.XqkStrToByteList("lllHello"),
			},
			want: xqkAll.XqkStrToByteList("lHeo"),
		},
		{
			name: "去重测试2",
			args: args{
				list: xqkAll.XqkStrToByteList("lllHHHHelloOOO"),
			},
			want: xqkAll.XqkStrToByteList("lHeoO"),
		},
		{
			name: "去重测试3",
			args: args{
				list: []byte{1, 2, 3, 1, 5, 4},
			},
			want: []byte{1, 2, 3, 5, 4},
		},
		{
			name: "去重测试4",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
			},
			want: xqkAll.XqkStrToByteList("Helo Xqk!"),
		},
		{
			name: "去重测试5-空去重",
			args: args{
				list: []byte{},
			},
			want: []byte{},
		},
		{
			name: "去重测试6-nil去重",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeduplicate(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeduplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDeleteByIndex(t *testing.T) {
	type args struct {
		list     []byte
		delIndex int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "删除测试1-正常删除",
			args: args{
				list:     xqkAll.XqkStrToByteList("Hello Xqk!"),
				delIndex: 0,
			},
			want: xqkAll.XqkStrToByteList("ello Xqk!"),
		},
		{
			name: "删除测试2-正常删除",
			args: args{
				list:     xqkAll.XqkStrToByteList("Hello Xqk!"),
				delIndex: len("Hello Xqk!") - 1,
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk"),
		},
		{
			name: "删除测试3-正常删除",
			args: args{
				list:     xqkAll.XqkStrToByteList("Hello Xqk!"),
				delIndex: strings.Index("Hello Xqk!", " "),
			},
			want: xqkAll.XqkStrToByteList("HelloXqk!"),
		},
		{
			name: "删除测试4-负值删除",
			args: args{
				list:     xqkAll.XqkStrToByteList("Hello Xqk!"),
				delIndex: -1,
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "删除测试5-越界删除",
			args: args{
				list:     xqkAll.XqkStrToByteList("Hello Xqk!"),
				delIndex: 999,
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "删除测试6-空删除",
			args: args{
				list:     []byte{},
				delIndex: 0,
			},
			want: []byte{},
		},
		{
			name: "删除测试7-nil删除",
			args: args{
				list:     nil,
				delIndex: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeleteByIndex(tt.args.list, tt.args.delIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeleteByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDeleteByRangeIndex(t *testing.T) {
	type args struct {
		list       []byte
		startIndex int
		endIndex   int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "范围删除1-正常删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: 0,
				endIndex:   0,
			},
			want: xqkAll.XqkStrToByteList("ello Xqk!"),
		},
		{
			name: "范围删除2-正常删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: 0,
				endIndex:   1,
			},
			want: xqkAll.XqkStrToByteList("llo Xqk!"),
		},
		{
			name: "范围删除3-正常删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: strings.Index("Hello Xqk!", " "),
				endIndex:   strings.Index("Hello Xqk!", "u"),
			},
			want: xqkAll.XqkStrToByteList("Hello!"),
		},
		{
			name: "范围删除4-负值删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: -1,
				endIndex:   strings.Index("Hello Xqk!", "u"),
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "范围删除5-负值删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: -1,
				endIndex:   -13,
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "范围删除6-负值删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: 2,
				endIndex:   -13,
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "范围删除7-越界删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: 999,
				endIndex:   3,
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "范围删除8-越界删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: 3,
				endIndex:   999,
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "范围删除9-越界删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: 998,
				endIndex:   999,
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "范围删除10-开始索引大于结束索引删除",
			args: args{
				list:       xqkAll.XqkStrToByteList("Hello Xqk!"),
				startIndex: 3,
				endIndex:   2,
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "范围删除11-空删除",
			args: args{
				list:       []byte{},
				startIndex: 0,
				endIndex:   0,
			},
			want: []byte{},
		},
		{
			name: "范围删除12-nil删除",
			args: args{
				list:       nil,
				startIndex: 0,
				endIndex:   0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeleteByRangeIndex(tt.args.list, tt.args.startIndex, tt.args.endIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeleteByRangeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFilter(t *testing.T) {
	type args struct {
		list []byte
		keep func(x byte) bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "过滤测试1",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				keep: func(x byte) bool {
					return x == 'Y' || x == 'i' || x == 'u'
				},
			},
			want: xqkAll.XqkStrToByteList("Xqk"),
		},
		{
			name: "过滤测试2",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				keep: func(x byte) bool {
					return false
				},
			},
			want: xqkAll.XqkStrToByteList(""),
		},
		{
			name: "过滤测试3",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				keep: func(x byte) bool {
					return true
				},
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFilter(tt.args.list, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPop(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		want1   []byte
		wantErr bool
	}{
		{
			name: "出栈测试1-正常出栈",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
			},
			want:    '!',
			want1:   xqkAll.XqkStrToByteList("Hello Xqk"),
			wantErr: false,
		},
		{
			name: "出栈测试2-空出栈",
			args: args{
				list: []byte{},
			},
			want:    0,
			want1:   []byte{},
			wantErr: true,
		},
		{
			name: "出栈测试2-nil出栈",
			args: args{
				list: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetPop(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetReverse(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "反序测试1",
			args: args{
				list: []byte{1, 2, 3, 4, 5, 6},
			},
			want: []byte{6, 5, 4, 3, 2, 1},
		},
		{
			name: "反序测试2",
			args: args{
				list: xqkAll.XqkStrToByteList("Xqk!"),
			},
			want: xqkAll.XqkStrToByteList("!uiY"),
		},
		{
			name: "反序测试3-空反序",
			args: args{
				list: []byte{},
			},
			want: []byte{},
		},
		{
			name: "反序测试4-nil反序",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReverse(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetShuffle(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "乱序测试1",
			args: args{
				list: []byte{1},
			},
			want: []byte{1},
		},
		{
			name: "乱序测试2-空乱序",
			args: args{
				list: []byte{},
			},
			want: []byte{},
		},
		{
			name: "乱序测试3-nil乱序",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetShuffle(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMap(t *testing.T) {
	type args struct {
		list   []byte
		opFunc func(int, byte) byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Map测试1",
			args: args{
				list: []byte{1, 2, 3, 4},
				opFunc: func(i int, b byte) byte {
					return byte(i+1) + b
				},
			},
			want: []byte{2, 4, 6, 8},
		},
		{
			name: "Map测试2-空测试",
			args: args{
				list: []byte{},
				opFunc: func(i int, b byte) byte {
					return byte(i+1) + b
				},
			},
			want: []byte{},
		},
		{
			name: "Map测试2-nil测试",
			args: args{
				list: nil,
				opFunc: func(i int, b byte) byte {
					return byte(i+1) + b
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMap(tt.args.list, tt.args.opFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCopy(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "拷贝测试1",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
			},
			want: xqkAll.XqkStrToByteList("Hello Xqk!"),
		},
		{
			name: "拷贝测试2-空测试",
			args: args{
				list: []byte{},
			},
			want: []byte{},
		},
		{
			name: "拷贝测试2-nil测试",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCopy(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndexByEl(t *testing.T) {
	type args struct {
		list []byte
		el   byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "索引元素测试1-正常索引",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				el:   'l',
			},
			want: strings.Index("Hello Xqk!", "l"),
		},
		{
			name: "索引元素测试2-无值索引",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				el:   '=',
			},
			want: -1,
		},
		{
			name: "索引元素测试3-空索引",
			args: args{
				list: []byte{},
				el:   'l',
			},
			want: -1,
		},
		{
			name: "索引元素测试4-nil索引",
			args: args{
				list: nil,
				el:   'l',
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIndexByEl(tt.args.list, tt.args.el); got != tt.want {
				t.Errorf("GetIndexByEl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndexByList(t *testing.T) {
	type args struct {
		list    []byte
		subList []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "多值索引1",
			args: args{
				list:    xqkAll.XqkStrToByteList("Hello Xqk!"),
				subList: xqkAll.XqkStrToByteList("ell"),
			},
			want: 1,
		},
		{
			name: "多值索引2",
			args: args{
				list:    xqkAll.XqkStrToByteList("Hello Xqk!"),
				subList: xqkAll.XqkStrToByteList("Xqk"),
			},
			want: 6,
		},
		{
			name: "多值索引3-空索引",
			args: args{
				list:    []byte{},
				subList: xqkAll.XqkStrToByteList("Xqk"),
			},
			want: -1,
		},
		{
			name: "多值索引4-空索引",
			args: args{
				list:    []byte{},
				subList: []byte{},
			},
			want: -1,
		},
		{
			name: "多值索引5-nil索引",
			args: args{
				list:    nil,
				subList: xqkAll.XqkStrToByteList("Xqk"),
			},
			want: -1,
		},
		{
			name: "多值索引6-nil索引",
			args: args{
				list:    nil,
				subList: nil,
			},
			want: -1,
		},
		{
			name: "多值索引7-索引空",
			args: args{
				list:    xqkAll.XqkStrToByteList("Xqk"),
				subList: []byte{},
			},
			want: -1,
		},
		{
			name: "多值索引8-索引nil",
			args: args{
				list:    xqkAll.XqkStrToByteList("Xqk"),
				subList: nil,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIndexByList(tt.args.list, tt.args.subList); got != tt.want {
				t.Errorf("GetIndexByList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndexBySList(t *testing.T) {
	type args struct {
		list       []byte
		subListArr [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "多值多组索引1",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("ll"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want: 2,
		},
		{
			name: "多值多组索引2",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want: 6,
		},
		{
			name: "多值多组索引3-索引空",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList(""),
					{},
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want: 6,
		},
		{
			name: "多值多组索引4-索引nil",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList(""),
					nil,
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want: 6,
		},
		{
			name: "多值多组索引5-空索引",
			args: args{
				list: []byte{},
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want: -1,
		},
		{
			name: "多值多组索引6-nil索引",
			args: args{
				list: nil,
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want: -1,
		},
		{
			name: "多值多组索引7-nil索引",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("fan"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIndexBySList(tt.args.list, tt.args.subListArr...); got != tt.want {
				t.Errorf("GetIndexBySList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndexAndSubBySList(t *testing.T) {
	type args struct {
		list       []byte
		subListArr [][]byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []byte
	}{
		{
			name: "多值多组索引元素测试1",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("ll"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want:  2,
			want1: xqkAll.XqkStrToByteList("ll"),
		},
		{
			name: "多值多组索引元素测试2",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want:  6,
			want1: xqkAll.XqkStrToByteList("Xqk"),
		},
		{
			name: "多值多组索引元素测试3-索引空",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList(""),
					{},
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want:  6,
			want1: xqkAll.XqkStrToByteList("Xqk"),
		},
		{
			name: "多值多组索引元素测试4-索引nil",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList(""),
					nil,
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want:  6,
			want1: xqkAll.XqkStrToByteList("Xqk"),
		},
		{
			name: "多值多组索引元素测试5-空索引",
			args: args{
				list: []byte{},
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want:  -1,
			want1: nil,
		},
		{
			name: "多值多组索引元素测试6-nil索引",
			args: args{
				list: nil,
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("iu"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want:  -1,
			want1: nil,
		},
		{
			name: "多值多组索引7-nil索引",
			args: args{
				list: xqkAll.XqkStrToByteList("Hello Xqk!"),
				subListArr: [][]byte{
					xqkAll.XqkStrToByteList("fan"),
					xqkAll.XqkStrToByteList("Xqk"),
					xqkAll.XqkStrToByteList("!"),
				},
			},
			want:  6,
			want1: xqkAll.XqkStrToByteList("Xqk"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetIndexAndSubBySList(tt.args.list, tt.args.subListArr...)
			if got != tt.want {
				t.Errorf("GetIndexAndSubBySList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetIndexAndSubBySList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetElByIndex(t *testing.T) {
	type args struct {
		list  []byte
		index int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "索引元素测试1",
			args: args{
				list:  []byte{1, 2, 3, 4, 5},
				index: 0,
			},
			want: 1,
		},
		{
			name: "索引元素测试2",
			args: args{
				list:  []byte{1, 2, 3, 4, 5},
				index: 2,
			},
			want: 3,
		},
		{
			name: "索引元素测试3-负值索引",
			args: args{
				list:  []byte{1, 2, 3, 4, 5},
				index: -2,
			},
			want: 0,
		},
		{
			name: "索引元素测试4-越界索引",
			args: args{
				list:  []byte{1, 2, 3, 4, 5},
				index: 999,
			},
			want: 0,
		},
		{
			name: "索引元素测试5-空索引",
			args: args{
				list:  []byte{},
				index: 2,
			},
			want: 0,
		},
		{
			name: "索引元素测试6-nil索引",
			args: args{
				list:  nil,
				index: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetElByIndex(tt.args.list, tt.args.index); got != tt.want {
				t.Errorf("GetElByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMax(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{
			name: "最大值测试1",
			args: args{
				list: []byte{1, 2, 3, 4, 5, 6},
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "最大值测试2-空最大值",
			args: args{
				list: []byte{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "最大值测试2-nil最大值",
			args: args{
				list: nil,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMax(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMax() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMin(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{
			name: "最小值测试1",
			args: args{
				list: []byte{1, 2, 3, 4, 5, 6},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "最小值测试2-空最小值",
			args: args{
				list: []byte{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "最小值测试3-nil最小值",
			args: args{
				list: nil,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMin(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSum(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "求和测试1",
			args: args{
				list: []byte{1, 2, 3, 4, 5, 6},
			},
			want: 21,
		},
		{
			name: "求和测试2-空求和",
			args: args{
				list: []byte{},
			},
			want: 0,
		},
		{
			name: "求和测试3-nil求和",
			args: args{
				list: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSum(tt.args.list); got != tt.want {
				t.Errorf("GetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
