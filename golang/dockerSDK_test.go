package dockersdk

import (
	"testing"
)

const (
	containerName = "abc"
)
func TestDocker_ContainerList(t *testing.T) {
	containers := GetClient().ContainerList()
	for _, container := range containers {
		t.Log(container)
	}
}
func TestDocker_ImageList(t *testing.T) {
	images := GetClient().ImageList()
	for _, image := range images {
		t.Log(image)
	}
}
func TestDocker_ImageInspect(t *testing.T) {
	a := GetClient().ImageInspect("redis:latest")
	t.Log("a", a)

}
func TestDocker_ImagePrune(t *testing.T) {
	result := GetClient().ImagePrune()
	t.Log(result)
}

func TestDocker_ImagePull(t *testing.T) {
	GetClient().ImagePull("ubuntu:latest", false)
}
func TestDocker_ContainerCreate(t *testing.T) {
	var name = "abc"
	//var config = ContainerConfig{
	//	name,
	//	"ubuntu:latest",
	//	[]string{},
	//	[]string{"date"},
	//	map[string]string{},
	//	map[int]int{},
	//	"",
	//	nil,
	//}
	GetClient().ContainerRemove(name)
	//var id = GetClient().ContainerCreate(config)
	//GetClient().ContainerStart(id)
}

func TestDocker_ContainerDeleteIfExist(t *testing.T) {
	var containerName = "abc"
	GetClient().ContainerDeleteIfExist(containerName)
}