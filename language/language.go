package language

import (
	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"

	cfg "github.com/cdumange/sharpazelle/config"
	"github.com/cdumange/sharpazelle/resolver"
)

type cSharp struct {
	resolve.Resolver
	config.Configurer
}

// Loads implements language.Language.
func (c cSharp) Loads() []rule.LoadInfo {
	panic("unimplemented")
}

// Fix implements language.Language.
func (cSharp) Fix(c *config.Config, f *rule.File) {
	panic("unimplemented")
}

// GenerateRules implements language.Language.
func (c cSharp) GenerateRules(args language.GenerateArgs) language.GenerateResult {
	panic("unimplemented")
}

// Kinds implements language.Language.
func (c cSharp) Kinds() map[string]rule.KindInfo {
	panic("unimplemented")
}

// CSharp returns the C# language for gazelle.
func CSharp() language.Language {
	return cSharp{
		Resolver:   resolver.NewResolver(),
		Configurer: cfg.NewConfigurer(),
	}
}
