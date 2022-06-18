FROM golang:1.18.1-alpine3.15

ENV TEMP_PATH=/app/temp
ENV DATABASE_NAME=etourism
ENV PUBLIC_PATH=prod
ENV PORT=8080
ENV CWD=/app
ENV DB_AUTH_SOURCE=etourism
ENV DB_USER_NAME=root
ENV DB_PASSWORD=password
ENV SECRET=secretJwtUbersnap001

RUN apk add git

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/api

EXPOSE ${PORT}/tcp
CMD ["api"]