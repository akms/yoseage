package main

import (
	"os"
	"yoseage/yose"
)

func main() {
	var dirPaths = []string{"/root/document"}
	yose.CheckTarget(dirPaths)
}
