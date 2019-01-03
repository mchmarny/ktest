FROM golang:latest as build
ENV GO111MODULE=on

WORKDIR /go/src/github.com/mchmarny/tellmeall/
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /server

FROM scratch
COPY --from=build /server /app/
ENTRYPOINT ["/app/server"]