ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine
ENV MYSQL_ADDRESS $MYSQL_ADDRESS
ENV MYSQL_NET $MYSQL_NET
ENV MYSQL_USER $MYSQL_USER
ENV MYSQL_PASSWORD $MYSQL_PASSWORD
ENV MYSQL_DBNAME $MYSQL_DBNAME

RUN apk add --no-cache git

WORKDIR /src

# COPY ./go.mod ./go.sum ./
# RUN go mod download

COPY ./ ./

ENV PATH="$PATH:$(go env GOPATH)/bin"

EXPOSE 80
