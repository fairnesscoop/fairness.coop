# Fairness.coop

## Install

You need [docker](https://docs.docker.com/get-docker/) on your machine.

Run the following command:

    make setup

## Usage

Run the server with hot reloading

    make run

Generate the public folder

    make build

## Deploy

The deployment is handle by a read-only [repository](https://github.com/fairnesscoop/fairness.coop.static) linked to the public dir.
You need to build the content of the blog and then commit it to the repository.

To do that, just run the following:

    make deploy

## Change Hugo version

If you want to change the Hugo version used, you need to modify the `HUGO_VERSION` env var from [docker/hugo/Dockerfile#L49](docker/hugo/Dockerfile#L49) and then run :

    make build-docker
