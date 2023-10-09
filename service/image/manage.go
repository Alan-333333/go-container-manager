package image

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/Alan-333333/go-container-manager/pkg/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func BuildImage(name string) (string, error) {

	// 初始化docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "fail", err
	}

	// 获取项目根目录
	repoRoot := config.GetProjectRoot()

	// 指定克隆目标目录
	repoPath := filepath.Join(repoRoot, "tmp/repo/", name)

	// 发送build请求
	buildContext, err := archive.TarWithOptions(repoPath, &archive.TarOptions{})
	if err != nil {
		return "fail", err
	}

	buildOptions := types.ImageBuildOptions{
		Tags: []string{name},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	buildResponse, err := cli.ImageBuild(ctx, buildContext, buildOptions)
	if err != nil {
		return "fail", err
	}

	buildOutput, err := io.ReadAll(buildResponse.Body)
	if err != nil {
		return "fail", err
	}
	// 拼接日志完整路径
	logPath := filepath.Join(repoRoot, "tmp/image/createlog", name)

	// 获取日志文件夹路径
	logDir := filepath.Dir(logPath)

	// 如果文件夹不存在,则创建文件夹
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			return "fail", err
		}
	}

	// 创建日志文件
	logFile, err := os.Create(logPath)
	if err != nil {
		return "fail", err
	}

	// 写入构建日志
	_, err = logFile.Write(buildOutput)
	if err != nil {
		return "fail", err
	}

	return "success", nil

}
