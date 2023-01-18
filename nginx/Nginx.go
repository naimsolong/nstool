package Nginx

import (
    "fmt"

	"naimsolong/nstool/helper"

	"github.com/manifoldco/promptui"
)

func Add() {
	prompt_1 := promptui.Prompt{
		Label: "URL",
	}
	url, err := prompt_1.Run()
	if err != nil {
		fmt.Printf("Prompt failed on URL %v\n", err)
		return
	}

	prompt_2 := promptui.Prompt{
		Label: "Path (without /public path)",
	}
	path, err := prompt_2.Run()
	if err != nil {
		fmt.Printf("Prompt failed on Path %v\n", err)
		return
	}

	prompt_3 := promptui.Select{
		Label: "Select PHP Version",
		Items: []string{"8.2", "8.1", "8.0", "7.4", "7.3",
			"7.2", "7.1"},
	}
	_, php_version, err := prompt_3.Run()
	if err != nil {
		fmt.Printf("Prompt failed on Select PHP Version %v\n", err)
		return
	}

	prompt_4 := promptui.Select{
		Label: "Select Laravel Version",
		Items: []string{"8"},
	}
	_, laravel_version, err := prompt_4.Run()
	if err != nil {
		fmt.Printf("Prompt failed on Select Laravel Version %v\n", err)
		return
	}

	prompt_5 := promptui.Prompt{
		Label: "Is it Laravel Octane",
		IsConfirm: true,
	}
	octane_flag, err := prompt_5.Run()
	if err != nil {
		fmt.Printf("Prompt failed on Is it Laravel Octane %v\n", err)
		return
	}

    Helper.Clear_screen()

	fmt.Printf("URL : %q\n", url)
	fmt.Printf("Project Path : %q\n", path)
	fmt.Printf("PHP Version : %q\n", php_version)
	fmt.Printf("Laravel Version : %q\n", laravel_version)
	fmt.Printf("Octane : %q\n", octane_flag)

	prompt_confirmation := promptui.Prompt{
		Label: "Please check detail above. Confirm?",
		IsConfirm: true,
	}
	confirmation, err := prompt_confirmation.Run()
	if err != nil {
		fmt.Printf("Prompt failed on confirmation %v\n", err)
		return
	}
	fmt.Printf("confirmation : %q\n", confirmation)

    Helper.Clear_screen()

    if(confirmation == "y") {
        Helper.Copy("../stub/nginx.laravel"+laravel_version+".stub", "/etc/nginx/sites-available/"+url)
        fmt.Println("nginx added")
    }
}

func Remove() {
    fmt.Println("nginx remove")
}