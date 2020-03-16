package ast

import "testing"

type _charTypeTestStruct struct {
	input    CharType
	expected string
}

type _symboleTypeTestStruct struct {
	input    SymbolType
	expected string
}

var charTypes = []_charTypeTestStruct{
	_charTypeTestStruct{
		input:    CharType(-10),
		expected: "unknown (-10)",
	},
	_charTypeTestStruct{
		input:    CharTypeUnsupported,
		expected: "unsupported",
	},
	_charTypeTestStruct{
		input:    CharTypeText,
		expected: "text",
	},
	_charTypeTestStruct{
		input:    CharTypeOpeningTag,
		expected: "tag open",
	},
	_charTypeTestStruct{
		input:    CharTypeClosingTag,
		expected: "tag close",
	},
	_charTypeTestStruct{
		input:    CharTypeSpace,
		expected: "space",
	},
	_charTypeTestStruct{
		input:    CharTypeTagName,
		expected: "tag name",
	},
	_charTypeTestStruct{
		input:    CharTypePropertyName,
		expected: "property name",
	},
	_charTypeTestStruct{
		input:    CharTypePropertyOpening,
		expected: "property value opening",
	},
	_charTypeTestStruct{
		input:    CharTypePropertyClosing,
		expected: "property value closing",
	},
	_charTypeTestStruct{
		input:    CharTypePropertyValue,
		expected: "property value",
	},
	_charTypeTestStruct{
		input:    CharTypeTagTextContent,
		expected: "#text",
	},
	_charTypeTestStruct{
		input:    CharTypeTagClosingSlash,
		expected: "closing slash",
	},
}

var symbolTypes = []_symboleTypeTestStruct{
	_symboleTypeTestStruct{
		input:    SymbolType(-10),
		expected: "unknown (-10)",
	},
	_symboleTypeTestStruct{
		input:    SymbolTypeText,
		expected: "text",
	},
	_symboleTypeTestStruct{
		input:    SymbolTypeUnsupported,
		expected: "unsupported",
	},
	_symboleTypeTestStruct{
		input:    SymbolTypeTagOpen,
		expected: "tag open",
	},
	_symboleTypeTestStruct{
		input:    SymbolTypeTagEnd,
		expected: "tag end",
	},
	_symboleTypeTestStruct{
		input:    SymbolTypeTagName,
		expected: "tag name",
	},
	_symboleTypeTestStruct{
		input:    SymbolTypePropertyName,
		expected: "property name",
	},
	_symboleTypeTestStruct{
		input:    SymbolTypePropertyValue,
		expected: "property value",
	},
	_symboleTypeTestStruct{
		input:    SymbolTypeTagInnerText,
		expected: "inner text",
	},
	_symboleTypeTestStruct{
		input:    SymbolTypeTagClosed,
		expected: "tag closed",
	},
}

func TestCharTypeString(t *testing.T) {
	for count, char := range charTypes {
		if char.input.String() != char.expected {
			t.Errorf("Test #%d expected: '%s', got '%s' (%d)",
				count, char.expected, char.input, char.input)
		}
	}
}

func TestSymbolTypeString(t *testing.T) {
	for count, symbol := range symbolTypes {
		if symbol.input.String() != symbol.expected {
			t.Errorf("Test #%d expected: '%s', got '%s' (%d)",
				count, symbol.expected, symbol.input, symbol.input)
		}
	}
}
