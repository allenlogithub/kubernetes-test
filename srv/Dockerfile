# arguments for Go version
ARG GO_VERSION=1.17

# STAGE 1: building the executable
FROM golang:${GO_VERSION}-alpine AS build
ARG PROJ_NAME
ENV MYSQL_ADDRESS $MYSQL_ADDRESS
ENV MYSQL_NET $MYSQL_NET
ENV MYSQL_USER $MYSQL_USER
ENV MYSQL_PASSWORD $MYSQL_PASSWORD
ENV MYSQL_DBNAME $MYSQL_DBNAME

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
# FROM gcr.io/distroless/static:debug AS final
ARG PROJ_NAME
ENV MYSQL_ADDRESS $MYSQL_ADDRESS
ENV MYSQL_NET $MYSQL_NET
ENV MYSQL_USER $MYSQL_USER
ENV MYSQL_PASSWORD $MYSQL_PASSWORD
ENV MYSQL_DBNAME $MYSQL_DBNAME

# add maintainer label
LABEL maintainer="yunhui"

# copy compiled app
COPY --from=build /usr/local/go/src/${PROJ_NAME}/${PROJ_NAME} app

# expose port
EXPOSE 80

# run binary
ENTRYPOINT ["./app"]
