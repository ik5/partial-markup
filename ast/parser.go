package ast

const (
	openingTag = '<'
)

// Parser parse each given char and generate a list of symbols
func Parser(text string) (*AST, error) {
	var result AST

	textRune := []rune(text)
	textLen := len(textRune)
	for pos := 0; pos < textLen; pos++ {
		switch textRune[pos] {
		case openingTag:
			result.Symbols = append(result.Symbols, Symbol{
				Chars: []Char{
					Char{
						Pos:      pos,
						Char:     []byte{byte(textRune[pos])},
						CharType: CharTypeOpeningTag,
					},
				},
				SymbolType: SymbolTypeTagOpen,
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
