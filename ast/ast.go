package ast

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() string
}

type Expression interface {
	Node
	expressionNode() string
}
