FROM golang:latest
WORKDIR /go/src/cin.ufpe.br/~if998/cittapet/mapservice
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on
RUN go get -d -v ./...
RUN go install -a -installsuffix cgo -ldflags '-extldflags "-static" -s -w' -tags netgo -v ./...

FROM scratch
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /go/bin/mapservice /init
CMD ["/init"]
