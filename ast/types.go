package ast

// CharType holds char type for the AST
type CharType int

// SymbolType holds information about a given block
type SymbolType int

// Char holds information regarding a given char
type Char struct {
	Pos      int      // Position of the "char". Rune position, rather than byte position
	Char     []byte   // The character that can be a multi-byte
	CharType CharType // The type/family of char that was found
}

// Symbol holds information about a block
type Symbol struct {
	Chars      []Char     // A slice of ASTChar that generate a symbol
	SymbolType SymbolType // The type of symbol that was found
}

// AST holds information about each char and it's requirements
type AST struct {
	Symbols []Symbol // A list of symbols that generate a document
}

// The type of char that currently processed
const (
	CharTypeUnsupported     CharType = -1       // Unsupported/known char
	CharTypeText            CharType = iota - 1 // part of text
	CharTypeOpeningTag                          // <
	CharTypeClosingTag                          // >
	CharTypeSpace                               // Space char
	CharTypeTagName                             // the tag name
	CharTypePropertyName                        // The name of the property
	CharTypePropertyEqual                       // =
	CharTypePropertyOpening                     // Double Quote
	CharTypePropertyClosing                     // Double Quote
	CharTypePropertyValue                       // The content of the property
	CharTypeTagTextContent                      // The inner text content of a tag
	CharTypeTagClosingSlash                     // the "/" char
)

// The type of symbol
const (
	SymbolTypeUnsupported   SymbolType = -1       // Unknown/supported tag
	SymbolTypeText          SymbolType = iota - 1 // External text of a tag
	SymbolTypeTagOpen                             // <tag>
	SymbolTypeTagName                             // tag
	SymbolTypePropertyName                        // a property name
	SymbolTypePropertyValue                       // property value
	SymbolTypeTagInnerText                        // tag inner text
	SymbolTypeTagClosed                           // closed tag </tag>
)
