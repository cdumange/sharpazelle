package parser

import (
	"context"
	"fmt"
	"os"
	"strings"
)

const (
	extSLN    = ".sln"
	extCSProj = ".csproj"
	extCS     = ".cs"
)

// ParseFolder retrieves the information of the folder and extract them to allow rules interpolation.
func ParseFolder(ctx context.Context, folder string) (Folder, error) {
	f, err := os.ReadDir(folder)
	if err != nil {
		return Folder{}, fmt.Errorf("error reading folder: %w", err)
	}
	return parserFolder(ctx, f), nil
}

func parserFolder(ctx context.Context, f []os.DirEntry) Folder {
	var folders []string

	for _, v := range f {
		if v.IsDir() {
			folders = append(folders, v.Name())
			continue
		}

		if strings.Contains(v.Name(), extCSProj) {

		}
	}

	return Folder{
		Folders: folders,
	}
}
