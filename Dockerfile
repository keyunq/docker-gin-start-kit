FROM golang:1.19 AS builder
 
ENV GO111MODULE=on \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64
ENV GOPROXY=https://goproxy.cn
 
WORKDIR /userInfoService
COPY go.mod ./
COPY go.sum ./
COPY routers/ ./routers/
COPY conf/ ./conf/
COPY docs/ ./docs/
COPY middlewares/ ./middlewares/
COPY models/ ./models/
COPY pkg/ ./pkg/
RUN mkdir logs
RUN go mod download
COPY main.go ./
RUN go build -o userInfoService ./main.go
 
FROM alpine:3.7
COPY --from=builder /userInfoService/userInfoService /
COPY --from=builder /userInfoService/conf/ /conf/
ENV TZ=Asia/Shanghai
RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories \
    && apk --no-cache add tzdata zeromq \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo '$TZ' > /etc/timezone
ENTRYPOINT ["./userInfoService"]