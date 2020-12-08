FROM golang
WORKDIR /root/deletecr
COPY * ./
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct
RUN go get k8s.io/client-go@kubernetes-1.19.2 && go build .
FROM alpine
WORKDIR /root/deletecr
COPY --from=0 /root/deletecr/deletecr /usr/local/bin/deletecr
CMD ["deletecr"]
