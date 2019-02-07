package cscope

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymbolsParsing(t *testing.T) {
	test := []byte("bs_file call_manager_init 573 void call_manager_init(){")
	symbols, err := parseSymbols(bytes.NewReader(test))
	assert.Nil(t, err)
	assert.Equal(t, len(symbols), 1)
}

func TestCantParse(t *testing.T) {
	test := []byte("bs_file call_manager_init573 void call_manager_init(){")
	symbols, err := parseSymbols(bytes.NewReader(test))
	assert.Error(t, err)
	assert.Nil(t, symbols)
}

func TestMultiLineSymbols(t *testing.T) {
	symbol := Symbol{
		File:        "bs_file",
		Name:        "call_manager_init",
		LineNumber:  573,
		LineContent: "void call_manager_init(){",
	}
	test := []byte("bs_file call_manager_init 573 void call_manager_init(){\r\nbs_file call_manager_init 573 void call_manager_init(){")
	symbols, err := parseSymbols(bytes.NewReader(test))
	assert.Nil(t, err)
	assert.Equal(t, symbols[0], symbol)
	assert.Equal(t, symbols[1], symbol)
}

func TestNotFoundDatabase(t *testing.T) {
	_, err := NewCscope("randomname")
	assert.Error(t, err)
}

func TestUnsupportedCommand(t *testing.T) {
	cscope, err := NewCscope("cscope.out")
	assert.NotNil(t, cscope)
	symbols, err := cscope.Cmd(30, "random")
	assert.Error(t, err)
	assert.Nil(t, symbols)
}

func TestCscopeMain(t *testing.T) {
	cscope, err := NewCscope("cscope.out")
	assert.NotNil(t, cscope)
	symbol := Symbol{
		File:        "main.c",
		Name:        "main",
		LineNumber:  3,
		LineContent: "int main()",
	}
	symbols, err := cscope.Cmd(0, "main")
	assert.Nil(t, err)
	assert.Equal(t, len(symbols), 1)
	assert.Equal(t, symbols[0], symbol)
}
