package ast

import "fmt"

func (c CharType) String() string {
	switch c {
	case CharTypeUnsupported:
		return "unsupported"
	case CharTypeText:
		return "text"
	case CharTypeOpeningTag:
		return "tag open"
	case CharTypeClosingTag:
		return "tag close"
	case CharTypeSpace:
		return "space"
	case CharTypeTagName:
		return "tag name"
	case CharTypePropertyName:
		return "property name"
	case CharTypePropertyOpening:
		return "property value opening"
	case CharTypePropertyClosing:
		return "property value closing"
	case CharTypePropertyValue:
		return "property value"
	case CharTypeTagTextContent:
		return "#text"
	case CharTypeTagClosingSlash:
		return "closing slash"
	default:
		return fmt.Sprintf("unknown (%d)", c)
	}
}

func (s SymbolType) String() string {
	switch s {
	case SymbolTypeUnsupported:
		return "unsupported"
	case SymbolTypeText:
		return "text"
	case SymbolTypeTagOpen:
		return "tag open"
	case SymbolTypeTagEnd:
		return "tag end"
	case SymbolTypeTagName:
		return "tag name"
	case SymbolTypePropertyName:
		return "property name"
	case SymbolTypePropertyValue:
		return "property value"
	case SymbolTypeTagInnerText:
		return "inner text"
	case SymbolTypeTagClosed:
		return "tag closed"
	case SymbolTypeTagWhitespace:
		return "whitespace"
	default:
		return fmt.Sprintf("unknown (%d)", s)
	}
}
