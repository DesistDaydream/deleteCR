FROM golang as builder
WORKDIR /root/deletecr

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download && go get k8s.io/client-go@kubernetes-1.19.2

COPY ./ /root/deletecr
RUN go build -o deletecr ./*.go

FROM alpine
RUN echo "hosts: files dns" > /etc/nsswitch.conf
WORKDIR /root/deletecr
COPY --from=builder /root/deletecr/deletecr /usr/local/bin/deletecr
ENTRYPOINT ["/usr/local/bin/deletecr"]
