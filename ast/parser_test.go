package ast

import (
	"fmt"
	"reflect"
	"testing"
)

type _tagTestStruct struct {
	desc     string
	input    string
	expected *AST
}

var simpleTags = []_tagTestStruct{
	_tagTestStruct{
		desc:  "opening tag char",
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
	_tagTestStruct{
		desc:  "opening tag char and tag name",
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
	_tagTestStruct{
		desc:  "opening tag char with tag name and number inside",
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
	_tagTestStruct{
		desc:  "Opening tag char and tag name with underscore",
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
	_tagTestStruct{
		desc:  "opening tag char, closing char and tag name",
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
	_tagTestStruct{
		desc:  "opening tag char, tag name and close tag char",
		input: "<tag>",
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
				Symbol{
					SymbolType: SymbolTypeTagEnd,
					Chars: []Char{
						Char{
							Pos:      4,
							Char:     []byte(string(">")),
							CharType: CharTypeClosingTag,
						},
					},
				},
			},
		},
	},
	_tagTestStruct{
		desc:  "opening tag char, tag name, space and close tag char ",
		input: "<tag >",
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
				Symbol{
					SymbolType: SymbolTypeTagWhitespace,
					Chars: []Char{
						Char{
							Pos:      4,
							Char:     []byte(string(" ")),
							CharType: CharTypeSpace,
						},
					},
				},
				Symbol{
					SymbolType: SymbolTypeTagEnd,
					Chars: []Char{
						Char{
							Pos:      5,
							Char:     []byte(string(">")),
							CharType: CharTypeClosingTag,
						},
					},
				},
			},
		},
	},
}

func TestSimpleTags(t *testing.T) {

	for _, astTag := range simpleTags {
		if testing.Verbose() {
			fmt.Printf("Going on '%s'\n\t%s\n", astTag.input, astTag.desc)
		}

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
