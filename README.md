# SGPT - Shell GPT

SGPT是一个命令行工具，可以将自然语言描述转换为Shell命令。通过大型语言模型的能力，SGPT让你用自己的语言描述你想要完成的任务，然后自动生成相应的Shell命令。

## 特点

- **简单易用**: 直接用自然语言描述你需要的命令
- **零依赖**: 编译后生成单个二进制文件，无需安装任何依赖
- **兼容性**: 支持标准OpenAI API接口格式
- **可配置性**: 支持多种命令行选项
- **自动执行**: 可选择自动执行生成的命令

## 安装

### 从源码编译

```bash
# 克隆仓库
git clone https://github.com/yourusername/sgpt.git
cd sgpt

# 编译
make build

# 安装到系统路径
sudo make install
```

## 使用方法

### 基本使用

```bash
# 设置API密钥
export OPENAI_API_KEY="你的OpenAI密钥"

# 生成命令
sgpt "查找超过100MB的文件并按大小排序"
```

### 高级选项

```bash
# 指定模型和其他参数
sgpt --model "gpt-4o" --max-tokens 200 --temp 0.5 "压缩所有PDF文件"

# 使用自定义API端点（例如本地部署的模型）
sgpt --url "http://localhost:8000/v1/chat/completions" "列出所有正在运行的进程"

# 自动执行生成的命令（会要求确认）
sgpt --exec "递归查找所有包含'error'的日志文件"
```

### 命令行选项

- `--key`: OpenAI API密钥 (也可通过OPENAI_API_KEY环境变量设置)
- `--url`: API端点URL (默认为OpenAI官方API)
- `--model`: 要使用的模型名称 (默认为gpt-3.5-turbo)
- `--max-tokens`: 生成的最大token数 (默认为100)
- `--temp`: 生成的随机性, 0-2之间 (默认为0.2)
- `--timeout`: API请求超时时间
- `--exec`: 生成命令后自动执行 (会要求确认)

## 安全注意事项

- 在执行生成的命令前，请先审查确保安全
- 特别是在生产环境或重要系统上更需谨慎
- 始终确保你了解命令的作用后再执行

## 许可证

MIT