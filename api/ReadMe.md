# booth-api

build image with

`docker build --tag booth-api:1.0 .`

create container with

`docker run -d -p 3001:80 --env-file .env -v ~/docker-volume/pronto-api/temp:/app/temp --name booth-api booth-api:1.0`