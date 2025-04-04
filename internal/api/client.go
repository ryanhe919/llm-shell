package api

import (
	"net/http"
	"time"
)

// Client 定义API客户端接口
type Client interface {
	GenerateShellCommand(prompt string) (string, error)
}

// ClientConfig 包含API客户端配置
type ClientConfig struct {
	APIKey      string
	APIURL      string
	Model       string
	MaxTokens   int
	Temperature float64
	Timeout     time.Duration
}

// HTTPClient 是HTTP请求的接口，便于测试
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// DefaultHTTPClient 返回默认的HTTP客户端
func DefaultHTTPClient(timeout time.Duration) HTTPClient {
	return &http.Client{
		Timeout: timeout,
	}
}
