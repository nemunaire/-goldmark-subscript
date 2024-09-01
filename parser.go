package subscript

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type subscriptDelimiterProcessor struct {
}

func (p *subscriptDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '~'
}

func (p *subscriptDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *subscriptDelimiterProcessor) OnMatch(consumes int) ast.Node {
	return &Node{}
}

var defaultEmphasisDelimiterProcessor = &subscriptDelimiterProcessor{}

type subscriptParser struct {
}

var defaultSubscriptParser = &subscriptParser{}

func (s *subscriptParser) Trigger() []byte {
	return []byte{'~'}
}

func (s *subscriptParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 1, defaultEmphasisDelimiterProcessor)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}
