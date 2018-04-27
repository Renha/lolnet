FROM golang

COPY ./ /go/src/github.com/lexfrei/lolnet/
WORKDIR /go/src/github.com/lexfrei/lolnet/cmd/

RUN go get ./
RUN go build -o lolnet

CMD ["./lolnet"]
