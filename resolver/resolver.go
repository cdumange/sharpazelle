package resolver

import (
	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/repo"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

const (
	cSharpLanguage = "csharp"
)

type cSharpResolver struct {
}

// Embeds implements resolve.Resolver.
func (c cSharpResolver) Embeds(r *rule.Rule, from label.Label) []label.Label {
	return nil
}

// Imports implements resolve.Resolver.
func (cSharpResolver) Imports(c *config.Config, r *rule.Rule, f *rule.File) []resolve.ImportSpec {
	return nil
}

// Name implements resolve.Resolver.
func (c cSharpResolver) Name() string {
	return cSharpLanguage
}

// Resolve implements resolve.Resolver.
func (cSharpResolver) Resolve(c *config.Config, ix *resolve.RuleIndex, rc *repo.RemoteCache, r *rule.Rule, imports interface{}, from label.Label) {
}

// NewResolver returns the Resolver for csharp language.
func NewResolver() resolve.Resolver {
	return cSharpResolver{}
}
