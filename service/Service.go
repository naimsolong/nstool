package Service

import (
	"os"

	"naimsolong/nstool/helper"
)

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