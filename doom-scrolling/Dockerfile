FROM golang:1.24 AS build

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
RUN go build -mod=vendor -o main main.go

FROM alpine:3.20

WORKDIR /app
COPY --from=build /app/main ./main
COPY config.yaml .

EXPOSE 8080

ENTRYPOINT [ "./main" ]