.PHONY : test format docker

docker: Dockerfile
	docker build -t service-order:latest .

order-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

order-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

test: ./middleware
	go test -race \
					./middleware
