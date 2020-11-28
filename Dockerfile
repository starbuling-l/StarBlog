FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/starbuling-l/StarBlog
COPY . $GOPATH/src/github.com/starbuling-l/StarBlog
RUN go build .

EXPOSE 9000
ENTRYPOINT ["./StarBlog"]

