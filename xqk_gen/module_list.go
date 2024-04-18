package xqkGen

type XqkModule struct {
	PackageName  string
	DirName      string
	MethodModule []XqkMethodModule
}

type XqkMethodModule struct {
	Type     MethodType
	FileName string
}

type MethodType string

const (
	NoType MethodType = "???"
	To     MethodType = "To"
	Get    MethodType = "Get"
	Op     MethodType = "Op"
	Is     MethodType = "Is"
	Do     MethodType = "Do"
)

func GetModuleList() []XqkModule {
	return []XqkModule{
		{
			PackageName: "xqkBool",
			DirName:     "bool",
			MethodModule: []XqkMethodModule{
				{
					Type:     To,
					FileName: "bool_to.go",
				},
			},
		},
		{
			PackageName: "xqkByte",
			DirName:     "byte",
			MethodModule: []XqkMethodModule{
				{
					Type:     Is,
					FileName: "byte_is.go",
				},
				{
					Type:     To,
					FileName: "byte_to.go",
				},
			},
		},
		{
			PackageName: "xqkByteList",
			DirName:     "byte_list",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "byte_list_get.go",
				},
				{
					Type:     Op,
					FileName: "byte_list_op.go",
				},
				{
					Type:     Is,
					FileName: "byte_list_is.go",
				},
				{
					Type:     To,
					FileName: "byte_list_to.go",
				},
			},
		},
		{
			PackageName: "xqkSByteList",
			DirName:     "byte_list_s",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "byte_list_s_get.go",
				},
			},
		},
		{
			PackageName: "xqkDir",
			DirName:     "dir",
			MethodModule: []XqkMethodModule{
				{
					Type:     Is,
					FileName: "dir_is.go",
				},
				{
					Type:     Do,
					FileName: "dir_do.go",
				},
			},
		},
		{
			PackageName: "xqkErr",
			DirName:     "error",
			MethodModule: []XqkMethodModule{
				{
					Type:     Is,
					FileName: "error_is.go",
				},
			},
		},
		{
			PackageName: "xqkSErr",
			DirName:     "error_s",
			MethodModule: []XqkMethodModule{
				{
					Type:     To,
					FileName: "error_s_to.go",
				},
			},
		},
		{
			PackageName: "xqkFile",
			DirName:     "file",
			MethodModule: []XqkMethodModule{
				{
					Type:     Is,
					FileName: "file_is.go",
				},
				{
					Type:     Do,
					FileName: "file_do.go",
				},
			},
		},
		{
			PackageName: "xqkInt",
			DirName:     "int",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "int_get.go",
				},
				{
					Type:     Is,
					FileName: "int_is.go",
				},
				{
					Type:     To,
					FileName: "int_to.go",
				},
			},
		},
		{
			PackageName: "xqkSInt",
			DirName:     "int_s",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "int_s_get.go",
				},
				{
					Type:     To,
					FileName: "int_s_to.go",
				},
			},
		},
		{
			PackageName: "xqkIntList",
			DirName:     "int_list",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "int_list_get.go",
				},
				{
					Type:     Op,
					FileName: "int_list_op.go",
				},
				{
					Type:     Is,
					FileName: "int_list_is.go",
				},
				{
					Type:     To,
					FileName: "int_list_to.go",
				},
			},
		},
		{
			PackageName: "xqkSIntList",
			DirName:     "int_list_s",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "int_list_s_get.go",
				},
			},
		},
		{
			PackageName: "xqkLog",
			DirName:     "log",
			MethodModule: []XqkMethodModule{
				{
					Type:     NoType,
					FileName: "log.go",
				},
			},
		},
		{
			PackageName: "xqkOs",
			DirName:     "os",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "os_get.go",
				},
				{
					Type:     Is,
					FileName: "os_is.go",
				},
				{
					Type:     Do,
					FileName: "os_do.go",
				},
			},
		},
		{
			PackageName: "xqkTime",
			DirName:     "time",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "time_get.go",
				},
				{
					Type:     To,
					FileName: "time_to.go",
				},
			},
		},
		{
			PackageName: "xqkBaseErr",
			DirName:     "xqk_error",
			MethodModule: []XqkMethodModule{
				{
					Type:     Is,
					FileName: "xqk_error_is.go",
				},
			},
		},
		{
			PackageName: "xqkStr",
			DirName:     "string",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "str_get.go",
				},
				{
					Type:     Is,
					FileName: "str_is.go",
				},
				{
					Type:     Op,
					FileName: "str_op.go",
				},
				{
					Type:     To,
					FileName: "str_to.go",
				},
			},
		},
		{
			PackageName: "xqkSStr",
			DirName:     "string_s",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "str_s_get.go",
				},
			},
		},
		{
			PackageName: "xqkStrList",
			DirName:     "string_list",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "str_list_get.go",
				},
				{
					Type:     Op,
					FileName: "str_list_op.go",
				},
				{
					Type:     Is,
					FileName: "str_list_is.go",
				},
				{
					Type:     To,
					FileName: "str_list_to.go",
				},
			},
		},
		{
			PackageName: "xqkSStrList",
			DirName:     "string_list_s",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "str_list_s_get.go",
				},
			},
		},
		{
			PackageName: "xqkRune",
			DirName:     "rune",
			MethodModule: []XqkMethodModule{
				{
					Type:     To,
					FileName: "rune_to.go",
				},
			},
		},
		{
			PackageName: "xqkRuneList",
			DirName:     "rune_list",
			MethodModule: []XqkMethodModule{
				{
					Type:     Get,
					FileName: "rune_list_get.go",
				},
				{
					Type:     Op,
					FileName: "rune_list_op.go",
				},
				{
					Type:     Is,
					FileName: "rune_list_is.go",
				},
			},
		},
	}
}
