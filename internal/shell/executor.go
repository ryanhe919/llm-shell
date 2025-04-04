package shell

import (
	"fmt"
	"os"
	"os/exec"
)

// CommandExecutor 提供shell命令执行功能
type CommandExecutor struct {
	// 可扩展为包含更多配置选项
}

// NewCommandExecutor 创建一个新的命令执行器
func NewCommandExecutor() *CommandExecutor {
	return &CommandExecutor{}
}

// Execute 执行给定的shell命令
func (e *CommandExecutor) Execute(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("命令执行失败: %v", err)
	}
	return nil
}

// ExecuteWithConfirmation 执行命令前请求用户确认
func (e *CommandExecutor) ExecuteWithConfirmation(command string) error {
	fmt.Printf("\n是否执行此命令? [y/N] ")
	var confirm string
	fmt.Scanln(&confirm)

	if confirm == "y" || confirm == "Y" {
		fmt.Println("执行命令...")
		return e.Execute(command)
	}

	fmt.Println("已取消执行")
	return nil
}
