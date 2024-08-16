package resolver

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
	"github.com/cdumange/sharpazelle/commonerrors"
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

	sdkBase = "Microsoft.NET.Sdk"
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

	project, err := c.fnParser(ctx, args.Dir)
	if err != nil {
		if !errors.Is(err, commonerrors.ErrNotToBeTreated) {
			slog.Error(fmt.Sprintf("error parsing project %s: %s", args.Dir, err))
			return
		}
		slog.Info(fmt.Sprintf("%s does not need to have a build file.", args.Dir))
	}

	if project.Project == nil {
		return
	}

	t := ruleFromSDK(project)
	res.Gen = append(res.Gen, rule.NewRule(t, args.Dir))

	res.Imports = append(res.Imports, "./**")

	return
}

func ruleFromSDK(project parser.Folder) string {
	if project.Project.IsTestProject {
		return RuleTest
	}

	if project.Project.ProjectSDK == sdkBase {
		return RuleLibrary
	}

	return RuleBinary
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
