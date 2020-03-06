package ast

import (
	"reflect"
	"testing"
)

type _checkStruct struct {
	input    string
	expected *AST
}

var simpleTags = []_checkStruct{
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
	_checkStruct{
		input: "<tag",
		expected: &AST{
			Symbols: []Symbol{
				Symbol{
					Chars: []Char{
						Char{
							Pos:      0,
							Char:     []byte("<"),
							CharType: CharTypeOpeningTag,
						},
					},
					SymbolType: SymbolTypeTagOpen,
				},
				Symbol{
					Chars: []Char{
						Char{
							Pos:  1,
							Char: []byte(string("t")),
						},
						Char{
							Pos:  2,
							Char: []byte(string("a")),
						},
						Char{
							Pos:  3,
							Char: []byte(string("g")),
						},
					},
					SymbolType: SymbolTypeTagName,
				},
			},
		},
	},
	_checkStruct{
		input: "<t5g",
		expected: &AST{
			Symbols: []Symbol{
				Symbol{
					Chars: []Char{
						Char{
							Pos:      0,
							Char:     []byte("<"),
							CharType: CharTypeOpeningTag,
						},
					},
					SymbolType: SymbolTypeTagOpen,
				},
				Symbol{
					Chars: []Char{
						Char{
							Pos:  1,
							Char: []byte(string("t")),
						},
						Char{
							Pos:  2,
							Char: []byte(string("5")),
						},
						Char{
							Pos:  3,
							Char: []byte(string("g")),
						},
					},
					SymbolType: SymbolTypeTagName,
				},
			},
		},
	},
	_checkStruct{
		input: "<t_g",
		expected: &AST{
			Symbols: []Symbol{
				Symbol{
					Chars: []Char{
						Char{
							Pos:      0,
							Char:     []byte("<"),
							CharType: CharTypeOpeningTag,
						},
					},
					SymbolType: SymbolTypeTagOpen,
				},
				Symbol{
					Chars: []Char{
						Char{
							Pos:  1,
							Char: []byte(string("t")),
						},
						Char{
							Pos:  2,
							Char: []byte(string("_")),
						},
						Char{
							Pos:  3,
							Char: []byte(string("g")),
						},
					},
					SymbolType: SymbolTypeTagName,
				},
			},
		},
	},
}

func TestSimpleTags(t *testing.T) {

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
