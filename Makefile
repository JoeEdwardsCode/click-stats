build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/clickEventHandler ./src/clickEventHandler
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/clickStatsService ./src/clickStatsService
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/utils ./src/utils