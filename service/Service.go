package Service

import (
	"os"
    "io/ioutil"
    "log"

	"naimsolong/nstool/helper"
)

var directory_array []string

func Create_symlink(path string, target string) {
    os.Symlink(target, target)
}

func Reload_restart_nginx() {
	Helper.Run_cmd("sudo systemctl reload nginx")
	Helper.Run_cmd("sudo systemctl restart nginx")
}

func Reload_restart_supervisor() {
	Helper.Run_cmd("sudo systemctl reload supervisor")
	Helper.Run_cmd("sudo systemctl restart supervisor")
}

func Read_directory(dirname string) ([]string, error) {
    files, err := ioutil.ReadDir(dirname)
    
	if err != nil {
        log.Fatal(err)
    } else {
		for _, file := range files {
			if(!file.IsDir()) {
				directory_array = append(directory_array, file.Name(),)
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