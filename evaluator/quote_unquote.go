package evaluator

import (
	"Monkey/ast"
	"Monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
