#
FROM golang

#
LABEL maintainer="ericghoubiguian@live.fr"

#
COPY . /go_world_server

#
WORKDIR /go_world_server

#
ENV GOPATH /go_world_server

#
EXPOSE 8080

#
ENTRYPOINT ["go", "run", "main.go"]
