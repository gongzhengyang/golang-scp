.PHONY: build
build:
	mkdir -p target
	CGO_ENABLED=0 go build -o target/scp
