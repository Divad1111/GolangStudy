package main

import (
	//tap "TestArchivePackage"
	tap "TestArchivePackage"
	//tos "TestOSPackage"
	"os"
)

func main() {
	//tos.Test()
	tap.Test()

	os.Exit(0)
}
