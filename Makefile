export CGO_ENABLED=0
#export GOOS=linux
export GOARCH=amd64

lint:
	golangci-lint -v run ./...

clean:
	goimports -local -w .
	go fmt ./...

test:
	time go test ./...

restart-dependencies:
	docker-compose down --volumes
	docker-compose up --build -d postgres

build:
	rm -rf "bin"
	mkdir "bin"
	go build -o "bin/"

test-integration:
	time go test ./... -tags=integration

restart: build
	docker-compose down --volumes
	docker-compose up --build

migrate: build
	./bin/boilerplate migrate --dir migrate --operation up