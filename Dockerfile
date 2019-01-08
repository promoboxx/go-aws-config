FROM golang:1.11 as builder
WORKDIR /go/src/github.com/promoboxx/go-aws-config/
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/go-aws-config src/main/*.go

FROM centurylink/ca-certs
WORKDIR /
COPY --from=builder /go/src/github.com/promoboxx/go-aws-config/bin/go-aws-config .
ENTRYPOINT ["./go-aws-config"]
