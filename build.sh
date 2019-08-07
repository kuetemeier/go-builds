go build -i -v -ldflags "-s -w -X github.com/kuetemeier/test/config.version=$(git describe --always --long --dirty)"
