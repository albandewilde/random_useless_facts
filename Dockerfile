FROM golang:1.14 as builder

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/ruf


FROM scratch

WORKDIR /bin/app

COPY --from=builder /bin/ruf .

CMD ["./ruf"]
