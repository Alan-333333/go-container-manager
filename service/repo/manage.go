package repo

import (
	"os"
	"path/filepath"

	"github.com/Alan-333333/go-container-manager/pkg/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func CloneRepository(repoURL string, name string) error {

	// 获取项目根目录
	repoRoot := config.GetProjectRoot()

	// 指定克隆目标目录
	repoPath := filepath.Join(repoRoot, "tmp/repo/", name)

	// 创建目录
	err := os.MkdirAll(repoPath, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:           repoURL,
		ReferenceName: plumbing.ReferenceName("refs/heads/master"),
	})

	if err != nil {
		return err
	}

	return nil
}

func RemoveRepository(name string) error {
	// 移除仓库目录

	// 获取项目根目录
	repoRoot := config.GetProjectRoot()

	// 指定克隆目标目录
	repoPath := filepath.Join(repoRoot, "tmp/repo/", name)

	err := os.RemoveAll(repoPath)
	if err != nil {
		return err
	}

	return nil
}
