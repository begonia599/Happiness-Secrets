# 构建阶段
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum* ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY main.go ./

# 构建应用
RUN go build -o happiness-secrets main.go

# 运行阶段
FROM alpine:latest

WORKDIR /app

# 复制编译好的二进制文件
COPY --from=builder /app/happiness-secrets .

# 复制页面文件
COPY pages/ ./pages/

# 暴露端口
EXPOSE 3000

# 运行应用
CMD ["./happiness-secrets"]
