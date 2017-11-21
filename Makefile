all:
		go build -o ws
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ws
	cp ws ~/projects/ibupro-clean-tools/app/Supports/
clean:
		rm -rf ws
