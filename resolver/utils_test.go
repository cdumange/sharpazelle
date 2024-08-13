package resolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isTreatableFolder(t *testing.T) {
	for folder, should := range map[string]bool{
		"/home/cdumange/code/zenika/sharpazelle/tests/":                                true,
		"/home/cdumange/code/zenika/sharpazelle/tests/simple/webapp/":                  true,
		"/home/cdumange/code/zenika/sharpazelle/tests/simple/webapp/bin/Debug/net8.0/": false,
		"/home/cdumange/code/zenika/sharpazelle/tests/simple/webapp/bin":               false,
	} {
		t.Run(folder, func(t *testing.T) {
			assert.Equal(t, should, isTreatableFolder(folder))
		})
	}
}
