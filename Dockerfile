#Put the golang image as image's base
FROM golang:latest

#
LABEL maintainer="ericghoubiguian@live.fr"

#Copy all the files and directories in the newly created directory go_world_server
COPY . /go_world_server

#Change work directory for the go_world_server project one
WORKDIR /go_world_server

#Edit the environment variable value GOPATH for the go_world_server directory
ENV GOPATH /go_world_server

#Expose the docker container listening port
EXPOSE 80

#Container instruction as entrypoint: 'go run main.go'
ENTRYPOINT ["go", "run", "main.go"]
