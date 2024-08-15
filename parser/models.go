package parser

import "encoding/xml"

// Folder holds the information about a folder
type Folder struct {
	Folders []string
}

type Project struct {
	Target string
}

type csproj struct {
	XMLName xml.Name `xml:"Project"`
	SDK     string   `xml:"Sdk,attr"`
}

type csprojProperty struct {
}
