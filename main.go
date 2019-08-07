// main.go
package main

import "github.com/kuetemeier/test/config"

func main() {
	println("Ba dum, tss!")
	println("")
	println("This is version: " + config.GetVersion())
	println("Commit Hash: " + config.GetGitHash())
	println("Build Time:" + config.GetBuildTime())
}
