FROM golang:1.20.5-buster as build-env

WORKDIR /opt/dockerize
ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go build main.go
CMD [ "./main" ]

FROM bitnami/minideb:latest
WORKDIR /opt/dockerize
COPY --from=build-env /opt/dockerize/main .
CMD [ "./main" ]