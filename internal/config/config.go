package config

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Config 存储应用程序的配置
type Config struct {
	APIKey      string
	APIURL      string
	Model       string
	MaxTokens   int
	Temperature float64
	Timeout     time.Duration
	Execute     bool
	Prompt      string
	System      string
}

// LoadConfig 从命令行参数和环境变量加载配置
func LoadConfig() (*Config, error) {
	config := &Config{}

	// 定义命令行选项
	apiKey := flag.String("key", "", "OpenAI API密钥 (也可通过OPENAI_API_KEY环境变量设置)")
	apiUrl := flag.String("url", "https://api.deepseek.com/chat/completions", "API端点URL")
	apiSystem := flag.String("system", "Windows/Powershell", "系统名")
	model := flag.String("model", "deepseek-chat", "要使用的模型名称")
	maxTokens := flag.Int("max-tokens", 100, "生成的最大token数")
	temperature := flag.Float64("temp", 0.2, "生成的随机性 (0-2)")
	timeout := flag.Duration("timeout", 30*time.Second, "API请求超时时间")
	execute := flag.Bool("exec", true, "自动执行生成的命令")

	// 解析命令行参数
	flag.Parse()

	// 获取用户描述
	args := flag.Args()
	if len(args) == 0 {
		return nil, nil // 触发帮助信息显示
	}

	// 设置配置
	config.APIURL = *apiUrl
	config.Model = *model
	config.MaxTokens = *maxTokens
	config.System = *apiSystem
	config.Temperature = *temperature
	config.Timeout = *timeout
	config.Execute = *execute
	config.Prompt = args[0]

	// 如果提供了多个参数，合并成一个字符串
	if len(args) > 1 {
		for _, arg := range args[1:] {
			config.Prompt += " " + arg
		}
	}

	// API密钥优先级：命令行参数 > 环境变量
	config.APIKey = *apiKey
	if config.APIKey == "" {
		config.APIKey = os.Getenv("DEEPSEEK_API_KEY")
	}

	return config, nil
}

// IsValid 检查配置是否有效
func (c *Config) IsValid() bool {
	return c.APIKey != "" && c.Prompt != ""
}

// PrintUsage 打印使用说明
func PrintUsage() {
	fmt.Println("用法: sgpt [选项] \"shell命令的中文描述\"")
	fmt.Println("\n选项:")
	flag.PrintDefaults()
}
