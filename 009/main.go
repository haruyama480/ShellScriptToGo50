package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func main() {

	reader, writer := io.Pipe()

	go func() {
		cmd := exec.Command("ps", "aux")
		cmd.Stdout = writer
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Wait()
		writer.Close()
	}()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		processName := " go "
		if strings.Contains(line, processName) {
			fmt.Println(line)
		}
	}
}
