// main.go
package main

import "github.com/kuetemeier/test/config"

// compile with: goreleaser --skip-publish --rm-dist --snapshot
func main() {
	println("Go Build Stack test")
	println("")
	println("This is my version: " + config.Version)
	println("Full Commit Hash: " + config.FullCommit)
	println("Short Commit Hash: " + config.ShortCommit)
	println("My Build Time:" + config.Buildtime)
	println("Go Version used to compile:" + config.GoVersion)
}
