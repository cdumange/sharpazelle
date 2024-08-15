package parser

import "encoding/xml"

// Folder holds the information about a folder
type Folder struct {
	Folders []string
}

type Project struct {
	ProjectSDK      string
	TargetFramework string `xml:"TargetFramework"`
}

type csproj struct {
	XMLName       xml.Name `xml:"Project"`
	SDK           string   `xml:"Sdk,attr"`
	PropertyGroup csprojProperty
}

type csprojProperty struct {
	XMLName         xml.Name `xml:"PropertyGroup"`
	TargetFramework string   `xml:"TargetFramework"`
}
