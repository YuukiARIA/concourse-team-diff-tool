.PHONY: default clean

default:
	go build -o bin/glanceable

clean:
	rm bin/*
