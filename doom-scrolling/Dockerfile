FROM golang:1.24 AS build


ENV PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /app

COPY . .

RUN make tools

ENV CGO_ENABLED=0
RUN go build -mod vendor -o /doom-scrolling cmd/app/main.go

FROM alpine:3.20

COPY --from=build /doom-scrolling /doom-scrolling

ENTRYPOINT [ "/doom-scrolling" ]
