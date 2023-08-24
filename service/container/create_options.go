package container

import (
	"os"
	"path/filepath"

	"github.com/Alan-333333/go-container-manager/pkg/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"gopkg.in/yaml.v2"
)

const YAML_PATH = "/configs/"

const YAML_FILE_EXTENSON = ".yaml"

type CreateOptions struct {
	Name        string   `yaml:"name"`
	Image       string   `yaml:"image"`
	RepoURL     string   `yaml:"repo_url"`
	Memory      int64    `yaml:"memory"` // 内存限制
	CPU         int64    `yaml:"cpu"`    // CPU限制
	NetworkMode string   `yaml:"network"`
	Binds       []string `yaml:"binds"` // 挂载卷
	Cmd         []string `yaml:"cmd"`
}

func LoadConfig(filename string) (CreateOptions, error) {

	var opts CreateOptions

	yamlFilePath := config.PORJECT_ROOT + YAML_PATH + filename + YAML_FILE_EXTENSON

	yamlFile, err := os.ReadFile(yamlFilePath)
	if err != nil {
		return opts, err
	}

	err = yaml.Unmarshal(yamlFile, &opts)
	if err != nil {
		return opts, err
	}

	return opts, nil

}

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
