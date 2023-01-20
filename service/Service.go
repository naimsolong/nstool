package Service

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var directory_array []string

func Create_symlink(path string, target string) {
	os.Symlink(target, target)
}

func Reload_restart_nginx() {
	c_1 := exec.Command("sudo", "systemctl", "reload", "nginx")
	c_1.Stdout = os.Stdout
	c_1.Run()

	c_2 := exec.Command("sudo", "systemctl", "restart", "nginx")
	c_2.Stdout = os.Stdout
	c_2.Run()
}

func Reload_restart_supervisor() {
	c_1 := exec.Command("sudo", "systemctl", "reload", "supervisor")
	c_1.Stdout = os.Stdout
	c_1.Run()

	c_2 := exec.Command("sudo", "systemctl", "restart", "supervisor")
	c_2.Stdout = os.Stdout
	c_2.Run()
}

func Read_directory(dirname string) ([]string, error) {
	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {
			if !file.IsDir() {
				directory_array = append(directory_array, file.Name())
			}
		}
	}

	return directory_array, err
}

func Remove_file(file_name string) error {
	e := os.Remove(file_name)
	if e != nil {
		log.Fatal(e)
	}

	return e
}
