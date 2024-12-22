package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Run() {
	filePath := ".codegate/pre-commit"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error : %s", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		cmd := exec.Command("bash", "-c", line)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %s", err)
			continue
		}
		fmt.Printf("%s", output)
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
