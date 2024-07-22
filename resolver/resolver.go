package resolver

import (
	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/repo"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

const (
	language = "csharp"
)

type cSharpResolver struct {
}

// Embeds implements resolve.Resolver.
func (c cSharpResolver) Embeds(r *rule.Rule, from label.Label) []label.Label {
	panic("unimplemented")
}

// Imports implements resolve.Resolver.
func (cSharpResolver) Imports(c *config.Config, r *rule.Rule, f *rule.File) []resolve.ImportSpec {
	panic("unimplemented")
}

// Name implements resolve.Resolver.
func (c cSharpResolver) Name() string {
	return language
}

// Resolve implements resolve.Resolver.
func (cSharpResolver) Resolve(c *config.Config, ix *resolve.RuleIndex, rc *repo.RemoteCache, r *rule.Rule, imports interface{}, from label.Label) {
	panic("unimplemented")
}

// NewResolver returns the Resolver for csharp language.
func NewResolver() resolve.Resolver {
	return cSharpResolver{}
}
