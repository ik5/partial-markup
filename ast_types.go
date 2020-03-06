package partialmarkup

// ASTCharType holds char type for the AST
type ASTCharType int

// ASTSymbolType holds information about a given block
type ASTSymbolType int

// ASTChar holds information regarding a given char
type ASTChar struct {
	Pos      int         // Position of the "char". Rune position, rather than byte position
	Char     []byte      // The character that can be a multi-byte
	CharType ASTCharType // The type/family of char that was found
}

// ASTSymbol holds information about a block
type ASTSymbol struct {
	Chars      []ASTChar     // A slice of ASTChar that generate a symbol
	SymbolType ASTSymbolType // The type of symbol that was found
}

// AST holds information about each char and it's requirements
type AST struct {
	Symbols []ASTSymbol // A list of symbols that generate a document
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
