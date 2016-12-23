#build go with binary versioning
#give binary new name
BINARY=mycrawl

#versioning
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

#setup ldflags
LDFLAGS=-ldflags "-X github.com/fakihariefnoto/gocrawlimage/main.Version=${VERSION} -X github.com/fakihariefnoto/gocrawlimage/main.BuildTime=${BUILD_TIME}"

#build
default:
	go build ${LDFLAGS} -o ${BINARY} main.go

run:
	./mycrawl

kick: default run

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY}; fi

fmt:
	go fmt ./src/..

lint:
	go lint ./src

help:
	@echo 'Welcome to GoCrawlImage'
	@echo 'Pilih option sebagi berikut'
	@echo '---------------------------------'
	@echo 'make       -untuk build'
	@echo 'make run	  -untuk run, binary -> mycrawl'
	@echo 'make fmt	  -go fmt semua file'
	@echo 'make kick  -build and run'
	@echo '-------------------------'
