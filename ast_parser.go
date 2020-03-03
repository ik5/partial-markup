package partialmarkup

// ASTParser parse each given char and generate a list of symbols
func ASTParser(text string) (*AST, error) {
	var result AST

	textLen := len(text)
	for pos := 0; pos < textLen; pos++ {
		switch text[pos] {
		case "<":
			var chars []ASTChar
			chars = append(chars, ASTChar{
				Pos:      pos,
				Char:     []byte(text[pos]),
				CharType: ASTCharTypeOpeningTag,
			})

			for pos = pos; pos < textLen; pos++ {
			}
		default:
		}
	}

	return &result, nil
}
