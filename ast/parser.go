package ast

const (
	openingTag = '<'
	closingTag = '>'
	space      = ' '
	tab        = '\t'
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
						Char:     []byte(string(textRune[pos])),
						CharType: CharTypeOpeningTag,
					},
				},
				SymbolType: SymbolTypeTagOpen,
			})

			var tagName []Char
		tagNameLoop:
			// Get the name of the tag
			for pos++; pos < textLen; pos++ {
				switch textRune[pos] {
				case space, tab, closingTag: // white space
					break tagNameLoop
				default:
					if (textRune[pos] >= 'a' && textRune[pos] <= 'z') ||
						(textRune[pos] >= 'A' && textRune[pos] <= 'Z') ||
						(textRune[pos] >= '0' && textRune[pos] <= '9') ||
						(textRune[pos] == '_') {
						tagName = append(tagName, Char{
							Pos:  pos,
							Char: []byte(string(textRune[pos])),
						})
					}
				}
			}
			if len(tagName) > 0 {
				result.Symbols = append(result.Symbols, Symbol{
					Chars:      tagName,
					SymbolType: SymbolTypeTagName,
				})
			}
		default:
		}
	}

	return &result, nil
}
