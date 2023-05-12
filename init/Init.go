package Init

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
    "path/filepath"

	Service "naimsolong/nstool/service"
	Validator "naimsolong/nstool/validation"

	"github.com/manifoldco/promptui"
)

type ConfigValue struct {
	Nginx_Sites_Available_Path string `json:"Nginx_Sites_Available_Path"`
	Nginx_Sites_Enable_Path string `json:"Nginx_Sites_Enable_Path"`
	Project_Path string `json:"Project_Path"`
}

var init_path = "/etc/nstool"
var init_file = init_path + "/init.json"

func Default() ConfigValue {
	return ConfigValue{
		Nginx_Sites_Available_Path: "/etc/nginx/sites-available",
		Nginx_Sites_Enable_Path:    "/etc/nginx/sites-enabled",
		Project_Path:               "/var/www",
	}
}

func Start(delete bool) bool {
	_, err := os.Stat(init_path)
	if err != nil {
		newpath := filepath.Join("/etc", "nstool")
		err = os.MkdirAll(newpath, os.ModePerm)
	}

	data := Default()

	prompt_1 := promptui.Select{
		Label:     "Using default value?",
		Items: []string{"Yes", "No"},
	}
	_, default_flag, err := prompt_1.Run()
	if err != nil {
		fmt.Println(err)
		return false
	}

	if default_flag == "No" {
		prompt_2 := promptui.Prompt{
			Label:    "Please state NGINX sites-available path (e.g.: /etc/nginx/sites-available) :",
			Validate: Validator.Not_empty_string,
		}
		Nginx_Sites_Available_Path, err := prompt_2.Run()
		if err != nil {
			fmt.Println(err)
			return false
		}

		prompt_3 := promptui.Prompt{
			Label:    "Please state NGINX sites-enable path (e.g.: /etc/nginx/sites-enabled) :",
			Validate: Validator.Not_empty_string,
		}
		Nginx_Sites_Enable_Path, err := prompt_3.Run()
		if err != nil {
			fmt.Println(err)
			return false
		}

		prompt_4 := promptui.Prompt{
			Label:    "Please state project base path (e.g.: /var/www) :",
			Validate: Validator.Not_empty_string,
		}
		Project_Path, err := prompt_4.Run()
		if err != nil {
			fmt.Println(err)
			return false
		}

		data = ConfigValue{
			Nginx_Sites_Available_Path: Nginx_Sites_Available_Path,
			Nginx_Sites_Enable_Path:    Nginx_Sites_Enable_Path,
			Project_Path:               Project_Path,
		}
	}

	_, err = os.Stat(init_file)
	if delete && err == nil {
		Service.Remove_file(init_file)
	}

	file, _ := json.Marshal(data)
	
	err = os.WriteFile(init_file, file, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func Get_file() ConfigValue {
	var data ConfigValue
	
	_, err := os.Stat(init_file)
	if err == nil {
		input, err := ioutil.ReadFile(init_file)
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(input, &data)
	} else {
		data = Default()
	}
	
    return data
}

func Get_value(variable string) string {
	data := Get_file()
	value := "test"

    switch variable {
		case "Nginx_Sites_Available_Path":
			value = data.Nginx_Sites_Available_Path
		case "Nginx_Sites_Enable_Path":
			value = data.Nginx_Sites_Enable_Path
		case "Project_Path":
			value = data.Project_Path
    }

	return value
}