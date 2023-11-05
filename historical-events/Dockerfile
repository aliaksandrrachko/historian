# syntax=docker/dockerfile:1
ARG GIT_COMMIT_HASH
ARG GIT_COMMIT_AUTHOR
ARG APP_VERSION

FROM golang:1.21

ARG GIT_COMMIT_HASH
ARG GIT_COMMIT_AUTHOR
ARG APP_VERSION

WORKDIR /usr/src/historical-events

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -v -o /usr/local/bin/historical-events  \
    -ldflags="-X 'github.com/aliaksandrrachko/historian/historical-events/internal/build.version=$APP_VERSION' \
        -X 'github.com/aliaksandrrachko/historian/historical-events/internal/build.gitCommit=$GIT_COMMIT_HASH' \
        -X 'github.com/aliaksandrrachko/historian/historical-events/internal/build.time=$(date)' \
        -X 'github.com/aliaksandrrachko/historian/historical-events/internal/build.user=$GIT_COMMIT_AUTHOR'" \
    ./cmd/

EXPOSE 8080

CMD ["/usr/local/bin/historical-events"]