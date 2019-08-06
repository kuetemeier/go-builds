// main.go
package main

var version = "unknown"
var commithash = "unknown"
var buildtime = "unknown"

func main() {
	println("Ba dum, tss!")
	println("")
	println("This is version: " + version)
	println("Commit Hash: " + commithash)
	println("Timestamp:" + buildtime)
}
