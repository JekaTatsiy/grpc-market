
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

run-test-http:
	@docker-compose -f test/search.yaml build
	@docker-compose -f test/search.yaml up -d 
	@echo --- HTTP TEST START ---
	@/usr/local/bin/go test github.com/JekaTatsiy/grpc-market/http/suggest --args -s 0.0.0.0:1000 || true
	@/usr/local/bin/go test github.com/JekaTatsiy/grpc-market/http/search --args -s 0.0.0.0:1000 || true
	@echo --- HTTP TEST END ---
	@docker-compose -f test/search.yaml down

run-test-search:
	@echo --- SEARCH TEST START ---
	@/usr/local/bin/go test github.com/JekaTatsiy/grpc-market/search/server --args -s 0.0.0.0:9200 || true
	@echo --- SEARCH TEST END ---


run-test: build
	make run-test-http -i
	make run-test-search -i