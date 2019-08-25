FROM golang:1.12.9

WORKDIR /go/src/app
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]