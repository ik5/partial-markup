package ast

const (
	openingTag = '<'
	closingTag = '>'
	space      = ' '
	tab        = '\t'
	closeSlash = '/'
)

// Parser parse each given char and generate a list of symbols
func Parser(text string) (*AST, error) {
	var result AST

	textRune := []rune(text)
	textLen := len(textRune)
	for pos := 0; pos < textLen; pos++ {
		switch textRune[pos] {
		case space:
			// TODO: Make this part smarter, if there is a text, it should not enter
			// here
			result.Symbols = append(result.Symbols, Symbol{
				SymbolType: SymbolTypeTagWhitespace,
				Chars: []Char{
					Char{
						Pos:      pos,
						Char:     []byte(string(textRune[pos])),
						CharType: CharTypeSpace,
					},
				},
			})
		case closingTag:
			result.Symbols = append(result.Symbols, Symbol{
				SymbolType: SymbolTypeTagEnd,
				Chars: []Char{
					Char{
						Pos:      pos,
						Char:     []byte(string(textRune[pos])),
						CharType: CharTypeClosingTag,
					},
				},
			})
			continue
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
				case closeSlash:
					result.Symbols = append(result.Symbols, Symbol{
						Chars: []Char{
							Char{
								Pos:      pos,
								Char:     []byte(string(textRune[pos])),
								CharType: CharTypeTagClosingSlash,
							},
						},
						SymbolType: SymbolTypeTagClosed,
					})
					continue
				case space, tab: // white space
					pos-- // do not handle this code here
					break tagNameLoop
				case closingTag:
					pos-- // do not handle this tag here
					break tagNameLoop
				default:
					if (textRune[pos] >= 'a' && textRune[pos] <= 'z') ||
						(textRune[pos] >= 'A' && textRune[pos] <= 'Z') ||
						(textRune[pos] >= '0' && textRune[pos] <= '9') ||
						(textRune[pos] == '_') {
						tagName = append(tagName, Char{
							Pos:      pos,
							Char:     []byte(string(textRune[pos])),
							CharType: CharTypeTagName,
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
