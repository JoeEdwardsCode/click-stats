build:
	go build -o bin/clickEventHandler ./src/clickEventHandler
	go build -o bin/clickStatsService ./src/clickStatsService
	go build -o bin/utils ./src/utils