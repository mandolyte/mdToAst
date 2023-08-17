/*
 * Markdown to AST (event) Table
 * Available at http://github.com/mandolyte/mdToAst
 *
 * Copyright © 2023 Cecil New <cecil.new@gmail.com>.
 * Distributed under the MIT License.
 * See README.md for details.
 *
 * Dependencies
 * This package depends on:
 *
 * gomarkdown Markdown Processor
 *   Available at http://github.com/gomarkdown/markdown
 *
 */

// Package mdToAst converts markdown to an Event Table
// created from the AST output of gomarkdown.
package mdToAst

import (
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

// Run takes the markdown content, parses it
func Run(content []byte) error {
	// Preprocess content by changing all CRLF to LF
	s := content
	s = markdown.NormalizeNewlines(s)

	// if unicodeTranslator != nil {
	// 	s = []byte(unicodeTranslator(string(s)))
	// }

	exts := parser.CommonExtensions // parser.OrderedListStart | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(exts)
	doc := markdown.Parse(s, p)
	_ = markdown.Render(doc, nil)

	return nil
}

// RenderNode is a default renderer of a single node of a syntax tree. For
// block nodes it will be called twice: first time with entering=true, second
// time with entering=false, so that it could know when it's working on an open
// tag and when on close. It writes the result to w.
//
// The return value is a way to tell the calling walker to adjust its walk
// pattern: e.g. it can terminate the traversal by returning Terminate. Or it
// can ask the walker to skip a subtree of this node by returning SkipChildren.
// The typical behavior is to return GoToNext, which asks for the usual
// traversal to the next node.
// (above taken verbatim from the blackfriday v2 package)
func RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	switch node := node.(type) {
	case *ast.Text:
		processText(node)
	case *ast.Softbreak:
	case *ast.Hardbreak:
	case *ast.Emph:
		processEmph(node, entering)
	case *ast.Strong:
		processStrong(node, entering)
	case *ast.Del:
		processDel(*node, entering)
	case *ast.HTMLSpan:
	case *ast.Link:
		processLink(*node, entering)
	case *ast.Image:
		processImage(*node, entering)
	case *ast.Code:
		processCode(node)
	case *ast.Document:
	case *ast.Paragraph:
		processParagraph(node, entering)
	case *ast.BlockQuote:
		processBlockQuote(node, entering)
	case *ast.HTMLBlock:
		processHTMLBlock(node)
	case *ast.Heading:
		processHeading(*node, entering)
	case *ast.HorizontalRule:
		processHorizontalRule(node)
	case *ast.List:
		processList(*node, entering)
	case *ast.ListItem:
		processItem(*node, entering)
	case *ast.CodeBlock:
		processCodeblock(*node)
	case *ast.Table:
		processTable(node, entering)
	case *ast.TableHeader:
		processTableHead(node, entering)
	case *ast.TableBody:
		processTableBody(node, entering)
	case *ast.TableRow:
		processTableRow(node, entering)
	case *ast.TableCell:
		processTableCell(*node, entering)
	default:
		// panic("Unknown node type " + reflect.TypeOf(node).Name())
	}
	return ast.GoToNext
}
