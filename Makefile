.PHONY: default clean

default:
	go build -o bin/glanceable

docker-build:
	docker build -t yuukiaria/glanceable .

clean:
	rm bin/*
