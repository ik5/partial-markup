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
							Pos:      1,
							Char:     []byte(string("t")),
							CharType: CharTypeTagName,
						},
						Char{
							Pos:      2,
							Char:     []byte(string("a")),
							CharType: CharTypeTagName,
						},
						Char{
							Pos:      3,
							Char:     []byte(string("g")),
							CharType: CharTypeTagName,
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
							Pos:      1,
							Char:     []byte(string("t")),
							CharType: CharTypeTagName,
						},
						Char{
							Pos:      2,
							Char:     []byte(string("5")),
							CharType: CharTypeTagName,
						},
						Char{
							Pos:      3,
							Char:     []byte(string("g")),
							CharType: CharTypeTagName,
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
							Pos:      1,
							Char:     []byte(string("t")),
							CharType: CharTypeTagName,
						},
						Char{
							Pos:      2,
							Char:     []byte(string("_")),
							CharType: CharTypeTagName,
						},
						Char{
							Pos:      3,
							Char:     []byte(string("g")),
							CharType: CharTypeTagName,
						},
					},
					SymbolType: SymbolTypeTagName,
				},
			},
		},
	},
	_checkStruct{
		input: "</tag",
		expected: &AST{
			Symbols: []Symbol{
				Symbol{
					SymbolType: SymbolTypeTagOpen,
					Chars: []Char{
						Char{
							Pos:      0,
							Char:     []byte(string("<")),
							CharType: CharTypeOpeningTag,
						},
					},
				},
				Symbol{
					SymbolType: SymbolTypeTagClosed,
					Chars: []Char{
						Char{
							Pos:      1,
							Char:     []byte(string("/")),
							CharType: CharTypeTagClosingSlash,
						},
					},
				},
				Symbol{
					SymbolType: SymbolTypeTagName,
					Chars: []Char{
						Char{
							Pos:      2,
							Char:     []byte(string("t")),
							CharType: CharTypeTagName,
						},
						Char{
							Pos:      3,
							Char:     []byte(string("a")),
							CharType: CharTypeTagName,
						},
						Char{
							Pos:      4,
							Char:     []byte(string("g")),
							CharType: CharTypeTagName,
						},
					},
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
			t.Errorf("input: %s\nExpected:\n%+v\ngot:\n%+v",
				astTag.input, *astTag.expected, *ast)
		}
	}
}
