FROM golang:1.22-alpine AS dev
WORKDIR /app
RUN apk add --no-cache ffmpeg
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM dev AS build
RUN CGO_ENABLED=0 go build -o /bin/api ./cmd/api
RUN CGO_ENABLED=0 go build -o /bin/worker ./cmd/worker

FROM alpine:3.19 as prod
RUN apk add --no-cache ffmpeg ca-certificates
COPY --from=build /bin/api /bin/worker /bin/
CMD ["/bin/api"]