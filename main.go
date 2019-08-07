// main.go
package main

import "github.com/kuetemeier/go-builds/config"

// compile with: goreleaser --skip-publish --rm-dist --snapshot
func main() {
	println("Go Build Stack test")
	println("")
	println("This is my version: " + config.Version)
	println("Full Commit Hash: " + config.FullCommit)
	println("Short Commit Hash: " + config.ShortCommit)
	println("My Build Time:" + config.Buildtime)
}
