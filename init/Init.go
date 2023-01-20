package Init

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	Service "naimsolong/nstool/service"
	Validator "naimsolong/nstool/validation"

	"github.com/manifoldco/promptui"
)

type JSONValue struct {
	nginx_sites_available_path, nginx_sites_enable_path, project_path string
}

func Default() JSONValue {
	data := JSONValue{
		nginx_sites_available_path: "/etc/nginx/sites-available",
		nginx_sites_enable_path:    "/etc/nginx/sites-enabled",
		project_path:               "/var/www",
	}

	return data
}

func Start(delete bool) bool {
	var data JSONValue

	prompt_1 := promptui.Prompt{
		Label:     "Using default value?",
		IsConfirm: true,
	}
	default_flag, err := prompt_1.Run()
	if err != nil {
		fmt.Println(err)
		return false
	}

	if default_flag == "y" {
		data = Default()
	} else {
		prompt_2 := promptui.Prompt{
			Label:    "Please state NGINX sites-available path (e.g.: /etc/nginx/sites-available) :",
			Validate: Validator.Not_empty_string,
		}
		nginx_sites_available_path, err := prompt_2.Run()
		if err != nil {
			fmt.Println(err)
			return false
		}

		prompt_3 := promptui.Prompt{
			Label:    "Please state NGINX sites-enable path (e.g.: /etc/nginx/sites-enabled) :",
			Validate: Validator.Not_empty_string,
		}
		nginx_sites_enable_path, err := prompt_3.Run()
		if err != nil {
			fmt.Println(err)
			return false
		}

		prompt_4 := promptui.Prompt{
			Label:    "Please state project base path (e.g.: /var/www) :",
			Validate: Validator.Not_empty_string,
		}
		project_path, err := prompt_4.Run()
		if err != nil {
			fmt.Println(err)
			return false
		}

		data = JSONValue{
			nginx_sites_available_path: nginx_sites_available_path,
			nginx_sites_enable_path:    nginx_sites_enable_path,
			project_path:               project_path,
		}
	}

	if delete {
		Service.Remove_file("./init.json")
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("./init.json", file, 0644)

	return true
}
