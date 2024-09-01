package subscript

import (
	"fmt"
	"strings"

	"github.com/yuin/goldmark/ast"
)

// Kind is the kind of hashtag AST nodes.
var Kind = ast.NewNodeKind("Subscript")

// Node is a parsed Subscript node.
type Node struct {
	ast.BaseInline
}

func (*Node) Kind() ast.NodeKind { return Kind }

func (n *Node) Dump(src []byte, level int) {
	fmt.Printf("%ssubscript: \"%s\"\n", strings.Repeat("    ", level), n.Text(src))
}
