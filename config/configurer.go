package config

import (
	"flag"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

type cSharpConfigurer struct{}

// CheckFlags implements config.Configurer.
func (cSharpConfigurer) CheckFlags(fs *flag.FlagSet, c *config.Config) error {
	return nil
}

// Configure implements config.Configurer.
func (cSharpConfigurer) Configure(c *config.Config, rel string, f *rule.File) {
}

// KnownDirectives implements config.Configurer.
func (c cSharpConfigurer) KnownDirectives() []string {
	return nil
}

// RegisterFlags implements config.Configurer.
func (cSharpConfigurer) RegisterFlags(fs *flag.FlagSet, cmd string, c *config.Config) {
}

// NewConfigurer returns the configurer for csharp language.
func NewConfigurer() config.Configurer {
	return cSharpConfigurer{}
}
