
up: go-build docker-up
down:
	docker-compose down


docker-up:
	docker-compose build
	docker-compose up

go-build:
	@GOOS=linux go build -o http/main http/main.go 
	@GOOS=linux go build -o search/main search/main.go 


