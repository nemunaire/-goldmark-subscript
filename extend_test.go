package subscript

import (
	"log"
	"os"
	"testing"

	"github.com/yuin/goldmark"
)

func TestAttributes(t *testing.T) {
	source := []byte(`
H~2~O is a liquid.  2^10^ is 1024.
`)

	var md = goldmark.New(Enable)
	err := md.Convert(source, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
