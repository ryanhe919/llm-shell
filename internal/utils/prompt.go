package utils

import "strings"

// SystemPromptForShellCommand 返回用于生成shell命令的系统提示
func SystemPromptForShellCommand() string {
	return `你是一个将自然语言转换为shell命令的工具。
请只返回shell命令，不需要任何解释或标记。
确保命令适用于Linux/macOS系统，并考虑命令的安全性。
如果无法确定用户的意图，可以提供最合理的猜测，并使用#添加注释说明。
你的任务是将用户的中文描述转换为有效的shell命令。`
}

// CleanCommandOutput 清理命令输出（去除代码块标记等）
func CleanCommandOutput(command string) string {
	// 移除常见的代码块标记
	commonPrefixes := []string{
		"```bash",
		"```sh",
		"```shell",
		"```",
	}

	for _, prefix := range commonPrefixes {
		if len(command) >= len(prefix) && command[:len(prefix)] == prefix {
			command = command[len(prefix):]
			break
		}
	}

	// 移除结尾的代码块标记
	if len(command) >= 3 && command[len(command)-3:] == "```" {
		command = command[:len(command)-3]
	}

	// 清理空白字符
	command = strings.TrimSpace(command)

	return command
}
