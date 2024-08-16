package parser

import "encoding/xml"

// Folder holds the information about a folder
type Folder struct {
	Folders []string
	Project *Project
}

type Project struct {
	ProjectSDK      string
	TargetFramework string
	IsTestProject   bool

	References   []string
	Dependencies []string
}

type csproj struct {
	XMLName       xml.Name `xml:"Project"`
	SDK           string   `xml:"Sdk,attr"`
	PropertyGroup csprojProperty
	References    csprojReferences
}

type csprojProperty struct {
	XMLName         xml.Name `xml:"PropertyGroup"`
	TargetFramework string   `xml:"TargetFramework"`
	IsTestProject   bool     `xml:"IsTestProject"`
}

type csprojReferences struct {
	XMLName xml.Name           `xml:"ItemGroup"`
	Refs    []projectReference `xml:"ProjectReference"`
	Deps    []projectReference `xml:"PackageReference"`
}

type projectReference struct {
	Ref string `xml:"Include,attr"`
}
