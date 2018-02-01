.PHONY : test format docker

docker: Dockerfile
	docker build -t service-auth:latest .

auth-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

auth-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

test: ./model ./query ./token ./usecase ./presenter
	go test -race \
					./model \
					./query \
					./token \
					./usecase \
					./presenter
