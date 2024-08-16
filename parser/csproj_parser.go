package parser

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

// ParseProject will parse a csproj file to retrieve project informations
func ParseProject(csprojPath string) (Project, error) {
	if !strings.HasSuffix(csprojPath, extCSProj) {
		return Project{}, fmt.Errorf("file %s is not a csproj", csprojPath)
	}

	b, err := os.ReadFile(csprojPath)
	if err != nil {
		return Project{}, fmt.Errorf("error opening the csproj %s: %w", csprojPath, err)
	}

	var file csproj
	if err := xml.Unmarshal(b, &file); err != nil {
		return Project{}, fmt.Errorf("error while parsing xml: %w", err)
	}

	var ref []string
	for _, v := range file.References.Refs {
		ref = append(ref, v.Ref)
	}

	var deps []string
	for _, v := range file.References.Deps {
		ref = append(ref, v.Ref)
	}

	return Project{
		ProjectSDK:      file.SDK,
		TargetFramework: file.PropertyGroup.TargetFramework,
		IsTestProject:   file.PropertyGroup.IsTestProject,
		References:      ref,
		Dependencies:    deps,
	}, nil
}
