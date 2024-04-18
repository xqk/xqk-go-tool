package xqkAll

import "testing"

func TestXqkOsDoOpenFileManagerByParent(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常测试",
			args: args{
				p: "E:\\学习文件\\前端\\Prismjs",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := XqkOsDoOpenFileManagerByParent(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("XqkOsDoOpenFileManagerByParent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
