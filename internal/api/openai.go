package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// OpenAIClient 实现OpenAI API的客户端
type OpenAIClient struct {
	config     ClientConfig
	httpClient HTTPClient
}

// NewOpenAIClient 创建一个新的OpenAI客户端
func NewOpenAIClient(config ClientConfig, httpClient HTTPClient) *OpenAIClient {
	if httpClient == nil {
		httpClient = DefaultHTTPClient(config.Timeout)
	}
	return &OpenAIClient{
		config:     config,
		httpClient: httpClient,
	}
}

// GenerateShellCommand 从自然语言描述生成shell命令
func (c *OpenAIClient) GenerateShellCommand(prompt string, system string) (string, error) {
	// 构建系统提示
	systemPrompt := `你是一个将自然语言转换为shell命令的工具。
请只返回shell命令，不需要任何解释或标记。
确保命令适用于` + system + `系统，并考虑命令的安全性。
如果无法确定用户的意图，可以提供最合理的猜测，并使用#添加注释说明。
你的任务是将用户的中文描述转换为有效的shell命令。`

	// 构建API请求
	request := OpenAIRequest{
		Model: c.config.Model,
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: prompt},
		},
		MaxTokens:   c.config.MaxTokens,
		Temperature: c.config.Temperature,
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("JSON编码错误: %v", err)
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", c.config.APIURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("API请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	// 检查状态码
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API错误 %d: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var response OpenAIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 提取命令
	if len(response.Choices) > 0 {
		command := response.Choices[0].Message.Content
		command = strings.TrimSpace(command)

		// 移除代码块标记（如果有）
		command = strings.TrimPrefix(command, "```bash")
		command = strings.TrimPrefix(command, "```zsh")
		command = strings.TrimPrefix(command, "```sh")
		command = strings.TrimPrefix(command, "```shell")
		command = strings.TrimPrefix(command, "```powershell")
		command = strings.TrimPrefix(command, "```")
		command = strings.TrimSuffix(command, "```")
		command = strings.TrimSpace(command)

		return command, nil
	}

	return "", fmt.Errorf("未能生成有效的命令")
}
