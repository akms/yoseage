package main

import (
	"github.com/akms/yoseage/yose"
)

func main () {
	var dirPaths = []string{"/root/document"}
	yose.CheckTarget(dirPaths)
}
