package partialmarkup

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// String returns the name of the BlockType.
// If the value is not part of the parser, "unknown" is returned
func (b BlockType) String() string {
	str := textBlockTypeUnknown
	switch b {
	case BlockTypeText:
		str = textBlockTypeText
	case BlockTypeTag:
		str = textBlockTypeTag
	case BlockTypeProperty:
		str = textBlockTypeProperty
	case BlockTypeInnerText:
		str = textBlockTypeInnerText
	}

	return str
}

// StrToBlockType convert a string to BlockType
// If string is not BlockType, then "BlockTypeUnknown" is returned
func StrToBlockType(str string) BlockType {
	result := BlockTypeUnknown
	switch strings.TrimSpace(strings.ToLower(str)) {
	case textBlockTypeText:
		result = BlockTypeText
	case textBlockTypeTag:
		result = BlockTypeTag
	case textBlockTypeProperty:
		result = BlockTypeProperty
	case textBlockTypeInnerText:
		result = BlockTypeInnerText
	}

	return result
}

// Value for the database interface
func (b BlockType) Value() (driver.Value, error) {
	return int(b), nil
}

// Scan a value from database to BlockType
func (b *BlockType) Scan(value interface{}) error {
	if value == nil {
		*b = BlockTypeUnknown
		return nil
	}
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		*b = BlockType(v.Int())
		return nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		*b = BlockType(v.Uint())
		return nil
	case reflect.Float32, reflect.Float64:
		*b = BlockType(v.Float())
		return nil
	case reflect.String:
		val, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return err
		}
		*b = BlockType(val)
		return nil
	}

	return fmt.Errorf("Unsupported %T for data %+v as BlockType", value, value)
}
