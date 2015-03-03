package main

import (
	"github.com/akms/yoseage/yose"
)

func main () {
	var dirPaths = []string{"/sbin"}
	yose.CheckTarget(dirPaths)
}
