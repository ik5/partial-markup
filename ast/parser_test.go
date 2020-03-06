package ast

import (
	"reflect"
	"testing"
)

type _checkStruct struct {
	input    string
	expected *AST
}

func TestSimpleTags(t *testing.T) {

	simpleTags := []_checkStruct{
		_checkStruct{
			input: "<",
			expected: &AST{
				Symbols: []Symbol{
					{
						Chars: []Char{
							Char{
								Pos:      0,
								Char:     []byte("<"),
								CharType: CharTypeOpeningTag,
							},
						},
						SymbolType: SymbolTypeTagOpen,
					},
				},
			},
		},
	}

	for _, astTag := range simpleTags {
		ast, err := Parser(astTag.input)
		if err != nil {
			t.Errorf("Error while running input: %s: %s", astTag.input, err)
		}

		if !reflect.DeepEqual(*astTag.expected, *ast) {
			t.Errorf("Expected: %+v, got: %+v", *astTag.expected, *ast)
		}
	}
}
