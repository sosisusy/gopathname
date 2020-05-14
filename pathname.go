package gopathname

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// get file templates
func GetTemplate(root string) *template.Template {
	t := template.New(root)

	paths := GetFileNames(root)
	for _, v := range paths {
		t, _ = t.ParseFiles(v)
	}

	return t
}

// root 폴더 내 전체 파일 이름 가져오기 (경로포함)
func GetFileNames(root string) []string {
	var (
		paths []string
	)

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			path = strings.ReplaceAll(path, "\\", "/")
			paths = append(paths, path)
		}
		return err
	})

	return paths
}
