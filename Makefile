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

windows:
	@echo "build for windows"
	@rm -f ./${BINARY}.exe
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe -ldflags ${LDFLAGS} cmd/*.go
	@upx ./${BINARY}.exe

release:
	@echo "build for mac"
	@rm -f ./${BINARY}
	@go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go
	@upx ./${BINARY}
	@zip -j ipcheck_darwin.zip ./${BINARY} ./assets/ip2region.db

	@echo "build for linux"
	@rm -f ./${BINARY}
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go
	@upx ./${BINARY}
	@zip -j ipcheck_linux.zip ./${BINARY} ./assets/ip2region.db

	@echo "build for windows"
	@rm -f ./${BINARY}.exe
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe -ldflags ${LDFLAGS} cmd/*.go
	@upx ./${BINARY}.exe
	@zip -j ipcheck_windows.zip ./${BINARY} ./assets/ip2region.db

	@rm -f ${BINARY}
	@rm -f ${BINARY}.exe

clean:
	@echo "clean binary"
	@rm -f ${BINARY}
	@rm -f ${BINARY}.exe
	@rm -f *.zip

.PHONY: mac linux windows release clean
