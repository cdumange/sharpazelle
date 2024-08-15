package parser

import (
	"context"
	"os"
	"testing"

	"github.com/cdumange/sharpazelle/utils"
	"github.com/stretchr/testify/assert"
)

func Test_parserFolder(t *testing.T) {
	ctx := context.Background()

	for name, values := range map[string]struct {
		folder   []os.DirEntry
		expected Folder
	}{
		"../tests/simple": {
			folder: []os.DirEntry{},
			expected: Folder{
				Folders: []string{"Tests", "lib", "webapp"},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, values.expected, utils.Must(ParseFolder(ctx, name)))
		})
	}
}
