# =====================
# Target 'base'
# =====================
ARG BUILD_ENV_IMAGE="build-env"

FROM golang:1.19 AS base

RUN apt-get install gcc git && \
    go install github.com/golang/mock/mockgen@v1.6.0 && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /src

EXPOSE 80

# =====================
# Target 'dev'
# =====================
FROM base AS dev

RUN go install github.com/githubnemo/CompileDaemon@latest
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /usr/local/bin/wait-for-it
RUN chmod +x /usr/local/bin/wait-for-it

CMD CompileDaemon --build="go get ./... && go build -o goapp -buildvcs=false" --command="./goapp http run"

# =====================
# Target 'build-env'
# =====================
FROM base AS build-env

COPY . /src

RUN go mod download && \
    swag init && \
    CGO_ENABLED=0 go build -o goapp

# =====================
# Target 'my-build'
# =====================
FROM $BUILD_ENV_IMAGE AS my-build

# =====================
# Target 'prod'
# =====================
FROM gcr.io/distroless/static-debian11 AS prod
ENV APP_ENV=prod

COPY --from=my-build /src/config /config
COPY --from=my-build /src/goapp /

ENTRYPOINT ["/goapp"]

CMD ["http", "run"]

EXPOSE 80
