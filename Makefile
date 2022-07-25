FUNCTIONS := clickEventHandler

.PHONY: build clean deploy

build:
		${MAKE} ${MAKEOPTS} $(foreach function,${FUNCTIONS}, build-${function})

build-%:
		cd functions/$* && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags="-s -w" -o ../../bin/${%}

clean:
	rm -rf ./bin ./vendor Gopkg.lock
