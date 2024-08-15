package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseProject(t *testing.T) {
	t.Run("webapp", func(t *testing.T) {
		p, err := ParseProject("../tests/simple/webapp/webapp.csproj")
		assert.NoError(t, err)

		assert.Equal(t, "Microsoft.NET.Sdk.Web", p.Target)
	})

	t.Run("not a csproj", func(t *testing.T) {
		_, err := ParseProject("../tests/simple/webapp/webapp.cs")
		assert.Error(t, err)
	})

	t.Run("csproj but not exists", func(t *testing.T) {
		_, err := ParseProject("../tests/simple/webapp/webapp2.csproj")
		assert.Error(t, err)
	})
}
