package resolver

import (
	"testing"

	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/rule"
	"github.com/stretchr/testify/assert"
)

func Test_GenerateRules(t *testing.T) {
	cs := NewLanguage()

	for folder, expected := range map[string]language.GenerateResult{
		"../parser/test_data/simple/webapp": {
			Gen:     []*rule.Rule{rule.NewRule(RuleBinary, "../parser/test_data/simple/webapp")},
			Empty:   []*rule.Rule(nil),
			Imports: []interface{}{"./**"},
		},
	} {
		t.Run(folder, func(t *testing.T) {
			v := cs.GenerateRules(language.GenerateArgs{
				Dir: folder,
			})

			assert.Equal(t, expected, v)
		})
	}
}
