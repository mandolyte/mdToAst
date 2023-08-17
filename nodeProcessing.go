/*
 * Markdown to PDF Converter
 * Available at http://github.com/mandolyte/mdtopdf
 *
 * Copyright © 2018 Cecil New <cecil.new@gmail.com>.
 * Distributed under the MIT License.
 * See README.md for details.
 *
 * Dependencies
 * This package depends on two other packages:
 *
 * Blackfriday Markdown Processor
 *   Available at http://github.com/russross/blackfriday
 *
 * fpdf - a PDF document generator with high level support for
 *   text, drawing and images.
 *   Available at https://github.com/go-pdf/fpdf
 */

package mdToAst

import (
	"github.com/gomarkdown/markdown/ast"
)

func processText(node *ast.Text) {

	switch node.Parent.(type) {

	case *ast.Link:
	case *ast.Heading:
	case *ast.TableCell:
	case *ast.BlockQuote:
	default:
	}
}

func processCodeblock(node ast.CodeBlock) {
	// lines := strings.Split(string(node.Literal), "\n")
}

func processList(node ast.List, entering bool) {
	// kind := unordered
	// if node.ListFlags&ast.ListTypeOrdered != 0 {
	// 	kind = ordered
	// }
	// if node.ListFlags&ast.ListTypeDefinition != 0 {
	// 	kind = definition
	// }
	// if entering {
	// } else {
	// }
}

// func isListItem(node ast.Node) bool {
// 	_, ok := node.(*ast.ListItem)
// 	return ok
// }

func processItem(node ast.ListItem, entering bool) {
	// if entering {
	// } else {
	// }
}

func processEmph(node ast.Node, entering bool) {
	// if entering {
	// } else {
	// }
}

func processStrong(node ast.Node, entering bool) {
	// if entering {
	// } else {
	// }
}

func processLink(node ast.Link, entering bool) {
	// destination := string(node.Destination)
	// if entering {
	// } else {
	// }
}

func processDel(node ast.Del, entering bool) {

}

func processImage(node ast.Image, entering bool) {
	// while this has entering and leaving states, it doesn't appear
	// to be useful except for other markup languages to close the tag
	// if entering {
	// } else {
	// }
}

func processCode(node ast.Node) {
}

func processParagraph(node *ast.Paragraph, entering bool) {
	// if entering {
	// } else {
	// }
}

func processBlockQuote(node ast.Node, entering bool) {
	// if entering {
	// } else {
	// }
}

func processHeading(node ast.Heading, entering bool) {
	// if entering {
	// } else {
	// }
}

func processHorizontalRule(node ast.Node) {
}

func processHTMLBlock(node ast.Node) {
}

func processTable(node ast.Node, entering bool) {
	// if entering {
	// } else {
	// }
}

func processTableHead(node ast.Node, entering bool) {
	// if entering {
	// } else {
	// }
}

func processTableBody(node ast.Node, entering bool) {
	// if entering {
	// } else {
	// }
}

func processTableRow(node ast.Node, entering bool) {
	// if entering {
	// } else {
	// }
}

func processTableCell(node ast.TableCell, entering bool) {
	// if entering {
	// } else {
	// }
}
