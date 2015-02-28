package yose

import (
	"testing"
	"archive/tar"
	"compress/gzip"
	"os"
)

func TestMakeFile(t *testing.T) {
	var (
		gw *gzip.Writer
		tw *tar.Writer
		file *os.File
	)
	gw, tw, file = MakeFile()
	if gw == nil {
		t.Errorf("make faild gzip writer")
	}
	if tw == nil {
		t.Errorf("make faild tar writer")
	}
	if file == nil {
		t.Errorf("make faild file")
	}
}
