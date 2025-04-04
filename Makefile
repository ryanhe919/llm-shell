.PHONY: build clean test install

# 默认模块名
MODULE := github.com/ryanhe919/sgpt

# 命令名
BINARY_NAME := sgpt

# 主要源文件
MAIN := cmd/sgpt/main.go

# 构建目标
build:
	go build -o $(BINARY_NAME) $(MAIN)

# 运行测试
test:
	go test ./...

# 安装到系统
install: build
	mv $(BINARY_NAME) /usr/local/bin/

# 清理编译文件
clean:
	rm -f $(BINARY_NAME)
	go clean

# 运行代码格式化和静态检查
lint:
	go fmt ./...
	go vet ./...