package subscript

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type subscriptRenderer struct{}

func (r *subscriptRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(Kind, r.Render)
}

func (r *subscriptRenderer) Render(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n, ok := node.(*Node)
	if !ok {
		return ast.WalkStop, fmt.Errorf("unexpected node %T, expected *Node", node)
	}

	if entering {
		if err := r.enter(w, n); err != nil {
			return ast.WalkStop, err
		}
	} else {
		r.exit(w, n)
	}

	return ast.WalkContinue, nil
}

func (r *subscriptRenderer) enter(w util.BufWriter, n *Node) error {
	w.WriteString(`<sub`)
	html.RenderAttributes(w, n, nil)
	w.WriteString(`>`)
	return nil
}

func (r *subscriptRenderer) exit(w util.BufWriter, n *Node) {
	w.WriteString("</sub>")
}
