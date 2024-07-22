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

// CSharp rules constants
const (
	RuleBinary  = "csharp_binary"
	RuleLibrary = "csharp_library"
	RuleTest    = "csharp_test"

	attributeName = "name"
	attributeDeps = "deps"
	attributeSrcs = "srcs"
)

var usedCSharpRules = []string{
	RuleBinary,
	RuleLibrary,
	RuleTest,
}

// Loads implements language.Language.
func (c cSharp) Loads() []rule.LoadInfo {
	return []rule.LoadInfo{
		{
			Name:    "@rules_dotnet//dotnet:defs.bzl",
			Symbols: usedCSharpRules,
		},
	}
}

// Fix implements language.Language.
func (cSharp) Fix(c *config.Config, f *rule.File) {
}

// GenerateRules implements language.Language.
func (c cSharp) GenerateRules(args language.GenerateArgs) language.GenerateResult {
	panic("unimplemented")
}

// Kinds implements language.Language.
func (c cSharp) Kinds() map[string]rule.KindInfo {
	return map[string]rule.KindInfo{
		RuleBinary: {
			MatchAny:   true,
			MatchAttrs: []string{attributeName},
			NonEmptyAttrs: map[string]bool{
				attributeName: true,
				attributeDeps: true,
				attributeSrcs: true,
			},
		},
		RuleLibrary: {
			MatchAny:   true,
			MatchAttrs: []string{attributeName},
			NonEmptyAttrs: map[string]bool{
				attributeName: true,
				attributeDeps: true,
				attributeSrcs: true,
			},
		},
		RuleTest: {
			MatchAny:   true,
			MatchAttrs: []string{attributeName},
			NonEmptyAttrs: map[string]bool{
				attributeName: true,
				attributeDeps: true,
				attributeSrcs: true,
			},
		},
	}
}

// CSharp returns the C# language for gazelle.
func CSharp() language.Language {
	return cSharp{
		Resolver:   resolver.NewResolver(),
		Configurer: cfg.NewConfigurer(),
	}
}
