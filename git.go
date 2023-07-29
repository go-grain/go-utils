package go_utils

import (
	"fmt"
	"os"
	"os/exec"
)

var dir string

func InitGit(d string) {
	dir = d
	gitInit()
}

func gitInit() {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	add()
}

func add() {
	cmd := exec.Command("git", "add", ".")
	cmd.Stdout = os.Stdout
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	commit()
}

func commit() {
	cmd := exec.Command("git", "commit", "-m", "init")
	cmd.Stdout = os.Stdout
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
