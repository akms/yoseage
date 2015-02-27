package yose

import (
	"testing"
	"os"
	."yoseage/yose"
)

func TestChangeDir(t *testing.T) {
	var dirPath string = "/var/log"
	ChangeDir(dirPath)
	workingDir, _ := os.Getwd()
	if workingDir != dirPath {
		t.Errorf("got %s\nwant %s", workingDir, dirPath)
	}
}
