setup: build-docker

build-docker:
	docker build -t fairnesscoop_hugo_server docker/hugo

start:
	go run cmd/victor/main.go

run:
	docker run --rm -it -v $(PWD):/srv/ -p 1313:1313 fairnesscoop_hugo_server hugo server -D -w --bind=0.0.0.0

build:
	docker run --rm -it -v $(PWD):/srv/ -p 1313:1313 fairnesscoop_hugo_server hugo
