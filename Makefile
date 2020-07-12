BINARY=ipcheck
VERSION=0.1
LDFLAGS='-w -s'

mac:
	@echo "build for mac"
	@rm -f ./${BINARY}
	@go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go

linux:
	@echo "build for linux"
	@rm -f ./${BINARY}
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go
	@upx ./${BINARY}

clean:
	@echo "clean binary"
	@rm -f ${BINARY}

.PHONY: mac linux clean
