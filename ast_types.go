package partialmarkup

// ASTCharType holds char type for the AST
type ASTCharType int

// ASTSymbolType holds information about a given block
type ASTSymbolType int

// ASTChar holds information regarding a given char
type ASTChar struct {
	Pos      int
	Char     []byte
	CharType ASTCharType
}

// ASTSymbol holds information about a block
type ASTSymbol struct {
	Chars      []ASTChar
	SymbolType ASTSymbolType
}

// AST holds information about each char and it's requirements
type AST struct {
	Symbols []ASTSymbol
}

// The type of char that currently processed
const (
	ASTCharTypeUnsupported     ASTCharType = -1   // Unsupported/known char
	ASTCharTypeText            ASTCharType = iota // part of text
	ASTCharTypeOpeningTag                         // <
	ASTCharTypeClosingTag                         // >
	ASTCharTYpeSpace                              // Space char
	ASTCharTypeTagName                            // the tag name
	ASTCharTypePropertyName                       // The name of the property
	ASTCharTypePropertyEqual                      // =
	ASTCharTypePropertyOpening                    // Double Quote
	ASTCharTypePropertyClosing                    // Double Quote
	ASTCharTypePropertyValue                      // The content of the property
	ASTCharTypeTagTextContent                     // The inner text content of a tag
	ASTCharTypeTagClosingSlash                    // the "/" char
)

// The type of symbol
const (
	ASTSymbolTypeUnsupported   ASTSymbolType = -1   // Unknown/supported tag
	ASTSymbolTypeText          ASTSymbolType = iota // External text of a tag
	ASTSymbolTypeTagOpen                            // <tag>
	ASTSymbolTypeTagName                            // tag
	ASTSymbolTypePropertyName                       // a property name
	ASTSymbolTypePropertyValue                      // property value
	ASTSymbolTypeTagInnerText                       // tag inner text
	ASTSymbolTypeTagClosed                          // closed tag </tag>
)
