FROM golang:alpine
COPY dab.go .
RUN apk add build-base && \
    go build -ldflags="-s -w" -trimpath dab.go

FROM alpine
COPY --from=0 /go/dab .
CMD ["./dab"]