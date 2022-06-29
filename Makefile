
up: docker

down:
	docker-compose down

docker:
	docker-compose build
	docker-compose up

build: proto
	GOOS=linux go build -o http/main http/main.go 
	GOOS=linux go build -o search/main search/main.go 

proto:
	@rm -rf suggest_proto/*pb.go
	protoc -I=suggest_proto/  --go_out=suggest_proto --go_opt=paths=source_relative --go-grpc_out=suggest_proto --go-grpc_opt=paths=source_relative suggest_proto/suggest.proto 

nettest:
	@curl 0.0.0.0:3000
	@curl 0.0.0.0:1000
	@curl 0.0.0.0:9200