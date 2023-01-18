package Helper

import (
    "fmt"
    "io"
	"io/ioutil"
	"os"
    "os/exec"
	
	"log"
	"strings"
)

func Clear_screen() {
	Run_cmd("clear")
}

func Run_cmd(command string) {
    c := exec.Command(command)
    c.Stdout = os.Stdout
    c.Run()
}

func Copy_file(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, 20)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}

func Replace_string_in_file(file string, search string, replace string) {
    input, err := ioutil.ReadFile(file)
    if err != nil {
		log.Fatalln(err)
    }

    lines := strings.Split(string(input), "\n")

    for i, line := range lines {
		if strings.Contains(line, search) {
			lines[i] = strings.ReplaceAll(lines[i], search, replace)
		}
    }
    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile(file, []byte(output), 0644)
    if err != nil {
		log.Fatalln(err)
    }
}