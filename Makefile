
up: go-build docker-up

down:
	docker-compose down

docker-up:
	docker-compose build
	docker-compose up

go-build: proto
	GOOS=linux go build -o http/main http/main.go 
	GOOS=linux go build -o search/main search/main.go 

proto:
	@rm -rf suggest_proto/*pb.go
	protoc -I=suggest_proto/  --go_out=suggest_proto --go_opt=paths=source_relative --go-grpc_out=suggest_proto --go-grpc_opt=paths=source_relative suggest_proto/suggest.proto 
