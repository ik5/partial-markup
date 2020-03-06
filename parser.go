package partialmarkup

import (
	"errors"
	"strings"

	"github.com/ik5/parital-markup/ast"
)

// Parse parses a given document and return a scan tree of Document
// if fails, an error is returned
func Parse(document string) (*Document, error) {

	if strings.TrimSpace(document) == "" {
		return nil, errors.New("document cannot be empty")
	}

	result := &Document{
		RawDocument: document,
		Root:        &Block{},
	}

	// TODO parse *AST and convert it to Document
	_, err := ast.Parser(document)

	return result, err
}
