FROM golang
WORKDIR /root/deleteCR
COPY * ./
ENV CGO_ENABLED=0 GO111MODULE=on GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct
RUN go get k8s.io/client-go@kubernetes-1.19.2 && go build .
FROM alpine
WORKDIR /root/deleteCR
COPY --from=0 /root/deleteCR/deletecr /usr/local/bin/deletecr
CMD ["deletecr"]
