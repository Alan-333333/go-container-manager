package container

type ContainerManager interface {
	Create(image string) (string, error)
	Remove(id string) error
	Start(id string) error
	Stop(id string) error
}
