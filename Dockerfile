FROM golang:alpine
COPY dab.go .
RUN go build dab.go

FROM alpine
COPY --from=0 /go/dab .
CMD ["./dab"]