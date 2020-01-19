build:
	docker build . -t demetra-dev

run:
	docker run --rm -v $(shell pwd):/app -it demetra-dev /bin/sh

build-and-run: build run
