FROM golang:1.13.1

WORKDIR /go/src/gin_todo
COPY . .
ENV GO111MODULE=on
ENV MYSQL_USER=root

RUN go get github.com/pilu/fresh
CMD ["fresh"]
