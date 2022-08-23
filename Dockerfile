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
RUN go mod download
COPY main.go ./
RUN go build -o userInfoService ./main.go
 
FROM scratch
COPY --from=builder /userInfoService/userInfoService /
COPY --from=builder /userInfoService/conf/ /conf/
RUN mkdir logs && chmod 755 logs
ENTRYPOINT ["./userInfoService"]