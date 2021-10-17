FROM golang:alpine AS build

RUN mkdir "/src" && mkdir "/build"
WORKDIR "/src"
COPY . .

RUN go test

RUN go build -o "/build/fabricator" .

FROM alpine:latest AS deployment

RUN mkdir -p "/opt/bin"
COPY --from=build "/build/fabricator" "/opt/bin"

ENTRYPOINT ["/opt/bin/fabricator"]