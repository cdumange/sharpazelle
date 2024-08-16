package parser

import (
	"context"
	"os"
	"testing"

	"github.com/cdumange/sharpazelle/commonerrors"
	"github.com/stretchr/testify/assert"
)

func Test_parserFolder(t *testing.T) {
	ctx := context.Background()

	for name, values := range map[string]struct {
		folder   []os.DirEntry
		expected Folder
		err      error
	}{
		"./test_data/simple/webapp": {
			folder: []os.DirEntry{},
			expected: Folder{
				Folders: []string{"Pages", "Properties"},
			},
		},
		"./test_data/simple": {
			folder:   []os.DirEntry{},
			expected: Folder{},
			err:      commonerrors.ErrNotToBeTreated,
		},
	} {
		t.Run(name, func(t *testing.T) {
			v, err := ParseFolder(ctx, name)
			assert.ErrorIs(t, err, values.err)
			assert.Equal(t, values.expected.Folders, v.Folders)
		})
	}
}
