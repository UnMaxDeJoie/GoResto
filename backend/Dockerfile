###############################
# DOCKER START STAGE
###############################
FROM golang:latest
WORKDIR /luc/go-resto/
USER ${USER}
ADD ./go.mod /luc/go-resto/
ADD ./src /luc/go-resto/src

###############################
# DOCKER ENVIRONMENT STAGE
###############################
ENV GO111MODULE="on" \
  CGO_ENABLED="0" \
  GO_GC="off"

###############################
# DOCKER UPGRADE STAGE
###############################
RUN apt-get autoremove \
  && apt-get autoclean \
  && apt-get update --fix-missing \
  && apt-get upgrade -y \
  && apt-get install curl build-essential -y
  
###############################
# DOCKER INSTALL & BUILD STAGE
###############################
RUN cd /luc/go-resto/src/main \
  && go mod download \
  && go mod tidy \
  && go mod verify \
  && go build -o main .

###############################
# DOCKER FINAL STAGE
###############################
EXPOSE 3000
CMD ["./src/main/main"]
