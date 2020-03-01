package partialmarkup

// Document holds the entire elements of Blocks
type Document struct {
	RawDocument string
	Root        *Block
}
