package test

import (
	"fmt"
	"pi_exchange/io"
	"testing"
)

func TestFullFlow(t *testing.T) {
	templateReader, err := io.NewJsonTemplateReader("data/template.json")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(templateReader)
}
