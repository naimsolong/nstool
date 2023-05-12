package Env

import (
    "fmt"
	"os"
    
	Helper "naimsolong/nstool/helper"
	Init "naimsolong/nstool/init"
	Service "naimsolong/nstool/service"

	"github.com/manifoldco/promptui"
)

var env_path = Init.Get_init_path() + "/env"
var env_backup_path = env_path + "/backup"

func Copy() bool {
	files, err := Service.Read_directory(Init.Get_value("Project_Path"))
	if err != nil {
		fmt.Printf("Process end\n")
		return false
	}

	prompt := promptui.Select{
		Label: "For which project?",
		Items: files,
	}
	_, selected_project, err := prompt.Run()

	if err != nil {
		fmt.Printf("%q\n", err)
		return false
	}
	
    selected_project_path := Init.Get_value("Project_Path") + "/" + selected_project
    selected_project_env_file := selected_project_path + "/.env"
    selected_project_env_example_file := selected_project_path + "/.env.example"

	if _, err := os.Stat(selected_project_env_file)
	err == nil {
		fmt.Printf(".env file exist!\n")
		return false
	}
	
	if _, err := os.Stat(selected_project_env_example_file)
	err == nil {
        prompt_confirmation := promptui.Select{
            Label:     "Do you want to copy from existing .env.example?",
            Items: []string{"Yes", "No"},
        }
        _, confirmation, err := prompt_confirmation.Run()
        if err != nil {
            fmt.Printf("Process end\n")
            return false
        }
        
        if confirmation == "Yes" {
            copy_err := Helper.Copy_file(selected_project_env_example_file, selected_project_env_file)
            if copy_err != nil {
                fmt.Printf("%q\n", copy_err)
            }
        }
		return false
	}

	prompt_laravel := promptui.Select{
		Label: "Which Laravel Version?",
		Items: []string{"10", "9", "8"},
	}
	_, laravel_version, err := prompt_laravel.Run()
	if err != nil {
		fmt.Println(err)
		return false
	}

	Helper.Clear_screen()

	fmt.Printf("Selected project : %q\n", selected_project)
	fmt.Printf("Laravel Version : %q\n", laravel_version)

	prompt_confirmation := promptui.Select{
		Label:     "Please check detail above. Confirm?",
		Items: []string{"Yes", "No"},
	}
	_, confirmation, err := prompt_confirmation.Run()
	if err != nil || confirmation == "No" {
		fmt.Printf("Process end\n")
		return false
	}

	Helper.Clear_screen()

	if confirmation == "Yes" {
		stub_file := "./stub/env.laravel-" + laravel_version + ".stub"

		copy_err := Helper.Copy_file(stub_file, selected_project_env_file)
		if copy_err != nil {
			fmt.Printf("%q\n", copy_err)
		}
	} else {
		fmt.Printf("Process end\n")
	}

	return true
}

func Change() {
    fmt.Println("env change")
}