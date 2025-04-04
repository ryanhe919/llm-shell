# 🚀 SGPT - Shell GPT

**SGPT** 是一款强大的命令行工具，利用大型语言模型（LLM）能力，快速将自然语言转换为 Shell 命令，让你的终端操作更简单直观！

![Stars](https://img.shields.io/github/stars/ryanhe919/llm-shell.svg?style=social)
![Forks](https://img.shields.io/github/forks/ryanhe919/llm-shell.svg?style=social)
![MIT License](https://img.shields.io/github/license/ryanhe919/llm-shell.svg)
![Release](https://img.shields.io/github/v/release/ryanhe919/llm-shell)

---

## ✨ 功能特点

- **自然语言交互**：直接使用日常语言生成 Shell 命令。
- **零依赖便捷**：编译后生成单个可执行文件，无需额外依赖。
- **高兼容性**：支持 OpenAI 标准 API 与自定义 API 端点。
- **可自动执行**：可选自动执行生成的命令，执行前确认。
- **安全优先**：执行命令前安全提示，确保操作可控。

---

## 📥 安装指南

### 编译安装

```bash
# 克隆项目
git clone https://github.com/ryanhe919/llm-shell.git
cd llm-shell

# 编译二进制文件
make build

# 安装到系统路径（可选）
sudo make install

# 验证安装
sgpt --version
```

---

## 🚩 快速开始

### 1. 设置 API 密钥

通过环境变量或命令参数指定 API 密钥：
```bash
export DEEPSEEK_API_KEY="你的API密钥"
```

### 2. 基本使用

```bash
sgpt "查找超过100MB的文件并按大小排序"
```

示例输出：

```bash
find . -type f -size +100M -exec ls -lh {} \; | sort -k 5 -hr
```

---

## 🔧 高级用法

### 指定模型和参数

```bash
sgpt --model "deepseek-chat" --max-tokens 200 --temp 0.5 "压缩所有PDF文件"
```

### 使用自定义 API

```bash
sgpt --url "http://localhost:8000/v1/chat/completions" "列出所有正在运行的进程"
```

### 自动执行（需确认）

```bash
sgpt --exec "递归查找所有包含'error'的日志文件"
```

---

## 📚 命令行选项

| 参数 | 描述 | 默认值                     |
|------|------|-------------------------|
| `--key` | OpenAI API 密钥 | 环境变量 `DEEPSEEK_API_KEY` |
| `--url` | API 端点 URL | DEEPSEEK 官方 API         |
| `--model` | 模型名称 | `deepseek-chat`         |
| `--max-tokens` | 最大 token 数 | `100`                   |
| `--temp` | 生成随机性 (0-2) | `0.2`                   |
| `--timeout` | 请求超时时间 | 无                       |
| `--exec` | 生成命令后自动执行（需确认） | 无                       |

---

## ⚠️ 安全须知

- **命令审查**：执行前务必检查生成的命令。
- **谨慎使用**：生产环境慎用，避免误操作。
- **命令确认**：理解命令具体作用再确认执行。

---

## 🛠️ 开发贡献

欢迎提交 Issue 和 Pull Request！

### 项目结构

```
llm-shell/
├── cmd/sgpt/main.go          # 程序入口
├── internal/api/client.go    # API交互
└── internal/shell/executor.go # 命令执行
```

### 贡献流程

1. Fork 仓库并创建新分支。
2. 提交更改并推送到仓库。
3. 创建 Pull Request。

---

## 📜 许可证

项目遵循 [MIT License](LICENSE) 开源协议。

---

## 💬 联系我们

欢迎通过 [GitHub Issues](https://github.com/ryanhe919/llm-shell/issues) 提出问题或建议！

**感谢你的 Star 和支持！** 🌟

