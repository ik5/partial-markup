package partialmarkup

const (
	openingTag = '<'
)

// ASTParser parse each given char and generate a list of symbols
func ASTParser(text string) (*AST, error) {
	var result AST

	textRune := []rune(text)
	textLen := len(textRune)
	for pos := 0; pos < textLen; pos++ {
		switch textRune[pos] {
		case openingTag:
			result.Symbols = append(result.Symbols, ASTSymbol{
				Chars: []ASTChar{
					ASTChar{
						Pos:      pos,
						Char:     []byte{byte(textRune[pos])},
						CharType: ASTCharTypeOpeningTag,
					},
				},
				SymbolType: ASTSymbolTypeTagOpen,
			})

			for pos++; pos < textLen; pos++ {
				switch textRune[pos] {

				}
			}
		default:
		}
	}

	return &result, nil
}
