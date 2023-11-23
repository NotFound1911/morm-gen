package main

import (
	_ "embed"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	file string
)

func main() {
	flag.StringVar(&file, "f", "", "input file")
	flag.Parse()
	fileName := filepath.Base(file)
	dir := filepath.Dir(file)
	idx := strings.LastIndexByte(fileName, '.')
	dst := filepath.Join(dir, fileName[:idx]+"_gen.go")
	f, err := os.Create(dst)
	if err != nil {
		fmt.Printf("create file failed,err:%v\n", err)
		return
	}
	err = gen(f, file)
	if err != nil {
		fmt.Printf("gen file failed,err:%v\n", err)
		return
	}
	fmt.Printf("gen file:%v successful!\n", dst)
}

//go:embed tpl.gohtml
var genOrm string

func gen(writer io.Writer, src string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, src, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	s := &SingleFileEntryVisitor{}
	ast.Walk(s, f)
	file := s.Get()
	tpl := template.New("gen-orm")
	tpl, err = tpl.Parse(genOrm)
	if err != nil {
		return err
	}
	return tpl.Execute(writer, Data{
		File: file,
		Ops:  []string{"LT", "GT", "EQ"},
	})
}

type Data struct {
	*File
	Ops []string
}
