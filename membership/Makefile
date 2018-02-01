.PHONY : test format docker

docker: Dockerfile
	docker build -t service-membership:latest .

membership-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

membership-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

test: ./model ./query ./usecase ./presenter
	go test -race \
					./model \
					./query \
					./usecase \
					./presenter
