package evelyn

//go:generate antlr4 -Dlanguage=Go -Werror -package parser -o parser Evelyn.g4

import (
	"github.com/evenfound/even-go/node/cmd/evec/compiler"
	"github.com/evenfound/even-go/node/cmd/evec/config"
	"github.com/evenfound/even-go/node/cmd/evec/implementation/evelyn/parser"
	"github.com/evenfound/even-go/node/cmd/evec/tool"
	"io/ioutil"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// evelynCompiler is an implementation of the compiler.Interface.
type evelynCompiler struct {
}

// Compiler translates a source code from a file into binary bytecode.
func (e evelynCompiler) Compile(filename string) (compiler.Bytecode, error) {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, tool.Wrap(err, "open file")
	}

	lexer := parser.NewEvelynLexer(input)
	// suppress noise from lexer:
	lexer.RemoveErrorListeners()

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewEvelynParser(stream)
	errList := newErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(errList)
	p.BuildParseTrees = true

	tree := p.SourceFile()
	if !errList.Empty() {
		return nil, tool.NewError(errList.FirstMessage())
	}

	tmpfile, err := ioutil.TempFile("", "evelyn.*"+config.TengoExt)
	if err != nil {
		return nil, tool.Wrap(err, "tempfile")
	}
	//defer func() { tool.Ignore(os.Remove(tmpfile.Name())) }()

	antlr.ParseTreeWalkerDefault.Walk(newListener(tmpfile), tree)

	return nil, nil
}
