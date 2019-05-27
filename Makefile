.PHONY: default clean

default:
	go build -o bin/glanceable

test:
	go test -v -cover ./...

docker-build:
	docker build -t yuukiaria/glanceable .

clean:
	rm bin/*
