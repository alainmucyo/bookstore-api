FROM golang:1.20 as build

WORKDIR /go/src/github.com/alaimucyo/bookstore

COPY go.mod go.sum  ./

RUN GO111MODULE=on GOPROXY="https://proxy.golang.org" go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go build -o /bin/bookstore .

FROM scratch
WORKDIR /
COPY .env /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/* /bin/


EXPOSE 3000

CMD ["/bin/bookstore"]
