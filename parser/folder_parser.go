package parser

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cdumange/sharpazelle/commonerrors"
	"github.com/cdumange/sharpazelle/utils"
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
	return parserFolder(ctx, folder, f)
}

func parserFolder(_ context.Context, parentFolder string, f []os.DirEntry) (Folder, error) {
	var folders []string
	var csproj string

	for _, v := range f {
		if v.IsDir() {
			if !utils.IsTreatableFolder(completePath(parentFolder, v.Name())) {
				continue
			}
			folders = append(folders, v.Name())
			continue
		}

		if !strings.Contains(v.Name(), extCSProj) {
			continue
		}

		csproj = v.Name()
		break
	}

	if csproj == "" {
		return Folder{}, commonerrors.ErrNotToBeTreated
	}

	p, err := ParseProject(completePath(parentFolder, csproj))
	if err != nil {
		return Folder{}, fmt.Errorf("error parsing %s: %w", csproj, err)
	}

	return Folder{
		Folders: folders,
		Project: &p,
	}, nil
}

func completePath(parent string, child string) string {
	return fmt.Sprintf("%s/%s", parent, child)
}
