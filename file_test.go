package main

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFileVisitor_Get(t *testing.T) {
	testCases := []struct {
		src  string
		want *File
	}{
		{
			src: `
package morm_gen
import (
	"fmt"
    sqlx "database/sql"
)

import (
	dri "database/sql/driver"
)
type (
	TestStructType struct {
		// Public is a field
		// @type string
		PublicField string
        Ptr *sqlx.NullString
		Struct sqlx.NullInt64
		Age *int8
		Slice []byte
	}
)
`,
			want: &File{
				Package: "morm_gen",
				Imports: []string{`"fmt"`, `sqlx "database/sql"`, `dri "database/sql/driver"`},
				Types: []Type{
					{
						Name: "TestStructType",
						Fields: []Field{
							{
								Name: "PublicField",
								Type: "string",
							},
							{
								Name: "Ptr",
								Type: "*sqlx.NullString",
							},
							{
								Name: "Struct",
								Type: "sqlx.NullInt64",
							},
							{
								Name: "Age",
								Type: "*int8",
							},
							{
								Name: "Slice",
								Type: "[]byte",
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		fset := token.NewFileSet() // 创建一个新的文件集合 fset，用于存储源代码解析的位置信息
		// 使用 parser.ParseFile 函数解析源代码，并将解析结果保存在变量 f 中
		f, err := parser.ParseFile(fset, "src.go", tc.src, parser.ParseComments)
		if err != nil {
			t.Fatal(err)
		}
		tv := &SingleFileEntryVisitor{}
		// 使用 ast.Walk 函数遍历抽象语法树（AST）
		ast.Walk(tv, f)
		file := tv.Get()
		assert.Equal(t, tc.want, file)
	}
}
