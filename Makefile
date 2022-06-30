
up: build docker

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

up-test:
	docker rm test-es || true
	docker run --name es-test -d -e discovery.type=single-node -e ELASTIC_USERNAME=elastic -e ELASTIC_PASSWORD=elastic -e xpack.security.enabled=true -p 9200:9200 --network bridge elasticsearch:8.2.3 

run-test:
	@docker-compose -f test/search.yaml build
	@docker-compose -f test/search.yaml up -d 
	@/usr/local/bin/go test github.com/JekaTatsiy/grpc-market/http/suggest --args -s 0.0.0.0:1000
#	@docker-compose -f test/search.yaml down

gotest:
	make run-test -i
