build-and-release:
	sbot release version
	sbot push version
	go build -o ./dist/errorutil -ldflags="-X 'github.com/fenixvlabs/meshkit/cmd/errorutil.version=0.0.3'" cmd/errorutil/main.go