package partialmarkup

// BlockType is a data type for what is the given block is
type BlockType int

// Types of blocks That found
const (
	blockTypeStart   BlockType = -1
	BlockTypeUnknown BlockType = 0
	BlockTypeText    BlockType = iota + 1
	BlockTypeTag
	BlockTypeProperty
	BlockTypeInnerText
)

// for internal use for not allocating many strings
const (
	textBlockTypeUnknown   = "unknown"
	textBlockTypeText      = "text"
	textBlockTypeTag       = "tag"
	textBlockTypeProperty  = "property"
	textBlockTypeInnerText = "#text"
)
