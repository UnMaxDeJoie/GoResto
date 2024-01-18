# Démarrer le projet
FROM golang:latest
WORKDIR /usr/src/goapp/
USER ${USER}
ADD ./src/go.mod /usr/src/goapp/
ADD ./src /usr/src/goapp/

# Environnement de docker
ENV GO111MODULE="on" \
  CGO_ENABLED="0" \
  GO_GC="off"

# Mise à jours de docker
RUN apt-get autoremove \
  && apt-get autoclean \
  && apt-get update --fix-missing \
  && apt-get upgrade -y \
  && apt-get install curl \
  build-essential -y

# Installation et build du docker
RUN go mod download \
  && go mod tidy \
  && go mod verify \
  && go build -o main .

EXPOSE 3000
CMD ["./main"]