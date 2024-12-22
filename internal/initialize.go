package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func Initialize() {
	dirName := ".codegate"
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Printf("Error :", err)
		return
	}
	filePath := dirName + "/pre-commit"
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		fmt.Printf("Error :", err)
		return
	}
	addIgnoredHook()
	createPreCommitFile()
	addPermission()
	fmt.Println("File created successfully")
}
func createPreCommitFile() {
	content := `#!/bin/bash
codegate --mode run`
	filePath := ".git/hooks/pre-commit"
	err := os.WriteFile(filePath, []byte(content), 0755)
	if err != nil {
		fmt.Printf("Failed to create pre-commit file: %v\n", err)
		return
	}
}
func addPermission() {
	_ = exec.Command("chmod", "+x", ".git/hooks/pre-commit")
}
func addIgnoredHook() {
	_ = exec.Command("git", "config", "advice.ignoredHook", "false")
}
