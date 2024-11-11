FROM golang:1.23-alpine

RUN mkdir -p /app/go/src
WORKDIR /app/go/src

ADD . .

# RUN git config --global url."https://developer:token@cimbgit.bitbucket.org/".insteadOf "https://cimbgit.bitbucket.org/"

RUN go install github.com/githubnemo/CompileDaemon@latest

EXPOSE 9999
ENTRYPOINT CompileDaemon -build="go build -o ./cmd/server/myapp ./cmd/server/main.go" -command="./cmd/server/myapp"
