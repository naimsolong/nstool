package Nginx

import (
	"fmt"
	"os"

	Helper "naimsolong/nstool/helper"
	Service "naimsolong/nstool/service"
	Validator "naimsolong/nstool/validation"

	"github.com/manifoldco/promptui"
)

func List() bool {
	files, err := Service.Read_directory("/etc/nginx/sites-available")
	if err != nil {
		fmt.Printf("Process end\n")
		return false
	}

	prompt := promptui.Select{
		Label: "Which NGINX Configuration files?",
		Items: files,
	}
	_, selected_file, err := prompt.Run()

	if err != nil {
		fmt.Printf("%q\n", err)
		return false
	}

	file_name := "/etc/nginx/sites-enabled/" + selected_file
	Helper.Show_content(file_name)

	return true
}

func Add() bool {
	prompt_1 := promptui.Prompt{
		Label:    "Using which URL?",
		Validate: Validator.Not_empty_string,
	}
	url, err := prompt_1.Run()
	if err != nil {
		fmt.Println(err)
		return false
	}
	
	if _, err := os.Stat("/etc/nginx/sites-available/" + url)
	err == nil {
		fmt.Printf("NGINX Configuration exist!\n")
		return false
	}

	prompt_2 := promptui.Prompt{
		Label:    "On which Project Path?",
		Validate: Validator.Not_empty_string,
	}
	path, err := prompt_2.Run()
	if err != nil {
		fmt.Println(err)
		return false
	}

	// prompt_3 := promptui.Select{
	// 	Label: "Which PHP Version?",
	// 	Items: []string{"8.2", "8.1", "8.0", "7.4", "7.3", "7.2", "7.1"},
	// }
	// _, php_version, err := prompt_3.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return false
	// }

	prompt_4 := promptui.Select{
		Label: "Which Laravel Version?",
		Items: []string{"10", "9", "8"},
	}
	_, laravel_version, err := prompt_4.Run()
	if err != nil {
		fmt.Println(err)
		return false
	}

	php_version := ""
    switch laravel_version {
		case "10":
			php_version = "8.1"
		case "9":
			php_version = "8.0"
		case "8":
			php_version = "7.4"
    }

	prompt_5 := promptui.Select{
		Label:     "Is it Laravel Octane?",
		Items: []string{"Yes", "No"},
	}
	_, octane_flag, err := prompt_5.Run()
	if err != nil {
		fmt.Println(err)
		return false
	}

	octane_port := "8000"
	if octane_flag == "Yes" {
		prompt_6 := promptui.Prompt{
			Label:    "On which port?",
			Validate: Validator.Not_empty_string,
		}
		octane_port, err = prompt_6.Run()
		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	Helper.Clear_screen()

	fmt.Printf("URL : %q\n", url)
	fmt.Printf("Project Path : %q\n", path)
	fmt.Printf("PHP Version : %q\n", php_version)
	fmt.Printf("Laravel Version : %q\n", laravel_version)
	fmt.Printf("Octane Support : %q\n", octane_flag)
	fmt.Printf("Octane Port : %q\n", octane_port)

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
		stub_file := ""
		if octane_flag == "Yes" {
			stub_file = "./stub/nginx.laravel-octane.stub"
		} else {
			stub_file = "./stub/nginx.laravel.stub"
		}
		destination := "/etc/nginx/sites-available/" + url
		symlink := "/etc/nginx/sites-enabled/" + url

		copy_err := Helper.Copy_file(stub_file, destination)
		if copy_err != nil {
			fmt.Printf("%q\n", copy_err)
		}

		Helper.Replace_string_in_file(destination, "{{URL}}", url)
		Helper.Replace_string_in_file(destination, "{{PROJECT_PATH}}", path+"/public")

		if octane_flag == "Yes" {
			fmt.Printf("Write Octane")
			Helper.Replace_string_in_file(destination, "{{OCTANE_PORT}}", octane_port)
		} else {
			fmt.Printf("Write PHP")
			Helper.Replace_string_in_file(destination, "{{PHP_VERSION}}", php_version)
		}

		Service.Create_symlink(destination, symlink)

		Service.Reload_restart_nginx()

		fmt.Println("NGINX configuration file added")
		fmt.Println(destination)
	} else {
		fmt.Printf("Process end\n")
	}

	return true
}

func Remove() bool {
	files, err := Service.Read_directory("/etc/nginx/sites-available")
	if err != nil {
		fmt.Printf("Process end\n")
		return false
	}

	prompt := promptui.Select{
		Label: "Which NGINX Configuration files?",
		Items: files,
	}
	_, selected_file, err := prompt.Run()

	if err != nil {
		fmt.Printf("%q\n", err)
		return false
	}

	prompt_confirmation := promptui.Select{
		Label:     "Are you double confirm sure to delete this file?",
		Items: []string{"Yes", "No"},
	}
	_, confirmation, err := prompt_confirmation.Run()
	if err != nil {
		fmt.Printf("Process end\n")
		return false
	}

	if confirmation == "Yes" {
		file_name := "/etc/nginx/sites-enabled/" + selected_file
		Service.Remove_file(file_name)

		file_name = "/etc/nginx/sites-available/" + selected_file
		Service.Remove_file(file_name)

		Service.Reload_restart_nginx()

		fmt.Printf("NGINX configuration file (%q) deleted\n", file_name)
	} else {
		fmt.Printf("Process end\n")
	}

	return true
}
