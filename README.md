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

### Adding pages

To add pages, add a Markdown file under `content/`. The file will be picked up by Hugo. Be sure to copy and adapt the [front matter](https://gohugo.io/content-management/front-matter/) from an existing page.

## Deploy

The deployment is launched automatically when a PR is merged into master

## Change Hugo version

If you want to change the Hugo version used, you need to modify the `HUGO_VERSION` env var from [docker/hugo/Dockerfile#L49](docker/hugo/Dockerfile#L49) and then run :

    make build-docker
