FROM golang:1.19
# not using golang-alpine image because it's not officially supported by go

RUN mkdir /rest
ADD . /rest
WORKDIR /rest
RUN go clean --modcache
RUN go build -o main .
CMD ["/rest/main"]