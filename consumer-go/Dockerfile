FROM golang:1.18-alpine AS build

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm64

COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o consumer

FROM scratch

WORKDIR /

COPY --from=build /app/consumer /consumer

EXPOSE 8080

ENTRYPOINT ["/consumer"]
