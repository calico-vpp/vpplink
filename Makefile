.PHONY: build
build:
	mkdir -p ./bin
	go build -buildmode=plugin -o ./bin/vpplink_plugin.so ./pkg/


.PHONY: test
test:
	gofmt -s -l . | grep -v binapi | diff -u /dev/null -
	go vet ./...
