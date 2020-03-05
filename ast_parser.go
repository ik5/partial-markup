package partialmarkup

const (
	openingTag byte = 0x3C // <
)

// ASTParser parse each given char and generate a list of symbols
func ASTParser(text string) (*AST, error) {
	var result AST

	textLen := len(text)
	for pos := 0; pos < textLen; pos++ {
		switch text[pos] {
		case openingTag:
			var chars []ASTChar
			var symbol ASTSymbol
			chars = append(chars, ASTChar{
				Pos:      pos,
				Char:     []byte{openingTag},
				CharType: ASTCharTypeOpeningTag,
			})
			symbol = ASTSymbol{
				Chars:      chars,
				SymbolType: ASTSymbolTypeTagOpen,
			}
			result.Symbols = append(result.Symbols, symbol)

			for pos = pos; pos < textLen; pos++ {
			}
		default:
		}
	}

	return &result, nil
}
