package Nginx

import (
    "fmt"

	"naimsolong/nstool/helper"
	"naimsolong/nstool/validation"
	"naimsolong/nstool/service"

	"github.com/manifoldco/promptui"
)

var octane_port string 


func List() {
	files, err := Service.Read_directory("/etc/nginx/sites-available")
	if err != nil {
		fmt.Printf("Process end")
		return
	}

	prompt := promptui.Select{
		Label: "Which NGINX Configuration files?",
		Items: files,
	}
	_, selected_file, err := prompt.Run()

	if err != nil {
		fmt.Printf("%q\n", err)
		return
	}
	
	file_name := "/etc/nginx/sites-enabled/"+selected_file
	Helper.Show_content(file_name)

	return
}

func Add() {
	prompt_1 := promptui.Prompt{
		Label: "Using which URL?",
		Validate: Validator.Not_empty_string,
	}
	url, err := prompt_1.Run()

	prompt_2 := promptui.Prompt{
		Label: "On which Project Path?",
		Validate: Validator.Not_empty_string,
	}
	path, err := prompt_2.Run()

	prompt_3 := promptui.Select{
		Label: "Which PHP Version?",
		Items: []string{"8.2", "8.1", "8.0", "7.4", "7.3",
			"7.2", "7.1"},
	}
	_, php_version, err := prompt_3.Run()

	prompt_4 := promptui.Select{
		Label: "Which Laravel Version?",
		Items: []string{"8"},
	}
	_, laravel_version, err := prompt_4.Run()

	prompt_5 := promptui.Prompt{
		Label: "Is it Laravel Octane?",
		IsConfirm: true,
	}
	octane_flag, err := prompt_5.Run()

	if(octane_flag == "y") {
		prompt_6 := promptui.Prompt{
			Label: "On which port?",
			Validate: Validator.Not_empty_string,
		}
		octane_port, err = prompt_6.Run()
	}

    Helper.Clear_screen()

	fmt.Printf("URL : %q\n", url)
	fmt.Printf("Project Path : %q\n", path)
	fmt.Printf("PHP Version : %q\n", php_version)
	fmt.Printf("Laravel Version : %q\n", laravel_version)
	fmt.Printf("Octane Support : %q\n", octane_flag)
	fmt.Printf("Octane Port : %q\n", octane_port)

	prompt_confirmation := promptui.Prompt{
		Label: "Please check detail above. Confirm?",
		IsConfirm: true,
	}
	confirmation, err := prompt_confirmation.Run()
	if err != nil {
		fmt.Printf("Process end")
		return
	}

    Helper.Clear_screen()

    if(confirmation == "y") {
		stub_file := "./stub/nginx.laravel"+laravel_version+".stub"
		destination := "/etc/nginx/sites-available/"+url
		symlink := "/etc/nginx/sites-enabled/"+url
		
        copy_err := Helper.Copy_file(stub_file, destination)
		if copy_err != nil {
			fmt.Printf("%q\n", copy_err)
		}

		Helper.Replace_string_in_file(destination, "{{URL}}", url)
		Helper.Replace_string_in_file(destination, "{{PROJECT_PATH}}", path+"/public")

		if(octane_flag == "y") {
			Helper.Replace_string_in_file(destination, "{{OCTANE_PORT}}", octane_port)
		} else {
			Helper.Replace_string_in_file(destination, "{{PHP_VERSION}}", php_version)
		}

		Service.Create_symlink(destination, symlink)

		Service.Reload_restart_nginx()

        fmt.Printf("NGINX configuration file (%q) added\n", destination)
    } else {
		fmt.Printf("Process end")
	}

	return
}

func Remove() {
	files, err := Service.Read_directory("/etc/nginx/sites-available")
	if err != nil {
		fmt.Printf("Process end")
		return
	}

	prompt := promptui.Select{
		Label: "Which NGINX Configuration files?",
		Items: files,
	}
	_, selected_file, err := prompt.Run()

	if err != nil {
		fmt.Printf("%q\n", err)
		return
	}

	prompt_confirmation := promptui.Prompt{
		Label: "Are you double confirm sure to delete this file?",
		IsConfirm: true,
	}
	confirmation, err := prompt_confirmation.Run()
	if err != nil {
		fmt.Printf("Process end")
		return
	}

    if(confirmation == "y") {
		file_name := "/etc/nginx/sites-enabled/"+selected_file
		Service.Remove_file(file_name)

		file_name = "/etc/nginx/sites-available/"+selected_file
		Service.Remove_file(file_name)
		
		Service.Reload_restart_nginx()
		
		fmt.Printf("NGINX configuration file (%q) deleted\n", file_name)
    } else {
		fmt.Printf("Process end")
	}
	
	return
}