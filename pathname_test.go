package gopathname

import "testing"

func TestGetFileNames(t *testing.T) {
	paths := GetFileNames("test")

	if len(paths) == 0 {
		t.Error("None paths")
	}
}
