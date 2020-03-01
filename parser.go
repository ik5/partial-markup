package partialmarkup

import (
	"errors"
	"strings"
)

// Parse parses a given document and return a scan tree of Document
// if fails, an error is returned
func Parse(document string) (*Document, error) {

	if strings.TrimSpace(document) == "" {
		return nil, errors.New("document cannot be empty")
	}

	// docLen := len(document)

	result := &Document{
		RawDocument: document,
		Root:        &Block{},
	}

	return result, nil
}
