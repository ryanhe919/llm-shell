package main

import (
	"fmt"
	"os"

	"github.com/ryanhe919/sgpt/internal/api"
	"github.com/ryanhe919/sgpt/internal/config"
	"github.com/ryanhe919/sgpt/internal/shell"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("配置加载错误: %v\n", err)
		os.Exit(1)
	}

	// 检查是否需要显示帮助
	if cfg == nil {
		config.PrintUsage()
		os.Exit(1)
	}

	// 验证配置
	if !cfg.IsValid() {
		if cfg.APIKey == "" {
			fmt.Println("错误: 必须提供OpenAI API密钥")
			fmt.Println("可以通过 --key 标志或设置 OPENAI_API_KEY 环境变量")
		}
		config.PrintUsage()
		os.Exit(1)
	}

	// 创建API客户端
	clientConfig := api.ClientConfig{
		APIKey:      cfg.APIKey,
		APIURL:      cfg.APIURL,
		Model:       cfg.Model,
		MaxTokens:   cfg.MaxTokens,
		Temperature: cfg.Temperature,
		Timeout:     cfg.Timeout,
	}
	client := api.NewOpenAIClient(clientConfig, nil)

	// 生成shell命令
	command, err := client.GenerateShellCommand(cfg.Prompt)
	if err != nil {
		fmt.Printf("命令生成失败: %v\n", err)
		os.Exit(1)
	}

	// 显示生成的命令
	fmt.Println(command)

	// 如果需要执行命令
	if cfg.Execute {
		executor := shell.NewCommandExecutor()
		if err := executor.ExecuteWithConfirmation(command); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
