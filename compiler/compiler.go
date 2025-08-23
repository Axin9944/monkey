package compiler

import (
	"Monkey/ast"
	"Monkey/code"
	"Monkey/object"
)

type Compiler struct {
	// 保存生成的字节码
	instructions code.Instructions
	// 表示常量池
	constants []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

// 传输给虚拟机的内容
type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}
