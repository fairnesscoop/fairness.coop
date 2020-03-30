setup: build-docker

build-docker:
	docker build -t fairnesscoop_hugo_server docker/hugo

run:
	docker run --rm -it -v $(PWD):/srv/ -p 1313:1313 fairnesscoop_hugo_server hugo server -D -w --bind=0.0.0.0 

build:
	docker run --rm -it -v $(PWD):/srv/ -p 1313:1313 fairnesscoop_hugo_server hugo

deploy: build
	cd public
	git checkout gh-pages || git checkout -b gh-pages
	git add .
	git commit -m "rebuilding site `date`"
	git push origin gh-pages