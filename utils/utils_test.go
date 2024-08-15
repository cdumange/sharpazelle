package utils

import (
	"errors"
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
			assert.Equal(t, should, IsTreatableFolder(folder))
		})
	}
}

func TestMust(t *testing.T) {
	testValue := "test"
	f := func(err error) (string, error) {
		return testValue, err
	}

	t.Run("ok", func(t *testing.T) {
		assert.NotPanics(t, func() {
			v := Must[string](f(nil))
			assert.Equal(t, testValue, v)
		})
	})

	t.Run("ko", func(t *testing.T) {
		assert.Panics(t, func() {
			v := Must[string](f(errors.New("an error")))
			assert.Equal(t, testValue, v)
		})
	})
}
