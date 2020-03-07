# partial-markup

[![GoDoc](https://godoc.org/github.com/ik5/partial-markup?status.svg)](https://godoc.org/github.com/ik5/partial-markup)

A parser for partial markup text using Golang.

The current implementation is a Proof of concept for having markup text inside
a normal text.

A partial markup looks as follows:

    Hello <span>world</span>.

The parser at this stage aim to support a simple text that can have one or more
markup structure.
The structure can have one single property, and text inside, there's a need to
support multiple properties per tag.


The structure of the markup **does not** support nested markup, and was not
designed to support such structure.

The structure does not support a single tag closing itself (as in XML).



