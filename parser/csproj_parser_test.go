package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseProject(t *testing.T) {
	t.Run("webapp", func(t *testing.T) {
		p, err := ParseProject("./test_data/simple/webapp/webapp.csproj")
		require.NoError(t, err)

		assert.Equal(t, "Microsoft.NET.Sdk.Web", p.ProjectSDK)
		assert.Equal(t, "net8.0", p.TargetFramework)
		assert.Contains(t, p.References, "..\\lib\\lib.csproj")
		assert.Contains(t, p.References, "JOS.Result")
		assert.Contains(t, p.References, "NSubstitute")
		assert.False(t, p.IsTestProject)
	})

	t.Run("Tests", func(t *testing.T) {
		p, err := ParseProject("./test_data/simple/Tests/Tests.csproj")
		require.NoError(t, err)

		assert.Equal(t, "Microsoft.NET.Sdk", p.ProjectSDK)
		assert.Equal(t, "net8.0", p.TargetFramework)
		assert.True(t, p.IsTestProject)
	})

	t.Run("not a csproj", func(t *testing.T) {
		_, err := ParseProject("../tests/simple/webapp/webapp.cs")
		require.Error(t, err)
	})

	t.Run("csproj but not exists", func(t *testing.T) {
		_, err := ParseProject("../tests/simple/webapp/webapp2.csproj")
		require.Error(t, err)
	})
}
