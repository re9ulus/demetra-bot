build:
	docker build . -t demetra-dev

run-dev: build
	docker run --rm -v $(shell pwd):/app -it demetra-dev /bin/sh

run-dev-redis:
	docker run -p 6379:6379 --rm --name dev-redis -d redis
