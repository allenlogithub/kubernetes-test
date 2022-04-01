ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine
ENV EXEC_ENV=dev

RUN apk add --no-cache git

WORKDIR /src

# COPY ./go.mod ./go.sum ./
# RUN go mod download

COPY ./ ./

ENV PATH="$PATH:$(go env GOPATH)/bin"

EXPOSE 80
