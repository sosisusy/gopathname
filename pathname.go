package gopathname

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFileNames(root string) []string {
	var (
		paths []string
	)

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			path = strings.ReplaceAll(path, "\\", "/")
			paths = append(paths, strings.Replace(path, fmt.Sprintf("%v/", root), "", 1))
		}
		return err
	})

	for _, v := range paths {
		fmt.Println(v)
	}

	return paths
}
