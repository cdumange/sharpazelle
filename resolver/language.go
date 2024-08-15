package resolver

import (
	"context"
	"fmt"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
	"github.com/cdumange/sharpazelle/parser"
	"github.com/cdumange/sharpazelle/utils"
)

type cSharp struct {
	resolve.Resolver
	config.Configurer

	fnParser
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

type fnParser = func(ctx context.Context, folder string) (parser.Folder, error)

// GenerateRules implements language.Language.
func (c cSharp) GenerateRules(args language.GenerateArgs) (res language.GenerateResult) {
	ctx := context.Background()
	if !utils.IsTreatableFolder(args.Dir) {
		return language.GenerateResult{}
	}

	_, err := c.fnParser(ctx, args.Dir)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(args.Dir))
	return
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

// NewLanguage returns the C# language implementation for gazelle.
func NewLanguage() language.Language {
	return cSharp{
		Resolver:   NewResolver(),
		Configurer: NewConfigurer(),
		fnParser:   parser.ParseFolder,
	}
}
