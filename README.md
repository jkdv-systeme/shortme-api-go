# short-me API

This is the API for the short-me project. It is a simple URL shortener.

## Used technologies
- [fiber](https://gofiber.io/)
- [gorm](https://gorm.io/)
- [postgres](https://www.postgresql.org/)
- [docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)
- [air](https://github.com/cosmtrek/air)

## How to run
- Install [docker](https://docs.docker.com/get-docker/) and [docker-compose](https://docs.docker.com/compose/install/)
- Copy `.env.example` to `.env` and fill the variables
- Run `docker-compose up -d`
- Run `air` to start the server with hot reloading
- The API will be available at `localhost:4682`