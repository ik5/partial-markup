package partialmarkup

import "testing"

var basicFixturesValid = []string{
	"simple text",
}

func TestBasicDocumentValid(t *testing.T) {
	for i, fixture := range basicFixturesValid {
		document, err := Parse(fixture)

		if err != nil {
			t.Errorf("Cannot parse fixture #%d: %s", i, err)
		}

		if document == nil {
			t.Errorf("Document is nil without an error for fixture #%d", i)
		}
	}
}
