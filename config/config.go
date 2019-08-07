package config

var (
	// Version of the binary, linked at build-time (@see goreleaser)
	Version = "dev"
	// ShortCommit is the short GitHub Hash of the commit it is build upon, linked at build-time (@see goreleaser)
	// First seven chars of the full hash
	ShortCommit = "none"
	// FullCommit is the full GitHub Hash of the commit it is build upon, linked at build-time (@see goreleaser)
	FullCommit = "none"
	// Buildtime is the RFC3339 timestamp, when this build was compiled, linked at build-time (@see goreleaser)
	Buildtime = "unknown"
	// GoVersion used to compile this binary
	GoVersion = "unknown"
)
