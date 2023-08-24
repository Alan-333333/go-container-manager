package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Alan-333333/go-container-manager/service/container"
)

type Input struct {
	command string
	params  []string
}

func main() {

	manager, err := container.NewContainerManager()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		help()
	}

	go startCLI(manager)

	forever := make(chan bool)
	<-forever
}

func startCLI(manager *container.ContainerManager) {

	for {
		input := parseInput()
		switch input.command {

		case "build":
			if len(input.params) < 1 {
				helpBuild()
				return
			}

			image := input.params[0]
			imageOption, err := container.LoadConfig(image)
			if err != nil {
				log.Fatal(err)
			}

			err = container.CloneRepository(imageOption.RepoURL, imageOption.Name)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("repo clone finish")
			id, err := manager.Create(imageOption)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Created container", id)
		case "create":
			if len(input.params) < 1 {
				helpCreate()
				return
			}

			image := input.params[0]
			imageOption, err := container.LoadConfig(image)
			if err != nil {
				log.Fatal(err)
			}

			err = container.CloneRepository(imageOption.RepoURL, imageOption.Name)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("repo clone finish")
			id, err := manager.Create(imageOption)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Created container", id)

		case "start":
			if len(input.params) < 1 {
				helpStart()
				return
			}
			id := input.params[0]
			err := manager.Start(id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Started container", id)

		case "remove":
			{
				if len(input.params) < 1 {
					helpRemove()
					return
				}
				id := input.params[0]
				err := manager.Remove(id)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Remove container", id)
			}

		default:
			help()
		}
	}

}

func help() {
	fmt.Println("Usage:")
	fmt.Println("  build <image>     Build a new Image")
	fmt.Println("  create <image>     Create a new container")
	fmt.Println("  start <id>         Start a container")
	fmt.Println("  stop <id>          Stop a running container")
	fmt.Println("  remove <id>        Remove a stopped container")
}

func helpBuild() {
	fmt.Println("Usage: build <image>")
	fmt.Println("Example: build nginx:latest")
}
func helpCreate() {
	fmt.Println("Usage: create <image>")
	fmt.Println("Example: create nginx:latest")
}

func helpStart() {
	fmt.Println("Usage: start <id>")
	fmt.Println("Example: start c01cd209fb12")
}

func helpRemove() {
	fmt.Println("Usage: remove <id>")
	fmt.Println("Example: remove c01cd209fb12")
}

// parseInput parses user input into command and parameters
func parseInput() Input {

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	cmd := strings.TrimSpace(input)

	params := strings.Split(cmd, " ")

	return Input{
		command: params[0],
		params:  params[1:],
	}
}
