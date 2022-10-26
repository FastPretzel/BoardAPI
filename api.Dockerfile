FROM golang:1.19-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git

COPY ./go.mod go.sum ./

RUN go mod download && go mod verify

#RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

COPY . .
COPY ./entrypoint.sh /entrypoint.sh

#RUN go run cmd/main.go
RUN chmod +rx /entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["sh", "/entrypoint.sh"]
