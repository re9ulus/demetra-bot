build:
	docker build . -t demetra-dev

run-redis:
	docker run -p 6379:6379 --rm --name dev-redis -d redis

run-dev: build run-redis
	docker run --rm -v $(shell pwd):/app -it demetra-dev /bin/sh
