# arguments for Go version
ARG GO_VERSION=1.17
ARG PROJ_NAME=srv

# STAGE 1: building the executable
FROM golang:${GO_VERSION}-alpine AS build
ARG PROJ_NAME

# git required for go mod
RUN apk add --no-cache git

# Import code
COPY . /usr/local/go/src/${PROJ_NAME}
WORKDIR /usr/local/go/src/${PROJ_NAME}

# Run tests
RUN CGO_ENABLED=0 go test -timeout 30s -v /usr/local/go/src/${PROJ_NAME}/controllers

# https://stackoverflow.com/a/68172023, skip
# RUN go mod download

# Build the executable
RUN CGO_ENABLED=0 go build \
	-installsuffix 'static' \
	-o ./ ./

# STAGE 2: build the container to run
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static AS final
ARG PROJ_NAME

# add maintainer label
LABEL maintainer="yunhui"

# copy compiled app
COPY --from=build /usr/local/go/src/${PROJ_NAME}/${PROJ_NAME} app

# expose port
EXPOSE 80

# run binary
ENTRYPOINT ["./app"]
