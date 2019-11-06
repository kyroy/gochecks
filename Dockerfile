FROM golang:1.11 AS builder

RUN mkdir /cmd
RUN mkdir /src

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o /cmd/gochecks


FROM golang:1.11

LABEL "repository"="http://github.com/kyroy/gochecks"
LABEL "homepage"="http://github.com/kyroy/gochecks"

WORKDIR /cmd
COPY --from=builder /cmd /cmd
