# Golang Makefile

# variable
bizLine=test
binaryName=kwina-demo
version=v0.1.0
outputPath=_output

all: build

build:
	@buildTime=`date "+%Y-%m-%d %H:%M:%S"`; \
	go build -ldflags " -X '${versionPath}.BuildTime=$$buildTime' \
						-X '${versionPath}.GoVersion=`go version`' \
						-X '${versionPath}.GitCommit=`git rev-parse --short HEAD`'" -o ${outputPath}/${binaryName}; \


clean:
	rm -rf _output
	rm -rf logs

.PHONY: all build clean
