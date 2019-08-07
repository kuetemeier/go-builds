package config

var version = "unknown"
var gitHash = "unknown"
var buildTime = "unknown"

func GetVersion() string {
	return version
}

func GetGitHash() string {
	return gitHash
}

func GetBuildTime() string {
	return buildTime
}
