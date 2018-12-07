APPNAME=sillybot

VERSION_TAG=`git describe 2>/dev/null | cut -f 1 -d '-' 2>/dev/null`
COMMIT_HASH=`git rev-parse --short=8 HEAD 2>/dev/null`
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-s -w \
    -X main.CommitHash=${COMMIT_HASH} \
    -X main.BuildTime=${BUILD_TIME} \
    -X main.Tag=${VERSION_TAG}"

all: fast

clean_all:
	go clean
	rm ./${APPNAME} || true
	rm -rf ./target || true

clean_binary:
	go clean
	rm -rf ./target || true

fast:
	go build -i -v -o ${APPNAME} ${LDFLAGS}

linux:
	GOOS=linux GOARCH=386 go build -i -v ${LDFLAGS} -o ./target/linux_386/${APPNAME}
	GOOS=linux GOARCH=amd64 go build -i -v ${LDFLAGS} -o ./target/linux_amd64/${APPNAME}

darwin:
	GOOS=darwin GOARCH=386 go build -i -v ${LDFLAGS} -o ./target/darwin_386/${APPNAME}
	GOOS=darwin GOARCH=amd64 go build -i -v ${LDFLAGS} -o ./target/darwin_amd64/${APPNAME}

windows:
	GOOS=windows GOARCH=386 go build -i -v ${LDFLAGS} -o ./target/windows_386/${APPNAME}.exe
	GOOS=windows GOARCH=amd64 go build -i -v ${LDFLAGS} -o ./target/windows_amd64/${APPNAME}.exe

build: clean_binary linux darwin windows
