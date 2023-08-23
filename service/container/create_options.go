package container

import (
	"os"

	"github.com/Alan-333333/go-container-manager/pkg/config"
	"gopkg.in/yaml.v2"
)

const YAML_PATH = "/configs/"

const YAML_FILE_EXTENSON = ".yaml"

type CreateOptions struct {
	Name        string
	Image       string
	Memory      int64 // 内存限制
	CPU         int64 // CPU限制
	NetworkMode string
	Binds       []string // 挂载卷
	Cmd         []string
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
