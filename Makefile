FUNCTIONS := clickEventHandler

.PHONY: build clean deploy

build:
		${MAKE} ${MAKEOPTS} $(foreach function,${FUNCTIONS}, build-${function})

build-%:
		cd functions/$* && GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ../../bin/${%}

clean:
	rm -rf ./bin ./vendor Gopkg.lock
