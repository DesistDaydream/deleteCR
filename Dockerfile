FROM golang
WORKDIR /root/deleteCR
COPY * ./
RUN export CGO_ENABLED=0 && \
    export GO111MODULE=on && \
    export GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct && \
    go get k8s.io/client-go@kubernetes-1.19.2 && go build .
FROM alpine
WORKDIR /root/deleteCR
COPY --from=0 /root/deleteCR/deletecr /usr/local/bin/deletecr
CMD ["deletecr"]