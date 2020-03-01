package partialmarkup

// Block holds information about specific element that was found
type Block struct {
	Type        BlockType
	RawText     string
	ElementName string
	Value       string
	Children    []*Block
}
